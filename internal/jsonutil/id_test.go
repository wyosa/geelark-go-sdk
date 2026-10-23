package jsonutil

import (
	"encoding/json"
	"testing"
)

func FuzzStringID(f *testing.F) {
	for _, seed := range []string{`"id"`, `9007199254740993`, `-1`, `null`, `1.25`, `true`, `{}`, `[]`, `"\u0000"`} {
		f.Add(seed)
	}
	f.Fuzz(func(t *testing.T, raw string) {
		if len(raw) > 65536 {
			t.Skip()
		}
		var id StringID
		if json.Unmarshal([]byte(raw), &id) != nil {
			return
		}
		encoded, err := json.Marshal(id)
		if err != nil {
			t.Fatal(err)
		}
		var text string
		if err := json.Unmarshal(encoded, &text); err != nil || text != string(id) {
			t.Fatalf("identifier cannot be encoded as a string: %s (%v)", encoded, err)
		}
	})
}
