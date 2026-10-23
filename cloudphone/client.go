package cloudphone

import (
	"context"

	"github.com/wyosa/geelark-go-sdk/internal/clientcore"
)

type Client struct {
	core *clientcore.Client
}

func New(core *clientcore.Client) *Client { return &Client{core: core} }

type CreateRequestDataItem struct {
	MobileType           string   `json:"mobileType,omitempty"`
	ProfileName          string   `json:"profileName,omitempty"`
	ProxyInformation     string   `json:"proxyInformation,omitempty"`
	RefreshURL           string   `json:"refreshUrl,omitempty"`
	ProxyQueryChannel    *int     `json:"proxyQueryChannel,omitempty"`
	ProxyNumber          *int     `json:"proxyNumber,omitempty"`
	DynamicProxy         string   `json:"dynamicProxy,omitempty"`
	DynamicProxyLocation string   `json:"dynamicProxyLocation,omitempty"`
	MobileRegion         string   `json:"mobileRegion,omitempty"`
	MobileProvince       string   `json:"mobileProvince,omitempty"`
	MobileCity           string   `json:"mobileCity,omitempty"`
	MobileLanguage       string   `json:"mobileLanguage,omitempty"`
	ProfileGroup         string   `json:"profileGroup,omitempty"`
	ProfileTags          []string `json:"profileTags,omitempty"`
	ProfileNote          string   `json:"profileNote,omitempty"`
	SurfaceBrandName     string   `json:"surfaceBrandName,omitempty"`
	SurfaceModelName     string   `json:"surfaceModelName,omitempty"`
	IsTeamBrand          *bool    `json:"isTeamBrand,omitempty"`
	NetType              *int     `json:"netType,omitempty"`
	PhoneNumber          string   `json:"phoneNumber,omitempty"`
	PhoneName            string   `json:"phoneName,omitempty"`
	WiFiID               string   `json:"wifiId,omitempty"`
	Protocol             *int     `json:"protocol,omitempty"`
	SessionType          *int     `json:"sessionType,omitempty"`
	FilterMode           *int     `json:"filterMode,omitempty"`
}

type CreateRequest struct {
	MobileType string                  `json:"mobileType"`
	ChargeMode *int                    `json:"chargeMode,omitempty"`
	Region     string                  `json:"region,omitempty"`
	Data       []CreateRequestDataItem `json:"data"`
}

type CreateResponseDetailsItemEquipmentInfo struct {
	CountryName  string `json:"countryName"`
	PhoneNumber  string `json:"phoneNumber"`
	EnableSim    int    `json:"enableSim"`
	IMEI         string `json:"imei"`
	OSVersion    string `json:"osVersion"`
	WiFiBssid    string `json:"wifiBssid"`
	MAC          string `json:"mac"`
	BluetoothMAC string `json:"bluetoothMac"`
	TimeZone     string `json:"timeZone"`
	DeviceBrand  string `json:"deviceBrand,omitempty"`
	DeviceModel  string `json:"deviceModel,omitempty"`
}

type CreateResponseDetailsItem struct {
	Index         int                                    `json:"index"`
	Code          int                                    `json:"code"`
	Msg           string                                 `json:"msg"`
	ID            string                                 `json:"id"`
	ProfileName   string                                 `json:"profileName"`
	EnvSerialNo   string                                 `json:"envSerialNo"`
	EquipmentInfo CreateResponseDetailsItemEquipmentInfo `json:"equipmentInfo"`
}

type CreateResponse struct {
	TotalAmount   int                         `json:"totalAmount"`
	SuccessAmount int                         `json:"successAmount"`
	FailAmount    int                         `json:"failAmount"`
	Details       []CreateResponseDetailsItem `json:"details"`
}

// Create calls /open/v1/phone/addNew.
func (c *Client) Create(ctx context.Context, req CreateRequest) (*CreateResponse, error) {
	var resp CreateResponse
	if err := c.core.Do(ctx, "/open/v1/phone/addNew", req, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

type ListRequest struct {
	Page       int      `json:"page,omitempty"`
	PageSize   int      `json:"pageSize,omitempty"`
	IDs        []string `json:"ids,omitempty"`
	SerialName string   `json:"serialName,omitempty"`
	Remark     *string  `json:"remark,omitempty"`
	GroupName  *string  `json:"groupName,omitempty"`
	Tags       []string `json:"tags,omitempty"`
	ChargeMode *int     `json:"chargeMode,omitempty"`
	OpenStatus *int     `json:"openStatus,omitempty"`
	ProxyIDs   []string `json:"proxyIds,omitempty"`
	SerialNos  []string `json:"serialNos,omitempty"`
}

type ListResponseItemsItemGroup struct {
	ID     string `json:"id"`
	Name   string `json:"name"`
	Remark string `json:"remark"`
}

type ListResponseItemsItemTagsItem struct {
	Name string `json:"name"`
}

type ListResponseItemsItemEquipmentInfo struct {
	CountryName  string `json:"countryName"`
	PhoneNumber  string `json:"phoneNumber"`
	EnableSim    int    `json:"enableSim"`
	IMEI         string `json:"imei"`
	OSVersion    string `json:"osVersion"`
	WiFiBssid    string `json:"wifiBssid"`
	MAC          string `json:"mac"`
	BluetoothMAC string `json:"bluetoothMac"`
	TimeZone     string `json:"timeZone"`
	DeviceName   string `json:"deviceName"`
	NetType      int    `json:"netType"`
	Language     string `json:"language"`
	Province     string `json:"province"`
	City         string `json:"city"`
	DeviceBrand  string `json:"deviceBrand,omitempty"`
	DeviceModel  string `json:"deviceModel,omitempty"`
}

type ListResponseItemsItemProxy struct {
	Type     string `json:"type"`
	Server   string `json:"server"`
	Port     int    `json:"port"`
	Username string `json:"username"`
	Password string `json:"password"`
}

type ListResponseItemsItem struct {
	ID            string                             `json:"id"`
	SerialName    string                             `json:"serialName"`
	SerialNo      string                             `json:"serialNo"`
	Group         ListResponseItemsItemGroup         `json:"group"`
	Remark        string                             `json:"remark"`
	Status        int                                `json:"status"`
	Tags          []ListResponseItemsItemTagsItem    `json:"tags"`
	EquipmentInfo ListResponseItemsItemEquipmentInfo `json:"equipmentInfo"`
	Proxy         ListResponseItemsItemProxy         `json:"proxy"`
	ChargeMode    int                                `json:"chargeMode,omitempty"`
	HasBind       bool                               `json:"hasBind,omitempty"`
	MonthlyExpire int64                              `json:"monthlyExpire,omitempty"`
	RPAStatus     int                                `json:"rpaStatus,omitempty"`
	CreateTime    int64                              `json:"createTime,omitempty"`
	SystemVer     string                             `json:"systemVer,omitempty"`
}

type ListResponse struct {
	Total    int                     `json:"total"`
	Page     int                     `json:"page"`
	PageSize int                     `json:"pageSize"`
	Items    []ListResponseItemsItem `json:"items"`
}

// List calls /open/v1/phone/list.
func (c *Client) List(ctx context.Context, req ListRequest) (*ListResponse, error) {
	var resp ListResponse
	if err := c.core.Do(ctx, "/open/v1/phone/list", req, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

type StatusRequest struct {
	IDs []string `json:"ids"`
}

type StatusResponseSuccessDetailsItem struct {
	ID         string `json:"id"`
	SerialName string `json:"serialName"`
	Status     int    `json:"status"`
}

type StatusResponseFailDetailsItem struct {
	Code int    `json:"code"`
	ID   string `json:"id"`
	Msg  string `json:"msg"`
}

type StatusResponse struct {
	TotalAmount    int                                `json:"totalAmount"`
	SuccessAmount  int                                `json:"successAmount"`
	FailAmount     int                                `json:"failAmount"`
	SuccessDetails []StatusResponseSuccessDetailsItem `json:"successDetails"`
	FailDetails    []StatusResponseFailDetailsItem    `json:"failDetails"`
}

// Status calls /open/v1/phone/status.
func (c *Client) Status(ctx context.Context, req StatusRequest) (*StatusResponse, error) {
	var resp StatusResponse
	if err := c.core.Do(ctx, "/open/v1/phone/status", req, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

type StartRequest struct {
	IDs                                []string `json:"ids"`
	Width                              *int     `json:"width,omitempty"`
	Center                             *int     `json:"center,omitempty"`
	EnableEnergySavingModeAfterMinutes *int     `json:"enableEnergySavingModeAfterMinutes,omitempty"`
	CleanupSDcardAfterDays             *int     `json:"cleanupSDcardAfterDays,omitempty"`
	MaterialTagIDs                     []string `json:"materialTagIds,omitempty"`
	Lang                               string   `json:"lang,omitempty"`
	DarkMode                           *bool    `json:"darkMode,omitempty"`
}

type StartResponseFailDetailsItem struct {
	Code int    `json:"code"`
	ID   string `json:"id"`
	Msg  string `json:"msg"`
}

type StartResponseSuccessDetailsItem struct {
	ID             string `json:"id"`
	URL            string `json:"url"`
	ChargingMethod string `json:"chargingMethod"`
}

type StartResponse struct {
	TotalAmount    int                               `json:"totalAmount"`
	SuccessAmount  int                               `json:"successAmount"`
	FailAmount     int                               `json:"failAmount"`
	FailDetails    []StartResponseFailDetailsItem    `json:"failDetails"`
	SuccessDetails []StartResponseSuccessDetailsItem `json:"successDetails"`
}

// Start calls /open/v1/phone/start.
func (c *Client) Start(ctx context.Context, req StartRequest) (*StartResponse, error) {
	var resp StartResponse
	if err := c.core.Do(ctx, "/open/v1/phone/start", req, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

type StopRequest struct {
	IDs []string `json:"ids"`
}

type StopResponseSuccessDetailsItem struct {
	ID         string `json:"id"`
	SerialName string `json:"serialName"`
	Status     int    `json:"status"`
}

type StopResponseFailDetailsItem struct {
	Code int    `json:"code"`
	ID   string `json:"id"`
	Msg  string `json:"msg"`
}

type StopResponse struct {
	TotalAmount    int                              `json:"totalAmount"`
	SuccessAmount  int                              `json:"successAmount"`
	FailAmount     int                              `json:"failAmount"`
	SuccessDetails []StopResponseSuccessDetailsItem `json:"successDetails"`
	FailDetails    []StopResponseFailDetailsItem    `json:"failDetails"`
}

// Stop calls /open/v1/phone/stop.
func (c *Client) Stop(ctx context.Context, req StopRequest) (*StopResponse, error) {
	var resp StopResponse
	if err := c.core.Do(ctx, "/open/v1/phone/stop", req, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

type RestartRequest struct {
	ID string `json:"id"`
}

type RestartResponse struct{}

// Restart calls /open/v1/phone/restart.
func (c *Client) Restart(ctx context.Context, req RestartRequest) (*RestartResponse, error) {
	var resp RestartResponse
	if err := c.core.Do(ctx, "/open/v1/phone/restart", req, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

type DeleteRequest struct {
	IDs []string `json:"ids"`
}

type DeleteResponseFailDetailsItem struct {
	Code int    `json:"code"`
	ID   string `json:"id"`
	Msg  string `json:"msg"`
}

type DeleteResponse struct {
	TotalAmount   int                             `json:"totalAmount"`
	SuccessAmount int                             `json:"successAmount"`
	FailAmount    int                             `json:"failAmount"`
	FailDetails   []DeleteResponseFailDetailsItem `json:"failDetails"`
}

// Delete calls /open/v1/phone/delete.
func (c *Client) Delete(ctx context.Context, req DeleteRequest) (*DeleteResponse, error) {
	var resp DeleteResponse
	if err := c.core.Do(ctx, "/open/v1/phone/delete", req, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

type GetGPSRequest struct {
	IDs []string `json:"ids"`
}

type GetGPSResponseListItem struct {
	ID        string  `json:"id"`
	Latitude  float64 `json:"latitude"`
	Longitude float64 `json:"longitude"`
}

type GetGPSResponse struct {
	TotalAmount   int                      `json:"totalAmount"`
	SuccessAmount int                      `json:"successAmount"`
	FailAmount    int                      `json:"failAmount"`
	List          []GetGPSResponseListItem `json:"list"`
}

// GetGPS calls /open/v1/phone/gps/get.
func (c *Client) GetGPS(ctx context.Context, req GetGPSRequest) (*GetGPSResponse, error) {
	var resp GetGPSResponse
	if err := c.core.Do(ctx, "/open/v1/phone/gps/get", req, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

type SetGPSRequestListItem struct {
	ID        string   `json:"id,omitempty"`
	Latitude  *float64 `json:"latitude,omitempty"`
	Longitude *float64 `json:"longitude,omitempty"`
}

type SetGPSRequest struct {
	List []SetGPSRequestListItem `json:"list"`
}

type SetGPSResponse struct {
	TotalAmount   int `json:"totalAmount"`
	SuccessAmount int `json:"successAmount"`
	FailAmount    int `json:"failAmount"`
}

// SetGPS calls /open/v1/phone/gps/set.
func (c *Client) SetGPS(ctx context.Context, req SetGPSRequest) (*SetGPSResponse, error) {
	var resp SetGPSResponse
	if err := c.core.Do(ctx, "/open/v1/phone/gps/set", req, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

type UpdateRequestProxyConfig struct {
	TypeID      *int    `json:"typeId,omitempty"`
	Server      *string `json:"server,omitempty"`
	Port        *int    `json:"port,omitempty"`
	Username    *string `json:"username,omitempty"`
	Password    *string `json:"password,omitempty"`
	UseProxyCfg *bool   `json:"useProxyCfg,omitempty"`
	Protocol    *int    `json:"protocol,omitempty"`
	Country     *string `json:"country,omitempty"`
	Region      *string `json:"region,omitempty"`
	City        *string `json:"city,omitempty"`
	SessionType *int    `json:"sessionType,omitempty"`
	FilterMode  *int    `json:"filterMode,omitempty"`
}

type UpdateRequest struct {
	ID                string                    `json:"id"`
	Name              *string                   `json:"name,omitempty"`
	Remark            *string                   `json:"remark,omitempty"`
	GroupID           *string                   `json:"groupID,omitempty"`
	TagIDs            []string                  `json:"tagIDs,omitempty"`
	ProxyConfig       *UpdateRequestProxyConfig `json:"proxyConfig,omitempty"`
	ProxyQueryChannel *int                      `json:"proxyQueryChannel,omitempty"`
	ProxyID           *string                   `json:"proxyId,omitempty"`
	ChargeMode        *int                      `json:"chargeMode,omitempty"`
	PhoneNumber       *string                   `json:"phoneNumber,omitempty"`
}

type UpdateResponse struct {
	FailDetails []UpdateResponseFailDetailsItem `json:"failDetails"`
}

type UpdateResponseFailDetailsItem struct {
	Code int    `json:"code"`
	ID   string `json:"id"`
	Msg  string `json:"msg"`
}

// Update calls /open/v1/phone/detail/update.
func (c *Client) Update(ctx context.Context, req UpdateRequest) (*UpdateResponse, error) {
	var resp UpdateResponse
	if err := c.core.Do(ctx, "/open/v1/phone/detail/update", req, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

type ResetRequest struct {
	ID               string `json:"id"`
	ChangeBrandModel *bool  `json:"changeBrandModel,omitempty"`
	KeepNetType      *bool  `json:"keepNetType,omitempty"`
	KeepPhoneNumber  *bool  `json:"keepPhoneNumber,omitempty"`
	KeepRegion       *bool  `json:"keepRegion,omitempty"`
	KeepLanguage     *bool  `json:"keepLanguage,omitempty"`
	MobileType       string `json:"mobileType,omitempty"`
}

type ResetResponse struct{}

// Reset calls /open/v2/phone/newOne.
func (c *Client) Reset(ctx context.Context, req ResetRequest) (*ResetResponse, error) {
	var resp ResetResponse
	if err := c.core.Do(ctx, "/open/v2/phone/newOne", req, &resp); err != nil {
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
	CloneNetType *bool  `json:"cloneNetType,omitempty"`
}

type CloneResponse struct {
	IDs []string `json:"ids"`
}

// Clone calls /open/v1/phone/clone.
func (c *Client) Clone(ctx context.Context, req CloneRequest) (*CloneResponse, error) {
	var resp CloneResponse
	if err := c.core.Do(ctx, "/open/v1/phone/clone", req, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

type ScreenshotRequest struct {
	ID string `json:"id"`
}

type ScreenshotResponse struct {
	TaskID string `json:"taskId"`
}

// Screenshot calls /open/v1/phone/screenShot.
func (c *Client) Screenshot(ctx context.Context, req ScreenshotRequest) (*ScreenshotResponse, error) {
	var resp ScreenshotResponse
	if err := c.core.Do(ctx, "/open/v1/phone/screenShot", req, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

type ScreenshotResultRequest struct {
	TaskID string `json:"taskId"`
}

type ScreenshotResultResponse struct {
	Status       int    `json:"status"`
	DownloadLink string `json:"downloadLink"`
}

// ScreenshotResult calls /open/v1/phone/screenShot/result.
func (c *Client) ScreenshotResult(ctx context.Context, req ScreenshotResultRequest) (*ScreenshotResultResponse, error) {
	var resp ScreenshotResultResponse
	if err := c.core.Do(ctx, "/open/v1/phone/screenShot/result", req, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

type SetRootRequest struct {
	IDs  []string `json:"ids"`
	Open bool     `json:"open"`
}

type SetRootResponseItemsItem struct {
	Code int    `json:"code"`
	Msg  string `json:"msg"`
	ID   string `json:"id"`
}

type SetRootResponse struct {
	Items []SetRootResponseItemsItem `json:"items"`
}

// SetRoot calls /open/v1/root/setStatus.
func (c *Client) SetRoot(ctx context.Context, req SetRootRequest) (*SetRootResponse, error) {
	var resp SetRootResponse
	if err := c.core.Do(ctx, "/open/v1/root/setStatus", req, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

type GetRootRequest struct {
	IDs []string `json:"ids"`
}

type GetRootResponseItemsItem struct {
	Code int    `json:"code"`
	Msg  string `json:"msg"`
	ID   string `json:"id"`
	Open bool   `json:"open"`
}

type GetRootResponse struct {
	Items []GetRootResponseItemsItem `json:"items"`
}

// GetRoot calls /open/v1/root/getStatus.
func (c *Client) GetRoot(ctx context.Context, req GetRootRequest) (*GetRootResponse, error) {
	var resp GetRootResponse
	if err := c.core.Do(ctx, "/open/v1/root/getStatus", req, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

type DeviceIDRequest struct {
	ID string `json:"id"`
}

type DeviceIDResponse struct {
	SerialNum string `json:"serialNum"`
}

// DeviceID calls /open/v1/phone/serialNum/get.
func (c *Client) DeviceID(ctx context.Context, req DeviceIDRequest) (*DeviceIDResponse, error) {
	var resp DeviceIDResponse
	if err := c.core.Do(ctx, "/open/v1/phone/serialNum/get", req, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

type SendSMSRequest struct {
	ID          string `json:"id"`
	PhoneNumber string `json:"phoneNumber"`
	Text        string `json:"text"`
}

type SendSMSResponse struct{}

// SendSMS calls /open/v1/phone/sendSms.
func (c *Client) SendSMS(ctx context.Context, req SendSMSRequest) (*SendSMSResponse, error) {
	var resp SendSMSResponse
	if err := c.core.Do(ctx, "/open/v1/phone/sendSms", req, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

type BrandsRequest struct {
	AndroidVer int `json:"androidVer"`
}

type BrandsResponseItem struct {
	SurfaceBrandName string `json:"surfaceBrandName"`
	SurfaceModelName string `json:"surfaceModelName"`
}

type BrandsResponse []BrandsResponseItem

// Brands calls /open/v1/phone/brand/list.
func (c *Client) Brands(ctx context.Context, req BrandsRequest) (*BrandsResponse, error) {
	var resp BrandsResponse
	if err := c.core.Do(ctx, "/open/v1/phone/brand/list", req, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

type TeamBrandsRequest struct {
	Page       int `json:"page"`
	PageSize   int `json:"pageSize"`
	AndroidVer int `json:"androidVer"`
}

type TeamBrandsResponseListItem struct {
	SurfaceBrandName string `json:"surfaceBrandName"`
	SurfaceModelName string `json:"surfaceModelName"`
}

type TeamBrandsResponse struct {
	Total    int                          `json:"total"`
	Page     int                          `json:"page"`
	PageSize int                          `json:"pageSize"`
	List     []TeamBrandsResponseListItem `json:"list"`
}

// TeamBrands calls /open/v1/phone/brand/teamList.
func (c *Client) TeamBrands(ctx context.Context, req TeamBrandsRequest) (*TeamBrandsResponse, error) {
	var resp TeamBrandsResponse
	if err := c.core.Do(ctx, "/open/v1/phone/brand/teamList", req, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

type TransferRequest struct {
	Account        string   `json:"account"`
	IDs            []string `json:"ids"`
	TransferOption []string `json:"transferOption,omitempty"`
}

type TransferResponseItem struct {
	SuccessCount int      `json:"successCount"`
	FailCount    int      `json:"failCount"`
	FailEnvIDs   []string `json:"failEnvIds"`
}

type TransferResponse []TransferResponseItem

// Transfer calls /open/v1/phone/transfer.
func (c *Client) Transfer(ctx context.Context, req TransferRequest) (*TransferResponse, error) {
	var resp TransferResponse
	if err := c.core.Do(ctx, "/open/v1/phone/transfer", req, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

type SetNetworkTypeRequest struct {
	ID      string `json:"id"`
	NetType int    `json:"netType"`
	WiFiID  string `json:"wifiId,omitempty"`
}

type SetNetworkTypeResponse struct{}

// SetNetworkType calls /open/v1/phone/net/set.
func (c *Client) SetNetworkType(ctx context.Context, req SetNetworkTypeRequest) (*SetNetworkTypeResponse, error) {
	var resp SetNetworkTypeResponse
	if err := c.core.Do(ctx, "/open/v1/phone/net/set", req, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

type HideAccessibilityRequest struct {
	IDs     []string `json:"ids"`
	PkgName []string `json:"pkgName"`
}

type HideAccessibilityResponseFailDetailsItem struct {
	ID   string `json:"id"`
	Code int    `json:"code"`
	Msg  string `json:"msg"`
}

type HideAccessibilityResponse struct {
	FailDetails []HideAccessibilityResponseFailDetailsItem `json:"failDetails"`
}

// HideAccessibility calls /open/v1/phone/hideAccessibility.
func (c *Client) HideAccessibility(ctx context.Context, req HideAccessibilityRequest) (*HideAccessibilityResponse, error) {
	var resp HideAccessibilityResponse
	if err := c.core.Do(ctx, "/open/v1/phone/hideAccessibility", req, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

type MoveGroupRequest struct {
	EnvIDs  []string `json:"envIds"`
	GroupID string   `json:"groupId"`
}

type MoveGroupResponse struct{}

// MoveGroup calls /open/v1/phone/moveGroup.
func (c *Client) MoveGroup(ctx context.Context, req MoveGroupRequest) (*MoveGroupResponse, error) {
	var resp MoveGroupResponse
	if err := c.core.Do(ctx, "/open/v1/phone/moveGroup", req, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

type ImportContactsRequestContactsItem struct {
	Email1    string `json:"email1,omitempty"`
	Email2    string `json:"email2,omitempty"`
	Fax       string `json:"fax,omitempty"`
	FirstName string `json:"firstName,omitempty"`
	LastName  string `json:"lastName,omitempty"`
	Mobile    string `json:"mobile,omitempty"`
	Work      string `json:"work,omitempty"`
}

type ImportContactsRequest struct {
	ID       string                              `json:"id"`
	Contacts []ImportContactsRequestContactsItem `json:"contacts"`
}

type ImportContactsResponse struct {
	TaskID string `json:"taskId"`
}

// ImportContacts calls /open/v1/phone/importContacts.
func (c *Client) ImportContacts(ctx context.Context, req ImportContactsRequest) (*ImportContactsResponse, error) {
	var resp ImportContactsResponse
	if err := c.core.Do(ctx, "/open/v1/phone/importContacts", req, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

type ImportContactsResultRequest struct {
	TaskID string `json:"taskId"`
}

type ImportContactsResultResponse struct {
	Status int `json:"status"`
}

// ImportContactsResult calls /open/v1/phone/importContactsResult.
func (c *Client) ImportContactsResult(ctx context.Context, req ImportContactsResultRequest) (*ImportContactsResultResponse, error) {
	var resp ImportContactsResultResponse
	if err := c.core.Do(ctx, "/open/v1/phone/importContactsResult", req, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

type GetNetworkSettingsRequest struct{}

type GetNetworkSettingsResponse struct {
	BlackList []string `json:"blackList"`
}

// GetNetworkSettings calls /open/v1/phone/netConfig/get.
func (c *Client) GetNetworkSettings(ctx context.Context, req GetNetworkSettingsRequest) (*GetNetworkSettingsResponse, error) {
	var resp GetNetworkSettingsResponse
	if err := c.core.Do(ctx, "/open/v1/phone/netConfig/get", req, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

type SetNetworkSettingsRequest struct {
	BlackList []string `json:"blackList,omitempty"`
}

type SetNetworkSettingsResponse struct{}

// SetNetworkSettings calls /open/v1/phone/netConfig/set.
func (c *Client) SetNetworkSettings(ctx context.Context, req SetNetworkSettingsRequest) (*SetNetworkSettingsResponse, error) {
	var resp SetNetworkSettingsResponse
	if err := c.core.Do(ctx, "/open/v1/phone/netConfig/set", req, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

type UIDumpRequest struct {
	ID string `json:"id"`
}

type UIDumpResponse struct {
	TaskID string `json:"taskId"`
}

// UIDump calls /open/v1/phone/uiDump.
func (c *Client) UIDump(ctx context.Context, req UIDumpRequest) (*UIDumpResponse, error) {
	var resp UIDumpResponse
	if err := c.core.Do(ctx, "/open/v1/phone/uiDump", req, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

type UIDumpResultRequest struct {
	TaskID string `json:"taskId"`
}

type UIDumpResultResponse struct {
	Status       int    `json:"status"`
	DownloadLink string `json:"downloadLink"`
}

// UIDumpResult calls /open/v1/phone/uiDump/result.
func (c *Client) UIDumpResult(ctx context.Context, req UIDumpResultRequest) (*UIDumpResultResponse, error) {
	var resp UIDumpResultResponse
	if err := c.core.Do(ctx, "/open/v1/phone/uiDump/result", req, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
