package webhook_test

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/wyosa/geelark-go-sdk/webhook"
)

func ExampleEvent() {
	// Header verification is one part of the receiving application's policy.
	// Apply your delivery deduplication policy before processing an event.
	headers := make(http.Header)
	headers.Set("Authorization", "Bearer example-token")
	if !webhook.VerifyToken(headers, "example-token") {
		panic("unauthorized callback")
	}
	var event webhook.Event
	if err := json.Unmarshal([]byte(`{"type":1,"id":"phone-id","result":true,"eventTime":1766567054}`), &event); err != nil {
		panic(err)
	}
	if event.Type == webhook.CloudPhoneStarted {
		var started bool
		if err := json.Unmarshal(event.Result, &started); err != nil {
			panic(err)
		}
		fmt.Println(event.ID, started, event.EventTime)
	}
	// Output: phone-id true 1766567054
}

func ExampleBrowserStartEvent() {
	headers := make(http.Header)
	headers.Set("Authorization", "Bearer example-token")
	if !webhook.VerifyToken(headers, "example-token") {
		panic("unauthorized callback")
	}
	// BrowserStartEvent uses a name, a millisecond timestamp, and nested data.
	// It is distinct from the cloud phone Event envelope.
	var event webhook.BrowserStartEvent
	raw := `{"event":"browser_start","timestamp":1776147008407,"data":{"id":"browser-id","status":"success","debugPort":11019,"ipCheckPass":true}}`
	if err := json.Unmarshal([]byte(raw), &event); err != nil {
		panic(err)
	}
	fmt.Println(event.Event, event.Data.ID, event.Data.Status, event.Data.DebugPort, event.Timestamp)
	// Output: browser_start browser-id success 11019 1776147008407
}
