package webhook

// BrowserStartEvent is the callback sent to browser.StartRequest.Webhook after
// a browser finishes starting. Unlike cloud callbacks, it uses an event name
// and a nested data object. VerifyToken can check its Authorization header.
//
// See https://open.geelark.com/api/launch-browser.
type BrowserStartEvent struct {
	// Event is "browser_start".
	Event string `json:"event"`
	// Timestamp is the callback time as a Unix timestamp in milliseconds.
	Timestamp int64 `json:"timestamp"`
	// Data contains the launch result for the requested browser profile.
	Data BrowserStartEventData `json:"data"`
}

// BrowserStartEventData describes the result of a browser launch.
type BrowserStartEventData struct {
	// ID is the browser profile ID.
	ID string `json:"id"`
	// Status is "success" or "failed".
	Status string `json:"status"`
	// DebugPort is the browser debugging port returned by the desktop client.
	DebugPort int `json:"debugPort"`
	// IPCheckPass reports whether the proxy check passed.
	IPCheckPass bool `json:"ipCheckPass"`
}
