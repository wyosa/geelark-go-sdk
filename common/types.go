// Package common provides optional-value and JSON helpers.
package common

import (
	"encoding/json"
	"fmt"
	"strconv"
)

// Ptr returns a pointer to v, allowing an optional JSON field to carry zero.
func Ptr[T any](v T) *T { return &v }

// PaginationRequest contains optional pagination parameters. Endpoint-specific
// defaults and limits are described by each service's request type.
type PaginationRequest struct {
	// Page is the page number; zero omits the parameter.
	Page int `json:"page,omitempty"`
	// PageSize is the requested number of records; zero omits the parameter.
	PageSize int `json:"pageSize,omitempty"`
}

// Pagination describes the page returned by an API endpoint.
type Pagination struct {
	// Page is the current page number.
	Page int `json:"page"`
	// PageSize is the number of records per page.
	PageSize int `json:"pageSize"`
	// Total is the total number of matching records.
	Total int `json:"total"`
}

// EmptyResponse represents an endpoint without documented response data.
type EmptyResponse struct{}

// Text accepts a JSON string, number, boolean, or null. Numbers keep their JSON
// spelling, booleans become "true" or "false", and null becomes an empty string.
// It marshals as a JSON string and does not validate identifier syntax.
type Text string

// UnmarshalJSON implements Text's permissive scalar conversion.
func (t *Text) UnmarshalJSON(data []byte) error {
	if string(data) == "null" {
		*t = ""
		return nil
	}
	var s string
	if err := json.Unmarshal(data, &s); err == nil {
		*t = Text(s)
		return nil
	}
	var n json.Number
	if err := json.Unmarshal(data, &n); err == nil {
		*t = Text(n.String())
		return nil
	}
	var b bool
	if err := json.Unmarshal(data, &b); err == nil {
		*t = Text(strconv.FormatBool(b))
		return nil
	}
	return fmt.Errorf("unsupported text value: %s", data)
}

// BoolToInt returns 1 for true and 0 for false.
func BoolToInt(v bool) int {
	if v {
		return 1
	}
	return 0
}
