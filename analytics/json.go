package analytics

import "encoding/json"

// MarshalJSON distinguishes an omitted tag list from an explicit empty array.
func (r UpdateAccountRequest) MarshalJSON() ([]byte, error) {
	type wire UpdateAccountRequest
	var tags *[]string
	if r.TagIDs != nil {
		tags = &r.TagIDs
	}
	return json.Marshal(struct {
		wire
		TagIDs *[]string `json:"tagIds,omitempty"`
	}{wire(r), tags})
}
