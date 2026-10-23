package clientcore

import (
	"context"
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"errors"
	"io"
	"net/http"
	"net/http/httptest"
	"regexp"
	"strings"
	"sync/atomic"
	"testing"
	"time"
)

func TestTokenRequest(t *testing.T) {
	var got struct {
		Path        string
		Method      string
		ContentType string
		TraceID     string
		Auth        string
		Body        string
	}
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		got.Path, got.Method = r.URL.Path, r.Method
		got.ContentType = r.Header.Get("Content-Type")
		got.TraceID = r.Header.Get("traceId")
		got.Auth = r.Header.Get("Authorization")
		body, _ := io.ReadAll(r.Body)
		got.Body = string(body)
		_, _ = w.Write([]byte(`{"traceId":"server-trace","code":0,"msg":"success","data":{"id":"phone-1"}}`))
	}))
	defer srv.Close()
	c, err := New(Config{BaseURL: srv.URL, Token: "secret", MinInterval: time.Nanosecond})
	if err != nil {
		t.Fatal(err)
	}
	var out struct {
		ID string `json:"id"`
	}
	if err := c.Do(context.Background(), "/open/v1/phone/list", map[string]int{"page": 1}, &out); err != nil {
		t.Fatal(err)
	}
	if got.Path != "/open/v1/phone/list" || got.Method != http.MethodPost || got.ContentType != "application/json" || got.Auth != "Bearer secret" || got.Body != `{"page":1}` || out.ID != "phone-1" {
		t.Fatalf("unexpected request or response: %+v, out=%+v", got, out)
	}
	uuidV4 := regexp.MustCompile(`^[0-9a-f]{8}-[0-9a-f]{4}-4[0-9a-f]{3}-[89ab][0-9a-f]{3}-[0-9a-f]{12}$`)
	if !uuidV4.MatchString(got.TraceID) {
		t.Fatalf("traceId = %q, want UUIDv4", got.TraceID)
	}
}

func TestKeySignature(t *testing.T) {
	if got := signature("eH6gQR4oHr3FsZpI36La01IW", "db6094ab-3797-4186-84d5-b0b58eebad56", "1716972892166", "db6094", "YjmFIUuq0oJgSDJ42fxLEb6R1qjjqf"); got != "6280C080AF7C3CCE168F15C913E3444A00A618CB0E16038EED9811D6E3366BDD" {
		t.Fatalf("documented signature = %q", got)
	}
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Header.Get("Authorization") != "" || r.Header.Get("appId") != "app" {
			t.Errorf("wrong auth headers: %v", r.Header)
		}
		trace, ts, nonce := r.Header.Get("traceId"), r.Header.Get("ts"), r.Header.Get("nonce")
		if len(trace) < 6 || nonce != trace[:6] {
			t.Errorf("bad traceId/nonce: %q, %q", trace, nonce)
		}
		if _, err := time.ParseDuration(ts + "ms"); err != nil {
			t.Errorf("bad millisecond timestamp %q: %v", ts, err)
		}
		sum := sha256.Sum256([]byte("app" + trace + ts + nonce + "key"))
		want := strings.ToUpper(hex.EncodeToString(sum[:]))
		if r.Header.Get("sign") != want {
			t.Errorf("sign = %q, want %q", r.Header.Get("sign"), want)
		}
		_, _ = w.Write([]byte(`{"code":0,"data":null}`))
	}))
	defer srv.Close()
	c, err := New(Config{BaseURL: srv.URL, AppID: "app", APIKey: "key"})
	if err != nil {
		t.Fatal(err)
	}
	if err := c.Do(context.Background(), "/open/v1/pay/wallet", nil, nil); err != nil {
		t.Fatal(err)
	}
}

func TestBusinessErrorPreservesPartialData(t *testing.T) {
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		_, _ = w.Write([]byte(`{"traceId":"trace-1","code":40006,"msg":"partial success","data":{"successAmount":1,"failAmount":1}}`))
	}))
	defer srv.Close()
	c, _ := New(Config{BaseURL: srv.URL, Token: "token"})
	err := c.Do(context.Background(), "/open/v1/phone/start", nil, nil)
	var apiErr *APIError
	if !errors.As(err, &apiErr) || apiErr.Code != 40006 || apiErr.TraceID != "trace-1" || !json.Valid(apiErr.Data) {
		t.Fatalf("got %v, want partial success APIError", err)
	}
}

func TestBaseURLAndRedirectSafety(t *testing.T) {
	for _, base := range []string{"http://example.com", "https://user:pass@example.com", "https://example.com?token=x", "https://example.com/#x"} {
		if _, err := New(Config{BaseURL: base, Token: "token"}); err == nil {
			t.Errorf("accepted insecure base URL %q", base)
		}
	}
	var targetCalls atomic.Int32
	target := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		targetCalls.Add(1)
	}))
	defer target.Close()
	redirect := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		http.Redirect(w, r, target.URL, http.StatusTemporaryRedirect)
	}))
	defer redirect.Close()
	c, err := New(Config{BaseURL: redirect.URL, Token: "token"})
	if err != nil {
		t.Fatal(err)
	}
	if err := c.Do(context.Background(), "/open/v1/phone/list", nil, nil); err == nil {
		t.Fatal("expected redirect refusal")
	}
	if targetCalls.Load() != 0 {
		t.Fatal("followed redirect with credentials")
	}
}

func TestUploadHasNoAPIAuth(t *testing.T) {
	var gotBody string
	storage := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPut || r.Header.Get("Authorization") != "" || r.Header.Get("appId") != "" || r.Header.Get("traceId") != "" {
			t.Errorf("unexpected upload method or auth headers: %s %v", r.Method, r.Header)
		}
		body, _ := io.ReadAll(r.Body)
		gotBody = string(body)
		w.WriteHeader(http.StatusOK)
	}))
	defer storage.Close()
	c, _ := New(Config{BaseURL: "https://openapi.geelark.com", Token: "secret"})
	if err := c.Upload(context.Background(), http.MethodPut, storage.URL+"/file?Signature=secret", nil, strings.NewReader("file"), 4); err != nil {
		t.Fatal(err)
	}
	if gotBody != "file" {
		t.Fatalf("uploaded %q", gotBody)
	}
}

func TestRateWaitRespectsContext(t *testing.T) {
	var calls atomic.Int32
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		calls.Add(1)
		_, _ = w.Write([]byte(`{"code":0,"data":null}`))
	}))
	defer srv.Close()
	c, _ := New(Config{BaseURL: srv.URL, Token: "token", MinInterval: 200 * time.Millisecond})
	if err := c.Do(context.Background(), "/open/v1/phone/list", nil, nil); err != nil {
		t.Fatal(err)
	}
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Millisecond)
	defer cancel()
	if err := c.Do(ctx, "/open/v1/phone/list", nil, nil); !errors.Is(err, context.DeadlineExceeded) {
		t.Fatalf("got %v, want deadline exceeded", err)
	}
	if calls.Load() != 1 {
		t.Fatalf("sent %d requests, want one", calls.Load())
	}
}
