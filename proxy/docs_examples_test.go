package proxy

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
		{"AddResponse", "{\"totalAmount\":2,\"successAmount\":1,\"failAmount\":1,\"failDetails\":[{\"index\":0,\"code\":45007,\"msg\":\"proxy already exists\"}],\"successDetails\":[{\"index\":1,\"id\":\"493188072704313353\"}]}", func() any { return new(AddResponse) }},
		{"UpdateResponse", "{\"totalAmount\":2,\"successAmount\":1,\"failAmount\":1,\"failDetails\":[{\"id\":\"493188072704313353\",\"code\":40005,\"msg\":\"proxy not found\"}]}", func() any { return new(UpdateResponse) }},
		{"DeleteResponse", "{\"totalAmount\":2,\"successAmount\":1,\"failAmount\":1,\"failDetails\":[{\"id\":\"493188072704313353\",\"code\":40005,\"msg\":\"proxy not found\"}]}", func() any { return new(DeleteResponse) }},
		{"ListResponse", "{\"total\":1,\"page\":1,\"pageSize\":1,\"list\":[{\"id\":\"493188072704313353\",\"serialNo\":1,\"scheme\":\"socks5\",\"server\":\"192.3.8.1\",\"port\":8000,\"username\":\"admin\",\"password\":\"admin\"}]}", func() any { return new(ListResponse) }},
		{"CheckRequest", "{\"proxyQueryChannel\":\"IP2Location\",\"proxyType\":\"socks5\",\"server\":\"185.162.130.86\",\"port\":10000,\"username\":\"username\",\"password\":\"pass\"}", func() any { return new(CheckRequest) }},
		{"CheckResponse", "{\"detectStatus\":true,\"message\":\"\",\"outboundIP\":\"223.135.25.196\",\"countryCode\":\"JP\",\"countryName\":\"Japan\",\"subdivision\":\"Tokyo\",\"city\":\"Tokyo\",\"timezone\":\"Asia/Tokyo\",\"isp\":\"Sony Network Communications Inc.\"}", func() any { return new(CheckResponse) }},
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
