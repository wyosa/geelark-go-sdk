package library

import "encoding/json"

// MarshalJSON preserves an explicitly empty TagsID slice to remove all tags.
// A nil slice omits the field.
func (r SetTagRequest) MarshalJSON() ([]byte, error) {
	type wire SetTagRequest
	var tags *[]string
	if r.TagsID != nil {
		tags = &r.TagsID
	}
	return json.Marshal(struct {
		wire
		TagsID *[]string `json:"tagsId,omitempty"`
	}{wire(r), tags})
}
