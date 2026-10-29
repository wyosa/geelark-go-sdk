package automation

import (
	"encoding/json"

	"github.com/wyosa/geelark-go-sdk/internal/jsonutil"
)

// UnmarshalJSON accepts the integer code from the response table and the quoted
// code shown in the documented partial-failure examples.
func (v *CancelTaskResponseFailDetailsItem) UnmarshalJSON(data []byte) error {
	type wire CancelTaskResponseFailDetailsItem
	value := struct {
		*wire
		Code jsonutil.Int `json:"code"`
	}{wire: (*wire)(v), Code: jsonutil.Int(v.Code)}
	if err := json.Unmarshal(data, &value); err != nil {
		return err
	}
	v.Code = int(value.Code)
	return nil
}

// UnmarshalJSON accepts the integer code from the response table and the quoted
// code shown in the documented partial-failure examples.
func (v *RetryTaskResponseFailDetailsItem) UnmarshalJSON(data []byte) error {
	type wire RetryTaskResponseFailDetailsItem
	value := struct {
		*wire
		Code jsonutil.Int `json:"code"`
	}{wire: (*wire)(v), Code: jsonutil.Int(v.Code)}
	if err := json.Unmarshal(data, &value); err != nil {
		return err
	}
	v.Code = int(value.Code)
	return nil
}
