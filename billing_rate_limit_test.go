package geelark_test

import (
	"context"
	"errors"
	"net/http"
	"net/http/httptest"
	"sync/atomic"
	"testing"
	"time"

	geelark "github.com/wyosa/geelark-go-sdk"
	"github.com/wyosa/geelark-go-sdk/billing"
	"github.com/wyosa/geelark-go-sdk/cloudphone"
)

func TestBillingRateLimits(t *testing.T) {
	for _, auth := range []string{"token", "key"} {
		for _, method := range []string{"PlanInfo", "Balance"} {
			t.Run(auth+"/"+method, func(t *testing.T) {
				var calls atomic.Int32
				srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
					calls.Add(1)
					_, _ = w.Write([]byte(`{"code":0,"data":{}}`))
				}))
				defer srv.Close()

				var c *geelark.Client
				var err error
				if auth == "key" {
					c, err = geelark.NewKeyClient("app", "key", geelark.WithBaseURL(srv.URL))
				} else {
					c, err = geelark.NewClient("token", geelark.WithBaseURL(srv.URL))
				}
				if err != nil {
					t.Fatal(err)
				}
				call := func(ctx context.Context) error {
					if method == "PlanInfo" {
						_, err := c.Billing.PlanInfo(ctx, billing.PlanInfoRequest{})
						return err
					}
					_, err := c.Billing.Balance(ctx, billing.BalanceRequest{})
					return err
				}
				if err := call(context.Background()); err != nil {
					t.Fatal(err)
				}

				// The former 310 ms interval allowed a second HTTP request here.
				ctx, cancel := context.WithTimeout(context.Background(), 500*time.Millisecond)
				defer cancel()
				errs := make(chan error, 3)
				for i := 0; i < cap(errs); i++ {
					go func() { errs <- call(ctx) }()
				}
				for i := 0; i < cap(errs); i++ {
					if err := <-errs; !errors.Is(err, context.DeadlineExceeded) {
						t.Errorf("request bypassed the billing limit: %v", err)
					}
				}
				if got := calls.Load(); got != 1 {
					t.Errorf("sent %d billing requests before the next slot, want 1", got)
				}

				// Waiting on Billing must not consume another endpoint's budget.
				otherCtx, otherCancel := context.WithTimeout(context.Background(), time.Second)
				defer otherCancel()
				if _, err := c.CloudPhones.List(otherCtx, cloudphone.ListRequest{}); err != nil {
					t.Fatalf("independent endpoint blocked: %v", err)
				}
			})
		}
	}
}
