package cloudphone

import (
	"encoding/json"

	"github.com/wyosa/geelark-go-sdk/internal/jsonutil"
)

// UnmarshalJSON accepts both string IDs from examples and integer IDs from the
// response table. Numeric IDs are preserved exactly; fractional IDs are rejected.
func (v *StopResponseFailDetailsItem) UnmarshalJSON(data []byte) error {
	type wire StopResponseFailDetailsItem
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
func (v *DeleteResponseFailDetailsItem) UnmarshalJSON(data []byte) error {
	type wire DeleteResponseFailDetailsItem
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
func (v *UpdateResponseFailDetailsItem) UnmarshalJSON(data []byte) error {
	type wire UpdateResponseFailDetailsItem
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
func (v *HideAccessibilityResponseFailDetailsItem) UnmarshalJSON(data []byte) error {
	type wire HideAccessibilityResponseFailDetailsItem
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
