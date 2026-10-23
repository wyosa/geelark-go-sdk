package webhook

import (
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"net/http"
	"os"
	"strings"
	"testing"
)

func auditHeaders() http.Header {
	h := make(http.Header)
	h.Set("appId", "audit-app")
	h.Set("traceId", "12345678-1234-4234-8234-123456789abc")
	h.Set("nonce", "123456")
	h.Set("ts", "1")
	sum := sha256.Sum256([]byte("audit-app" + h.Get("traceId") + h.Get("ts") + h.Get("nonce") + "audit-key"))
	h.Set("sign", strings.ToUpper(hex.EncodeToString(sum[:])))
	return h
}
func TestAuditKeyMutations(t *testing.T) {
	h := auditHeaders()
	if !VerifyKeySignature(h, "audit-app", "audit-key") {
		t.Fatal("correct signature rejected")
	}
	for _, key := range []string{"appId", "traceId", "nonce", "ts", "sign"} {
		t.Run(key, func(t *testing.T) {
			h := auditHeaders()
			h.Set(key, h.Get(key)+"x")
			if VerifyKeySignature(h, "audit-app", "audit-key") {
				t.Error("tampered signature accepted")
			}
		})
	}
	if VerifyKeySignature(h, "audit-app", "wrong-key") || VerifyKeySignature(h, "wrong-app", "audit-key") {
		t.Fatal("wrong secret accepted")
	}
	if VerifyKeySignature(h, "audit-app", "audit-key") {
		t.Log("OBSERVED_PROTOCOL_LIMIT: timestamp freshness and replay are not checked")
	}
}
func TestAuditAllCallbackExamples(t *testing.T) {
	raw, err := os.ReadFile("../testdata/official/callbacks.json")
	if err != nil {
		t.Fatal(err)
	}
	var examples []json.RawMessage
	if err := json.Unmarshal(raw, &examples); err != nil {
		t.Fatal(err)
	}
	for _, raw := range examples {
		var event Event
		if err := json.Unmarshal(raw, &event); err != nil {
			t.Errorf("callback %s: %v", raw, err)
		}
		if event.Type == 0 {
			t.Errorf("event type missing: %s", raw)
		}
		if event.Type == 16 && event.TaskID == "" {
			t.Error("capitalized TaskId lost")
		}
	}
}
func FuzzAuditVerification(f *testing.F) {
	f.Add("Bearer audit-token", "12345678-1234-4234-8234-123456789abc", "123456", "1", "abc")
	f.Add("", "", "", "", "")
	f.Fuzz(func(t *testing.T, auth, trace, nonce, ts, sign string) {
		if len(auth)+len(trace)+len(nonce)+len(ts)+len(sign) > 65536 {
			t.Skip()
		}
		h := make(http.Header)
		h.Set("Authorization", auth)
		h.Set("appId", "audit-app")
		h.Set("traceId", trace)
		h.Set("nonce", nonce)
		h.Set("ts", ts)
		h.Set("sign", sign)
		_ = VerifyToken(h, "audit-token")
		_ = VerifyKeySignature(h, "audit-app", "audit-key")
	})
}
