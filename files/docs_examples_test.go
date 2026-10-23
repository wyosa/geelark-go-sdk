package files

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
		{"GetUploadURLRequest", "{\"fileType\":\"mp4\"}", func() any { return new(GetUploadURLRequest) }},
		{"GetUploadURLResponse", "{\"uploadUrl\":\"http://42-studio-prod.oss-cn-hangzhou.aliyuncs.com/open-upload%2F497521349346987872%2F20240730%2Fe2u5amyB.mp4?Expires=1722310832&OSSAccessKeyId=LTAI5t7JzQBfi1nV3HbsKojG&Signature=HGBVqTUfXcUAthLPnO3ZYIVnAxg%3D\",\"resourceUrl\":\"https://material-prod.geelark.cn/open-upload/497521349346987872/20240730/e2u5amyB.mp4\"}", func() any { return new(GetUploadURLResponse) }},
		{"UploadToPhoneRequest", "{\"id\":\"528715748189668352\",\"fileUrl\":\"https://material-prod.geelark.cn/app/icon/20240506/nFrUEcRc9I.jpg\"}", func() any { return new(UploadToPhoneRequest) }},
		{"UploadToPhoneResponse", "{\"taskId\":\"1850726441252569088\"}", func() any { return new(UploadToPhoneResponse) }},
		{"UploadResultRequest", "{\"taskId\":\"528715748189668352\"}", func() any { return new(UploadResultRequest) }},
		{"UploadResultResponse", "{\"status\":1}", func() any { return new(UploadResultResponse) }},
		{"UploadKeyboxRequest", "{\"id\":\"528715748189668352\",\"fileUrl\":\"https://material.geelark.cn/client-img/oakendn.xml\"}", func() any { return new(UploadKeyboxRequest) }},
		{"UploadKeyboxResponse", "{\"taskId\":\"1850726441252569088\"}", func() any { return new(UploadKeyboxResponse) }},
		{"KeyboxResultRequest", "{\"taskId\":\"528715748189668352\"}", func() any { return new(KeyboxResultRequest) }},
		{"KeyboxResultResponse", "{\"status\":1}", func() any { return new(KeyboxResultResponse) }},
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
