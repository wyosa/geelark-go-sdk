package cloudphone

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
		TagIDs *[]string `json:"tagIDs,omitempty"`
	}{wire(r), tags})
}

// MarshalJSON distinguishes an omitted blacklist from an empty blacklist.
func (r SetNetworkSettingsRequest) MarshalJSON() ([]byte, error) {
	type wire SetNetworkSettingsRequest
	var blacklist *[]string
	if r.BlackList != nil {
		blacklist = &r.BlackList
	}
	return json.Marshal(struct {
		wire
		BlackList *[]string `json:"blackList,omitempty"`
	}{wire(r), blacklist})
}
