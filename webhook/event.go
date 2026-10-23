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
	CloudPhoneStarted     = 1
	PatchInstalled        = 2
	FileUploaded          = 4
	ScreenshotReady       = 5
	TaskCompleted         = 6
	PaymentCompleted      = 7
	CloudPhoneStopped     = 8
	CloudPhoneRenamed     = 9
	CloudPhoneDeleted     = 10
	CloudPhoneTagsChanged = 11
	TaskCreated           = 12
	TaskCancelled         = 13
	ContactsImported      = 14
	AppInstalled          = 15
	UIDumpReady           = 16
)

// Event contains the common fields across GeeLark callback types.
// Result varies by event and is kept as JSON so bool and numeric results decode.
type Event struct {
	Type         int             `json:"type"`
	ID           string          `json:"id,omitempty"`
	TaskID       string          `json:"taskId,omitempty"`
	TaskIDs      []string        `json:"taskIds,omitempty"`
	Result       json.RawMessage `json:"result,omitempty"`
	Status       *int            `json:"status,omitempty"`
	Username     string          `json:"username,omitempty"`
	IP           string          `json:"ip,omitempty"`
	EventTime    int64           `json:"eventTime,omitempty"`
	UseMin       int             `json:"useMin,omitempty"`
	DownloadLink string          `json:"downloadLink,omitempty"`
	Items        []EventItem     `json:"items,omitempty"`
	EnvIDs       []string        `json:"envIds,omitempty"`
	Channel      string          `json:"channel,omitempty"`
	PackageName  string          `json:"packageName,omitempty"`
	AppVersionID string          `json:"appVersionId,omitempty"`
}

type EventItem struct {
	ID      string `json:"id"`
	Name    string `json:"name,omitempty"`
	Recycle bool   `json:"recycle,omitempty"`
}

// VerifyKeySignature validates a callback signed with the team's API key.
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
