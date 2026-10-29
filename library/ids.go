package library

import (
	"encoding/json"

	"github.com/wyosa/geelark-go-sdk/common"
	"github.com/wyosa/geelark-go-sdk/internal/jsonutil"
)

// UnmarshalJSON accepts both string IDs from examples and integer IDs from the
// response table. Numeric IDs are preserved exactly; fractional IDs are rejected.
func (v *SearchResponseListItem) UnmarshalJSON(data []byte) error {
	type wire SearchResponseListItem
	value := struct {
		*wire
		ID jsonutil.StringID `json:"id"`
	}{wire: (*wire)(v), ID: jsonutil.StringID(v.ID)}
	if err := json.Unmarshal(data, &value); err != nil {
		return err
	}
	v.ID = string(value.ID)
	return nil
}

// UnmarshalJSON accepts both string IDs from examples and integer IDs from the
// response table. Numeric IDs are preserved exactly; fractional IDs are rejected.
// The tag name is also accepted as a number, as the response table describes it.
func (v *SearchResponseListItemTagsItem) UnmarshalJSON(data []byte) error {
	type wire SearchResponseListItemTagsItem
	value := struct {
		*wire
		ID   jsonutil.StringID `json:"id"`
		Name common.Text       `json:"name"`
	}{wire: (*wire)(v), ID: jsonutil.StringID(v.ID), Name: common.Text(v.Name)}
	if err := json.Unmarshal(data, &value); err != nil {
		return err
	}
	v.ID = string(value.ID)
	v.Name = string(value.Name)
	return nil
}

// UnmarshalJSON accepts both string IDs from examples and integer IDs from the
// response table. Numeric IDs are preserved exactly; fractional IDs are rejected.
func (v *SearchTagsResponseListItem) UnmarshalJSON(data []byte) error {
	type wire SearchTagsResponseListItem
	value := struct {
		*wire
		ID jsonutil.StringID `json:"id"`
	}{wire: (*wire)(v), ID: jsonutil.StringID(v.ID)}
	if err := json.Unmarshal(data, &value); err != nil {
		return err
	}
	v.ID = string(value.ID)
	return nil
}
