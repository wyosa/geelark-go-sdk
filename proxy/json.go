package proxy

import (
	"encoding/json"

	"github.com/wyosa/geelark-go-sdk/common"
)

// UnmarshalJSON accepts the undocumented proxy group ID as a string or a number,
// so an unexpected numeric ID cannot fail the whole List response.
func (v *ListResponseListItem) UnmarshalJSON(data []byte) error {
	type wire ListResponseListItem
	value := struct {
		*wire
		ProxyGroupID common.Text `json:"proxyGroupId"`
	}{wire: (*wire)(v), ProxyGroupID: common.Text(v.ProxyGroupID)}
	if err := json.Unmarshal(data, &value); err != nil {
		return err
	}
	v.ProxyGroupID = string(value.ProxyGroupID)
	return nil
}
