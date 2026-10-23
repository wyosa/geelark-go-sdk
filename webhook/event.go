package webhook

import (
	"crypto/sha256"
	"crypto/subtle"
	"encoding/hex"
	"encoding/json"
	"net/http"
	"strings"
)

const (
	// CloudPhoneStarted identifies cloud phone startup callbacks.
	CloudPhoneStarted = 1
	// PatchInstalled identifies patch installation callbacks.
	PatchInstalled = 2
	// FileUploaded identifies cloud phone file upload callbacks.
	FileUploaded = 4
	// ScreenshotReady identifies cloud phone screenshot completion callbacks.
	ScreenshotReady = 5
	// TaskCompleted identifies cloud phone task completion callbacks.
	TaskCompleted = 6
	// PaymentCompleted identifies payment completion callbacks.
	PaymentCompleted = 7
	// CloudPhoneStopped identifies cloud phone shutdown callbacks.
	CloudPhoneStopped = 8
	// CloudPhoneRenamed identifies cloud phone name changes callbacks.
	CloudPhoneRenamed = 9
	// CloudPhoneDeleted identifies cloud phone deletion callbacks.
	CloudPhoneDeleted = 10
	// CloudPhoneTagsChanged identifies cloud phone tag changes callbacks.
	CloudPhoneTagsChanged = 11
	// TaskCreated identifies cloud phone task creation callbacks.
	TaskCreated = 12
	// TaskCancelled identifies cloud phone task cancellation callbacks.
	TaskCancelled = 13
	// ContactsImported identifies batch contacts import callbacks.
	ContactsImported = 14
	// AppInstalled identifies application installation callbacks.
	AppInstalled = 15
	// UIDumpReady identifies UI dump completion callbacks.
	UIDumpReady = 16
)

// Event contains the common fields across GeeLark cloud phone callback types.
// Result varies by event and is kept as JSON so bool and numeric results decode.
// Use BrowserStartEvent for callbacks from the local browser API.
type Event struct {
	// Cloud callback type, identified by the constants in this package.
	Type int `json:"type"`
	// Cloud phone ID, when the event relates to one phone.
	ID string `json:"id,omitempty"`
	// Task ID for completion events. Decoding also accepts the documented TaskId spelling.
	TaskID string `json:"taskId,omitempty"`
	// Task IDs for task creation and cancellation events.
	TaskIDs []string `json:"taskIds,omitempty"`
	// Execution result: a JSON boolean for most events, or 0/1 for app installation.
	Result json.RawMessage `json:"result,omitempty"`
	// Contacts import result for type 14: 0 means failed, 1 means successful.
	Status *int `json:"status,omitempty"`
	// User who triggered the cloud phone event.
	Username string `json:"username,omitempty"`
	// User IP address reported with cloud phone lifecycle events.
	IP string `json:"ip,omitempty"`
	// Cloud event time as a Unix timestamp in seconds.
	EventTime int64 `json:"eventTime,omitempty"`
	// Cloud phone usage duration in minutes for stop events.
	UseMin int `json:"useMin,omitempty"`
	// Screenshot or UI dump download URL.
	DownloadLink string `json:"downloadLink,omitempty"`
	// Affected phones for rename and deletion events.
	Items []EventItem `json:"items,omitempty"`
	// Cloud phone IDs whose tags changed.
	EnvIDs []string `json:"envIds,omitempty"`
	// Task triggering channel: api or client.
	Channel string `json:"channel,omitempty"`
	// Android package name for an app installation event.
	PackageName string `json:"packageName,omitempty"`
	// Installed application version ID.
	AppVersionID string `json:"appVersionId,omitempty"`
}

// EventItem describes a cloud phone affected by a rename or deletion event.
type EventItem struct {
	// Cloud phone ID, when the event relates to one phone.
	ID string `json:"id"`
	// New cloud phone name for a rename event.
	Name string `json:"name,omitempty"`
	// Whether a deleted cloud phone was moved to the trash.
	Recycle bool `json:"recycle,omitempty"`
}

// VerifyKeySignature validates callback header signatures using the team's API key.
// GeeLark's signature does not cover the request body. This function does not
// check timestamp freshness or deduplicate trace IDs; applications must apply
// their own delivery and replay policy before processing an event.
func VerifyKeySignature(h http.Header, appID, apiKey string) bool {
	if appID == "" || apiKey == "" || h.Get("appId") != appID {
		return false
	}
	traceID, ts, nonce, got := h.Get("traceId"), h.Get("ts"), h.Get("nonce"), h.Get("sign")
	if len(traceID) < 6 || nonce != traceID[:6] || ts == "" || got == "" {
		return false
	}
	sum := sha256.Sum256([]byte(appID + traceID + ts + nonce + apiKey))
	want := strings.ToUpper(hex.EncodeToString(sum[:]))
	return subtle.ConstantTimeCompare([]byte(want), []byte(strings.ToUpper(got))) == 1
}

// VerifyToken validates a callback carrying the team's Bearer token.
// It accepts both "Bearer token" and the cloud documentation's "Bearer: token".
// It does not validate the event body or prevent repeated delivery.
func VerifyToken(h http.Header, token string) bool {
	if token == "" {
		return false
	}
	auth := strings.TrimSpace(h.Get("Authorization"))
	if strings.HasPrefix(auth, "Bearer:") {
		auth = "Bearer " + strings.TrimSpace(strings.TrimPrefix(auth, "Bearer:"))
	}
	want := "Bearer " + token
	return subtle.ConstantTimeCompare([]byte(auth), []byte(want)) == 1
}
