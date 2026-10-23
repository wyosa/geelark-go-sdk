package browser

import "encoding/json"

// MarshalJSON preserves an explicitly empty TagIDs slice to clear tags.
// A nil slice omits the field and leaves the existing tags unchanged.
func (r UpdateRequest) MarshalJSON() ([]byte, error) {
	type wire UpdateRequest
	var tags *[]string
	if r.TagIDs != nil {
		tags = &r.TagIDs
	}
	return json.Marshal(struct {
		wire
		TagIDs *[]string `json:"tagIds,omitempty"`
	}{wire(r), tags})
}
