package geelark

import (
	"encoding/json"
	"testing"

	"github.com/wyosa/geelark-go-sdk/automation"
	"github.com/wyosa/geelark-go-sdk/browser"
	"github.com/wyosa/geelark-go-sdk/cloudphone"
	"github.com/wyosa/geelark-go-sdk/common"
	"github.com/wyosa/geelark-go-sdk/oem"
	"github.com/wyosa/geelark-go-sdk/proxy"
)

func TestNestedRequestFields(t *testing.T) {
	tests := []struct {
		name string
		in   any
		want string
	}{
		{
			"proxy group and detection channel",
			proxy.AddRequest{List: []proxy.AddItem{{Scheme: "socks5", Server: "proxy.example", Port: 1080, ProxyGroupID: "group", ProxyQueryChannel: common.Ptr(1)}}},
			`{"list":[{"scheme":"socks5","server":"proxy.example","port":1080,"proxyGroupId":"group","proxyQueryChannel":1}]}`,
		},
		{
			"GProxy browser configuration",
			browser.CreateRequest{SerialName: "browser", BrowserOS: 2, ProxyConfig: &browser.CreateRequestProxyConfig{TypeID: common.Ptr(30), SessionType: common.Ptr(2), FilterMode: common.Ptr(0)}},
			`{"serialName":"browser","browserOs":2,"proxyConfig":{"typeId":30,"sessionType":2,"filterMode":0}}`,
		},
		{
			"GProxy cloud phone update",
			cloudphone.UpdateRequest{ID: "phone", ProxyConfig: &cloudphone.UpdateRequestProxyConfig{TypeID: common.Ptr(30), UseProxyCfg: common.Ptr(true), SessionType: common.Ptr(2)}},
			`{"id":"phone","proxyConfig":{"typeId":30,"useProxyCfg":true,"sessionType":2}}`,
		},
		{
			"warmup probability and latest task fields",
			automation.AddTaskRequest{TaskType: 2, List: []automation.AddTaskItem{{ScheduleAt: 1760000000, EnvID: "phone", Action: "search video", Duration: common.Ptr(10), QRCode: "qr", Index: common.Ptr(0), Probability: &automation.AddTaskProbability{HasKey: &automation.AddTaskBehavior{Like: common.Ptr(0)}}}}},
			`{"taskType":2,"list":[{"scheduleAt":1760000000,"envId":"phone","action":"search video","duration":10,"qrcode":"qr","index":0,"probability":{"hasKey":{"like":0}}}]}`,
		},
		{
			"OEM toolbar icon",
			oem.CustomizationRequest{ToolBarSettings: []oem.CustomizationRequestToolBarSettingsItem{{ToolBar: "networkQuality", Visible: common.Ptr(true), IconURL: "https://example.com/icon.svg"}}},
			`{"toolBarSettings":[{"toolBar":"networkQuality","visible":true,"iconUrl":"https://example.com/icon.svg"}]}`,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := json.Marshal(tt.in)
			if err != nil {
				t.Fatal(err)
			}
			assertJSON(t, got, tt.want)
		})
	}
}
