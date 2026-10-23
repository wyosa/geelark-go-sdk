package webhook

import (
	"encoding/json"
	"net/http"
	"testing"
)

func TestVerifyKeySignature(t *testing.T) {
	h := http.Header{}
	h.Set("appId", "eH6gQR4oHr3FsZpI36La01IW")
	h.Set("traceId", "db6094ab-3797-4186-84d5-b0b58eebad56")
	h.Set("ts", "1716972892166")
	h.Set("nonce", "db6094")
	h.Set("sign", "6280C080AF7C3CCE168F15C913E3444A00A618CB0E16038EED9811D6E3366BDD")
	if !VerifyKeySignature(h, "eH6gQR4oHr3FsZpI36La01IW", "YjmFIUuq0oJgSDJ42fxLEb6R1qjjqf") {
		t.Fatal("valid signature rejected")
	}
	h.Set("nonce", "wrong")
	if VerifyKeySignature(h, "eH6gQR4oHr3FsZpI36La01IW", "YjmFIUuq0oJgSDJ42fxLEb6R1qjjqf") {
		t.Fatal("invalid signature accepted")
	}
}

func TestEventHandlesBoolAndNumberResults(t *testing.T) {
	for _, raw := range []string{
		`{"type":16,"id":"phone","TaskId":"task","result":true,"downloadLink":"https://example.com/ui.xml"}`,
		`{"type":15,"id":"phone","result":1,"packageName":"com.example.app"}`,
	} {
		var event Event
		if err := json.Unmarshal([]byte(raw), &event); err != nil {
			t.Fatal(err)
		}
		if event.ID != "phone" || !json.Valid(event.Result) {
			t.Fatalf("event = %+v", event)
		}
	}
}

func TestVerifyToken(t *testing.T) {
	h := http.Header{}
	h.Set("Authorization", "Bearer token")
	if !VerifyToken(h, "token") || VerifyToken(h, "other") {
		t.Fatal("Bearer verification failed")
	}
	h.Set("Authorization", "Bearer: token")
	if !VerifyToken(h, "token") {
		t.Fatal("documented Bearer callback form rejected")
	}
}
