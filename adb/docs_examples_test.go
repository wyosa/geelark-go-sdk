package adb

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
		{"SetStatusRequest", "{\"ids\":[\"526209711930868736\"],\"open\":true}", func() any { return new(SetStatusRequest) }},
		{"GetInfoRequest", "{\"ids\":[\"526806961778328576\",\"524798337208026112\",\"524783756767134720\"]}", func() any { return new(GetInfoRequest) }},
		{"GetInfoResponse", "{\"items\":[{\"code\":0,\"id\":\"524783756767134720\",\"ip\":\"124.71.210.176\",\"pwd\":\"8c1da4\",\"port\":\"25219\"},{\"code\":42002,\"id\":\"524798337208026112\",\"ip\":\"\",\"pwd\":\"\",\"port\":\"\"}]}", func() any { return new(GetInfoResponse) }},
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
