package cloudphone

import (
	"context"
)

type SetRootRequest struct {
	// List of cloud phone IDs ( currently supports Android 12,13,14,15,16), maximum 100.
	// Required.
	IDs []string `json:"ids"`
	// open/close. Required.
	Open bool `json:"open"`
}

type SetRootResponseItemsItem struct {
	// Per-item result code; zero indicates success.
	Code int `json:"code"`
	// Per-item result or failure message.
	Msg string `json:"msg"`
	// Cloud phone ID.
	ID string `json:"id"`
}

type SetRootResponse struct {
	// Per-phone operation results.
	Items []SetRootResponseItemsItem `json:"items"`
}

// SetRoot: https://open.geelark.com/api/root-setStatus
//
// Start the cloud phone before changing its root status.
func (c *Client) SetRoot(ctx context.Context, req SetRootRequest) (*SetRootResponse, error) {
	var resp SetRootResponse
	if err := c.core.Do(ctx, "/open/v1/root/setStatus", req, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

type GetRootRequest struct {
	// List of cloud phone IDs, minimum 1, maximum 100. Required.
	IDs []string `json:"ids"`
}

type GetRootResponseItemsItem struct {
	// Error code. 0 means success; for others, refer to the error code table.
	Code int `json:"code"`
	// Error message.
	Msg string `json:"msg"`
	// Cloud phone ID.
	ID string `json:"id"`
	// Whether root is enabled.
	Open bool `json:"open"`
}

type GetRootResponse struct {
	// Per-phone operation results.
	Items []GetRootResponseItemsItem `json:"items"`
}

// GetRoot: https://open.geelark.com/api/root-getStatus
func (c *Client) GetRoot(ctx context.Context, req GetRootRequest) (*GetRootResponse, error) {
	var resp GetRootResponse
	if err := c.core.Do(ctx, "/open/v1/root/getStatus", req, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

type DeviceIDRequest struct {
	// clound phone id. Required.
	ID string `json:"id"`
}

type DeviceIDResponse struct {
	// cloud phone device ID.
	SerialNum string `json:"serialNum"`
}

// DeviceID: https://open.geelark.com/api/phone-serialNum-get
//
// Retrieve the device ID again after Reset because the hardware identity changes. This is
// different from the persistent cloud phone profile ID.
func (c *Client) DeviceID(ctx context.Context, req DeviceIDRequest) (*DeviceIDResponse, error) {
	var resp DeviceIDResponse
	if err := c.core.Do(ctx, "/open/v1/phone/serialNum/get", req, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

type BrandsRequest struct {
	// android version, 10-16. Required.
	AndroidVer int `json:"androidVer"`
}

type BrandsResponseItem struct {
	// mobile phone brand.
	SurfaceBrandName string `json:"surfaceBrandName"`
	// mobile phone model.
	SurfaceModelName string `json:"surfaceModelName"`
}

type BrandsResponse []BrandsResponseItem

// Brands: https://open.geelark.com/api/phone-brand-list
func (c *Client) Brands(ctx context.Context, req BrandsRequest) (*BrandsResponse, error) {
	var resp BrandsResponse
	if err := c.core.Do(ctx, "/open/v1/phone/brand/list", req, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

type TeamBrandsRequest struct {
	// Page number, minimum is 1. Required.
	Page int `json:"page"`
	// Page size, minimum is 1, maximum is 100. Required.
	PageSize int `json:"pageSize"`
	// Android version, 9-16. Required.
	AndroidVer int `json:"androidVer"`
}

type TeamBrandsResponseListItem struct {
	// mobile phone brand.
	SurfaceBrandName string `json:"surfaceBrandName"`
	// mobile phone model.
	SurfaceModelName string `json:"surfaceModelName"`
}

type TeamBrandsResponse struct {
	// Total number.
	Total int `json:"total"`
	// Page number.
	Page     int `json:"page"`
	PageSize int `json:"pageSize"`
	// Array of brand models.
	List []TeamBrandsResponseListItem `json:"list"`
}

// TeamBrands: https://open.geelark.com/api/phone-brand-teamList
func (c *Client) TeamBrands(ctx context.Context, req TeamBrandsRequest) (*TeamBrandsResponse, error) {
	var resp TeamBrandsResponse
	if err := c.core.Do(ctx, "/open/v1/phone/brand/teamList", req, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

type HideAccessibilityRequest struct {
	// clound phone id array. Required.
	IDs []string `json:"ids"`
	// App package name array. If your App's package name is com.example.app, please pass
	// com.example.app directly on Android 14. On Android 12/13/15, please pass the package
	// name of the App you want to hide from, for example, to hide from TikTok, pass
	// com.zhiliaoapp.musically. Required.
	PkgName []string `json:"pkgName"`
}

type HideAccessibilityResponseFailDetailsItem struct {
	// cloud phone id.
	ID string `json:"id"`
	// error code.
	Code int `json:"code"`
	// error msg.
	Msg string `json:"msg"`
}

type HideAccessibilityResponse struct {
	// Details of individual failures, including the affected ID and error code.
	FailDetails []HideAccessibilityResponseFailDetailsItem `json:"failDetails"`
}

// HideAccessibility: https://open.geelark.com/api/phone-hideAccessibility
//
// Supports Android 12, 13, 14, and 15. Each call replaces the previous configuration.
func (c *Client) HideAccessibility(ctx context.Context, req HideAccessibilityRequest) (*HideAccessibilityResponse, error) {
	var resp HideAccessibilityResponse
	if err := c.core.Do(ctx, "/open/v1/phone/hideAccessibility", req, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
