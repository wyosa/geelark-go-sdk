package common

import (
	"encoding/json"
	"testing"
)

func FuzzAuditText(f *testing.F) {
	for _, s := range []string{`"9007199254740993"`, `9007199254740993`, `-2147483649`, `0.125`, `null`, `true`, `{}`, `[]`, `"\u0000"`} {
		f.Add(s)
	}
	f.Fuzz(func(t *testing.T, s string) {
		if len(s) > 65536 {
			t.Skip()
		}
		var out Text
		if err := json.Unmarshal([]byte(s), &out); err != nil {
			return
		}
		raw, err := json.Marshal(out)
		if err != nil || !json.Valid(raw) {
			t.Fatalf("invalid roundtrip: %s %v", raw, err)
		}
	})
}
