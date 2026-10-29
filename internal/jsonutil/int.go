package jsonutil

import (
	"bytes"
	"encoding/json"
	"errors"
	"strconv"
)

// Int accepts a JSON integer or a string holding a decimal integer, as in
// documented examples that quote numeric codes. Null and the empty string leave
// an existing value unchanged.
type Int int

// UnmarshalJSON rejects fractions, booleans, arrays, objects, and values that
// overflow int.
func (n *Int) UnmarshalJSON(data []byte) error {
	data = bytes.TrimSpace(data)
	if bytes.Equal(data, []byte("null")) {
		return nil
	}
	if len(data) > 0 && data[0] == '"' {
		var text string
		if err := json.Unmarshal(data, &text); err != nil {
			return err
		}
		if text == "" {
			return nil
		}
		data = []byte(text)
	}
	value, err := strconv.ParseInt(string(data), 10, strconv.IntSize)
	if err != nil {
		return errors.New("decode integer: expected a number or a decimal string")
	}
	*n = Int(value)
	return nil
}
