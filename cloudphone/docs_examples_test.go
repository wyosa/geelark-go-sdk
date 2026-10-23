package cloudphone

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
		{"CreateRequest", "{\"mobileType\":\"Android 12\",\"chargeMode\":0,\"data\":[{\"profileName\":\"myPhone\",\"proxyInformation\":\"socks5://AD00xx004:3000xxx0002@100.200.200.100:30000\",\"proxyQueryChannel\":1,\"mobileLanguage\":\"default\",\"profileGroup\":\"myGroup\",\"profileTags\":[\"myTag\"],\"profileNote\":\"remark\"}]}", func() any { return new(CreateRequest) }},
		{"CreateResponse", "{\"totalAmount\":1,\"successAmount\":1,\"failAmount\":0,\"details\":[{\"index\":1,\"code\":0,\"msg\":\"success\",\"id\":\"497652752864775437\",\"profileName\":\"22 ungrouped\",\"envSerialNo\":\"22\",\"equipmentInfo\":{\"countryName\":\"Thailand\",\"phoneNumber\":\"+66877382166\",\"enableSim\":1,\"imei\":\"863406055475987\",\"osVersion\":\"Android 11.0\",\"wifiBssid\":\"1C:1D:67:B1:C1:76\",\"mac\":\"9C:A5:C0:5F:C5:AD\",\"bluetoothMac\":\"D0:15:4A:5B:7E:AE\",\"timeZone\":\"Asia/Bangkok\"}}]}", func() any { return new(CreateResponse) }},
		{"ListRequest", "{\"page\":1,\"pageSize\":10,\"serialName\":\"test\",\"remark\":\"\",\"groupName\":\"\",\"tags\":[\"tag1\",\"tag2\"],\"openStatus\":1,\"proxyIds\":[\"602943680722009770\"],\"serialNos\":[\"238\"]}", func() any { return new(ListRequest) }},
		{"ListResponse", "{\"total\":1,\"page\":1,\"pageSize\":10,\"items\":[{\"id\":\"123456ABCDEF\",\"serialName\":\"test\",\"serialNo\":\"1\",\"group\":{\"id\":\"123456ABCDEF\",\"name\":\"test group\",\"remark\":\"group remark\"},\"remark\":\"env remark\",\"status\":0,\"tags\":[{\"name\":\"hi\"},{\"name\":\"test\"}],\"equipmentInfo\":{\"countryName\":\"Thailand\",\"phoneNumber\":\"+66877382166\",\"enableSim\":1,\"imei\":\"863406055475987\",\"osVersion\":\"Android 11.0\",\"wifiBssid\":\"1C:1D:67:B1:C1:76\",\"mac\":\"9C:A5:C0:5F:C5:AD\",\"bluetoothMac\":\"D0:15:4A:5B:7E:AE\",\"timeZone\":\"Asia/Bangkok\",\"deviceName\":\"\",\"netType\":1,\"language\":\"en-US\",\"province\":\"\",\"city\":\"\"},\"proxy\":{\"type\":\"socks5\",\"server\":\"129.129.129.129\",\"port\":30000,\"username\":\"user\",\"password\":\"pass\"}}]}", func() any { return new(ListResponse) }},
		{"StatusRequest", "{\"ids\":[\"123456ABCDEF\",\"123456ABCDEF\",\"123456ABCDEF\",\"123456ABCDEF\"]}", func() any { return new(StatusRequest) }},
		{"StatusResponse", "{\"totalAmount\":4,\"successAmount\":3,\"failAmount\":1,\"successDetails\":[{\"id\":\"123456ABCDEF\",\"serialName\":\"name1\",\"status\":0},{\"id\":\"123456ABCDEF\",\"serialName\":\"name2\",\"status\":1},{\"id\":\"123456ABCDEF\",\"serialName\":\"name3\",\"status\":1}],\"failDetails\":[{\"code\":42001,\"id\":\"123456ABCDEF\",\"msg\":\"env not found\"}]}", func() any { return new(StatusResponse) }},
		{"StartRequest", "{\"ids\":[\"123456ABCDEF\",\"123456ABCDEF\",\"123456ABCDEF\",\"123456ABCDEF\"]}", func() any { return new(StartRequest) }},
		{"StartResponse", "{\"totalAmount\":3,\"successAmount\":1,\"failAmount\":2,\"failDetails\":[{\"code\":43004,\"id\":\"12345678ABCDEFG\",\"msg\":\"env is expired\"},{\"code\":42001,\"id\":\"12345678ABCDEFG\",\"msg\":\"env not found\"}],\"successDetails\":[{\"id\":\"12345678ABCDEFG\",\"url\":\"https://speedup.geelark.com/phone-api\",\"chargingMethod\":\"Per-minute usage\"}]}", func() any { return new(StartResponse) }},
		{"StopRequest", "{\"ids\":[\"123456ABCDEF\",\"123456ABCDEF\",\"123456ABCDEF\",\"123456ABCDEF\"]}", func() any { return new(StopRequest) }},
		{"StopResponse", "{\"totalAmount\":4,\"successAmount\":3,\"failAmount\":1,\"successDetails\":[{\"id\":\"123456ABCDEF\",\"serialName\":\"name1\",\"status\":0},{\"id\":\"123456ABCDEF\",\"serialName\":\"name2\",\"status\":1},{\"id\":\"123456ABCDEF\",\"serialName\":\"name3\",\"status\":1}],\"failDetails\":[{\"code\":42001,\"id\":\"123456ABCDEF\",\"msg\":\"env not found\"}]}", func() any { return new(StopResponse) }},
		{"RestartRequest", "{\"id\":\"631490227545875981\"}", func() any { return new(RestartRequest) }},
		{"DeleteRequest", "{\"ids\":[\"123456ABCDEF\",\"123456ABCDEF\"]}", func() any { return new(DeleteRequest) }},
		{"DeleteResponse", "{\"totalAmount\":4,\"successAmount\":2,\"failAmount\":2,\"failDetails\":[{\"code\":42001,\"id\":\"12345ABCDEF\",\"msg\":\"env not found\"},{\"code\":43009,\"id\":\"12345ABCDEF\",\"msg\":\"env is started\"}]}", func() any { return new(DeleteResponse) }},
		{"GetGPSRequest", "{\"ids\":[\"528086321789535232\"]}", func() any { return new(GetGPSRequest) }},
		{"GetGPSResponse", "{\"totalAmount\":1,\"successAmount\":1,\"failAmount\":0,\"list\":[{\"id\":\"528086321789535232\",\"latitude\":1.3024300336837769,\"longitude\":103.87545776367188}]}", func() any { return new(GetGPSResponse) }},
		{"SetGPSRequest", "{\"list\":[{\"id\":\"528086321789535232\",\"latitude\":1.30243,\"longitude\":103.87546},{\"id\":\"530011895768286208\",\"latitude\":11.30243,\"longitude\":104.87546}]}", func() any { return new(SetGPSRequest) }},
		{"SetGPSResponse", "{\"totalAmount\":2,\"successAmount\":2,\"failAmount\":0}", func() any { return new(SetGPSResponse) }},
		{"UpdateRequest", "{\"id\":\"528086284158239744\",\"name\":\"api update\",\"remark\":\"api remark\",\"tagIDs\":[\"528989565877355520\",\"528989565877289984\"],\"groupID\":\"528995439832269824\",\"proxyQueryChannel\":1,\"proxyConfig\":{\"typeId\":1,\"server\":\"123.123.123.123\",\"port\":32080,\"username\":\"username\",\"password\":\"password\"},\"proxyId\":\"528989565877355520\"}", func() any { return new(UpdateRequest) }},
		{"ResetRequest", "{\"id\":\"528715748189668352\"}", func() any { return new(ResetRequest) }},
		{"CloneRequest", "{\"envId\":\"590711571886417452\",\"amount\":1,\"groupId\":\"590711571886417453\"}", func() any { return new(CloneRequest) }},
		{"CloneResponse", "{\"ids\":[\"590711571886417454\"]}", func() any { return new(CloneResponse) }},
		{"ScreenshotRequest", "{\"id\":\"528715748189668352\"}", func() any { return new(ScreenshotRequest) }},
		{"ScreenshotResponse", "{\"taskId\":\"1850726441252569088\"}", func() any { return new(ScreenshotResponse) }},
		{"ScreenshotResultRequest", "{\"taskId\":\"528715748189668352\"}", func() any { return new(ScreenshotResultRequest) }},
		{"ScreenshotResultResponse", "{\"status\":2,\"downloadLink\":\"https://zx-cloud-phone-pre.obs.cn-southwest-2.myhuaweicloud.com/envirFileExport/1851511129017700352/IMG_20241122160248.png?AccessKeyId=UFNIPAPFJX2MAGMFGRYZ&Expires=1732264377&Signature=dqV5JYzYDdm0wwAgkZIpDrs%2FL%2FE%3D\"}", func() any { return new(ScreenshotResultResponse) }},
		{"SetRootRequest", "{\"ids\":[\"526209711930868736\"],\"open\":true}", func() any { return new(SetRootRequest) }},
		{"SetRootResponse", "{\"items\":[{\"code\":42002,\"msg\":\"phone is not running\",\"id\":\"543483007558772199\"},{\"code\":0,\"msg\":\"success\",\"id\":\"543483063829554663\"}]}", func() any { return new(SetRootResponse) }},
		{"GetRootRequest", "{\"ids\":[\"526209711930868736\"]}", func() any { return new(GetRootRequest) }},
		{"GetRootResponse", "{\"items\":[{\"code\":0,\"msg\":\"success\",\"id\":\"543483063829554663\",\"open\":true},{\"code\":42001,\"msg\":\"phone is not exist\",\"id\":\"543483007558772199\",\"open\":false}]}", func() any { return new(GetRootResponse) }},
		{"DeviceIDRequest", "{\"id\":\"528715748189668352\"}", func() any { return new(DeviceIDRequest) }},
		{"DeviceIDResponse", "{\"serialNum\":\"r2cbvzlx5bs\"}", func() any { return new(DeviceIDResponse) }},
		{"SendSMSRequest", "{\"id\":\"526209711930868736\",\"phoneNumber\":\"+17723504471\",\"text\":\"your tk code: 6666\"}", func() any { return new(SendSMSRequest) }},
		{"BrandsRequest", "{\"androidVer\":10}", func() any { return new(BrandsRequest) }},
		{"BrandsResponse", "[{\"surfaceBrandName\":\"samsung\",\"surfaceModelName\":\"Galaxy S20+\"}]", func() any { return new(BrandsResponse) }},
		{"TeamBrandsRequest", "{\"page\":1,\"pageSize\":10,\"androidVer\":10}", func() any { return new(TeamBrandsRequest) }},
		{"TeamBrandsResponse", "{\"total\":1,\"page\":1,\"pageSize\":10,\"list\":[{\"surfaceBrandName\":\"HUAWEI\",\"surfaceModelName\":\"SPN-AL00\"}]}", func() any { return new(TeamBrandsResponse) }},
		{"TransferRequest", "{\"ids\":[\"539893235657500146\"],\"account\":\"Anna@geelark.com\",\"transferOption\":[\"name\",\"proxy\",\"tag\",\"remark\"]}", func() any { return new(TransferRequest) }},
		{"TransferResponse", "[{\"successCount\":10,\"failCount\":2,\"failEnvIds\":[\"539893235657500146\"]}]", func() any { return new(TransferResponse) }},
		{"SetNetworkTypeRequest", "{\"id\":\"528086284158239744\",\"netType\":0}", func() any { return new(SetNetworkTypeRequest) }},
		{"HideAccessibilityRequest", "{\"ids\":[\"599257164413959866\"],\"pkgName\":[\"com.zhiliaoapp.musically\"]}", func() any { return new(HideAccessibilityRequest) }},
		{"HideAccessibilityResponse", "{\"failDetails\":[{\"id\":\"599257164413959866\",\"code\":42001,\"msg\":\"env not found\"}]}", func() any { return new(HideAccessibilityResponse) }},
		{"MoveGroupRequest", "{\"envIds\":[\"601876382020062634\"],\"groupId\":\"590711571886417453\"}", func() any { return new(MoveGroupRequest) }},
		{"ImportContactsRequest", "{\"id\":\"557536075321468390\",\"contacts\":[{\"firstName\":\"jay\",\"mobile\":\"13288888888\"}]}", func() any { return new(ImportContactsRequest) }},
		{"ImportContactsResponse", "{\"taskId\":\"558017255909123564\"}", func() any { return new(ImportContactsResponse) }},
		{"ImportContactsResultRequest", "{\"taskId\":\"528715748189668352\"}", func() any { return new(ImportContactsResultRequest) }},
		{"ImportContactsResultResponse", "{\"status\":2}", func() any { return new(ImportContactsResultResponse) }},
		{"GetNetworkSettingsResponse", "{\"blackList\":[\"c.com\"]}", func() any { return new(GetNetworkSettingsResponse) }},
		{"SetNetworkSettingsRequest", "{\"blackList\":[\"c.com\"]}", func() any { return new(SetNetworkSettingsRequest) }},
		{"UIDumpRequest", "{\"id\":\"528715748189668352\"}", func() any { return new(UIDumpRequest) }},
		{"UIDumpResponse", "{\"taskId\":\"1850726441252569088\"}", func() any { return new(UIDumpResponse) }},
		{"UIDumpResultRequest", "{\"taskId\":\"1850726441252569088\"}", func() any { return new(UIDumpResultRequest) }},
		{"UIDumpResultResponse", "{\"status\":2,\"downloadLink\":\"https://test.com/test.xml\"}", func() any { return new(UIDumpResultResponse) }},
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
			if err := json.Unmarshal(got, &b); err != nil {
				t.Fatal(err)
			}
			if !reflect.DeepEqual(a, b) {
				t.Errorf("roundtrip mismatch: want %s, got %s", tc.raw, got)
			}
		})
	}
}
