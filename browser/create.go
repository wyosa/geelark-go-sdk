package browser

import (
	"context"
)

type CreateRequestSimulateConfigTimeZone struct {
	// 1: Matched by IP address, 2: Custom, 3: Local time zone. Required.
	Switcher *int `json:"switcher,omitempty"`
	// Custom value, refer to Browser Timezone
	// (https://singapore-upgrade.geelark.cn/apiResource/browser_timezone.txt); GMT-12:00
	// Etc/GMT+12. Required.
	Value string `json:"value,omitempty"`
}

type CreateRequestSimulateConfigLanguage struct {
	// 1: Matched by IP address, 2: Custom. Required.
	Switcher *int `json:"switcher,omitempty"`
	// Custom value, multiple values separated by commas, refer to Browser Language
	// (https://singapore-upgrade.geelark.cn/apiResource/browser_language.txt). Required.
	Value string `json:"value,omitempty"`
}

type CreateRequestSimulateConfigResolution struct {
	// 1: Random, 2: Custom, 3: Default. Required.
	Switcher *int `json:"switcher,omitempty"`
	// Custom value, refer to browser resolution
	// (https://singapore-upgrade.geelark.cn/apiResource/browser_resolution.txt). Required.
	Value string `json:"value,omitempty"`
}

type CreateRequestSimulateConfigWebRTC struct {
	// 1: Privacy, 2: Replace, 3: True, 4: Disable. Required.
	Switcher *int `json:"switcher,omitempty"`
}

type CreateRequestSimulateConfigGeoLocation struct {
	// 1: Ask, 2: Disable, 3: Allow. Required.
	Switcher *int `json:"switcher,omitempty"`
	// Match based on IP. Required.
	BaseOnIP *bool `json:"baseOnIp,omitempty"`
	// Longitude in degrees. The browser API documents integer coordinates. Required.
	Longitude *int `json:"longitude,omitempty"`
	// Latitude in degrees. The browser API documents integer coordinates. Required.
	Latitude *int `json:"latitude,omitempty"`
	// Accuracy (meters). Required.
	Accuracy *int `json:"accuracy,omitempty"`
}

type CreateRequestSimulateConfigCanvas struct {
	// 1: Noise, 2: True. Required.
	Switcher *int `json:"switcher,omitempty"`
}

type CreateRequestSimulateConfigWebGLImage struct {
	// 1: Noise, 2: Real. Required.
	Switcher *int `json:"switcher,omitempty"`
}

type CreateRequestSimulateConfigHardware struct {
	// 1: Default, 2: Enabled, 3: Disabled. Required.
	Switcher *int `json:"switcher,omitempty"`
}

type CreateRequestSimulateConfigAudioContext struct {
	// 1: Noise, 2: Disabled. Required.
	Switcher *int `json:"switcher,omitempty"`
}

type CreateRequestSimulateConfigMediaDevice struct {
	// 1: Noise, 2: Disabled. Required.
	Switcher *int `json:"switcher,omitempty"`
}

type CreateRequestSimulateConfigClientRects struct {
	// 1: Noise, 2: Disabled. Required.
	Switcher *int `json:"switcher,omitempty"`
}

type CreateRequestSimulateConfigSpeechVoise struct {
	// 1: Noise, 2: Off. Required.
	Switcher *int `json:"switcher,omitempty"`
}

type CreateRequestSimulateConfigBluetooth struct {
	// 1: Privacy, 2: True. Required.
	Switcher *int `json:"switcher,omitempty"`
}

type CreateRequestSimulateConfigBattery struct {
	// 1: Privacy, 2: True. Required.
	Switcher *int `json:"switcher,omitempty"`
}

type CreateRequestSimulateConfigPortScanProtection struct {
	// 1: Enable, 2: Disable. Required.
	Switcher *int `json:"switcher,omitempty"`
	// Ports allowed to be scanned, comma separated. Required.
	Value string `json:"value,omitempty"`
}

type CreateRequestSimulateConfig struct {
	// Time zone. Required.
	TimeZone *CreateRequestSimulateConfigTimeZone `json:"timeZone,omitempty"`
	// Language. Required.
	Language *CreateRequestSimulateConfigLanguage `json:"language,omitempty"`
	// Resolution. Required.
	Resolution *CreateRequestSimulateConfigResolution `json:"resolution,omitempty"`
	// WebRTC. Required.
	WebRTC *CreateRequestSimulateConfigWebRTC `json:"webRtc,omitempty"`
	// Geolocation. Required.
	GeoLocation *CreateRequestSimulateConfigGeoLocation `json:"geoLocation,omitempty"`
	// Canvas. Required.
	Canvas *CreateRequestSimulateConfigCanvas `json:"canvas,omitempty"`
	// WebGL Image. Required.
	WebGLImage *CreateRequestSimulateConfigWebGLImage `json:"webglImage,omitempty"`
	// Hardware Acceleration. Required.
	Hardware *CreateRequestSimulateConfigHardware `json:"hardware,omitempty"`
	// AudioContext. Required.
	AudioContext *CreateRequestSimulateConfigAudioContext `json:"audioContext,omitempty"`
	// Media Device. Required.
	MediaDevice *CreateRequestSimulateConfigMediaDevice `json:"mediaDevice,omitempty"`
	// ClientRects. Required.
	ClientRects *CreateRequestSimulateConfigClientRects `json:"clientRects,omitempty"`
	// SpeechVoices. Required.
	SpeechVoise *CreateRequestSimulateConfigSpeechVoise `json:"speechVoise,omitempty"`
	// Hardware concurrency. Required.
	HardwareConcurrency *int `json:"hardwareConcurrency,omitempty"`
	// Legacy misspelling of MemoryDevice. Use MemoryDevice for new requests; this field
	// remains for compatibility with older payloads.
	MemeryDevice *int `json:"memeryDevice,omitempty"`
	// Device memory. Required.
	MemoryDevice *int `json:"memoryDevice,omitempty"`
	// Do Not Track 0: Default, 1: On, 2: Off. Required.
	DoNotTrack *int `json:"doNotTrack,omitempty"`
	// Bluetooth. Required.
	Bluetooth *CreateRequestSimulateConfigBluetooth `json:"bluetooth,omitempty"`
	// Battery. Required.
	Battery *CreateRequestSimulateConfigBattery `json:"battery,omitempty"`
	// Port scan protection. Required.
	PortScanProtection *CreateRequestSimulateConfigPortScanProtection `json:"portScanProtection,omitempty"`
}

type CreateRequestProxyConfig struct {
	// Proxy type ID. -1 represents a direct connection. Required.
	TypeID *int `json:"typeId,omitempty"`
	// Proxy host.
	Server string `json:"server,omitempty"`
	// Proxy port.
	Port *int `json:"port,omitempty"`
	// Proxy username.
	Username string `json:"username,omitempty"`
	// Proxy password.
	Password string `json:"password,omitempty"`
	// Dynamic proxy country.
	Country string `json:"country,omitempty"`
	// Dynamic proxy region.
	Region string `json:"region,omitempty"`
	// Dynamic proxy city.
	City string `json:"city,omitempty"`
	// Whether to use the configured dynamic proxy.
	UseProxyCfg *bool `json:"useProxyCfg,omitempty"`
	// Dynamic proxy protocol. 1 represents socks5, 2 represents http.
	Protocol *int `json:"protocol,omitempty"`
	// GProxy Session Type: 1 Sticky 2 Rotating, Default 1.
	SessionType *int `json:"sessionType,omitempty"`
	// GProxy FilterMode: 0 MaxPoolSize 1 Quality 2 QualitySpeed 3 Speed, Default.
	FilterMode *int `json:"filterMode,omitempty"`
}

type CreateRequest struct {
	// Profile name, up to 100 characters. Required.
	SerialName string `json:"serialName"`
	// Profile group ID.
	GroupID string `json:"groupId,omitempty"`
	// Extension category ID. If omitted or passed as an empty string, the team extensions will
	// be used.See Get Browser Extension Category List
	// (https://open.geelark.com/api/browser-extgroup-list).
	ExtGroup string `json:"extGroup,omitempty"`
	// Profile tag IDs.
	TagIDs []string `json:"tagIds,omitempty"`
	// Profile remark, up to 1500 characters.
	Remark string `json:"remark,omitempty"`
	// Cookies in JSON or Netscape format.
	Cookie string `json:"cookie,omitempty"`
	// Account platform. Only supports https://www.tiktok.com/ and https://www.facebook.com/.
	AccountPlatform string `json:"accountPlatform,omitempty"`
	AccountUsername string `json:"accountUsername,omitempty"`
	AccountPassword string `json:"accountPassword,omitempty"`
	// Account 2FA key.
	AccountTOTPSecret string `json:"accountTOTPSecret,omitempty"`
	// Whether to restore the last visit, 1-Yes, 2-No, default is No.
	OpenLastPage *int `json:"openLastPage,omitempty"`
	// Whether to open the specified URL, 1-Yes, 2-No, default is No.
	OpenSpecPage *int `json:"openSpecPage,omitempty"`
	// URLs to open, separated by semicolons (;).
	OpenTabs string `json:"openTabs,omitempty"`
	// Operating system, 1-win, 2-mac. Required.
	BrowserOS int `json:"browserOs"`
	// userAgent, if not passed, a random value will be used.
	BrowserUa string `json:"browserUa,omitempty"`
	// Simulation configuration, if not passed, a random value will be used.
	SimulateConfig *CreateRequestSimulateConfig `json:"simulateConfig,omitempty"`
	// Proxy ID, preferred.
	ProxyID string `json:"proxyId,omitempty"`
	// Proxy information.
	ProxyConfig *CreateRequestProxyConfig `json:"proxyConfig,omitempty"`
	// Startup parameters.
	BrowserStartArg string `json:"browserStartArg,omitempty"`
	// Kernel version. You may select a kernel version via this API: Query kernel
	// (https://open.geelark.com/api/browser-get-kernel) Or fill in "auto" for automatic
	// selection; leaving it blank also defaults to "auto".
	BrowserKernelVer string `json:"browserKernelVer,omitempty"`
}

type CreateResponse struct {
	// Profile ID.
	ID string `json:"id"`
}

// Create: https://open.geelark.com/api/create-new-browser
func (c *Client) Create(ctx context.Context, req CreateRequest) (*CreateResponse, error) {
	var resp CreateResponse
	if err := c.core.Do(ctx, "/api/v1/browser/create", req, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
