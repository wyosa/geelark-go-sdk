package common

import (
	"encoding/json"
	"fmt"
	"strconv"
)

// Ptr returns a pointer to v, allowing an optional JSON field to carry zero.
func Ptr[T any](v T) *T { return &v }

type PaginationRequest struct {
	Page     int `json:"page,omitempty"`
	PageSize int `json:"pageSize,omitempty"`
}

type Pagination struct {
	Page     int `json:"page"`
	PageSize int `json:"pageSize"`
	Total    int `json:"total"`
}

type EmptyResponse struct{}

// Text accepts an API identifier supplied as a JSON string or number.
type Text string

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

func BoolToInt(v bool) int {
	if v {
		return 1
	}
	return 0
}
