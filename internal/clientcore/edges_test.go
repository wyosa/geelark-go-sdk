package clientcore

import (
	"bytes"
	"context"
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"sort"
	"strings"
	"sync"
	"sync/atomic"
	"testing"
	"time"
)

type auditReadClose struct {
	io.Reader
	closed bool
}

func (r *auditReadClose) Close() error { r.closed = true; return nil }

type auditErrorReader struct{ err error }

func (r auditErrorReader) Read([]byte) (int, error) { return 0, r.err }

func TestAuditBodyLifecycle(t *testing.T) {
	sentinel := errors.New("audit read failure")
	cases := []struct {
		name                   string
		status                 int
		body                   string
		readErr, errorExpected bool
	}{
		{"success", 200, `{"code":0,"data":{"id":"9007199254740993"}}`, false, false},
		{"business_error", 200, `{"code":40006,"data":{"id":"9007199254740993"}}`, false, true},
		{"http_error", 503, `upstream failure`, false, true},
		{"bad_json", 200, `{broken`, false, true},
		{"wrong_data_type", 200, `{"code":0,"data":[1]}`, false, true},
		{"read_error", 200, ``, true, true},
		{"limit", 200, strings.Repeat("x", maxResponseBytes+1), false, true},
	}
	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			reader := &auditReadClose{Reader: strings.NewReader(tc.body)}
			if tc.readErr {
				reader.Reader = auditErrorReader{sentinel}
			}
			calls := 0
			c, err := New(Config{BaseURL: DefaultBaseURL, Token: "audit", HTTPClient: &http.Client{Transport: roundTripFunc(func(r *http.Request) (*http.Response, error) {
				calls++
				return &http.Response{StatusCode: tc.status, Body: reader, Header: make(http.Header)}, nil
			})}})
			if err != nil {
				t.Fatal(err)
			}
			var out struct{ ID string }
			err = c.Do(context.Background(), "/open/v1/audit", nil, &out)
			if (err != nil) != tc.errorExpected {
				t.Fatalf("error = %v", err)
			}
			if tc.readErr && !errors.Is(err, sentinel) {
				t.Fatalf("read error lost: %v", err)
			}
			if !reader.closed || calls != 1 {
				t.Errorf("body closed=%v calls=%d", reader.closed, calls)
			}
			if !tc.errorExpected && out.ID != "9007199254740993" {
				t.Error("large string identifier lost")
			}
		})
	}
}

func TestAuditTransportFailureNoRetry(t *testing.T) {
	sentinel := errors.New("audit transport failure")
	calls := 0
	c, _ := New(Config{BaseURL: DefaultBaseURL, Token: "audit", HTTPClient: &http.Client{Transport: roundTripFunc(func(r *http.Request) (*http.Response, error) { calls++; return nil, sentinel })}})
	if err := c.Do(context.Background(), "/open/v1/audit", nil, nil); !errors.Is(err, sentinel) {
		t.Fatal(err)
	}
	if calls != 1 {
		t.Fatalf("unexpected retries: %d", calls)
	}
}

func TestAuditFreshSignedRequests(t *testing.T) {
	traces := map[string]bool{}
	c, _ := New(Config{BaseURL: DefaultBaseURL, AppID: "audit-app", APIKey: "audit-key", MinInterval: time.Nanosecond, HTTPClient: &http.Client{Transport: roundTripFunc(func(r *http.Request) (*http.Response, error) {
		trace, nonce, ts := r.Header.Get("traceId"), r.Header.Get("nonce"), r.Header.Get("ts")
		sum := sha256.Sum256([]byte("audit-app" + trace + ts + nonce + "audit-key"))
		if len(trace) != 36 || nonce != trace[:6] || traces[trace] || r.Header.Get("sign") != strings.ToUpper(hex.EncodeToString(sum[:])) || r.Header.Get("Authorization") != "" {
			t.Error("invalid/reused signed headers")
		}
		traces[trace] = true
		return &http.Response{StatusCode: 200, Body: io.NopCloser(strings.NewReader(`{"code":0}`)), Header: make(http.Header)}, nil
	})}})
	for i := 0; i < 200; i++ {
		if err := c.Do(context.Background(), "/open/v1/audit", nil, nil); err != nil {
			t.Fatal(err)
		}
	}
}

func TestAuditAllRedirectStatuses(t *testing.T) {
	for _, status := range []int{301, 302, 303, 307, 308} {
		t.Run(fmt.Sprint(status), func(t *testing.T) {
			var targets atomic.Int32
			target := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) { targets.Add(1) }))
			defer target.Close()
			redirect := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) { http.Redirect(w, r, target.URL, status) }))
			defer redirect.Close()
			c, _ := New(Config{BaseURL: redirect.URL, Token: "audit-token"})
			if err := c.Do(context.Background(), "/open/v1/audit", nil, nil); err == nil {
				t.Error("API redirect accepted")
			}
			if err := c.Upload(context.Background(), "", redirect.URL+"?Signature=audit-secret", nil, strings.NewReader("x"), 1); err == nil || strings.Contains(err.Error(), "audit-secret") {
				t.Errorf("upload redirect error %v", err)
			}
			if targets.Load() != 0 {
				t.Error("cross-host redirect followed")
			}
		})
	}
}

func TestAuditUploadStreamEdges(t *testing.T) {
	sentinel := errors.New("audit stream failure")
	cases := []struct {
		name          string
		body          func() io.Reader
		size          int64
		want          string
		errorExpected bool
	}{
		{"empty", func() io.Reader { return nil }, 0, "", false},
		{"exact", func() io.Reader { return strings.NewReader("payload") }, 7, "payload", false},
		{"short", func() io.Reader { return strings.NewReader("payload") }, 10, "", true},
		{"long", func() io.Reader { return strings.NewReader("payload") }, 3, "", true},
		{"reader_error", func() io.Reader { return auditErrorReader{sentinel} }, 3, "", true},
		{"negative_size", func() io.Reader { return strings.NewReader("x") }, -1, "", true},
		{"missing_reader", func() io.Reader { return nil }, 1, "", true},
	}
	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			var received string
			srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
				if r.Method != "PUT" {
					t.Error("wrong storage method")
				}
				for _, h := range []string{"Authorization", "appId", "sign", "traceId", "ts", "nonce"} {
					if r.Header.Get(h) != "" {
						t.Errorf("API header %s sent to storage", h)
					}
				}
				body, err := io.ReadAll(r.Body)
				if err != nil {
					w.WriteHeader(400)
					return
				}
				received = string(body)
				w.WriteHeader(200)
			}))
			defer srv.Close()
			c, _ := New(Config{BaseURL: DefaultBaseURL, AppID: "audit-app", APIKey: "audit-key"})
			ctx, cancel := context.WithTimeout(context.Background(), time.Second)
			defer cancel()
			err := c.Upload(ctx, "", srv.URL+"?Signature=audit-secret", nil, tc.body(), tc.size)
			if (err != nil) != tc.errorExpected {
				t.Fatalf("error=%v", err)
			}
			if !tc.errorExpected && received != tc.want {
				t.Errorf("received %q", received)
			}
			if err != nil && strings.Contains(err.Error(), "audit-secret") {
				t.Errorf("signed URL leaked: %v", err)
			}
		})
	}
}

func TestAuditUploadResponses(t *testing.T) {
	for _, tc := range []struct {
		status  int
		body    string
		failure bool
	}{{200, "", false}, {204, "", false}, {200, `{"code":0}`, false}, {200, `{"code":40006,"data":{"id":"x"}}`, true}, {403, `<Error>Denied</Error>`, true}, {500, `oops`, true}} {
		t.Run(fmt.Sprintf("%d_%s", tc.status, tc.body), func(t *testing.T) {
			reader := &auditReadClose{Reader: strings.NewReader(tc.body)}
			c, _ := New(Config{BaseURL: DefaultBaseURL, Token: "audit", HTTPClient: &http.Client{Transport: roundTripFunc(func(r *http.Request) (*http.Response, error) {
				return &http.Response{StatusCode: tc.status, Body: reader, Header: make(http.Header)}, nil
			})}})
			err := c.Upload(context.Background(), "PUT", "https://storage.example/audit", nil, nil, 0)
			if (err != nil) != tc.failure || !reader.closed {
				t.Fatalf("err=%v closed=%v", err, reader.closed)
			}
		})
	}
}

func TestAuditConcurrentCancellationIsolation(t *testing.T) {
	const interval = 2 * time.Millisecond
	var mu sync.Mutex
	var starts []time.Time
	c, _ := New(Config{BaseURL: DefaultBaseURL, Token: "audit", MinInterval: interval, HTTPClient: &http.Client{Transport: roundTripFunc(func(r *http.Request) (*http.Response, error) {
		mu.Lock()
		starts = append(starts, time.Now())
		mu.Unlock()
		return &http.Response{StatusCode: 200, Body: io.NopCloser(strings.NewReader(`{"code":0}`)), Header: make(http.Header)}, nil
	})}})
	var wg sync.WaitGroup
	errs := make(chan error, 80)
	for i := 0; i < 80; i++ {
		wg.Add(1)
		go func(i int) {
			defer wg.Done()
			ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
			defer cancel()
			if i%2 == 0 {
				cancel()
			}
			err := c.Do(ctx, "/open/v1/audit", nil, nil)
			if i%2 == 0 {
				if !errors.Is(err, context.Canceled) {
					errs <- fmt.Errorf("cancelled: %v", err)
				}
			} else if err != nil {
				errs <- err
			}
		}(i)
	}
	wg.Wait()
	close(errs)
	for err := range errs {
		t.Error(err)
	}
	if len(starts) != 40 {
		t.Fatalf("sent %d calls", len(starts))
	}
	sort.Slice(starts, func(i, j int) bool { return starts[i].Before(starts[j]) })
	if elapsed := starts[len(starts)-1].Sub(starts[0]); elapsed < interval*38 {
		t.Errorf("limiter lost spacing: %v", elapsed)
	}
	c.mu.Lock()
	c.nextRequest["/open/v1/audit"] = time.Now().Add(time.Hour)
	c.mu.Unlock()
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	if err := c.Do(ctx, "/open/v1/independent", nil, nil); err != nil {
		t.Fatal(err)
	}
}

func FuzzAuditEnvelope(f *testing.F) {
	for _, body := range []string{`{"code":0,"data":{}}`, `{"code":40006,"data":{"x":1}}`, `{}`, `null`, `{"code":null}`, `{"code":0,"code":1}`, `[]`, `<html>`, `{"code":0,"data":"x"}`} {
		f.Add(body)
	}
	f.Fuzz(func(t *testing.T, body string) {
		if len(body) > 65536 {
			t.Skip()
		}
		reader := &auditReadClose{Reader: bytes.NewBufferString(body)}
		c, _ := New(Config{BaseURL: DefaultBaseURL, Token: "audit", HTTPClient: &http.Client{Transport: roundTripFunc(func(r *http.Request) (*http.Response, error) {
			return &http.Response{StatusCode: 200, Body: reader, Header: make(http.Header)}, nil
		})}})
		var out json.RawMessage
		err := c.Do(context.Background(), "/open/v1/audit", nil, &out)
		if !reader.closed {
			t.Fatal("response body leaked")
		}
		if err == nil {
			var expected envelope
			if json.Unmarshal([]byte(body), &expected) != nil || expected.Code == nil || *expected.Code != 0 ||
				len(expected.Data) == 0 || bytes.Equal(bytes.TrimSpace(expected.Data), []byte("null")) || !json.Valid(out) {
				t.Fatal("accepted an incomplete success envelope")
			}
		}
	})
}
