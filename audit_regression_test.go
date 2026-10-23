package geelark

import (
	"context"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"reflect"
	"testing"

	"github.com/wyosa/geelark-go-sdk/analytics"
	"github.com/wyosa/geelark-go-sdk/app"
	"github.com/wyosa/geelark-go-sdk/automation"
	"github.com/wyosa/geelark-go-sdk/billing"
	"github.com/wyosa/geelark-go-sdk/browser"
	"github.com/wyosa/geelark-go-sdk/browserautomation"
	"github.com/wyosa/geelark-go-sdk/cloudphone"
	"github.com/wyosa/geelark-go-sdk/group"
	"github.com/wyosa/geelark-go-sdk/tag"
)

func TestBillingFractionalAmounts(t *testing.T) {
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		switch r.URL.Path {
		case "/open/v1/pay/wallet":
			_, _ = w.Write([]byte(`{"code":0,"data":{"balance":12.5,"giftMoney":0.75,"availableTimeAddOn":10}}`))
		case "/open/v1/pay/profiles/list":
			_, _ = w.Write([]byte(`{"code":0,"data":[{"id":"plan","price":19.99}]}`))
		case "/open/v1/billing/transaction/detail":
			_, _ = w.Write([]byte(`{"code":0,"data":{"list":[{"id":"transaction","amount":0.05}]}}`))
		default:
			t.Errorf("unexpected path: %s", r.URL.Path)
		}
	}))
	defer srv.Close()
	c, err := NewClient("token", WithBaseURL(srv.URL))
	if err != nil {
		t.Fatal(err)
	}
	t.Run("balance", func(t *testing.T) {
		got, err := c.Billing.Balance(context.Background(), billing.BalanceRequest{})
		if err != nil {
			t.Fatal(err)
		}
		if float64(got.GiftMoney) != 0.75 {
			t.Fatalf("giftMoney = %v", got.GiftMoney)
		}
	})
	t.Run("plans", func(t *testing.T) {
		got, err := c.Billing.Plans(context.Background(), billing.PlansRequest{})
		if err != nil {
			t.Fatal(err)
		}
		if len(*got) != 1 || float64((*got)[0].Price) != 19.99 {
			t.Fatalf("plans = %+v", got)
		}
	})
	t.Run("transactions", func(t *testing.T) {
		got, err := c.Billing.Transactions(context.Background(), billing.TransactionsRequest{})
		if err != nil {
			t.Fatal(err)
		}
		if len(got.List) != 1 || float64(got.List[0].Amount) != 0.05 {
			t.Fatalf("transactions = %+v", got)
		}
	})
}

func TestOptionalListsCanBeCleared(t *testing.T) {
	cases := []struct {
		name, key string
		request   func([]string) any
	}{
		{"cloud phone tags", "tagIDs", func(v []string) any { return cloudphone.UpdateRequest{ID: "phone", TagIDs: v} }},
		{"browser tags", "tagIds", func(v []string) any { return browser.UpdateRequest{ID: "browser", TagIDs: v} }},
		{"analytics tags", "tagIds", func(v []string) any { return analytics.UpdateAccountRequest{ID: "account", TagIDs: v} }},
		{"network blacklist", "blackList", func(v []string) any { return cloudphone.SetNetworkSettingsRequest{BlackList: v} }},
	}
	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			for _, value := range [][]string{nil, {}, {"one"}} {
				body, err := json.Marshal(tc.request(value))
				if err != nil {
					t.Fatal(err)
				}
				var fields map[string]json.RawMessage
				if err := json.Unmarshal(body, &fields); err != nil {
					t.Fatal(err)
				}
				got, present := fields[tc.key]
				if value == nil {
					if present {
						t.Errorf("nil must omit the optional field: %s", body)
					}
					continue
				}
				want, _ := json.Marshal(value)
				if !present || string(got) != string(want) {
					t.Errorf("explicit %s must be sent, got %s", want, body)
				}
			}
		})
	}
}

func TestDocumentedResponseFieldsArePreserved(t *testing.T) {
	cases := []struct {
		name, body string
		out        any
	}{
		{"query task result", `{"failCode":29999,"failDesc":"some reason","cost":60,"shareLink":"https://example.com/video"}`, new(automation.QueryTaskResponseItemsItem)},
		{"task history result", `{"failCode":29999,"failDesc":"some reason","cost":60}`, new(automation.ListTasksResponseItemsItem)},
		{"task details failure", `{"failCode":29999,"failDesc":"some reason"}`, new(automation.TaskDetailResponse)},
		{"tag lookup failures", `{"failDetails":[{"code":43022,"type":1,"param":"tag","msg":"missing"}]}`, new(tag.ListResponse)},
		{"group lookup failures", `{"failDetails":[{"code":43032,"type":1,"param":"group","msg":"missing"}]}`, new(group.ListResponse)},
		{"browser memory spelling", `{"memoryDevice":8}`, new(browser.ListResponseItemsItemSimulateInfo)},
		{"browser legacy memory spelling", `{"memeryDevice":8}`, new(browser.ListResponseItemsItemSimulateInfo)},
	}
	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			if err := json.Unmarshal([]byte(tc.body), tc.out); err != nil {
				t.Fatal(err)
			}
			body, err := json.Marshal(tc.out)
			if err != nil {
				t.Fatal(err)
			}
			var want, got map[string]any
			_ = json.Unmarshal([]byte(tc.body), &want)
			_ = json.Unmarshal(body, &got)
			for key, value := range want {
				if !reflect.DeepEqual(got[key], value) {
					t.Errorf("lost %s: want %v, got %v", key, value, got[key])
				}
			}
		})
	}
}

func TestAppVersionCodeAcceptsDocumentedForms(t *testing.T) {
	for _, raw := range []string{`21120300`, `"21120300"`} {
		var value app.ListAvailableResponseItemsItemAppVersionInfoListItem
		if err := json.Unmarshal([]byte(`{"id":"version","versionCode":`+raw+`}`), &value); err != nil {
			t.Errorf("versionCode %s: %v", raw, err)
			continue
		}
		if value.VersionCode != 21120300 || value.ID != "version" {
			t.Errorf("versionCode %s: %+v", raw, value)
		}
	}
}

func TestCustomTasksPreserveArbitraryParameters(t *testing.T) {
	const params = `{"Title":"video","Caption":"custom caption","Count":0,"Ratio":0.75,"Enabled":false,"Lines":["first","second"],"Attachments":["https://example.com/file"]}`
	cases := []struct {
		name, body string
		out        any
	}{
		{"cloud", `{"scheduleAt":1790000000,"id":"phone","flowId":"flow","paramMap":` + params + `}`, new(automation.CreateCustomTaskRequest)},
		{"browser", `{"scheduleAt":1790000000,"eid":"browser","flowId":"flow","paramMap":` + params + `}`, new(browserautomation.AddRequest)},
	}
	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			if err := json.Unmarshal([]byte(tc.body), tc.out); err != nil {
				t.Fatal(err)
			}
			got, err := json.Marshal(tc.out)
			if err != nil {
				t.Fatal(err)
			}
			assertJSON(t, got, tc.body)
		})
	}
}

func TestTaskLogCursorPreservesMillisecondTimestamps(t *testing.T) {
	var response automation.TaskDetailResponse
	if err := json.Unmarshal([]byte(`{"id":"task","searchAfter":[1749004889853]}`), &response); err != nil {
		t.Fatal(err)
	}
	if len(response.SearchAfter) != 1 || response.SearchAfter[0] != 1749004889853 {
		t.Fatalf("cursor = %v", response.SearchAfter)
	}
	body, err := json.Marshal(automation.TaskDetailRequest{ID: response.ID, SearchAfter: response.SearchAfter})
	if err != nil {
		t.Fatal(err)
	}
	assertJSON(t, body, `{"id":"task","searchAfter":[1749004889853]}`)
}
