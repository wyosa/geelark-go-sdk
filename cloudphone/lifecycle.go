package cloudphone

import (
	"context"
)

type StatusRequest struct {
	// List of cloud phone IDs, Limit to 100 elements. Required.
	IDs []string `json:"ids"`
}

type StatusResponseSuccessDetailsItem struct {
	// ID of the successful cloud phone.
	ID string `json:"id"`
	// Name of the successful cloud phone.
	SerialName string `json:"serialName"`
	// Cloud phone status code; 0 - Started; 1 - Starting; 2 - Shut down; 3 - Expired.
	Status int `json:"status"`
}

type StatusResponseFailDetailsItem struct {
	// Failure code 42001: Cloud phone does not exist.
	Code int `json:"code"`
	// ID of the failed cloud phone.
	ID string `json:"id"`
	// Failure message.
	Msg string `json:"msg"`
}

type StatusResponse struct {
	// Total number of requested IDs.
	TotalAmount int `json:"totalAmount"`
	// Total number of successful responses.
	SuccessAmount int `json:"successAmount"`
	// Total number of failed responses.
	FailAmount int `json:"failAmount"`
	// Information about successful responses.
	SuccessDetails []StatusResponseSuccessDetailsItem `json:"successDetails"`
	// Information about failed responses.
	FailDetails []StatusResponseFailDetailsItem `json:"failDetails"`
}

// Status: https://open.geelark.com/api/phone-status
func (c *Client) Status(ctx context.Context, req StatusRequest) (*StatusResponse, error) {
	var resp StatusResponse
	if err := c.core.Do(ctx, "/open/v1/phone/status", req, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

type StartRequest struct {
	// List of cloud phone IDs, maximum 200. Required.
	IDs []string `json:"ids"`
	// Cloud phone display width in px; Default: 336 (allowed range: 200<=width<=600).
	Width *int `json:"width,omitempty"`
	// Whether the cloud phone display is centered; 0: not centered, 1: centered; default is 1
	// if not provided.
	Center *int `json:"center,omitempty"`
	// Enables energy-saving mode after the specified number of minutes; the cloud phone will
	// automatically shut down after being idle for that duration. Allowed values: 30, 60, 90;
	// passing 0 disables energy-saving mode.
	EnableEnergySavingModeAfterMinutes *int `json:"enableEnergySavingModeAfterMinutes,omitempty"`
	// Automatically cleans up data on the SD card older than the specified number of days,
	// which speeds up cloud phone startup. Allowed values: 3, 7; passing 0 disables the
	// cleanup feature, while omitting the parameter keeps the team's default settings.
	CleanupSDcardAfterDays *int `json:"cleanupSDcardAfterDays,omitempty"`
	// Material tag ID array, supporting up to 10 elements. When this parameter is provided,
	// the Cloud Phone Library opened by the API will only display materials corresponding to
	// the specified tags. If omitted or empty, all materials are shown by default. This
	// parameter requires OEM permissions (https://open.geelark.com/api/phone-customization) to
	// take effect.
	MaterialTagIDs []string `json:"materialTagIds,omitempty"`
	// Language. If not specified, the default is "en-US". Supported values:
	// "zh-CN","en-US","vi-VN","pt-BR","ru-RU","es-LA","uk-UA","ja-JP","zh-TW","id-ID","tr-TR","fr-FR".
	Lang string `json:"lang,omitempty"`
	// Dark mode. If omitted or set to false, dark mode is disabled.
	DarkMode *bool `json:"darkMode,omitempty"`
}

type StartResponseFailDetailsItem struct {
	// Failure code.
	Code int `json:"code"`
	// ID of the failed cloud phone.
	ID string `json:"id"`
	// Failure message.
	Msg string `json:"msg"`
}

type StartResponseSuccessDetailsItem struct {
	// ID of the cloud phone.
	ID string `json:"id"`
	// remote url, can be opened and visit cloud phone via Browser directly; To configure a
	// custom domain, please refer to OEM Custom Domain
	// (https://open.geelark.com/api/OEM-custom-domain).
	URL string `json:"url"`
	// Billing type, Per-minute usage, Monthly rental, Parallels.
	ChargingMethod string `json:"chargingMethod"`
}

type StartResponse struct {
	// Total number of requested IDs.
	TotalAmount int `json:"totalAmount"`
	// Number of successful starts.
	SuccessAmount int `json:"successAmount"`
	// Number of failed starts.
	FailAmount int `json:"failAmount"`
	// Information about failures.
	FailDetails []StartResponseFailDetailsItem `json:"failDetails"`
	// Information about successed.
	SuccessDetails []StartResponseSuccessDetailsItem `json:"successDetails"`
}

// Start: https://open.geelark.com/api/phone-start
func (c *Client) Start(ctx context.Context, req StartRequest) (*StartResponse, error) {
	var resp StartResponse
	if err := c.core.Do(ctx, "/open/v1/phone/start", req, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

type StopRequest struct {
	// List of cloud phone IDs, maximum 200. Required.
	IDs []string `json:"ids"`
}

type StopResponseSuccessDetailsItem struct {
	// Cloud phone ID.
	ID string `json:"id"`
	// Cloud phone name.
	SerialName string `json:"serialName"`
	// Cloud phone status reported after stopping. See the Status API for status values.
	Status int `json:"status"`
}

type StopResponseFailDetailsItem struct {
	// Error code.
	Code int `json:"code"`
	// Cloud phone ID.
	ID string `json:"id"`
	// Error message.
	Msg string `json:"msg"`
}

type StopResponse struct {
	// Total number of requested IDs.
	TotalAmount int `json:"totalAmount"`
	// Number of successfully shut down IDs.
	SuccessAmount int `json:"successAmount"`
	// Number of failed IDs.
	FailAmount int `json:"failAmount"`
	// Cloud phones successfully stopped.
	SuccessDetails []StopResponseSuccessDetailsItem `json:"successDetails"`
	// Information about failures.
	FailDetails []StopResponseFailDetailsItem `json:"failDetails"`
}

// Stop: https://open.geelark.com/api/phone-stop
func (c *Client) Stop(ctx context.Context, req StopRequest) (*StopResponse, error) {
	var resp StopResponse
	if err := c.core.Do(ctx, "/open/v1/phone/stop", req, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

type RestartRequest struct {
	// Cloud phone ID. Required.
	ID string `json:"id"`
}

type RestartResponse struct{}

// Restart: https://open.geelark.com/api/restart-cloud-phone
//
// Wait until the cloud phone startup callback has been received before calling Restart.
func (c *Client) Restart(ctx context.Context, req RestartRequest) (*RestartResponse, error) {
	var resp RestartResponse
	if err := c.core.Do(ctx, "/open/v1/phone/restart", req, nil); err != nil {
		return nil, err
	}
	return &resp, nil
}

type ResetRequest struct {
	// Cloud Phone ID. Required.
	ID string `json:"id"`
	// Whether to randomize the brand and model. true: randomize, false: keep unchanged; if not
	// provided, it remains randomize.
	ChangeBrandModel *bool `json:"changeBrandModel,omitempty"`
	// Preserve network connection type, defaults to false, random network connection type.
	KeepNetType *bool `json:"keepNetType,omitempty"`
	// Preserve phone number, defaults to false, random phone number.
	KeepPhoneNumber *bool `json:"keepPhoneNumber,omitempty"`
	// Preserve region, defaults to false, follow proxy.
	KeepRegion *bool `json:"keepRegion,omitempty"`
	// Preserve language, defaults to false, use English.
	KeepLanguage *bool `json:"keepLanguage,omitempty"`
	// Change mobile type: Android 9 Android 10 Android 11 Android 12 Android 13 Android 14
	// Android 15 Android 16.
	MobileType string `json:"mobileType,omitempty"`
}

type ResetResponse struct{}

// Reset: https://open.geelark.com/api/phone-newOne
//
// Reset deletes the applications and data on the cloud phone and generates a new device
// identity. Retrieve DeviceID again after the operation. The Keep fields preserve only the
// named settings.
func (c *Client) Reset(ctx context.Context, req ResetRequest) (*ResetResponse, error) {
	var resp ResetResponse
	if err := c.core.Do(ctx, "/open/v2/phone/newOne", req, nil); err != nil {
		return nil, err
	}
	return &resp, nil
}
