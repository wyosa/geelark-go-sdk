package clientcore

import (
	"context"
	"errors"
	"io"
	"net"
	"net/http"
	"net/http/httptest"
	"sort"
	"strings"
	"sync/atomic"
	"testing"
	"time"
)

func TestRejectsInvalidSuccessEnvelope(t *testing.T) {
	for _, body := range []string{`{}`, `null`, `{"code":null}`, `{"data":{"id":"phone"}}`, `{"code":"0"}`, `[]`, `<html>Bad gateway</html>`} {
		t.Run(body, func(t *testing.T) {
			srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
				_, _ = w.Write([]byte(body))
			}))
			defer srv.Close()
			c, err := New(Config{BaseURL: srv.URL, Token: "token"})
			if err != nil {
				t.Fatal(err)
			}
			if err := c.Do(context.Background(), "/open/v1/phone/list", nil, nil); err == nil {
				t.Fatalf("accepted malformed success envelope %s", body)
			}
		})
	}
}

func TestNon2xxStatusIsAlwaysAnError(t *testing.T) {
	for _, status := range []int{300, 304, 400, 429, 500} {
		srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			w.WriteHeader(status)
			_, _ = w.Write([]byte(`{"code":0,"data":null}`))
		}))
		c, _ := New(Config{BaseURL: srv.URL, Token: "token"})
		err := c.Do(context.Background(), "/open/v1/phone/list", nil, nil)
		srv.Close()
		var apiErr *APIError
		if !errors.As(err, &apiErr) || apiErr.HTTPStatus != status {
			t.Errorf("HTTP %d: got %v, want APIError with HTTP status", status, err)
		}
	}
}

func TestCancelledWaitsDoNotReserveRateSlots(t *testing.T) {
	c, _ := New(Config{BaseURL: DefaultBaseURL, Token: "token", MinInterval: time.Hour})
	path := "/open/v1/phone/list"
	if err := c.waitTurn(context.Background(), path); err != nil {
		t.Fatal(err)
	}
	want := c.nextRequest[path]
	ctx, cancel := context.WithCancel(context.Background())
	cancel()
	for i := 0; i < 5; i++ {
		if err := c.waitTurn(ctx, path); !errors.Is(err, context.Canceled) {
			t.Fatalf("got %v, want cancellation", err)
		}
	}
	ctx, cancel = context.WithTimeout(context.Background(), time.Millisecond)
	defer cancel()
	if err := c.waitTurn(ctx, path); !errors.Is(err, context.DeadlineExceeded) {
		t.Fatalf("got %v, want deadline exceeded", err)
	}
	if got := c.nextRequest[path]; !got.Equal(want) {
		t.Fatalf("cancelled calls added %v to the rate limit queue", got.Sub(want))
	}
}

func TestMarshalFailureDoesNotReserveRateSlot(t *testing.T) {
	c, _ := New(Config{BaseURL: DefaultBaseURL, Token: "token"})
	path := "/open/v1/phone/list"
	if err := c.Do(context.Background(), path, make(chan int), nil); err == nil {
		t.Fatal("expected JSON marshal error")
	}
	if !c.nextRequest[path].IsZero() {
		t.Fatal("invalid request consumed a rate limit slot")
	}
}

func TestUploadHonorsCustomClientTimeout(t *testing.T) {
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		time.Sleep(150 * time.Millisecond)
		w.WriteHeader(http.StatusOK)
	}))
	defer srv.Close()
	hc := &http.Client{Timeout: 20 * time.Millisecond}
	c, _ := New(Config{BaseURL: DefaultBaseURL, Token: "token", HTTPClient: hc})
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	err := c.Upload(ctx, http.MethodPut, srv.URL, nil, strings.NewReader("file"), 4)
	var timeoutErr net.Error
	if !errors.As(err, &timeoutErr) || !timeoutErr.Timeout() || ctx.Err() != nil {
		t.Fatalf("upload ignored HTTPClient.Timeout: %v", err)
	}
	if hc.Timeout != 20*time.Millisecond || hc.CheckRedirect != nil {
		t.Fatal("SDK modified the supplied HTTP client")
	}
}

func TestBaseURLRejectsEmptyQueryAndFragment(t *testing.T) {
	for _, base := range []string{"https://example.com?", "https://example.com#", "https://:443"} {
		if _, err := New(Config{BaseURL: base, Token: "token"}); err == nil {
			t.Errorf("accepted malformed base URL %q", base)
		}
	}
}

type roundTripFunc func(*http.Request) (*http.Response, error)

func (f roundTripFunc) RoundTrip(r *http.Request) (*http.Response, error) { return f(r) }

func TestConcurrentRequestsRespectPerPathLimit(t *testing.T) {
	const count = 6
	const interval = 20 * time.Millisecond
	arrivals := make(chan time.Time, count)
	hc := &http.Client{Transport: roundTripFunc(func(r *http.Request) (*http.Response, error) {
		arrivals <- time.Now()
		return &http.Response{StatusCode: 200, Body: io.NopCloser(strings.NewReader(`{"code":0}`)), Header: make(http.Header)}, nil
	})}
	c, _ := New(Config{BaseURL: DefaultBaseURL, Token: "token", HTTPClient: hc, MinInterval: interval})
	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()
	start := make(chan struct{})
	errorsCh := make(chan error, count)
	for i := 0; i < count; i++ {
		go func() {
			<-start
			errorsCh <- c.Do(ctx, "/open/v1/phone/list", nil, nil)
		}()
	}
	close(start)
	for i := 0; i < count; i++ {
		if err := <-errorsCh; err != nil {
			t.Fatal(err)
		}
	}
	times := make([]time.Time, count)
	for i := range times {
		times[i] = <-arrivals
	}
	sort.Slice(times, func(i, j int) bool { return times[i].Before(times[j]) })
	for i := 1; i < len(times); i++ {
		if gap := times[i].Sub(times[i-1]); gap < interval/2 {
			t.Fatalf("concurrent calls bypassed the rate limiter: gap %s", gap)
		}
	}
	// A different interface has its own budget, even while the first is blocked.
	c.mu.Lock()
	c.nextRequest["/open/v1/phone/list"] = time.Now().Add(time.Hour)
	c.mu.Unlock()
	if err := c.Do(ctx, "/open/v1/phone/status", nil, nil); err != nil {
		t.Fatalf("independent endpoint blocked: %v", err)
	}
}

func TestRequestsHonorInFlightCancellation(t *testing.T) {
	for _, upload := range []bool{false, true} {
		hc := &http.Client{Transport: roundTripFunc(func(r *http.Request) (*http.Response, error) {
			<-r.Context().Done()
			return nil, r.Context().Err()
		})}
		c, _ := New(Config{BaseURL: DefaultBaseURL, Token: "token", HTTPClient: hc})
		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Millisecond)
		var err error
		if upload {
			err = c.Upload(ctx, http.MethodPut, "https://storage.example/upload?Signature=private", nil, strings.NewReader("a"), 1)
		} else {
			err = c.Do(ctx, "/open/v1/phone/list", nil, nil)
		}
		cancel()
		if !errors.Is(err, context.DeadlineExceeded) {
			t.Errorf("upload=%v: cancellation lost: %v", upload, err)
		}
		if upload && err != nil && strings.Contains(err.Error(), "private") {
			t.Errorf("signed URL leaked into upload error: %v", err)
		}
	}
}

func TestUploadRefusesRedirectsAndCredentialHeaders(t *testing.T) {
	var calls atomic.Int32
	target := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		calls.Add(1)
	}))
	defer target.Close()
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		http.Redirect(w, r, target.URL, http.StatusTemporaryRedirect)
	}))
	defer srv.Close()
	c, _ := New(Config{BaseURL: DefaultBaseURL, Token: "token"})
	if err := c.Upload(context.Background(), http.MethodPut, srv.URL+"?Signature=private", nil, strings.NewReader("a"), 1); err == nil || strings.Contains(err.Error(), "private") {
		t.Fatalf("redirect must fail without exposing the signed URL: %v", err)
	}
	for _, key := range []string{"Authorization", "APPID", "sign", "traceId"} {
		if err := c.Upload(context.Background(), http.MethodPut, target.URL, map[string]string{key: "secret"}, strings.NewReader("a"), 1); err == nil {
			t.Errorf("accepted credential header %s", key)
		}
	}
	if calls.Load() != 0 {
		t.Fatalf("unsafe upload reached the target %d times", calls.Load())
	}
}

func TestResponseSizeLimit(t *testing.T) {
	hc := &http.Client{Transport: roundTripFunc(func(r *http.Request) (*http.Response, error) {
		return &http.Response{StatusCode: 200, Body: io.NopCloser(strings.NewReader(strings.Repeat(" ", maxResponseBytes+1)))}, nil
	})}
	c, _ := New(Config{BaseURL: DefaultBaseURL, Token: "token", HTTPClient: hc})
	if err := c.Do(context.Background(), "/open/v1/phone/list", nil, nil); err == nil || !strings.Contains(err.Error(), "body exceeds") {
		t.Fatalf("oversized API response accepted: %v", err)
	}
	if err := c.Upload(context.Background(), http.MethodPut, "https://storage.example/upload", nil, nil, 0); err == nil || !strings.Contains(err.Error(), "body exceeds") {
		t.Fatalf("oversized upload response accepted: %v", err)
	}
}
