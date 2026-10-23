package geelark

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"reflect"
	"strings"
	"testing"
	"time"

	"github.com/wyosa/geelark-go-sdk/browser"
	"github.com/wyosa/geelark-go-sdk/cloudphone"
	"github.com/wyosa/geelark-go-sdk/common"
	"github.com/wyosa/geelark-go-sdk/files"
	"github.com/wyosa/geelark-go-sdk/library"
)

func TestConstructors(t *testing.T) {
	if _, err := NewClient(""); err == nil {
		t.Fatal("accepted an empty token")
	}
	if _, err := NewKeyClient("app", ""); err == nil {
		t.Fatal("accepted an empty API key")
	}
	if _, err := NewBrowserClient(""); err == nil {
		t.Fatal("accepted an empty browser token")
	}
	cloud, err := NewClient("token")
	if err != nil || cloud.CloudPhones == nil || cloud.Automation == nil || cloud.ProxyGroups == nil || cloud.Billing == nil {
		t.Fatalf("cloud services unavailable: %v", err)
	}
	local, err := NewBrowserClient("token")
	if err != nil || local.Browsers == nil || local.Automation == nil {
		t.Fatalf("browser services unavailable: %v", err)
	}
}

func TestLatestCloudPhoneFields(t *testing.T) {
	var bodies = make(chan struct {
		path string
		body []byte
	}, 2)
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		body, _ := io.ReadAll(r.Body)
		bodies <- struct {
			path string
			body []byte
		}{r.URL.Path, body}
		_, _ = w.Write([]byte(`{"code":0,"data":{}}`))
	}))
	defer srv.Close()
	c, err := NewClient("token", WithBaseURL(srv.URL), WithMinInterval(time.Nanosecond))
	if err != nil {
		t.Fatal(err)
	}
	_, err = c.CloudPhones.Create(context.Background(), cloudphone.CreateRequest{
		MobileType: "Android 15",
		Data: []cloudphone.CreateRequestDataItem{{
			MobileType: "Android 16", ProfileName: "phone", WiFiID: "office",
			SessionType: common.Ptr(2), FilterMode: common.Ptr(0),
		}},
	})
	if err != nil {
		t.Fatal(err)
	}
	first := <-bodies
	assertJSON(t, first.body, `{"mobileType":"Android 15","data":[{"mobileType":"Android 16","profileName":"phone","wifiId":"office","sessionType":2,"filterMode":0}]}`)
	if first.path != "/open/v1/phone/addNew" {
		t.Fatalf("create path = %q", first.path)
	}
	_, err = c.CloudPhones.Start(context.Background(), cloudphone.StartRequest{
		IDs: []string{"p1"}, EnableEnergySavingModeAfterMinutes: common.Ptr(0),
		CleanupSDcardAfterDays: common.Ptr(0), DarkMode: common.Ptr(false),
	})
	if err != nil {
		t.Fatal(err)
	}
	second := <-bodies
	assertJSON(t, second.body, `{"ids":["p1"],"enableEnergySavingModeAfterMinutes":0,"cleanupSDcardAfterDays":0,"darkMode":false}`)
}

func TestBrowserExtensionGroupsUseRemoteHost(t *testing.T) {
	localCalls := make(chan string, 1)
	remoteCalls := make(chan string, 1)
	local := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		localCalls <- r.URL.Path
		_, _ = w.Write([]byte(`{"code":0,"data":{"total":0,"items":[]}}`))
	}))
	defer local.Close()
	remote := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		remoteCalls <- r.URL.Path
		_, _ = w.Write([]byte(`{"code":0,"data":{"total":0,"items":[]}}`))
	}))
	defer remote.Close()
	c, err := NewBrowserClient("token", WithBaseURL(local.URL), WithRemoteBaseURL(remote.URL))
	if err != nil {
		t.Fatal(err)
	}
	if _, err := c.Browsers.ExtensionGroups(context.Background(), browser.ExtensionGroupsRequest{Page: 1}); err != nil {
		t.Fatal(err)
	}
	if got := <-remoteCalls; got != "/open/v1/browser/extGroup/list" {
		t.Fatalf("extension path = %q", got)
	}
	if _, err := c.Browsers.List(context.Background(), browser.ListRequest{Page: 1}); err != nil {
		t.Fatal(err)
	}
	if got := <-localCalls; got != "/api/v1/browser/list" {
		t.Fatalf("local path = %q", got)
	}
}

func TestSignedUploadWorkflows(t *testing.T) {
	var uploads []string
	storage := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPut || r.Header.Get("Authorization") != "" || r.Header.Get("traceId") != "" {
			t.Errorf("signed upload sent API credentials: %v", r.Header)
		}
		body, _ := io.ReadAll(r.Body)
		uploads = append(uploads, string(body))
		w.WriteHeader(http.StatusOK)
	}))
	defer storage.Close()
	api := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path != "/open/v1/upload/getUrl" && r.URL.Path != "/open/v1/material/getUploadUrl" {
			t.Errorf("unexpected API path %q", r.URL.Path)
		}
		_, _ = fmt.Fprintf(w, `{"code":0,"data":{"uploadUrl":%q,"resourceUrl":"https://material.geelark.com/file"}}`, storage.URL+"/upload?Signature=test")
	}))
	defer api.Close()
	c, err := NewClient("token", WithBaseURL(api.URL))
	if err != nil {
		t.Fatal(err)
	}
	temp, err := c.Files.Upload(context.Background(), files.GetUploadURLRequest{FileType: "mp4"}, strings.NewReader("video"), 5)
	if err != nil || temp.ResourceURL == "" {
		t.Fatalf("temporary upload: %v, %+v", err, temp)
	}
	stored, err := c.Library.Upload(context.Background(), library.GetUploadURLRequest{FileType: "jpg"}, strings.NewReader("image"), 5)
	if err != nil || stored.ResourceURL == "" {
		t.Fatalf("library upload: %v, %+v", err, stored)
	}
	if !reflect.DeepEqual(uploads, []string{"video", "image"}) {
		t.Fatalf("uploads = %q", uploads)
	}
}

func assertJSON(t *testing.T, got []byte, want string) {
	t.Helper()
	var a, b any
	if err := json.Unmarshal(got, &a); err != nil {
		t.Fatal(err)
	}
	if err := json.Unmarshal([]byte(want), &b); err != nil {
		t.Fatal(err)
	}
	if !reflect.DeepEqual(a, b) {
		t.Errorf("JSON = %s, want %s", got, want)
	}
}
