package app

import (
	"encoding/json"

	"github.com/wyosa/geelark-go-sdk/internal/jsonutil"
)

// UnmarshalJSON accepts versionCode as a number (the documented example) or
// a decimal string (the documented response schema), keeping the public int field.
// An empty string is treated like an absent value.
func (v *ListAvailableResponseItemsItemAppVersionInfoListItem) UnmarshalJSON(data []byte) error {
	type wire ListAvailableResponseItemsItemAppVersionInfoListItem
	value := struct {
		*wire
		VersionCode jsonutil.Int `json:"versionCode"`
	}{wire: (*wire)(v), VersionCode: jsonutil.Int(v.VersionCode)}
	if err := json.Unmarshal(data, &value); err != nil {
		return err
	}
	v.VersionCode = int(value.VersionCode)
	return nil
}
