package oem

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
		{"CustomizationRequest", "{\"logo\":\"https://material.geelark.cn/user-upload/banner0903_cn.jpg\",\"title\":\"GeeLark\",\"hideHeader\":false,\"mirrorUrl\":\"https://www.abcd.com/mirror/url\",\"toolBarSettings\":[{\"toolBar\":\"networkQuality\",\"visible\":false},{\"toolBar\":\"rotate\",\"visible\":false},{\"toolBar\":\"screenshot\",\"visible\":false},{\"toolBar\":\"upload\",\"visible\":true},{\"toolBar\":\"library\",\"visible\":false},{\"toolBar\":\"volumeUp\",\"visible\":false}]}", func() any { return new(CustomizationRequest) }},
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
