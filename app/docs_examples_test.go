package app

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
		{"ListInstalledRequest", "{\"envId\":\"1809135651036667904\",\"page\":1,\"pageSize\":5}", func() any { return new(ListInstalledRequest) }},
		{"ListInstalledResponse", "{\"items\":[{\"appIcon\":\"http://cmp1-prod.zxpcloud.com/apps/io.tm.k.drama/K-DRAMA_1716451323126.png\",\"appId\":\"1793552962123993090\",\"appName\":\"K-DRAMA\",\"appVersionId\":\"1793552962140770305\",\"installStatus\":1,\"installTime\":\"2024-07-10 23:07:56\",\"packageName\":\"io.tm.k.drama\",\"versionCode\":\"21120300\",\"versionName\":\"1.0.1\"}],\"total\":1,\"page\":1,\"pageSize\":5}", func() any { return new(ListInstalledResponse) }},
		{"ListAvailableRequest", "{\"name\":\"tiktok\",\"envId\":\"1809135651036667904\",\"getUploadApp\":false,\"page\":1,\"pageSize\":5}", func() any { return new(ListAvailableRequest) }},
		{"ListAvailableResponse", "{\"items\":[{\"appIcon\":\"http://cmp1-prod.zxpcloud.com/apps/io.tm.k.drama/K-DRAMA_1716451323126.png\",\"appName\":\"K-DRAMA\",\"appVersionInfoList\":[{\"id\":\"1793552962140770305\",\"installStatus\":1,\"versionCode\":21120300,\"versionName\":\"1.0.1\"}],\"id\":\"1793552962123993090\",\"installStatus\":1,\"packageName\":\"io.tm.k.drama\"}],\"total\":1,\"page\":1,\"pageSize\":5}", func() any { return new(ListAvailableResponse) }},
		{"InstallRequest", "{\"envId\":\"1809135651036667904\",\"appVersionId\":\"1793552962140770305\"}", func() any { return new(InstallRequest) }},
		{"StartRequest", "{\"envId\":\"1809135651036667904\",\"appVersionId\":\"1793552962140770305\"}", func() any { return new(StartRequest) }},
		{"StopRequest", "{\"envId\":\"1809135651036667904\",\"appVersionId\":\"1793552962140770305\"}", func() any { return new(StopRequest) }},
		{"UninstallRequest", "{\"envId\":\"1809135651036667904\",\"packageName\":\"com.zhiliaoapp.musically\"}", func() any { return new(UninstallRequest) }},
		{"UploadRequest", "{\"fileUrl\":\"https://material.geelark.cn/user-upload/495622401615203830/apk/uCeztwvXpswsHPH5WFht.apk\",\"desc\":\"app description\"}", func() any { return new(UploadRequest) }},
		{"UploadResponse", "{\"taskId\":\"1830906144634757120\"}", func() any { return new(UploadResponse) }},
		{"UploadStatusRequest", "{\"taskId\":\"1830906144634757120\"}", func() any { return new(UploadStatusRequest) }},
		{"UploadStatusResponse", "{\"status\":1,\"appName\":\"\u30d1\u30ef\u30b5\u30ab\",\"appIcon\":\"https://material.geelark.cn/app/icon/20240903/223wgZeZq3.jpg\",\"appId\":\"1813124121385549826\",\"versionId\":\"1813124121435881473\"}", func() any { return new(UploadStatusResponse) }},
		{"ListShopRequest", "{\"key\":\"tiktok\",\"getUploadApp\":false,\"page\":1,\"pageSize\":5}", func() any { return new(ListShopRequest) }},
		{"ListShopResponse", "{\"items\":[{\"appIcon\":\"http://cmp1-prod.zxpcloud.com/apps/io.tm.k.drama/K-DRAMA_1716451323126.png\",\"appName\":\"K-DRAMA\",\"appVersionList\":[{\"id\":\"1793552962140770305\",\"versionCode\":21120300,\"versionName\":\"1.0.1\"}],\"id\":\"1793552962123993090\"}],\"total\":1,\"page\":1,\"pageSize\":5}", func() any { return new(ListShopResponse) }},
		{"AddToTeamRequest", "{\"id\":\"1793552962123993090\",\"versionId\":\"1793552962140770305\",\"installGroupIds\":[\"0\"]}", func() any { return new(AddToTeamRequest) }},
		{"ListTeamRequest", "{\"page\":1,\"pageSize\":5}", func() any { return new(ListTeamRequest) }},
		{"ListTeamResponse", "{\"total\":4,\"page\":1,\"pageSize\":1,\"items\":[{\"id\":\"497652752864775437\",\"appName\":\"TikTok\",\"appIcon\":\"https://material.geelark.com/app/icon/20251026/kVAQ8OuTNF.png\",\"versionId\":\"1793552962123993090\",\"versionCode\":410903,\"versionName\":\"41.9.3\",\"status\":0,\"isUpload\":false,\"uploadStatus\":0,\"appAuth\":0,\"appRoot\":0,\"hideApp\":false,\"envGroups\":[]}]}", func() any { return new(ListTeamResponse) }},
		{"RemoveTeamRequest", "{\"id\":\"497652752864775437\"}", func() any { return new(RemoveTeamRequest) }},
		{"SetAutoInstallRequest", "{\"id\":\"497652752864775437\",\"status\":1}", func() any { return new(SetAutoInstallRequest) }},
		{"SetAuthorizationRequest", "{\"id\":\"497652752864775437\",\"appAuth\":1}", func() any { return new(SetAuthorizationRequest) }},
		{"SetRootRequest", "{\"id\":\"497652752864775437\",\"opera\":1}", func() any { return new(SetRootRequest) }},
		{"BatchOperateRequest", "{\"action\":1,\"groupIds\":[\"595681312937911830\"],\"packageName\":\"io.tm.k.drama\"}", func() any { return new(BatchOperateRequest) }},
		{"BatchOperateResponse", "{\"items\":[{\"id\":\"583502967211075086\",\"errCode\":2}]}", func() any { return new(BatchOperateResponse) }},
		{"SetKeepAliveRequest", "{\"id\":\"497652752864775437\",\"opera\":1}", func() any { return new(SetKeepAliveRequest) }},
		{"SetAutoStartRequest", "{\"id\":\"497652752864775437\",\"opera\":1}", func() any { return new(SetAutoStartRequest) }},
		{"SetHiddenRequest", "{\"id\":\"627129068445269681\",\"opera\":1}", func() any { return new(SetHiddenRequest) }},
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
