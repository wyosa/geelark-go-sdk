package livestream

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
		{"StartRequest", "{\"id\":\"628708996374619612\",\"fileUrl\":\"https://material.geelark.com/open-upload/497521349346987872/20260828/HCgv7sfC.mp4\",\"isLoop\":1}", func() any { return new(StartRequest) }},
		{"StartResponse", "{\"taskId\":\"2093234110406909952\"}", func() any { return new(StartResponse) }},
		{"ResultRequest", "{\"id\":\"2093234110406909952\"}", func() any { return new(ResultRequest) }},
		{"ResultResponse", "{\"status\":1}", func() any { return new(ResultResponse) }},
		{"StopRequest", "{\"id\":\"628708996374619612\"}", func() any { return new(StopRequest) }},
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
