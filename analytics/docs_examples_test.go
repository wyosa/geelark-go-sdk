package analytics

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
		{"DataRequest", "{\"account\":\"tk_acc\",\"dataDate\":1764137986,\"createdId\":\"497521349330210656\",\"channel\":1,\"page\":1,\"pageSize\":10}", func() any { return new(DataRequest) }},
		{"DataResponse", "{\"total\":2,\"page\":1,\"pageSize\":10,\"items\":[{\"id\":\"581465531480044822\",\"channel\":0,\"account\":\"khian.kb\",\"playCount\":-1,\"followerCount\":-1,\"diggCount\":-1,\"commentCount\":-1,\"collectCount\":-1,\"shareCount\":-1,\"dataDate\":-1,\"addAccDate\":1756110064,\"remark\":\"\",\"createdId\":\"497521349330210656\",\"username\":\"exlrhoo@foxmail.com\"}]}", func() any { return new(DataResponse) }},
		{"AddAccountsRequest", "{\"channel\":0,\"accountsData\":[{\"account\":\"account-a\",\"remark\":\"remark-a\",\"tagIds\":[\"628384126517345281\",\"628384126517345282\"]},{\"account\":\"account-b\",\"remark\":\"remark-b\",\"tagIds\":[\"628384126517345283\"]}]}", func() any { return new(AddAccountsRequest) }},
		{"AddAccountsResponse", "{\"bizCode\":0,\"successCount\":2,\"failCount\":0,\"repeatCount\":0}", func() any { return new(AddAccountsResponse) }},
		{"DeleteAccountRequest", "{\"channel\":0,\"account\":\"xxxx\"}", func() any { return new(DeleteAccountRequest) }},
		{"ListAccountsRequest", "{\"page\":1,\"pageSize\":10,\"channel\":1}", func() any { return new(ListAccountsRequest) }},
		{"ListAccountsResponse", "{\"items\":[{\"id\":\"565523829426802069\",\"account\":\"xxx\",\"channel\":1,\"remark\":\"remark\",\"operator\":\"xxx\",\"created_time\":1746608069,\"updated_time\":1746608069,\"tags\":[{\"id\":\"123\",\"name\":\"vip\",\"color\":1}]}],\"total\":1,\"page\":1}", func() any { return new(ListAccountsResponse) }},
		{"UpdateAccountRequest", "{\"id\":\"565523829426802069\",\"account\":\"abc\",\"tagIds\":[\"628384126517345281\",\"628384126517345282\"]}", func() any { return new(UpdateAccountRequest) }},
		{"CreateTagRequest", "{\"name\":\"Important Account\",\"color\":1}", func() any { return new(CreateTagRequest) }},
		{"ListTagsRequest", "{\"name\":\"Important\",\"page\":1,\"pageSize\":200}", func() any { return new(ListTagsRequest) }},
		{"ListTagsResponse", "{\"total\":2,\"page\":1,\"pageSize\":200,\"list\":[{\"id\":\"628384126517345281\",\"name\":\"Important Account\",\"color\":1},{\"id\":\"628384126517345282\",\"name\":\"Pending Follow-up\",\"color\":2}]}", func() any { return new(ListTagsResponse) }},
		{"UpdateTagRequest", "{\"id\":\"628384126517345281\",\"name\":\"Core Account\",\"color\":2}", func() any { return new(UpdateTagRequest) }},
		{"DeleteTagsRequest", "{\"ids\":[\"628384126517345281\",\"628384126517345282\"]}", func() any { return new(DeleteTagsRequest) }},
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
