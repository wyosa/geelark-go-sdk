package jsonutil

import (
	"encoding/json"
	"testing"
)

func TestInt(t *testing.T) {
	for _, tc := range []struct {
		raw  string
		want Int
		ok   bool
	}{
		{`48001`, 48001, true},
		{`"48001"`, 48001, true},
		{`-1`, -1, true},
		{`null`, 7, true},
		{`""`, 7, true},
		{`1.5`, 0, false},
		{`"1.5"`, 0, false},
		{`"abc"`, 0, false},
		{`true`, 0, false},
		{`{}`, 0, false},
		{`[]`, 0, false},
		{`99999999999999999999`, 0, false},
	} {
		value := Int(7)
		err := json.Unmarshal([]byte(tc.raw), &value)
		if (err == nil) != tc.ok || tc.ok && value != tc.want {
			t.Errorf("%s: got %d, %v", tc.raw, value, err)
		}
	}
}
