package browser

import (
	"context"
)

type UpdateRequestSimulateConfigTimeZone struct {
	// 1: Matched by IP address, 2: Custom, 3: Local time zone. Required.
	Switcher *int `json:"switcher,omitempty"`
	// Custom value, refer to Browser Timezone
	// (https://singapore-upgrade.geelark.cn/apiResource/browser_timezone.txt); GMT-12:00
	// Etc/GMT+12. Required.
	Value *string `json:"value,omitempty"`
}

type UpdateRequestSimulateConfigLanguage struct {
	// 1: Matched by IP address, 2: Custom. Required.
	Switcher *int `json:"switcher,omitempty"`
	// Custom value, multiple values separated by commas, refer to Browser Language
	// (https://singapore-upgrade.geelark.cn/apiResource/browser_language.txt). Required.
	Value *string `json:"value,omitempty"`
}

type UpdateRequestSimulateConfigResolution struct {
	// 1: Random, 2: Custom, 3: Default. Required.
	Switcher *int `json:"switcher,omitempty"`
	// Custom value, refer to browser resolution
	// (https://singapore-upgrade.geelark.cn/apiResource/browser_resolution.txt). Required.
	Value *string `json:"value,omitempty"`
}

type UpdateRequestSimulateConfigWebRTC struct {
	// 1: Privacy, 2: Replace, 3: True, 4: Disable. Required.
	Switcher *int `json:"switcher,omitempty"`
}

type UpdateRequestSimulateConfigGeoLocation struct {
	// 1: Ask, 2: Disable, 3: Allow. Required.
	Switcher *int `json:"switcher,omitempty"`
	// Match based on IP. Required.
	BaseOnIP *bool `json:"baseOnIp,omitempty"`
	// Longitude in degrees. The parameter table says integer, but fractional degrees are
	// accepted here so locations are not rounded to about 100 km. Required.
	Longitude *float64 `json:"longitude,omitempty"`
	// Latitude in degrees. The parameter table says integer, but fractional degrees are
	// accepted here so locations are not rounded to about 100 km. Required.
	Latitude *float64 `json:"latitude,omitempty"`
	// Accuracy (meters). Required.
	Accuracy *int `json:"accuracy,omitempty"`
}

type UpdateRequestSimulateConfigCanvas struct {
	// 1: Noise, 2: True. Required.
	Switcher *int `json:"switcher,omitempty"`
}

type UpdateRequestSimulateConfigWebGLImage struct {
	// 1: Noise, 2: Real. Required.
	Switcher *int `json:"switcher,omitempty"`
}

type UpdateRequestSimulateConfigHardware struct {
	// 1: Default, 2: Enabled, 3: Disabled. Required.
	Switcher *int `json:"switcher,omitempty"`
}

type UpdateRequestSimulateConfigAudioContext struct {
	// 1: Noise, 2: Disabled. Required.
	Switcher *int `json:"switcher,omitempty"`
}

type UpdateRequestSimulateConfigMediaDevice struct {
	// 1: Noise, 2: Disabled. Required.
	Switcher *int `json:"switcher,omitempty"`
}

type UpdateRequestSimulateConfigClientRects struct {
	// 1: Noise, 2: Disabled. Required.
	Switcher *int `json:"switcher,omitempty"`
}

type UpdateRequestSimulateConfigSpeechVoise struct {
	// 1: Noise, 2: Off. Required.
	Switcher *int `json:"switcher,omitempty"`
}

type UpdateRequestSimulateConfigBluetooth struct {
	// 1: Privacy, 2: True. Required.
	Switcher *int `json:"switcher,omitempty"`
}

type UpdateRequestSimulateConfigBattery struct {
	// 1: Privacy, 2: True. Required.
	Switcher *int `json:"switcher,omitempty"`
}

type UpdateRequestSimulateConfigPortScanProtection struct {
	// 1: Enable, 2: Disable. Required.
	Switcher *int `json:"switcher,omitempty"`
	// Ports allowed to be scanned, comma separated. Required.
	Value *string `json:"value,omitempty"`
}

type UpdateRequestSimulateConfig struct {
	// Time zone. Required.
	TimeZone *UpdateRequestSimulateConfigTimeZone `json:"timeZone,omitempty"`
	// Language. Required.
	Language *UpdateRequestSimulateConfigLanguage `json:"language,omitempty"`
	// Resolution. Required.
	Resolution *UpdateRequestSimulateConfigResolution `json:"resolution,omitempty"`
	// WebRTC. Required.
	WebRTC *UpdateRequestSimulateConfigWebRTC `json:"webRtc,omitempty"`
	// Geolocation. Required.
	GeoLocation *UpdateRequestSimulateConfigGeoLocation `json:"geoLocation,omitempty"`
	// Canvas. Required.
	Canvas *UpdateRequestSimulateConfigCanvas `json:"canvas,omitempty"`
	// WebGL Image. Required.
	WebGLImage *UpdateRequestSimulateConfigWebGLImage `json:"webglImage,omitempty"`
	// Hardware Acceleration. Required.
	Hardware *UpdateRequestSimulateConfigHardware `json:"hardware,omitempty"`
	// AudioContext. Required.
	AudioContext *UpdateRequestSimulateConfigAudioContext `json:"audioContext,omitempty"`
	// Media Device. Required.
	MediaDevice *UpdateRequestSimulateConfigMediaDevice `json:"mediaDevice,omitempty"`
	// ClientRects. Required.
	ClientRects *UpdateRequestSimulateConfigClientRects `json:"clientRects,omitempty"`
	// SpeechVoices. Required.
	SpeechVoise *UpdateRequestSimulateConfigSpeechVoise `json:"speechVoise,omitempty"`
	// Hardware concurrency. Required.
	HardwareConcurrency *int `json:"hardwareConcurrency,omitempty"`
	// Device memory under the spelling used by the documented JSON examples. Set either
	// field: MarshalJSON sends the value under both memeryDevice and memoryDevice.
	MemeryDevice *int `json:"memeryDevice,omitempty"`
	// Device memory under the spelling used by the parameter table. Required.
	MemoryDevice *int `json:"memoryDevice,omitempty"`
	// Do Not Track 0: Default, 1: On, 2: Off. Required.
	DoNotTrack *int `json:"doNotTrack,omitempty"`
	// Bluetooth. Required.
	Bluetooth *UpdateRequestSimulateConfigBluetooth `json:"bluetooth,omitempty"`
	// Battery. Required.
	Battery *UpdateRequestSimulateConfigBattery `json:"battery,omitempty"`
	// Port scan protection. Required.
	PortScanProtection *UpdateRequestSimulateConfigPortScanProtection `json:"portScanProtection,omitempty"`
}

type UpdateRequest struct {
	// Profile id. Required.
	ID string `json:"id"`
	// Profile name. If left blank, no change. Maximum 100 characters.
	SerialName *string `json:"serialName,omitempty"`
	// Profile group id. If omitted, no change. If left blank, the Profile is marked as
	// ungrouped.
	GroupID *string `json:"groupId,omitempty"`
	// Extension category ID. If omitted or passed as an empty string, the team extensions will
	// be used.please refer to Get Browser Extension Category List
	// (https://open.geelark.com/api/browser-extgroup-list).
	ExtGroup *string `json:"extGroup,omitempty"`
	// Browser profile tag IDs. Nil omits the field and preserves existing tags; an empty
	// non-nil slice sends an empty array to clear them.
	TagIDs []string `json:"tagIds,omitempty"`
	// Profile remark. Maximum 1500 characters. If left blank, no change.
	Remark *string `json:"remark,omitempty"`
	// Cookies in JSON or Netscape format.
	Cookie *string `json:"cookie,omitempty"`
	// Account platform. Will not be changed if not passed. Only supports
	// https://www.tiktok.com/, https://www.facebook.com/.
	AccountPlatform *string `json:"accountPlatform,omitempty"`
	// Account username. Valid only when accountPlatform is valid.
	AccountUsername *string `json:"accountUsername,omitempty"`
	// Account password. Valid only when accountPlatform is valid.
	AccountPassword *string `json:"accountPassword,omitempty"`
	// Account 2FA key, only effective if accountPlatform is valid. If not provided, it will
	// not be changed; if empty, it will be left blank.
	AccountTOTPSecret *string `json:"accountTOTPSecret,omitempty"`
	// Restore last access, 1-Yes, 2-No.
	OpenLastPage *int `json:"openLastPage,omitempty"`
	// Open specified URL, 1-Yes, 2-No.
	OpenSpecPage *int `json:"openSpecPage,omitempty"`
	// URLs to open, separated by semicolons (;). Nil leaves the setting unchanged.
	OpenTabs *string `json:"openTabs,omitempty"`
	// Operating system. If not passed, no changes will be made. 1-win, 2-mac.
	BrowserOS *int `json:"browserOs,omitempty"`
	// UserAgent. If not passed, no changes will be made.
	BrowserUa *string `json:"browserUa,omitempty"`
	// Simulate configuration. If not passed, no changes will be made.
	SimulateConfig *UpdateRequestSimulateConfig `json:"simulateConfig,omitempty"`
	// Proxy ID. Used first. If neither proxyId nor proxyConfig is passed, no changes will be
	// made to the proxy.
	ProxyID *string `json:"proxyId,omitempty"`
	// Proxy information. If neither proxyId nor proxyConfig is passed, the proxy will not be
	// modified.
	ProxyConfig *UpdateRequestProxyConfig `json:"proxyConfig,omitempty"`
	// Startup parameters; if not passed, they will not be modified.
	BrowserStartArg *string `json:"browserStartArg,omitempty"`
}

type UpdateRequestProxyConfig struct {
	// Proxy type ID. -1 represents a direct connection. Required.
	TypeID *int `json:"typeId,omitempty"`
	// Proxy host.
	Server *string `json:"server,omitempty"`
	// Proxy port.
	Port *int `json:"port,omitempty"`
	// Proxy username.
	Username *string `json:"username,omitempty"`
	// Proxy password.
	Password *string `json:"password,omitempty"`
	// Dynamic proxy country.
	Country *string `json:"country,omitempty"`
	// Dynamic proxy region.
	Region *string `json:"region,omitempty"`
	// Dynamic proxy city.
	City *string `json:"city,omitempty"`
	// Whether to use the configured dynamic proxy.
	UseProxyCfg *bool `json:"useProxyCfg,omitempty"`
	// Dynamic proxy protocol. 1 represents socks5, 2 represents http.
	Protocol *int `json:"protocol,omitempty"`
	// GProxy Session Type: 1 Sticky 2 Rotating, Default 1.
	SessionType *int `json:"sessionType,omitempty"`
	// GProxy FilterMode: 0 MaxPoolSize 1 Quality 2 QualitySpeed 3 Speed, Default.
	FilterMode *int `json:"filterMode,omitempty"`
}

type UpdateResponse struct{}

// Update: https://open.geelark.com/api/edit-browser
func (c *Client) Update(ctx context.Context, req UpdateRequest) (*UpdateResponse, error) {
	var resp UpdateResponse
	if err := c.core.Do(ctx, "/api/v1/browser/update", req, nil); err != nil {
		return nil, err
	}
	return &resp, nil
}
