package group

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
		{"CreateRequest", "{\"list\":[{\"name\":\"group\"},{\"name\":\"groupRemark\",\"remark\":\"remark\"},{\"name\":\"groupInvalid\"}]}", func() any { return new(CreateRequest) }},
		{"CreateResponse", "{\"totalAmount\":3,\"successAmount\":2,\"failAmount\":1,\"successDetails\":[{\"id\":\"528994851237135360\",\"name\":\"group\"},{\"id\":\"528994851237200896\",\"name\":\"groupRemark\",\"remark\":\"remark\"}],\"failDetails\":[{\"code\":43031,\"name\":\"groupInvalid\",\"msg\":\"group existed\"}]}", func() any { return new(CreateResponse) }},
		{"DeleteRequest", "{\"ids\":[\"528994851237200896\",\"528994851237135360\",\"528984285433037824\"]}", func() any { return new(DeleteRequest) }},
		{"DeleteResponse", "{\"totalAmount\":3,\"successAmount\":2,\"failAmount\":1,\"failDetails\":[{\"code\":43032,\"id\":\"528984285433037824\",\"msg\":\"group not found\"}]}", func() any { return new(DeleteResponse) }},
		{"ListRequest", "{\"page\":1,\"pageSize\":5}", func() any { return new(ListRequest) }},
		{"ListResponse", "{\"total\":2,\"page\":1,\"pageSize\":5,\"list\":[{\"id\":\"528995439832269824\",\"name\":\"groupRemark\",\"remark\":\"remark\"},{\"id\":\"528985080069096448\",\"name\":\"testRemark\",\"remark\":\"test\"}]}", func() any { return new(ListResponse) }},
		{"UpdateRequest", "{\"list\":[{\"id\":\"528995439832269824\",\"name\":\"newGroupRemark\",\"remark\":\"update remark\"}]}", func() any { return new(UpdateRequest) }},
		{"UpdateResponse", "{\"totalAmount\":1,\"successAmount\":1,\"failAmount\":0}", func() any { return new(UpdateResponse) }},
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
