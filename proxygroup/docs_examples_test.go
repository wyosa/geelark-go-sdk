package proxygroup

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
		{"CreateRequest", "{\"name\":\"Business Group A\"}", func() any { return new(CreateRequest) }},
		{"CreateResponse", "{\"id\":\"123456789012345678\"}", func() any { return new(CreateResponse) }},
		{"ListRequest", "{\"page\":1,\"pageSize\":20,\"name\":\"Business Group\"}", func() any { return new(ListRequest) }},
		{"ListResponse", "{\"total\":2,\"page\":1,\"pageSize\":20,\"list\":[{\"id\":\"0\",\"name\":\"Ungrouped Proxies\"},{\"id\":\"123456789012345678\",\"name\":\"Business Group A\"}]}", func() any { return new(ListResponse) }},
		{"UpdateRequest", "{\"id\":\"123456789012345678\",\"name\":\"New Group Name\"}", func() any { return new(UpdateRequest) }},
		{"DeleteRequest", "{\"ids\":[\"123456789012345678\",\"223456789012345678\"]}", func() any { return new(DeleteRequest) }},
		{"DeleteResponse", "{\"totalAmount\":2,\"successAmount\":2,\"failAmount\":0,\"failDetails\":[]}", func() any { return new(DeleteResponse) }},
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
