package app

import (
	"encoding/json"
	"fmt"
	"strconv"
)

// UnmarshalJSON accepts versionCode as a number (the documented example) or
// a decimal string (the documented response schema), keeping the public int field.
func (v *ListAvailableResponseItemsItemAppVersionInfoListItem) UnmarshalJSON(data []byte) error {
	type wire ListAvailableResponseItemsItemAppVersionInfoListItem
	value := struct {
		*wire
		VersionCode json.Number `json:"versionCode"`
	}{wire: (*wire)(v)}
	if err := json.Unmarshal(data, &value); err != nil {
		return err
	}
	if value.VersionCode != "" {
		code, err := strconv.ParseInt(value.VersionCode.String(), 10, strconv.IntSize)
		if err != nil {
			return fmt.Errorf("decode app versionCode: %w", err)
		}
		v.VersionCode = int(code)
	}
	return nil
}
