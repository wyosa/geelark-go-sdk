package browser

import (
	"context"
	"encoding/json"

	"github.com/wyosa/geelark-go-sdk/internal/clientcore"
)

type Client struct {
	core   *clientcore.Client
	remote *clientcore.Client
}

func New(core, remote *clientcore.Client) *Client { return &Client{core: core, remote: remote} }

type CreateRequestSimulateConfigTimeZone struct {
	Switcher *int   `json:"switcher,omitempty"`
	Value    string `json:"value,omitempty"`
}

type CreateRequestSimulateConfigLanguage struct {
	Switcher *int   `json:"switcher,omitempty"`
	Value    string `json:"value,omitempty"`
}

type CreateRequestSimulateConfigResolution struct {
	Switcher *int   `json:"switcher,omitempty"`
	Value    string `json:"value,omitempty"`
}

type CreateRequestSimulateConfigWebRTC struct {
	Switcher *int `json:"switcher,omitempty"`
}

type CreateRequestSimulateConfigGeoLocation struct {
	Switcher  *int  `json:"switcher,omitempty"`
	BaseOnIP  *bool `json:"baseOnIp,omitempty"`
	Longitude *int  `json:"longitude,omitempty"`
	Latitude  *int  `json:"latitude,omitempty"`
	Accuracy  *int  `json:"accuracy,omitempty"`
}

type CreateRequestSimulateConfigCanvas struct {
	Switcher *int `json:"switcher,omitempty"`
}

type CreateRequestSimulateConfigWebGLImage struct {
	Switcher *int `json:"switcher,omitempty"`
}

type CreateRequestSimulateConfigHardware struct {
	Switcher *int `json:"switcher,omitempty"`
}

type CreateRequestSimulateConfigAudioContext struct {
	Switcher *int `json:"switcher,omitempty"`
}

type CreateRequestSimulateConfigMediaDevice struct {
	Switcher *int `json:"switcher,omitempty"`
}

type CreateRequestSimulateConfigClientRects struct {
	Switcher *int `json:"switcher,omitempty"`
}

type CreateRequestSimulateConfigSpeechVoise struct {
	Switcher *int `json:"switcher,omitempty"`
}

type CreateRequestSimulateConfigBluetooth struct {
	Switcher *int `json:"switcher,omitempty"`
}

type CreateRequestSimulateConfigBattery struct {
	Switcher *int `json:"switcher,omitempty"`
}

type CreateRequestSimulateConfigPortScanProtection struct {
	Switcher *int   `json:"switcher,omitempty"`
	Value    string `json:"value,omitempty"`
}

type CreateRequestSimulateConfig struct {
	TimeZone            *CreateRequestSimulateConfigTimeZone           `json:"timeZone,omitempty"`
	Language            *CreateRequestSimulateConfigLanguage           `json:"language,omitempty"`
	Resolution          *CreateRequestSimulateConfigResolution         `json:"resolution,omitempty"`
	WebRTC              *CreateRequestSimulateConfigWebRTC             `json:"webRtc,omitempty"`
	GeoLocation         *CreateRequestSimulateConfigGeoLocation        `json:"geoLocation,omitempty"`
	Canvas              *CreateRequestSimulateConfigCanvas             `json:"canvas,omitempty"`
	WebGLImage          *CreateRequestSimulateConfigWebGLImage         `json:"webglImage,omitempty"`
	Hardware            *CreateRequestSimulateConfigHardware           `json:"hardware,omitempty"`
	AudioContext        *CreateRequestSimulateConfigAudioContext       `json:"audioContext,omitempty"`
	MediaDevice         *CreateRequestSimulateConfigMediaDevice        `json:"mediaDevice,omitempty"`
	ClientRects         *CreateRequestSimulateConfigClientRects        `json:"clientRects,omitempty"`
	SpeechVoise         *CreateRequestSimulateConfigSpeechVoise        `json:"speechVoise,omitempty"`
	HardwareConcurrency *int                                           `json:"hardwareConcurrency,omitempty"`
	MemeryDevice        *int                                           `json:"memeryDevice,omitempty"`
	MemoryDevice        *int                                           `json:"memoryDevice,omitempty"`
	DoNotTrack          *int                                           `json:"doNotTrack,omitempty"`
	Bluetooth           *CreateRequestSimulateConfigBluetooth          `json:"bluetooth,omitempty"`
	Battery             *CreateRequestSimulateConfigBattery            `json:"battery,omitempty"`
	PortScanProtection  *CreateRequestSimulateConfigPortScanProtection `json:"portScanProtection,omitempty"`
}

type CreateRequestProxyConfig struct {
	TypeID      *int   `json:"typeId,omitempty"`
	Server      string `json:"server,omitempty"`
	Port        *int   `json:"port,omitempty"`
	Username    string `json:"username,omitempty"`
	Password    string `json:"password,omitempty"`
	Country     string `json:"country,omitempty"`
	Region      string `json:"region,omitempty"`
	City        string `json:"city,omitempty"`
	UseProxyCfg *bool  `json:"useProxyCfg,omitempty"`
	Protocol    *int   `json:"protocol,omitempty"`
	SessionType *int   `json:"sessionType,omitempty"`
	FilterMode  *int   `json:"filterMode,omitempty"`
}

type CreateRequest struct {
	SerialName        string                       `json:"serialName"`
	GroupID           string                       `json:"groupId,omitempty"`
	ExtGroup          string                       `json:"extGroup,omitempty"`
	TagIDs            []string                     `json:"tagIds,omitempty"`
	Remark            string                       `json:"remark,omitempty"`
	Cookie            string                       `json:"cookie,omitempty"`
	AccountPlatform   string                       `json:"accountPlatform,omitempty"`
	AccountUsername   string                       `json:"accountUsername,omitempty"`
	AccountPassword   string                       `json:"accountPassword,omitempty"`
	AccountTOTPSecret string                       `json:"accountTOTPSecret,omitempty"`
	OpenLastPage      *int                         `json:"openLastPage,omitempty"`
	OpenSpecPage      *int                         `json:"openSpecPage,omitempty"`
	OpenTabs          string                       `json:"openTabs,omitempty"`
	BrowserOS         int                          `json:"browserOs"`
	BrowserUa         string                       `json:"browserUa,omitempty"`
	SimulateConfig    *CreateRequestSimulateConfig `json:"simulateConfig,omitempty"`
	ProxyID           string                       `json:"proxyId,omitempty"`
	ProxyConfig       *CreateRequestProxyConfig    `json:"proxyConfig,omitempty"`
	BrowserStartArg   string                       `json:"browserStartArg,omitempty"`
	BrowserKernelVer  string                       `json:"browserKernelVer,omitempty"`
}

type CreateResponse struct {
	ID string `json:"id"`
}

// Create calls /api/v1/browser/create.
func (c *Client) Create(ctx context.Context, req CreateRequest) (*CreateResponse, error) {
	var resp CreateResponse
	if err := c.core.Do(ctx, "/api/v1/browser/create", req, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

type UpdateRequestSimulateConfigTimeZone struct {
	Switcher *int    `json:"switcher,omitempty"`
	Value    *string `json:"value,omitempty"`
}

type UpdateRequestSimulateConfigLanguage struct {
	Switcher *int    `json:"switcher,omitempty"`
	Value    *string `json:"value,omitempty"`
}

type UpdateRequestSimulateConfigResolution struct {
	Switcher *int    `json:"switcher,omitempty"`
	Value    *string `json:"value,omitempty"`
}

type UpdateRequestSimulateConfigWebRTC struct {
	Switcher *int `json:"switcher,omitempty"`
}

type UpdateRequestSimulateConfigGeoLocation struct {
	Switcher  *int  `json:"switcher,omitempty"`
	BaseOnIP  *bool `json:"baseOnIp,omitempty"`
	Longitude *int  `json:"longitude,omitempty"`
	Latitude  *int  `json:"latitude,omitempty"`
	Accuracy  *int  `json:"accuracy,omitempty"`
}

type UpdateRequestSimulateConfigCanvas struct {
	Switcher *int `json:"switcher,omitempty"`
}

type UpdateRequestSimulateConfigWebGLImage struct {
	Switcher *int `json:"switcher,omitempty"`
}

type UpdateRequestSimulateConfigHardware struct {
	Switcher *int `json:"switcher,omitempty"`
}

type UpdateRequestSimulateConfigAudioContext struct {
	Switcher *int `json:"switcher,omitempty"`
}

type UpdateRequestSimulateConfigMediaDevice struct {
	Switcher *int `json:"switcher,omitempty"`
}

type UpdateRequestSimulateConfigClientRects struct {
	Switcher *int `json:"switcher,omitempty"`
}

type UpdateRequestSimulateConfigSpeechVoise struct {
	Switcher *int `json:"switcher,omitempty"`
}

type UpdateRequestSimulateConfigBluetooth struct {
	Switcher *int `json:"switcher,omitempty"`
}

type UpdateRequestSimulateConfigBattery struct {
	Switcher *int `json:"switcher,omitempty"`
}

type UpdateRequestSimulateConfigPortScanProtection struct {
	Switcher *int    `json:"switcher,omitempty"`
	Value    *string `json:"value,omitempty"`
}

type UpdateRequestSimulateConfig struct {
	TimeZone            *UpdateRequestSimulateConfigTimeZone           `json:"timeZone,omitempty"`
	Language            *UpdateRequestSimulateConfigLanguage           `json:"language,omitempty"`
	Resolution          *UpdateRequestSimulateConfigResolution         `json:"resolution,omitempty"`
	WebRTC              *UpdateRequestSimulateConfigWebRTC             `json:"webRtc,omitempty"`
	GeoLocation         *UpdateRequestSimulateConfigGeoLocation        `json:"geoLocation,omitempty"`
	Canvas              *UpdateRequestSimulateConfigCanvas             `json:"canvas,omitempty"`
	WebGLImage          *UpdateRequestSimulateConfigWebGLImage         `json:"webglImage,omitempty"`
	Hardware            *UpdateRequestSimulateConfigHardware           `json:"hardware,omitempty"`
	AudioContext        *UpdateRequestSimulateConfigAudioContext       `json:"audioContext,omitempty"`
	MediaDevice         *UpdateRequestSimulateConfigMediaDevice        `json:"mediaDevice,omitempty"`
	ClientRects         *UpdateRequestSimulateConfigClientRects        `json:"clientRects,omitempty"`
	SpeechVoise         *UpdateRequestSimulateConfigSpeechVoise        `json:"speechVoise,omitempty"`
	HardwareConcurrency *int                                           `json:"hardwareConcurrency,omitempty"`
	MemeryDevice        *int                                           `json:"memeryDevice,omitempty"`
	MemoryDevice        *int                                           `json:"memoryDevice,omitempty"`
	DoNotTrack          *int                                           `json:"doNotTrack,omitempty"`
	Bluetooth           *UpdateRequestSimulateConfigBluetooth          `json:"bluetooth,omitempty"`
	Battery             *UpdateRequestSimulateConfigBattery            `json:"battery,omitempty"`
	PortScanProtection  *UpdateRequestSimulateConfigPortScanProtection `json:"portScanProtection,omitempty"`
}

type UpdateRequest struct {
	ID                string                       `json:"id"`
	SerialName        *string                      `json:"serialName,omitempty"`
	GroupID           *string                      `json:"groupId,omitempty"`
	ExtGroup          *string                      `json:"extGroup,omitempty"`
	TagIDs            []string                     `json:"tagIds,omitempty"`
	Remark            *string                      `json:"remark,omitempty"`
	Cookie            *string                      `json:"cookie,omitempty"`
	AccountPlatform   *string                      `json:"accountPlatform,omitempty"`
	AccountUsername   *string                      `json:"accountUsername,omitempty"`
	AccountPassword   *string                      `json:"accountPassword,omitempty"`
	AccountTOTPSecret *string                      `json:"accountTOTPSecret,omitempty"`
	OpenLastPage      *int                         `json:"openLastPage,omitempty"`
	OpenSpecPage      *int                         `json:"openSpecPage,omitempty"`
	OpenTabs          *string                      `json:"openTabs,omitempty"`
	BrowserOS         *int                         `json:"browserOs,omitempty"`
	BrowserUa         *string                      `json:"browserUa,omitempty"`
	SimulateConfig    *UpdateRequestSimulateConfig `json:"simulateConfig,omitempty"`
	ProxyID           *string                      `json:"proxyId,omitempty"`
	ProxyConfig       *UpdateRequestProxyConfig    `json:"proxyConfig,omitempty"`
	BrowserStartArg   *string                      `json:"browserStartArg,omitempty"`
}

type UpdateRequestProxyConfig struct {
	TypeID      *int    `json:"typeId,omitempty"`
	Server      *string `json:"server,omitempty"`
	Port        *int    `json:"port,omitempty"`
	Username    *string `json:"username,omitempty"`
	Password    *string `json:"password,omitempty"`
	Country     *string `json:"country,omitempty"`
	Region      *string `json:"region,omitempty"`
	City        *string `json:"city,omitempty"`
	UseProxyCfg *bool   `json:"useProxyCfg,omitempty"`
	Protocol    *int    `json:"protocol,omitempty"`
	SessionType *int    `json:"sessionType,omitempty"`
	FilterMode  *int    `json:"filterMode,omitempty"`
}

type UpdateResponse struct{}

// Update calls /api/v1/browser/update.
func (c *Client) Update(ctx context.Context, req UpdateRequest) (*UpdateResponse, error) {
	var resp UpdateResponse
	if err := c.core.Do(ctx, "/api/v1/browser/update", req, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

type DeleteRequest struct {
	EnvIDs []string `json:"envIds"`
}

type DeleteResponse struct {
	BusyIDs      []string `json:"busyIds"`
	ServerErrIDs []string `json:"serverErrIds"`
	SuccessIDs   []string `json:"successIds"`
}

// Delete calls /api/v1/browser/delete.
func (c *Client) Delete(ctx context.Context, req DeleteRequest) (*DeleteResponse, error) {
	var resp DeleteResponse
	if err := c.core.Do(ctx, "/api/v1/browser/delete", req, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

type ListRequest struct {
	Page       int      `json:"page,omitempty"`
	PageSize   int      `json:"pageSize,omitempty"`
	IDs        []string `json:"ids,omitempty"`
	SerialName string   `json:"serialName,omitempty"`
	Remark     string   `json:"remark,omitempty"`
	GroupName  string   `json:"groupName,omitempty"`
	Tags       []string `json:"tags,omitempty"`
}

type ListResponseItemsItemGroup struct {
	ID     string `json:"id"`
	Name   string `json:"name"`
	Remark string `json:"remark"`
}

type ListResponseItemsItemTagsItem struct {
	Name string `json:"name"`
}

type ListResponseItemsItemProxy struct {
	Type     string `json:"type"`
	Server   string `json:"server"`
	Port     int    `json:"port"`
	Username string `json:"username"`
	Password string `json:"password"`
}

type ListResponseItemsItemAccountInfo struct {
	URL          string   `json:"url"`
	UserName     string   `json:"userName"`
	PassWord     string   `json:"passWord"`
	TotpSecret   string   `json:"totpSecret"`
	AfterStartup int      `json:"afterStartup"`
	OpenLastPage int      `json:"openLastPage"`
	OpenSpecPage int      `json:"openSpecPage"`
	OpenSiteURL  bool     `json:"openSiteUrl"`
	AutoOpenURLs []string `json:"autoOpenUrls"`
}

type ListResponseItemsItemSimulateInfoTimeZone struct {
	Switcher int    `json:"switcher"`
	Value    string `json:"value"`
}

type ListResponseItemsItemSimulateInfoWebRTC struct {
	Switcher int `json:"switcher"`
}

type ListResponseItemsItemSimulateInfoGeoLocation struct {
	Switcher  int  `json:"switcher"`
	BaseOnIP  bool `json:"baseOnIp"`
	Longitude int  `json:"longitude"`
	Latitude  int  `json:"latitude"`
	Accuracy  int  `json:"accuracy"`
}

type ListResponseItemsItemSimulateInfoLanguage struct {
	Switcher int    `json:"switcher"`
	Value    string `json:"value"`
}

type ListResponseItemsItemSimulateInfoResolution struct {
	Switcher int    `json:"switcher"`
	Value    string `json:"value"`
}

type ListResponseItemsItemSimulateInfoFont struct {
	Switcher int `json:"switcher"`
}

type ListResponseItemsItemSimulateInfoCanvas struct {
	Switcher int `json:"switcher"`
}

type ListResponseItemsItemSimulateInfoWebGLImage struct {
	Switcher int `json:"switcher"`
}

type ListResponseItemsItemSimulateInfoWebGLMetadata struct {
	Switcher int    `json:"switcher"`
	Provider string `json:"provider"`
	Render   string `json:"render"`
}

type ListResponseItemsItemSimulateInfoHardware struct {
	Switcher int `json:"switcher"`
}

type ListResponseItemsItemSimulateInfoAudioContext struct {
	Switcher int `json:"switcher"`
}

type ListResponseItemsItemSimulateInfoMediaDevice struct {
	Switcher int `json:"switcher"`
}

type ListResponseItemsItemSimulateInfoClientRects struct {
	Switcher int `json:"switcher"`
}

type ListResponseItemsItemSimulateInfoSpeechVoise struct {
	Switcher int `json:"switcher"`
}

type ListResponseItemsItemSimulateInfoBluetooth struct {
	Switcher int `json:"switcher"`
}

type ListResponseItemsItemSimulateInfoBattery struct {
	Switcher int `json:"switcher"`
}

type ListResponseItemsItemSimulateInfoPortScanProtection struct {
	Switcher int    `json:"switcher"`
	Value    string `json:"value"`
}

type ListResponseItemsItemSimulateInfo struct {
	OS                  int                                                 `json:"os"`
	Vendor              int                                                 `json:"vendor"`
	MixtureKey          string                                              `json:"mixtureKey"`
	Ua                  string                                              `json:"ua"`
	UaVersion           string                                              `json:"uaVersion"`
	TimeZone            ListResponseItemsItemSimulateInfoTimeZone           `json:"timeZone"`
	WebRTC              ListResponseItemsItemSimulateInfoWebRTC             `json:"webRtc"`
	GeoLocation         ListResponseItemsItemSimulateInfoGeoLocation        `json:"geoLocation"`
	Language            ListResponseItemsItemSimulateInfoLanguage           `json:"language"`
	Resolution          ListResponseItemsItemSimulateInfoResolution         `json:"resolution"`
	Font                ListResponseItemsItemSimulateInfoFont               `json:"font"`
	Canvas              ListResponseItemsItemSimulateInfoCanvas             `json:"canvas"`
	WebGLImage          ListResponseItemsItemSimulateInfoWebGLImage         `json:"webglImage"`
	WebGLMetadata       ListResponseItemsItemSimulateInfoWebGLMetadata      `json:"webglMetadata"`
	Hardware            ListResponseItemsItemSimulateInfoHardware           `json:"hardware"`
	AudioContext        ListResponseItemsItemSimulateInfoAudioContext       `json:"audioContext"`
	MediaDevice         ListResponseItemsItemSimulateInfoMediaDevice        `json:"mediaDevice"`
	ClientRects         ListResponseItemsItemSimulateInfoClientRects        `json:"clientRects"`
	SpeechVoise         ListResponseItemsItemSimulateInfoSpeechVoise        `json:"speechVoise"`
	HardwareConcurrency int                                                 `json:"hardwareConcurrency"`
	MemeryDevice        int                                                 `json:"memeryDevice"`
	MemoryDevice        int                                                 `json:"memoryDevice,omitempty"`
	DoNotTrack          int                                                 `json:"doNotTrack"`
	Bluetooth           ListResponseItemsItemSimulateInfoBluetooth          `json:"bluetooth"`
	Battery             ListResponseItemsItemSimulateInfoBattery            `json:"battery"`
	PortScanProtection  ListResponseItemsItemSimulateInfoPortScanProtection `json:"portScanProtection"`
}

type ListResponseItemsItem struct {
	ID           string                            `json:"id"`
	SerialName   string                            `json:"serialName"`
	SerialNo     string                            `json:"serialNo"`
	Group        ListResponseItemsItemGroup        `json:"group"`
	ExtGroup     string                            `json:"extGroup"`
	Remark       string                            `json:"remark"`
	Tags         []ListResponseItemsItemTagsItem   `json:"tags"`
	Proxy        ListResponseItemsItemProxy        `json:"proxy"`
	AccountInfo  ListResponseItemsItemAccountInfo  `json:"accountInfo"`
	SimulateInfo ListResponseItemsItemSimulateInfo `json:"simulateInfo"`
}

type ListResponse struct {
	Total    int                     `json:"total"`
	Page     int                     `json:"page"`
	PageSize int                     `json:"pageSize"`
	Items    []ListResponseItemsItem `json:"items"`
}

// List calls /api/v1/browser/list.
func (c *Client) List(ctx context.Context, req ListRequest) (*ListResponse, error) {
	var resp ListResponse
	if err := c.core.Do(ctx, "/api/v1/browser/list", req, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

type StartRequest struct {
	ID      string `json:"id"`
	Webhook string `json:"webhook,omitempty"`
}

type StartResponse struct {
	DebugPort int `json:"debugPort"`
}

// Start calls /api/v1/browser/start.
func (c *Client) Start(ctx context.Context, req StartRequest) (*StartResponse, error) {
	var resp StartResponse
	if err := c.core.Do(ctx, "/api/v1/browser/start", req, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

type StopRequest struct {
	ID string `json:"id"`
}

type StopResponse struct{}

// Stop calls /api/v1/browser/stop.
func (c *Client) Stop(ctx context.Context, req StopRequest) (*StopResponse, error) {
	var resp StopResponse
	if err := c.core.Do(ctx, "/api/v1/browser/stop", req, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

type TransferRequest struct {
	Username       string   `json:"username"`
	EnvIDs         []string `json:"envIds"`
	TransferOption []string `json:"transferOption,omitempty"`
}

type TransferResponse struct {
	SuccessCount int      `json:"successCount"`
	FailCount    int      `json:"failCount"`
	FailEnvIDs   []string `json:"failEnvIds"`
}

// Transfer calls /api/v1/browser/transfer.
func (c *Client) Transfer(ctx context.Context, req TransferRequest) (*TransferResponse, error) {
	var resp TransferResponse
	if err := c.core.Do(ctx, "/api/v1/browser/transfer", req, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

type APIStatusRequest struct{}

type APIStatusResponse struct{}

// APIStatus calls /api/v1/status.
func (c *Client) APIStatus(ctx context.Context, req APIStatusRequest) (*APIStatusResponse, error) {
	var resp APIStatusResponse
	if err := c.core.Do(ctx, "/api/v1/status", req, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

type StatusRequest struct {
	ID string `json:"id"`
}

type StatusResponse struct {
	Status    string `json:"status"`
	DebugPort int    `json:"debugPort"`
}

// Status calls /api/v1/browser/status.
func (c *Client) Status(ctx context.Context, req StatusRequest) (*StatusResponse, error) {
	var resp StatusResponse
	if err := c.core.Do(ctx, "/api/v1/browser/status", req, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

type ClearCacheRequest struct {
	IDs  []string `json:"ids"`
	Type []string `json:"type"`
}

type ClearCacheResponse struct {
	SuccessCount int               `json:"success_count"`
	ErrorCount   int               `json:"error_count"`
	ErrorInfo    []json.RawMessage `json:"error_info"`
}

// ClearCache calls /api/v1/browser/deleteCache.
func (c *Client) ClearCache(ctx context.Context, req ClearCacheRequest) (*ClearCacheResponse, error) {
	var resp ClearCacheResponse
	if err := c.core.Do(ctx, "/api/v1/browser/deleteCache", req, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

type MoveGroupRequest struct {
	EnvIDs  []string `json:"envIds,omitempty"`
	GroupID string   `json:"groupId,omitempty"`
}

type MoveGroupResponse struct{}

// MoveGroup calls /api/v1/browser/moveGroup.
func (c *Client) MoveGroup(ctx context.Context, req MoveGroupRequest) (*MoveGroupResponse, error) {
	var resp MoveGroupResponse
	if err := c.core.Do(ctx, "/api/v1/browser/moveGroup", req, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

type CookiesRequest struct {
	ID string `json:"id"`
}

type CookiesResponse struct {
	Cookies string `json:"cookies"`
}

// Cookies calls /api/v1/browser/getCookie.
func (c *Client) Cookies(ctx context.Context, req CookiesRequest) (*CookiesResponse, error) {
	var resp CookiesResponse
	if err := c.core.Do(ctx, "/api/v1/browser/getCookie", req, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

type SetBookmarksRequestBrowserBookmark struct {
	Type     *int    `json:"type,omitempty"`
	FileAddr *string `json:"fileAddr,omitempty"`
	Text     string  `json:"text,omitempty"`
}

type SetBookmarksRequest struct {
	BrowserBookmark SetBookmarksRequestBrowserBookmark `json:"browserBookmark"`
}

type SetBookmarksResponse struct{}

// SetBookmarks calls /api/v1/browser/setBookmark.
func (c *Client) SetBookmarks(ctx context.Context, req SetBookmarksRequest) (*SetBookmarksResponse, error) {
	var resp SetBookmarksResponse
	if err := c.core.Do(ctx, "/api/v1/browser/setBookmark", req, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

type GetBookmarksRequest struct{}

type GetBookmarksResponseBrowserBookmark struct {
	Type     int    `json:"type"`
	FileAddr string `json:"fileAddr"`
	Text     string `json:"text"`
}

type GetBookmarksResponse struct {
	BrowserBookmark GetBookmarksResponseBrowserBookmark `json:"browserBookmark"`
}

// GetBookmarks calls /api/v1/browser/getBookmark.
func (c *Client) GetBookmarks(ctx context.Context, req GetBookmarksRequest) (*GetBookmarksResponse, error) {
	var resp GetBookmarksResponse
	if err := c.core.Do(ctx, "/api/v1/browser/getBookmark", req, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

type CloneRequest struct {
	EnvID        string `json:"envId"`
	Amount       int    `json:"amount"`
	GroupID      string `json:"groupId,omitempty"`
	CloneName    *bool  `json:"cloneName,omitempty"`
	CloneRemark  *bool  `json:"cloneRemark,omitempty"`
	CloneTag     *bool  `json:"cloneTag,omitempty"`
	CloneProxy   *bool  `json:"cloneProxy,omitempty"`
	CloneCookie  *bool  `json:"cloneCookie,omitempty"`
	CloneAccount *bool  `json:"cloneAccount,omitempty"`
}

type CloneResponse struct {
	IDs []string `json:"ids"`
}

// Clone calls /api/v1/browser/clone.
func (c *Client) Clone(ctx context.Context, req CloneRequest) (*CloneResponse, error) {
	var resp CloneResponse
	if err := c.core.Do(ctx, "/api/v1/browser/clone", req, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

type KernelsRequest struct{}

type KernelsResponseItem struct {
	Kernel       string `json:"kernel"`
	IsDownloaded bool   `json:"isDownloaded"`
}

type KernelsResponse []KernelsResponseItem

// Kernels calls /api/v1/browser/getKernelsList.
func (c *Client) Kernels(ctx context.Context, req KernelsRequest) (*KernelsResponse, error) {
	var resp KernelsResponse
	if err := c.core.Do(ctx, "/api/v1/browser/getKernelsList", req, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

type UpdateKernelsRequest struct {
	KernelVersion string `json:"kernel_version"`
}

type UpdateKernelsResponse struct {
	Status   string `json:"status"`
	Progress int    `json:"progress"`
}

// UpdateKernels calls /api/v1/browser/updateKernels.
func (c *Client) UpdateKernels(ctx context.Context, req UpdateKernelsRequest) (*UpdateKernelsResponse, error) {
	var resp UpdateKernelsResponse
	if err := c.core.Do(ctx, "/api/v1/browser/updateKernels", req, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

type ExtensionGroupsRequest struct {
	Page     int `json:"page,omitempty"`
	PageSize int `json:"pageSize,omitempty"`
}

type ExtensionGroupsResponseItemsItem struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

type ExtensionGroupsResponse struct {
	Total    int                                `json:"total"`
	Page     int                                `json:"page"`
	PageSize int                                `json:"pageSize"`
	Items    []ExtensionGroupsResponseItemsItem `json:"items"`
}

// ExtensionGroups calls /open/v1/browser/extGroup/list.
func (c *Client) ExtensionGroups(ctx context.Context, req ExtensionGroupsRequest) (*ExtensionGroupsResponse, error) {
	var resp ExtensionGroupsResponse
	if err := c.remote.Do(ctx, "/open/v1/browser/extGroup/list", req, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
