package webhook_test

import (
	"encoding/json"
	"testing"

	"github.com/wyosa/geelark-go-sdk/webhook"
)

func TestBrowserStartEvent(t *testing.T) {
	for _, status := range []string{"success", "failed"} {
		raw := `{"event":"browser_start","timestamp":1776147008407,"data":{"id":"612342716134614477","status":"` + status + `","debugPort":11019,"ipCheckPass":false}}`
		var event webhook.BrowserStartEvent
		if err := json.Unmarshal([]byte(raw), &event); err != nil {
			t.Fatal(err)
		}
		if event.Event != "browser_start" || event.Timestamp != 1776147008407 || event.Data.ID != "612342716134614477" || event.Data.Status != status || event.Data.DebugPort != 11019 || event.Data.IPCheckPass {
			t.Fatalf("callback fields lost: %+v", event)
		}
		encoded, err := json.Marshal(event)
		if err != nil || string(encoded) != raw {
			t.Fatalf("callback roundtrip: %s (%v)", encoded, err)
		}
	}
}
