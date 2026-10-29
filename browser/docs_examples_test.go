package browser

import (
	"encoding/json"
	"reflect"
	"testing"
)

func TestDocumentationExamples(t *testing.T) {
	tests := []struct {
		name     string
		raw      string
		newValue func() any
	}{
		{"CreateRequest", "{\"serialName\":\"myBrowserName\",\"extGroup\":\"497548067550006541\",\"browserOs\":1,\"proxyConfig\":{\"typeId\":-1},\"simulateConfig\":{\"timeZone\":{\"switcher\":2,\"value\":\"GMT-12:00 Etc/GMT+12\"},\"language\":{\"switcher\":2,\"value\":\"Albanian\"},\"resolution\":{\"switcher\":2,\"value\":\"750*1334\"},\"webRtc\":{\"switcher\":1},\"geoLocation\":{\"switcher\":1,\"baseOnIp\":true,\"longitude\":20,\"latitude\":10,\"accuracy\":1},\"canvas\":{\"switcher\":1},\"webglImage\":{\"switcher\":1},\"hardware\":{\"switcher\":1},\"audioContext\":{\"switcher\":1},\"mediaDevice\":{\"switcher\":1},\"clientRects\":{\"switcher\":1},\"speechVoise\":{\"switcher\":1},\"hardwareConcurrency\":26,\"memeryDevice\":8,\"doNotTrack\":2,\"bluetooth\":{\"switcher\":1},\"battery\":{\"switcher\":1},\"portScanProtection\":{\"switcher\":1,\"value\":\"80\"}}}", func() any { return new(CreateRequest) }},
		{"CreateResponse", "{\"id\":\"497548067550006541\"}", func() any { return new(CreateResponse) }},
		{"UpdateRequest", "{\"id\":\"497548067550006541\",\"extGroup\":\"497548067550006541\",\"simulateConfig\":{\"timeZone\":{\"switcher\":2,\"value\":\"GMT-12:00 Etc/GMT+12\"},\"language\":{\"switcher\":2,\"value\":\"Albanian\"},\"resolution\":{\"switcher\":2,\"value\":\"750*1334\"},\"webRtc\":{\"switcher\":1},\"geoLocation\":{\"switcher\":1,\"baseOnIp\":true,\"longitude\":20,\"latitude\":10,\"accuracy\":1},\"canvas\":{\"switcher\":1},\"webglImage\":{\"switcher\":1},\"hardware\":{\"switcher\":1},\"audioContext\":{\"switcher\":1},\"mediaDevice\":{\"switcher\":1},\"clientRects\":{\"switcher\":1},\"speechVoise\":{\"switcher\":1},\"hardwareConcurrency\":26,\"memeryDevice\":8,\"doNotTrack\":2,\"bluetooth\":{\"switcher\":1},\"battery\":{\"switcher\":1},\"portScanProtection\":{\"switcher\":1,\"value\":\"80\"}}}", func() any { return new(UpdateRequest) }},
		{"DeleteRequest", "{\"envIds\":[\"539893235657500146\"]}", func() any { return new(DeleteRequest) }},
		{"DeleteResponse", "{\"busyIds\":[\"539893235657500146\"],\"serverErrIds\":[\"539893235657500147\"],\"successIds\":[\"539893235657500148\"]}", func() any { return new(DeleteResponse) }},
		{"ListRequest", "{\"page\":1,\"pageSize\":10}", func() any { return new(ListRequest) }},
		{"ListResponse", "{\"total\":1,\"page\":1,\"pageSize\":10,\"items\":[{\"id\":\"5213214343124321\",\"serialName\":\"myEnv\",\"serialNo\":\"3\",\"group\":{\"id\":\"5213214343124321\",\"name\":\"myGroup\",\"remark\":\"myRemark\"},\"extGroup\":\"497548067550006541\",\"remark\":\"myRemark\",\"tags\":[{\"name\":\"myTag\"}],\"proxy\":{\"type\":\"\",\"server\":\"\",\"port\":0,\"username\":\"\",\"password\":\"\"},\"accountInfo\":{\"url\":\"https://www.tiktok.com/\",\"userName\":\"jay\",\"passWord\":\"password\",\"totpSecret\":\"\",\"afterStartup\":3,\"openLastPage\":2,\"openSpecPage\":1,\"openSiteUrl\":true,\"autoOpenUrls\":[\"http://www.b.com\"]},\"simulateInfo\":{\"os\":2,\"vendor\":1,\"mixtureKey\":\"87da5186e1feabc1\",\"ua\":\"Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/133.0.6943.141 Safari/537.36\",\"uaVersion\":\"133\",\"timeZone\":{\"switcher\":2,\"value\":\"GMT-12:00 Etc/GMT+12\"},\"webRtc\":{\"switcher\":1},\"geoLocation\":{\"switcher\":1,\"baseOnIp\":true,\"longitude\":20,\"latitude\":10,\"accuracy\":1},\"language\":{\"switcher\":2,\"value\":\"Albanian\"},\"resolution\":{\"switcher\":2,\"value\":\"750*1334\"},\"font\":{\"switcher\":2},\"canvas\":{\"switcher\":1},\"webglImage\":{\"switcher\":1},\"webglMetadata\":{\"switcher\":3,\"provider\":\"Google Inc. (Intel Inc.)\",\"render\":\"ANGLE (Intel Inc., Intel(R) HD Graphics 6000, OpenGL 4.1)\"},\"hardware\":{\"switcher\":1},\"audioContext\":{\"switcher\":1},\"mediaDevice\":{\"switcher\":1},\"clientRects\":{\"switcher\":1},\"speechVoise\":{\"switcher\":1},\"hardwareConcurrency\":26,\"memeryDevice\":8,\"doNotTrack\":2,\"bluetooth\":{\"switcher\":1},\"battery\":{\"switcher\":1},\"portScanProtection\":{\"switcher\":1,\"value\":\"80\"}}}]}", func() any { return new(ListResponse) }},
		{"StartRequest", "{\"id\":\"612342716134614477\",\"webhook\":\"http://localhost:3001\"}", func() any { return new(StartRequest) }},
		{"StartResponse", "{\"debugPort\":11001}", func() any { return new(StartResponse) }},
		{"StopRequest", "{\"id\":\"539893235657500146\"}", func() any { return new(StopRequest) }},
		{"TransferRequest", "{\"envIds\":[\"539893235657500146\"],\"username\":\"Anna@geelark.com\"}", func() any { return new(TransferRequest) }},
		{"TransferResponse", "{\"successCount\":10,\"failCount\":2,\"failEnvIds\":[\"539893235657500146\"]}", func() any { return new(TransferResponse) }},
		{"StatusRequest", "{\"id\":\"539893235657500146\"}", func() any { return new(StatusRequest) }},
		{"StatusResponse", "{\"status\":\"open\",\"debugPort\":11000}", func() any { return new(StatusResponse) }},
		{"ClearCacheRequest", "{\"ids\":[\"123456789\",\"987654321\"],\"type\":[\"local_storage\",\"indexeddb\",\"extension_cache\",\"cookie\",\"history\",\"image_file\"]}", func() any { return new(ClearCacheRequest) }},
		{"ClearCacheResponse", "{\"success_count\":2,\"error_count\":0,\"error_info\":[]}", func() any { return new(ClearCacheResponse) }},
		{"MoveGroupRequest", "{\"envIds\":[\"123456789\",\"987654321\"],\"groupId\":\"0\"}", func() any { return new(MoveGroupRequest) }},
		{"CookiesRequest", "{\"id\":\"123456789\"}", func() any { return new(CookiesRequest) }},
		{"CookiesResponse", "{\"cookies\":\"[]\"}", func() any { return new(CookiesResponse) }},
		{"SetBookmarksRequest", "{\"browserBookmark\":{\"type\":2,\"fileAddr\":\"\",\"text\":\"http://a.com\\nhttp://b.com\"}}", func() any { return new(SetBookmarksRequest) }},
		{"GetBookmarksResponse", "{\"browserBookmark\":{\"type\":2,\"fileAddr\":\"\",\"text\":\"http://a.com\\nhttp://b.com\"}}", func() any { return new(GetBookmarksResponse) }},
		{"CloneRequest", "{\"envId\":\"590711571886417452\",\"amount\":1,\"groupId\":\"590711571886417453\"}", func() any { return new(CloneRequest) }},
		{"CloneResponse", "{\"ids\":[\"590711571886417454\"]}", func() any { return new(CloneResponse) }},
		{"KernelsResponse", "[{\"kernel\":\"147\",\"isDownloaded\":true}]", func() any { return new(KernelsResponse) }},
		{"UpdateKernelsRequest", "{\"kernel_version\":\"143\"}", func() any { return new(UpdateKernelsRequest) }},
		{"UpdateKernelsResponse", "{\"status\":\"complete\",\"progress\":0}", func() any { return new(UpdateKernelsResponse) }},
		{"ExtensionGroupsRequest", "{\"page\":1,\"pageSize\":10}", func() any { return new(ExtensionGroupsRequest) }},
		{"ExtensionGroupsResponse", "{\"total\":2,\"page\":1,\"pageSize\":10,\"items\":[{\"id\":\"504916330621701622\",\"name\":\"Category A\"},{\"id\":\"509711117119982738\",\"name\":\"Category B\"}]}", func() any { return new(ExtensionGroupsResponse) }},
	}
	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			value := tc.newValue()
			if err := json.Unmarshal([]byte(tc.raw), value); err != nil {
				t.Fatal(err)
			}
			got, err := json.Marshal(value)
			if err != nil {
				t.Fatal(err)
			}
			var a, b any
			if err := json.Unmarshal([]byte(tc.raw), &a); err != nil {
				t.Fatal(err)
			}
			mirrorDeviceMemory(a)
			if err := json.Unmarshal(got, &b); err != nil {
				t.Fatal(err)
			}
			if !reflect.DeepEqual(a, b) {
				t.Errorf("roundtrip mismatch: want %s, got %s", tc.raw, got)
			}
		})
	}
}

// mirrorDeviceMemory adds memoryDevice wherever an example has only memeryDevice,
// because the SDK sends and returns device memory under both documented spellings.
func mirrorDeviceMemory(v any) {
	switch v := v.(type) {
	case map[string]any:
		if memory, ok := v["memeryDevice"]; ok {
			if _, ok := v["memoryDevice"]; !ok {
				v["memoryDevice"] = memory
			}
		}
		for _, child := range v {
			mirrorDeviceMemory(child)
		}
	case []any:
		for _, child := range v {
			mirrorDeviceMemory(child)
		}
	}
}

func TestDeviceMemorySpellings(t *testing.T) {
	for _, tc := range []struct {
		name string
		in   any
	}{
		{"create table spelling", CreateRequestSimulateConfig{MemoryDevice: ptr(8)}},
		{"create example spelling", CreateRequestSimulateConfig{MemeryDevice: ptr(8)}},
		{"update table spelling", &UpdateRequestSimulateConfig{MemoryDevice: ptr(8)}},
		{"update example spelling", UpdateRequestSimulateConfig{MemeryDevice: ptr(8)}},
	} {
		got, err := json.Marshal(tc.in)
		if err != nil {
			t.Fatal(err)
		}
		if string(got) != `{"memeryDevice":8,"memoryDevice":8}` {
			t.Errorf("%s: %s", tc.name, got)
		}
	}
	if got, _ := json.Marshal(CreateRequestSimulateConfig{}); string(got) != `{}` {
		t.Errorf("unset memory must be omitted: %s", got)
	}
	for _, raw := range []string{`{"memeryDevice":8}`, `{"memoryDevice":8}`} {
		var info ListResponseItemsItemSimulateInfo
		if err := json.Unmarshal([]byte(raw), &info); err != nil {
			t.Fatal(err)
		}
		if info.MemoryDevice != 8 || info.MemeryDevice != 8 {
			t.Errorf("%s decoded as %d/%d", raw, info.MemoryDevice, info.MemeryDevice)
		}
	}
}

func TestFractionalGeolocation(t *testing.T) {
	var resp ListResponse
	raw := `{"items":[{"simulateInfo":{"geoLocation":{"switcher":3,"longitude":103.87546,"latitude":1.30243}}}]}`
	if err := json.Unmarshal([]byte(raw), &resp); err != nil {
		t.Fatal(err)
	}
	geo := resp.Items[0].SimulateInfo.GeoLocation
	if geo.Longitude != 103.87546 || geo.Latitude != 1.30243 {
		t.Fatalf("geolocation = %+v", geo)
	}
	got, err := json.Marshal(CreateRequestSimulateConfigGeoLocation{Longitude: ptr(103.87546), Latitude: ptr(20.0)})
	if err != nil {
		t.Fatal(err)
	}
	if string(got) != `{"longitude":103.87546,"latitude":20}` {
		t.Fatalf("request geolocation = %s", got)
	}
}

func ptr[T any](v T) *T { return &v }
