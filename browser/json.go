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

// The parameter tables spell device memory memoryDevice, while every JSON
// example uses memeryDevice. Requests send the value under both names, and
// responses fill both fields, so either spelling works with the desktop client.

// MarshalJSON sends device memory as both memoryDevice and memeryDevice.
func (c CreateRequestSimulateConfig) MarshalJSON() ([]byte, error) {
	type wire CreateRequestSimulateConfig
	w := wire(c)
	w.MemoryDevice, w.MemeryDevice = deviceMemory(w.MemoryDevice, w.MemeryDevice)
	return json.Marshal(w)
}

// MarshalJSON sends device memory as both memoryDevice and memeryDevice.
func (c UpdateRequestSimulateConfig) MarshalJSON() ([]byte, error) {
	type wire UpdateRequestSimulateConfig
	w := wire(c)
	w.MemoryDevice, w.MemeryDevice = deviceMemory(w.MemoryDevice, w.MemeryDevice)
	return json.Marshal(w)
}

// UnmarshalJSON copies device memory from whichever spelling is present.
func (v *ListResponseItemsItemSimulateInfo) UnmarshalJSON(data []byte) error {
	type wire ListResponseItemsItemSimulateInfo
	if err := json.Unmarshal(data, (*wire)(v)); err != nil {
		return err
	}
	if v.MemoryDevice == 0 {
		v.MemoryDevice = v.MemeryDevice
	}
	if v.MemeryDevice == 0 {
		v.MemeryDevice = v.MemoryDevice
	}
	return nil
}

// deviceMemory fills a missing spelling from the other one. Explicitly
// different values are sent unchanged.
func deviceMemory(memory, memery *int) (*int, *int) {
	if memory == nil {
		memory = memery
	}
	if memery == nil {
		memery = memory
	}
	return memory, memery
}
