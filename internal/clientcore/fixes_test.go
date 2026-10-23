package clientcore

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"strings"
	"sync/atomic"
	"testing"
)

func TestOptionalDataStillDecodesFailures(t *testing.T) {
	for _, body := range []string{`{"code":0}`, `{"code":0,"data":null}`, `{"code":0,"data":{"failCount":2}}`} {
		c, _ := New(Config{BaseURL: DefaultBaseURL, Token: "test", HTTPClient: &http.Client{Transport: roundTripFunc(func(*http.Request) (*http.Response, error) {
			return &http.Response{StatusCode: 200, Body: io.NopCloser(strings.NewReader(body))}, nil
		})}})
		var out struct{ FailCount int }
		if err := c.DoOptionalData(context.Background(), "/open/v1/phone/detail/update", nil, &out); err != nil {
			t.Fatal(err)
		}
		if strings.Contains(body, "failCount") && out.FailCount != 2 {
			t.Fatal("optional response data was discarded")
		}
	}
}

func TestRedirectErrorsKeepHTTPResponse(t *testing.T) {
	for _, status := range []int{301, 302, 303, 307, 308} {
		t.Run(fmt.Sprint(status), func(t *testing.T) {
			var followed atomic.Int32
			target := httptest.NewServer(http.HandlerFunc(func(http.ResponseWriter, *http.Request) { followed.Add(1) }))
			defer target.Close()
			body := `{"code":40006,"msg":"moved","traceId":"redirect-trace","data":{"successAmount":1}}`
			origin := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
				w.Header().Set("Location", target.URL)
				w.WriteHeader(status)
				_, _ = io.WriteString(w, body)
			}))
			defer origin.Close()
			c, _ := New(Config{BaseURL: origin.URL, Token: "test"})
			err := c.Do(context.Background(), "/open/v1/phone/list", nil, &struct{}{})
			var apiErr *APIError
			if !errors.As(err, &apiErr) || apiErr.HTTPStatus != status || apiErr.Code != 40006 || apiErr.TraceID != "redirect-trace" || apiErr.Message != "moved" || apiErr.Body != body {
				t.Fatalf("redirect response lost: %v", err)
			}
			var partial struct{ SuccessAmount int }
			if err := apiErr.DecodeData(&partial); err != nil || partial.SuccessAmount != 1 {
				t.Fatalf("redirect partial data lost: %v", err)
			}
			err = c.Upload(context.Background(), "PUT", origin.URL+"?Signature=secret", nil, nil, 0)
			if !errors.As(err, &apiErr) || apiErr.HTTPStatus != status || strings.Contains(err.Error(), "secret") {
				t.Fatalf("upload redirect status lost or signed URL leaked: %v", err)
			}
			if followed.Load() != 0 {
				t.Fatal("redirect was followed")
			}
		})
	}
}

func TestZeroSizeUploads(t *testing.T) {
	sentinel := errors.New("reader failed")
	for _, tc := range []struct {
		name string
		body io.Reader
		fail bool
	}{
		{"nil", nil, false},
		{"empty", strings.NewReader(""), false},
		{"empty_stream", io.LimitReader(strings.NewReader(""), 1), false},
		{"nonempty", strings.NewReader("payload"), true},
		{"nonempty_stream", io.LimitReader(strings.NewReader("payload"), 1), true},
		{"reader_error", auditErrorReader{sentinel}, true},
	} {
		t.Run(tc.name, func(t *testing.T) {
			calls := 0
			c, _ := New(Config{BaseURL: DefaultBaseURL, Token: "test", HTTPClient: &http.Client{Transport: roundTripFunc(func(r *http.Request) (*http.Response, error) {
				calls++
				raw, err := io.ReadAll(r.Body)
				if err != nil || len(raw) != 0 || r.ContentLength != 0 {
					t.Fatalf("unexpected empty upload: %q %v", raw, err)
				}
				return &http.Response{StatusCode: 204, Body: http.NoBody}, nil
			})}})
			err := c.Upload(context.Background(), "PUT", "https://storage.example/file", nil, tc.body, 0)
			if (err != nil) != tc.fail || tc.fail && calls != 0 || !tc.fail && calls != 1 {
				t.Fatalf("err=%v calls=%d", err, calls)
			}
			if tc.name == "reader_error" && !errors.Is(err, sentinel) {
				t.Fatalf("reader error lost: %v", err)
			}
		})
	}
	ctx, cancel := context.WithCancel(context.Background())
	cancel()
	c, _ := New(Config{BaseURL: DefaultBaseURL, Token: "test"})
	if err := c.Upload(ctx, "PUT", "https://storage.example/file", nil, auditErrorReader{sentinel}, 0); !errors.Is(err, context.Canceled) {
		t.Fatalf("cancelled upload read its reader: %v", err)
	}
}

func TestAPIErrorDetails(t *testing.T) {
	err := &APIError{HTTPStatus: 503, Code: 40006, Message: "partial success", TraceID: "trace", Data: json.RawMessage(`{"count":1}`)}
	if got := err.Error(); got != "geelark api error: http_status=503, code=40006, trace_id=trace, partial success" {
		t.Fatal(got)
	}
	for _, tc := range []struct {
		err  APIError
		want string
	}{
		{APIError{}, "geelark api error"},
		{APIError{Body: "unavailable"}, "geelark api error: unavailable"},
		{APIError{HTTPStatus: 403, Body: "private storage details"}, "geelark api error: http_status=403"},
	} {
		if got := tc.err.Error(); got != tc.want {
			t.Fatalf("got %q want %q", got, tc.want)
		}
	}
	for _, raw := range []json.RawMessage{nil, json.RawMessage(`null`), json.RawMessage(`{"x":`)} {
		if err := (&APIError{Data: raw}).DecodeData(&struct{}{}); err == nil {
			t.Fatalf("accepted absent or malformed error data: %s", raw)
		}
	}
}
