package clientcore

import (
	"context"
	"io"
	"net/http"
	"strings"
	"testing"
	"time"
)

func TestEndpointRateLimitIntervals(t *testing.T) {
	cases := []struct {
		name, path string
		configured time.Duration
		want       time.Duration
	}{
		{"default plan info", "/open/v1/pay/plan/info", 0, time.Minute},
		{"default balance", "/open/v1/pay/wallet", 0, 6 * time.Second},
		{"default other endpoint", "/open/v1/phone/list", 0, DefaultMinInterval},
		{"shorter plan info", "/open/v1/pay/plan/info", time.Nanosecond, time.Minute},
		{"shorter balance", "/open/v1/pay/wallet", time.Nanosecond, 6 * time.Second},
		{"shorter other endpoint", "/open/v1/phone/list", time.Nanosecond, time.Nanosecond},
		{"longer plan info", "/open/v1/pay/plan/info", 2 * time.Minute, 2 * time.Minute},
		{"longer balance", "/open/v1/pay/wallet", 10 * time.Second, 10 * time.Second},
	}
	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			c, err := New(Config{
				BaseURL: DefaultBaseURL, Token: "token", MinInterval: tc.configured,
				HTTPClient: &http.Client{Transport: roundTripFunc(func(r *http.Request) (*http.Response, error) {
					return &http.Response{
						StatusCode: http.StatusOK,
						Body:       io.NopCloser(strings.NewReader(`{"code":0}`)),
						Header:     make(http.Header),
					}, nil
				})},
			})
			if err != nil {
				t.Fatal(err)
			}
			before := time.Now()
			if err := c.Do(context.Background(), tc.path, nil, nil); err != nil {
				t.Fatal(err)
			}
			after := time.Now()
			// Inspect the booked slot to verify minute-long limits without sleeping.
			next := c.nextRequest[tc.path]
			if next.Before(before.Add(tc.want)) || next.After(after.Add(tc.want)) {
				t.Fatalf("next slot is in %s, want an interval of %s", next.Sub(before), tc.want)
			}
		})
	}
}
