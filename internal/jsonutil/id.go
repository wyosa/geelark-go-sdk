// Package jsonutil implements compatibility decoding for inconsistent API schemas.
package jsonutil

import (
	"bytes"
	"encoding/json"
	"errors"
)

// StringID accepts a JSON string or a decimal integer without losing precision.
// Null leaves an existing value unchanged, like encoding/json decoding a string.
type StringID string

// UnmarshalJSON rejects fractional numbers, booleans, arrays, and objects.
func (id *StringID) UnmarshalJSON(data []byte) error {
	data = bytes.TrimSpace(data)
	if bytes.Equal(data, []byte("null")) {
		return nil
	}
	if len(data) > 0 && data[0] == '"' {
		var value string
		if err := json.Unmarshal(data, &value); err != nil {
			return err
		}
		*id = StringID(value)
		return nil
	}
	if !json.Valid(data) || len(data) == 0 {
		return errors.New("decode identifier: expected a string or decimal integer")
	}
	for i, b := range data {
		if b == '-' && i == 0 {
			continue
		}
		if b < '0' || b > '9' {
			return errors.New("decode identifier: expected a string or decimal integer")
		}
	}
	*id = StringID(data)
	return nil
}
