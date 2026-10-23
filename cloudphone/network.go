package cloudphone

import (
	"context"
)

type GetGPSRequest struct {
	// List of cloud phone IDs. Required.
	IDs []string `json:"ids"`
}

type GetGPSResponseListItem struct {
	// Cloud phone ID.
	ID        string  `json:"id"`
	Latitude  float64 `json:"latitude"`
	Longitude float64 `json:"longitude"`
}

type GetGPSResponse struct {
	// Total number of requested IDs.
	TotalAmount int `json:"totalAmount"`
	// Total number of successful requests.
	SuccessAmount int `json:"successAmount"`
	// Total number of failed requests.
	FailAmount int `json:"failAmount"`
	// GPS information.
	List []GetGPSResponseListItem `json:"list"`
}

// GetGPS: https://open.geelark.com/api/phone-gps-get
func (c *Client) GetGPS(ctx context.Context, req GetGPSRequest) (*GetGPSResponse, error) {
	var resp GetGPSResponse
	if err := c.core.Do(ctx, "/open/v1/phone/gps/get", req, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

type SetGPSRequestListItem struct {
	// Cloud phone ID. Required.
	ID string `json:"id,omitempty"`
	// Latitude. Required.
	Latitude *float64 `json:"latitude,omitempty"`
	// Longitude. Required.
	Longitude *float64 `json:"longitude,omitempty"`
}

type SetGPSRequest struct {
	// Cloud phone GPS info. Required.
	List []SetGPSRequestListItem `json:"list"`
}

type SetGPSResponse struct {
	// Total number of requested IDs.
	TotalAmount int `json:"totalAmount"`
	// Total number of successful requests.
	SuccessAmount int `json:"successAmount"`
	// Total number of failed requests.
	FailAmount int `json:"failAmount"`
}

// SetGPS: https://open.geelark.com/api/phone-gps-set
//
// Android 16 is not currently supported. Longitude must be within [-180, 180] and latitude
// within [-90, 90], in degrees.
func (c *Client) SetGPS(ctx context.Context, req SetGPSRequest) (*SetGPSResponse, error) {
	var resp SetGPSResponse
	if err := c.core.Do(ctx, "/open/v1/phone/gps/set", req, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

type SetNetworkTypeRequest struct {
	// Cloud Phone ID. Required.
	ID string `json:"id"`
	// Networking. 0-Wi-Fi, 1-Mobile. Only supported on Android 12-16. Required.
	NetType int `json:"netType"`
	// wifi name, limit: 16 characters.
	WiFiID string `json:"wifiId,omitempty"`
}

type SetNetworkTypeResponse struct{}

// SetNetworkType: https://open.geelark.com/api/phone-net-set
func (c *Client) SetNetworkType(ctx context.Context, req SetNetworkTypeRequest) (*SetNetworkTypeResponse, error) {
	var resp SetNetworkTypeResponse
	if err := c.core.Do(ctx, "/open/v1/phone/net/set", req, nil); err != nil {
		return nil, err
	}
	return &resp, nil
}

// GetNetworkSettingsRequest contains no parameters for [Client.GetNetworkSettings].
type GetNetworkSettingsRequest struct{}

type GetNetworkSettingsResponse struct {
	// Accessing blacklisted domains.
	BlackList []string `json:"blackList"`
}

// GetNetworkSettings: https://open.geelark.com/api/phone-netConfig-get
func (c *Client) GetNetworkSettings(ctx context.Context, req GetNetworkSettingsRequest) (*GetNetworkSettingsResponse, error) {
	var resp GetNetworkSettingsResponse
	if err := c.core.Do(ctx, "/open/v1/phone/netConfig/get", req, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

type SetNetworkSettingsRequest struct {
	// Network blacklist. Nil omits the field; an empty non-nil slice sends an empty array to
	// clear the blacklist.
	BlackList []string `json:"blackList,omitempty"`
}

type SetNetworkSettingsResponse struct{}

// SetNetworkSettings: https://open.geelark.com/api/phone-netConfig-set
//
// At most three domains may be blacklisted. Settings take effect immediately. Only Android
// 9, 10, 11, 12, 13, and 15 cloud phones are supported.
func (c *Client) SetNetworkSettings(ctx context.Context, req SetNetworkSettingsRequest) (*SetNetworkSettingsResponse, error) {
	var resp SetNetworkSettingsResponse
	if err := c.core.Do(ctx, "/open/v1/phone/netConfig/set", req, nil); err != nil {
		return nil, err
	}
	return &resp, nil
}
