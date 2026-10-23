package geelark_test

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	geelark "github.com/wyosa/geelark-go-sdk"
	"github.com/wyosa/geelark-go-sdk/analytics"
	"github.com/wyosa/geelark-go-sdk/automation"
	"github.com/wyosa/geelark-go-sdk/browser"
	"github.com/wyosa/geelark-go-sdk/browserautomation"
	"github.com/wyosa/geelark-go-sdk/cloudphone"
	"github.com/wyosa/geelark-go-sdk/common"
	"github.com/wyosa/geelark-go-sdk/files"
	"github.com/wyosa/geelark-go-sdk/group"
	"github.com/wyosa/geelark-go-sdk/library"
	"github.com/wyosa/geelark-go-sdk/proxy"
	"io"
	"net/http"
	"os"
	"path/filepath"
	"reflect"
	"strings"
	"testing"
	"time"
)

func TestAuditExplicitZeroValues(t *testing.T) {
	cases := []struct {
		name     string
		request  any
		expected string
	}{
		{"startup", cloudphone.StartRequest{IDs: []string{"p"}, Center: common.Ptr(0), EnableEnergySavingModeAfterMinutes: common.Ptr(0), CleanupSDcardAfterDays: common.Ptr(0), DarkMode: common.Ptr(false)}, `{"ids":["p"],"center":0,"enableEnergySavingModeAfterMinutes":0,"cleanupSDcardAfterDays":0,"darkMode":false}`},
		{"gps", cloudphone.SetGPSRequest{List: []cloudphone.SetGPSRequestListItem{{ID: "p", Longitude: common.Ptr(0.0), Latitude: common.Ptr(-0.125)}}}, `{"list":[{"id":"p","longitude":0,"latitude":-0.125}]}`},
		{"phone_clear", cloudphone.UpdateRequest{ID: "p", Name: common.Ptr(""), Remark: common.Ptr(""), GroupID: common.Ptr(""), TagIDs: []string{}}, `{"id":"p","name":"","remark":"","groupID":"","tagIDs":[]}`},
		{"browser_clear", browser.UpdateRequest{ID: "b", GroupID: common.Ptr(""), AccountTOTPSecret: common.Ptr(""), TagIDs: []string{}}, `{"id":"b","groupId":"","accountTOTPSecret":"","tagIds":[]}`},
		{"analytics_zero", analytics.UpdateAccountRequest{ID: "a", Channel: common.Ptr(0), TagIDs: []string{}}, `{"id":"a","channel":0,"tagIds":[]}`},
		{"group_remark", group.UpdateRequest{List: []group.UpdateRequestListItem{{ID: common.Ptr("g"), Remark: common.Ptr("")}}}, `{"list":[{"id":"g","remark":""}]}`},
		{"proxy_credentials", proxy.UpdateRequest{List: []proxy.UpdateItem{{ID: "x", Scheme: "http", Server: "proxy.invalid", Port: 80, Username: common.Ptr(""), Password: common.Ptr(""), ProxyGroupID: common.Ptr("")}}}, `{"list":[{"id":"x","scheme":"http","server":"proxy.invalid","port":80,"username":"","password":"","proxyGroupId":""}]}`},
		{"blacklist", cloudphone.SetNetworkSettingsRequest{BlackList: []string{}}, `{"blackList":[]}`},
		{"flow_int64", automation.CreateCustomTaskRequest{ID: "p", FlowID: "f", ScheduleAt: 1790000000, ParamMap: automation.CreateCustomTaskRequestParamMap{"large": int64(9007199254740993), "zero": 0, "bool": false, "nested": map[string]any{"x": []string{}}}}, `{"id":"p","flowId":"f","scheduleAt":1790000000,"paramMap":{"large":9007199254740993,"zero":0,"bool":false,"nested":{"x":[]}}}`},
		{"browser_flow", browserautomation.AddRequest{Eid: "b", FlowID: "f", ScheduleAt: 1790000000, ParamMap: browserautomation.AddRequestParamMap{"fraction": 0.125, "false": false}}, `{"eid":"b","flowId":"f","scheduleAt":1790000000,"paramMap":{"fraction":0.125,"false":false}}`},
	}
	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			raw, err := json.Marshal(tc.request)
			if err != nil {
				t.Fatal(err)
			}
			if !reflect.DeepEqual(auditJSON(raw), auditJSON([]byte(tc.expected))) {
				t.Fatalf("want %s, got %s", tc.expected, raw)
			}
		})
	}
}

func TestAuditUploadFileOwnership(t *testing.T) {
	for _, service := range []string{"files", "library"} {
		for _, fail := range []bool{false, true} {
			t.Run(fmt.Sprintf("%s_fail=%v", service, fail), func(t *testing.T) {
				path := filepath.Join(t.TempDir(), "file.txt")
				if err := os.WriteFile(path, []byte("exact content"), 0600); err != nil {
					t.Fatal(err)
				}
				var fileBody io.ReadCloser
				var bytesSent []byte
				puts := 0
				sentinel := errors.New("audit storage rejection")
				hc := &http.Client{Transport: auditTransport(func(r *http.Request) (*http.Response, error) {
					if r.Method == "POST" {
						return &http.Response{StatusCode: 200, Body: io.NopCloser(strings.NewReader(`{"code":0,"data":{"uploadUrl":"https://storage.example/file?Signature=audit","resourceUrl":"https://storage.example/public"}}`)), Header: make(http.Header)}, nil
					}
					puts++
					fileBody = r.Body
					defer fileBody.Close()
					bytesSent, _ = io.ReadAll(r.Body)
					for _, key := range []string{"Authorization", "appId", "traceId", "sign"} {
						if r.Header.Get(key) != "" {
							t.Errorf("storage credential %s", key)
						}
					}
					if fail {
						return nil, sentinel
					}
					return &http.Response{StatusCode: 200, Body: http.NoBody, Header: make(http.Header)}, nil
				})}
				c, _ := geelark.NewClient("audit", geelark.WithHTTPClient(hc))
				var err error
				if service == "files" {
					_, err = c.Files.UploadFile(context.Background(), path, files.GetUploadURLRequest{FileType: "txt"})
				} else {
					_, err = c.Library.UploadFile(context.Background(), path, library.GetUploadURLRequest{FileType: "txt"})
				}
				if fail && !errors.Is(err, sentinel) || !fail && err != nil {
					t.Fatalf("error=%v", err)
				}
				if puts != 1 || !bytes.Equal(bytesSent, []byte("exact content")) {
					t.Fatalf("puts=%d content=%q", puts, bytesSent)
				}
				_, readErr := fileBody.Read(make([]byte, 1))
				if !errors.Is(readErr, os.ErrClosed) {
					t.Errorf("file not closed: %v", readErr)
				}
			})
		}
	}
}

func TestAuditMissingUploadFile(t *testing.T) {
	calls := 0
	c, _ := geelark.NewClient("audit", geelark.WithHTTPClient(&http.Client{Transport: auditTransport(func(r *http.Request) (*http.Response, error) { calls++; return nil, errors.New("unexpected network") })}))
	path := filepath.Join(t.TempDir(), "missing")
	if _, err := c.Files.UploadFile(context.Background(), path, files.GetUploadURLRequest{}); !errors.Is(err, os.ErrNotExist) {
		t.Fatal(err)
	}
	if _, err := c.Library.UploadFile(context.Background(), path, library.GetUploadURLRequest{}); !errors.Is(err, os.ErrNotExist) {
		t.Fatal(err)
	}
	if calls != 0 {
		t.Errorf("missing file made %d API requests", calls)
	}
}

func TestAuditNilAndCancelledContext(t *testing.T) {
	var calls int
	c, _ := geelark.NewClient("audit", geelark.WithMinInterval(time.Hour), geelark.WithHTTPClient(&http.Client{Transport: auditTransport(func(r *http.Request) (*http.Response, error) {
		calls++
		return &http.Response{StatusCode: 200, Body: io.NopCloser(strings.NewReader(`{"code":0,"data":{"items":[]}}`)), Header: make(http.Header)}, nil
	})}))
	//lint:ignore SA1012 Exercise the SDK's existing nil-context compatibility path.
	if _, err := c.CloudPhones.List(nil, cloudphone.ListRequest{}); err != nil {
		t.Fatal(err)
	}
	ctx, cancel := context.WithCancel(context.Background())
	cancel()
	if _, err := c.CloudPhones.List(ctx, cloudphone.ListRequest{}); !errors.Is(err, context.Canceled) {
		t.Fatal(err)
	}
	if calls != 1 {
		t.Errorf("calls %d", calls)
	}
}
