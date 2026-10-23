package library

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
		{"GetUploadURLRequest", "{\"fileType\":\"jpg\"}", func() any { return new(GetUploadURLRequest) }},
		{"GetUploadURLResponse", "{\"uploadUrl\":\"http://42-studio-singapore.oss-ap-southeast-3.aliyuncs.com/open-material-upload%2F503609206784396150%2F20260309%2FVwujVZdl.jpg?Expires=1773046044&OSSAccessKeyId=LTAI5t7JzQBfi1nV3HbsKojG&Signature=TZbYGwfcad4XFiyh4mZ1gdVI%2FzI%3D\",\"resourceUrl\":\"https://singapore-upgrade.geelark.com/open-material-upload/503609206784396150/20260309/VwujVZdl.jpg\"}", func() any { return new(GetUploadURLResponse) }},
		{"CreateRequest", "{\"url\":\"https://material.geelark.cn/client-img/banner0903_cn.gif\",\"tagsId\":[\"569577509402731994\",\"569577514586891738\"],\"fileName\":\"a.jpg\"}", func() any { return new(CreateRequest) }},
		{"CreateResponse", "{\"id\":\"570606523940605924\",\"failDetails\":[{\"id\":\"5695775094027319941\",\"code\":43022,\"msg\":\"tag not found\"}]}", func() any { return new(CreateResponse) }},
		{"DeleteRequest", "{\"ids\":[\"569569510948864567\"]}", func() any { return new(DeleteRequest) }},
		{"DeleteResponse", "{\"totalAmount\":1,\"successAmount\":0,\"failAmount\":1,\"failDetails\":[{\"id\":\"5695695109488645671\",\"code\":60005,\"msg\":\"material not found\"}]}", func() any { return new(DeleteResponse) }},
		{"SetTagRequest", "{\"materialsId\":[\"570457374221926935\"],\"tagsId\":[\"570461738663674391\"]}", func() any { return new(SetTagRequest) }},
		{"SetTagResponse", "{\"failDetails\":[{\"id\":\"5704573742219269351\",\"code\":60005,\"msg\":\"material not found\"},{\"id\":\"5704617386636743911\",\"code\":43022,\"msg\":\"tag not found\"}]}", func() any { return new(SetTagResponse) }},
		{"SearchRequest", "{\"fileType\":[1],\"fileName\":\"demo\",\"tagIds\":[\"569577514586891738\"],\"source\":1,\"page\":1,\"pageSize\":50,\"ids\":[\"608127302118696420\"]}", func() any { return new(SearchRequest) }},
		{"SearchResponse", "{\"total\":12,\"page\":1,\"pageSize\":50,\"list\":[{\"id\":\"569546671000653787\",\"createdTime\":1749005870,\"fileName\":\"2025_06_02_11_31_IMG_9607.MOV\",\"fileSize\":120713028,\"fileUrl\":\"https://material.geelark.cn/user-upload/497521349346987872/material-center/Ryzt3vkK14T5asJQPx2W.MOV\",\"fileType\":2,\"width\":2160,\"height\":3840,\"source\":0,\"tags\":[{\"id\":\"569577514586891738\",\"name\":\"2\",\"color\":2}],\"userName\":\"Tom\"}]}", func() any { return new(SearchResponse) }},
		{"CreateTagRequest", "{\"name\":\"test\",\"color\":1}", func() any { return new(CreateTagRequest) }},
		{"CreateTagResponse", "{\"id\":\"570461738663674391\"}", func() any { return new(CreateTagResponse) }},
		{"DeleteTagRequest", "{\"ids\":[\"570461738663674391\"]}", func() any { return new(DeleteTagRequest) }},
		{"DeleteTagResponse", "{\"totalAmount\":2,\"successAmount\":1,\"failAmount\":1,\"failDetails\":[{\"code\":43022,\"id\":\"528953724308030464\",\"msg\":\"tag not found\"}]}", func() any { return new(DeleteTagResponse) }},
		{"SearchTagsRequest", "{\"page\":1,\"pageSize\":50,\"name\":\"\"}", func() any { return new(SearchTagsRequest) }},
		{"SearchTagsResponse", "{\"total\":2,\"page\":1,\"pageSize\":50,\"list\":[{\"id\":\"569577514586891738\",\"name\":\"2\",\"color\":4}]}", func() any { return new(SearchTagsResponse) }},
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
