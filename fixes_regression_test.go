package geelark_test

import (
	"context"
	"encoding/json"
	"io"
	"net/http"
	"reflect"
	"strings"
	"testing"

	"github.com/wyosa/geelark-go-sdk/analytics"
	"github.com/wyosa/geelark-go-sdk/automation"
	"github.com/wyosa/geelark-go-sdk/billing"
	"github.com/wyosa/geelark-go-sdk/browserautomation"
	"github.com/wyosa/geelark-go-sdk/common"
	"github.com/wyosa/geelark-go-sdk/library"
)

func TestMissingResponseData(t *testing.T) {
	for _, fixture := range auditFixtures(t) {
		for _, body := range []string{`{"code":0}`, `{"code":0,"data":null}`} {
			t.Run(fixture.Package+"."+fixture.Method+"/"+body, func(t *testing.T) {
				hc := &http.Client{Transport: auditTransport(func(*http.Request) (*http.Response, error) {
					return &http.Response{StatusCode: 200, Header: make(http.Header), Body: io.NopCloser(strings.NewReader(body))}, nil
				})}
				method, err := auditMethod(fixture, hc)
				if err != nil {
					t.Fatal(err)
				}
				responseType := method.Type().Out(0).Elem()
				// Update documents a data-free full success; HideAccessibility returns the same
				// failDetails-only shape and is treated the same way.
				optional := responseType.Kind() == reflect.Struct && responseType.NumField() == 0 ||
					fixture.Package == "cloudphone" && (fixture.Method == "Update" || fixture.Method == "HideAccessibility")
				result := method.Call([]reflect.Value{reflect.ValueOf(context.Background()), reflect.Zero(method.Type().In(1))})
				if optional {
					if !result[1].IsNil() || result[0].IsNil() {
						t.Fatalf("documented no-data success rejected: %v", result[1].Interface())
					}
				} else if result[1].IsNil() || !strings.Contains(result[1].Interface().(error).Error(), "missing or null data") || !result[0].IsNil() {
					t.Fatalf("missing data must fail without a zero-value result: %v", result)
				}
			})
		}
	}
}

func TestLargeCountersAndTimestamps(t *testing.T) {
	const large int64 = 3_000_000_000
	var metrics analytics.DataResponseItemsItem
	if err := json.Unmarshal([]byte(`{"playCount":3000000000,"followerCount":3000000000,"diggCount":3000000000,"commentCount":3000000000,"collectCount":3000000000,"shareCount":3000000000,"dataDate":3000000000,"addAccDate":3000000000}`), &metrics); err != nil {
		t.Fatal(err)
	}
	for _, got := range []int64{metrics.PlayCount, metrics.FollowerCount, metrics.DiggCount, metrics.CommentCount, metrics.CollectCount, metrics.ShareCount, metrics.DataDate, metrics.AddAccDate} {
		if got != large {
			t.Fatalf("large metric or date lost: %d", got)
		}
	}
	var material library.SearchResponseListItem
	if err := json.Unmarshal([]byte(`{"id":9007199254740993,"fileSize":3000000000,"createdTime":3000000000}`), &material); err != nil || material.FileSize != large || material.CreatedTime != large || material.ID != "9007199254740993" {
		t.Fatalf("large material fields lost: %+v, %v", material, err)
	}
	var task browserautomation.DetailResponse
	if err := json.Unmarshal([]byte(`{"startAt":3000000000,"finishAt":3000000000,"scheduleAt":3000000000}`), &task); err != nil || task.StartAt != large || task.FinishAt != large || task.ScheduleAt != large {
		t.Fatalf("large task dates lost: %+v, %v", task, err)
	}
	var plan billing.PlanInfoResponse
	if err := json.Unmarshal([]byte(`{"expirationTime":3000000000}`), &plan); err != nil || plan.ExpirationTime != large {
		t.Fatalf("large plan expiration lost: %+v, %v", plan, err)
	}
	requests := []struct {
		request any
		field   string
	}{
		{automation.CreateCustomTaskRequest{ScheduleAt: large}, "scheduleAt"},
		{automation.AddTaskItem{ScheduleAt: large}, "scheduleAt"},
		{browserautomation.AddRequest{ScheduleAt: large}, "scheduleAt"},
		{analytics.DataRequest{DataDate: common.Ptr(large)}, "dataDate"},
		{billing.TransactionsRequest{StartAt: common.Ptr(large)}, "startAt"},
	}
	for _, tc := range requests {
		raw, err := json.Marshal(tc.request)
		if err != nil {
			t.Fatal(err)
		}
		if auditJSON(raw).(map[string]any)[tc.field] != json.Number("3000000000") {
			t.Fatalf("timestamp changed in %T: %s", tc.request, raw)
		}
	}
}
