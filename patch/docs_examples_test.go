package patch

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
		{"InstallRequest", "{\"apkURL\":\"https://material.geelark.com/user-upload/495622401615203830/apk/uCeztwvXpswsHPH5WFht.apk\",\"id\":\"493188646250219877\"}", func() any { return new(InstallRequest) }},
		{"InstallResponse", "{\"taskId\":\"1830906144634757120\"}", func() any { return new(InstallResponse) }},
		{"InstallStatusRequest", "{\"type\":1,\"patchId\":\"sdktest\",\"id\":\"532930887177274368\"}", func() any { return new(InstallStatusRequest) }},
		{"InstallStatusResponse", "{\"status\":1}", func() any { return new(InstallStatusResponse) }},
		{"UninstallRequest", "{\"patchId\":\"sdktest\",\"id\":\"532930887177274368\"}", func() any { return new(UninstallRequest) }},
		{"UninstallResponse", "{\"status\":1}", func() any { return new(UninstallResponse) }},
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
