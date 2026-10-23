package browser

import (
	"context"
)

type ListRequest struct {
	// Page number, minimum is 1, defaults to 1 if left blank.
	Page int `json:"page,omitempty"`
	// Number of records per page, minimum is 1, maximum is 100, defaults to 10 if left blank.
	PageSize int `json:"pageSize,omitempty"`
	// Profile IDs, maximum is 100. Pagination parameters are used only if left blank.
	IDs []string `json:"ids,omitempty"`
	// Profile name.
	SerialName string `json:"serialName,omitempty"`
	Remark     string `json:"remark,omitempty"`
	GroupName  string `json:"groupName,omitempty"`
	// Array of tag names.
	Tags []string `json:"tags,omitempty"`
}

type ListResponseItemsItemGroup struct {
	// Profile group id.
	ID string `json:"id"`
	// Profile group name.
	Name string `json:"name"`
	// Profile group remark.
	Remark string `json:"remark"`
}

type ListResponseItemsItemTagsItem struct {
	// Tag name.
	Name string `json:"name"`
}

type ListResponseItemsItemProxy struct {
	// Proxy type.
	Type string `json:"type"`
	// Proxy host.
	Server string `json:"server"`
	// Proxy port.
	Port int `json:"port"`
	// Proxy username.
	Username string `json:"username"`
	// Proxy password.
	Password string `json:"password"`
}

type ListResponseItemsItemAccountInfo struct {
	// Platform address.
	URL string `json:"url"`
	// Platform account.
	UserName string `json:"userName"`
	// Platform password.
	PassWord string `json:"passWord"`
	// 2FA key.
	TotpSecret string `json:"totpSecret"`
	// Page to open after startup. 1 - Restore last access, 2 - Open the specified URL, 3 -
	// Open the specified URL and the platform page simultaneously, 4 - Restore last access and
	// the platform page simultaneously. This field is obsolete.
	AfterStartup int `json:"afterStartup"`
	// Restore last visit 1 Yes 2 No.
	OpenLastPage int `json:"openLastPage"`
	// Open specified URL 1 Yes 2 No.
	OpenSpecPage int `json:"openSpecPage"`
	// Open platform page.
	OpenSiteURL bool `json:"openSiteUrl"`
	// Specified URL.
	AutoOpenURLs []string `json:"autoOpenUrls"`
}

type ListResponseItemsItemSimulateInfoTimeZone struct {
	// 1: IP-based matching, 2: Custom, 3: Local timezone.
	Switcher int `json:"switcher"`
	// Custom value.
	Value string `json:"value"`
}

type ListResponseItemsItemSimulateInfoWebRTC struct {
	// 1: Privacy, 2: Replace, 3: Real, 4: Disable.
	Switcher int `json:"switcher"`
}

type ListResponseItemsItemSimulateInfoGeoLocation struct {
	// 1: Ask, 2: Disable, 3: Allow.
	Switcher int `json:"switcher"`
	// Match based on IP address.
	BaseOnIP bool `json:"baseOnIp"`
	// Longitude in degrees. The browser API documents integer coordinates.
	Longitude int `json:"longitude"`
	// Latitude in degrees. The browser API documents integer coordinates.
	Latitude int `json:"latitude"`
	// Accuracy (meters).
	Accuracy int `json:"accuracy"`
}

type ListResponseItemsItemSimulateInfoLanguage struct {
	// 1: IP-based matching, 2: Custom.
	Switcher int `json:"switcher"`
	// A custom value, multiple values separated by commas.
	Value string `json:"value"`
}

type ListResponseItemsItemSimulateInfoResolution struct {
	// 1: random, 2: custom, 3: default.
	Switcher int `json:"switcher"`
	// custom value.
	Value string `json:"value"`
}

type ListResponseItemsItemSimulateInfoFont struct {
	// 1: Default, 2: Custom.
	Switcher int `json:"switcher"`
}

type ListResponseItemsItemSimulateInfoCanvas struct {
	// 1: Noise, 2: Real.
	Switcher int `json:"switcher"`
}

type ListResponseItemsItemSimulateInfoWebGLImage struct {
	// 1: Noise, 2: Real.
	Switcher int `json:"switcher"`
}

type ListResponseItemsItemSimulateInfoWebGLMetadata struct {
	// 2: Disabled, 3: Custom.
	Switcher int `json:"switcher"`
	// WebGL Provider.
	Provider string `json:"provider"`
	// WebGL Rendering.
	Render string `json:"render"`
}

type ListResponseItemsItemSimulateInfoHardware struct {
	// 1: Default, 2: Enabled, 3: Disabled.
	Switcher int `json:"switcher"`
}

type ListResponseItemsItemSimulateInfoAudioContext struct {
	// 1: Noise, 2: Disabled.
	Switcher int `json:"switcher"`
}

type ListResponseItemsItemSimulateInfoMediaDevice struct {
	// 1: Noise, 2: Disabled.
	Switcher int `json:"switcher"`
}

type ListResponseItemsItemSimulateInfoClientRects struct {
	// 1: Noise, 2: Disabled.
	Switcher int `json:"switcher"`
}

type ListResponseItemsItemSimulateInfoSpeechVoise struct {
	// 1: Noise, 2: Disabled.
	Switcher int `json:"switcher"`
}

type ListResponseItemsItemSimulateInfoBluetooth struct {
	// 1: Private, 2: True.
	Switcher int `json:"switcher"`
}

type ListResponseItemsItemSimulateInfoBattery struct {
	// 1: Private, 2: True.
	Switcher int `json:"switcher"`
}

type ListResponseItemsItemSimulateInfoPortScanProtection struct {
	// 1: Enable, 2: Disable.
	Switcher int `json:"switcher"`
	// Comma-separated list of ports allowed to be scanned.
	Value string `json:"value"`
}

type ListResponseItemsItemSimulateInfo struct {
	// Operating system, 1: Win, 2: Mac.
	OS int `json:"os"`
	// Browser type, 1: Chrome.
	Vendor int `json:"vendor"`
	// Fingerprint algorithm ID.
	MixtureKey string `json:"mixtureKey"`
	// User agent.
	Ua string `json:"ua"`
	// Browser version, 0 represents all.
	UaVersion     string                                         `json:"uaVersion"`
	TimeZone      ListResponseItemsItemSimulateInfoTimeZone      `json:"timeZone"`
	WebRTC        ListResponseItemsItemSimulateInfoWebRTC        `json:"webRtc"`
	GeoLocation   ListResponseItemsItemSimulateInfoGeoLocation   `json:"geoLocation"`
	Language      ListResponseItemsItemSimulateInfoLanguage      `json:"language"`
	Resolution    ListResponseItemsItemSimulateInfoResolution    `json:"resolution"`
	Font          ListResponseItemsItemSimulateInfoFont          `json:"font"`
	Canvas        ListResponseItemsItemSimulateInfoCanvas        `json:"canvas"`
	WebGLImage    ListResponseItemsItemSimulateInfoWebGLImage    `json:"webglImage"`
	WebGLMetadata ListResponseItemsItemSimulateInfoWebGLMetadata `json:"webglMetadata"`
	// Hardware acceleration.
	Hardware     ListResponseItemsItemSimulateInfoHardware     `json:"hardware"`
	AudioContext ListResponseItemsItemSimulateInfoAudioContext `json:"audioContext"`
	MediaDevice  ListResponseItemsItemSimulateInfoMediaDevice  `json:"mediaDevice"`
	ClientRects  ListResponseItemsItemSimulateInfoClientRects  `json:"clientRects"`
	// SpeechVoices.
	SpeechVoise         ListResponseItemsItemSimulateInfoSpeechVoise `json:"speechVoise"`
	HardwareConcurrency int                                          `json:"hardwareConcurrency"`
	// Legacy misspelling of MemoryDevice. Use MemoryDevice for new requests; this field
	// remains for compatibility with older payloads.
	MemeryDevice int `json:"memeryDevice"`
	// Device memory.
	MemoryDevice int `json:"memoryDevice,omitempty"`
	// Do Not Track 0: Default, 1: Enabled, 2: Disabled.
	DoNotTrack         int                                                 `json:"doNotTrack"`
	Bluetooth          ListResponseItemsItemSimulateInfoBluetooth          `json:"bluetooth"`
	Battery            ListResponseItemsItemSimulateInfoBattery            `json:"battery"`
	PortScanProtection ListResponseItemsItemSimulateInfoPortScanProtection `json:"portScanProtection"`
}

type ListResponseItemsItem struct {
	// Browser profile ID.
	ID string `json:"id"`
	// name.
	SerialName string `json:"serialName"`
	// Profile number.
	SerialNo string `json:"serialNo"`
	// Profile group.
	Group ListResponseItemsItemGroup `json:"group"`
	// Browser extension category ID.
	ExtGroup string `json:"extGroup"`
	// Browser profile remark.
	Remark string `json:"remark"`
	// Profile tags.
	Tags []ListResponseItemsItemTagsItem `json:"tags"`
	// Proxy information.
	Proxy ListResponseItemsItemProxy `json:"proxy"`
	// Account information.
	AccountInfo ListResponseItemsItemAccountInfo `json:"accountInfo"`
	// Simulation information.
	SimulateInfo ListResponseItemsItemSimulateInfo `json:"simulateInfo"`
}

type ListResponse struct {
	// Total number of data.
	Total int `json:"total"`
	// Current page number.
	Page int `json:"page"`
	// Number of data items per page.
	PageSize int `json:"pageSize"`
	// Data list.
	Items []ListResponseItemsItem `json:"items"`
}

// List: https://open.geelark.com/api/query-browsers
func (c *Client) List(ctx context.Context, req ListRequest) (*ListResponse, error) {
	var resp ListResponse
	if err := c.core.Do(ctx, "/api/v1/browser/list", req, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
