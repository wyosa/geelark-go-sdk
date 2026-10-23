package browser

// CreateFingerprint configures the fingerprint of a new browser profile.
type CreateFingerprint = CreateRequestSimulateConfig

// UpdateFingerprint changes fingerprint settings on an existing browser profile.
type UpdateFingerprint = UpdateRequestSimulateConfig

// Fingerprint is the fingerprint information returned for a browser profile.
type Fingerprint = ListResponseItemsItemSimulateInfo

// Profile is a browser profile returned by List.
type Profile = ListResponseItemsItem
