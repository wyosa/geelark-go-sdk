// Package adb provides ADB access to cloud phones.
package adb

import (
	"context"

	"github.com/wyosa/geelark-go-sdk/common"
)

type Client struct {
	core common.Transport
}

func New(core common.Transport) *Client { return &Client{core: core} }

type SetStatusRequest struct {
	// Array of cloud phone profile IDs, maximum 200. Required.
	IDs []string `json:"ids"`
	// Open/Close. Required.
	Open bool `json:"open"`
}

type SetStatusResponse struct{}

// SetStatus: https://open.geelark.com/api/adb-setStatus
//
// Start the cloud phone first. ADB is supported on Android 9, 11, 12, 13, 14, 15, and 16.
// Enabling ADB is asynchronous: GeeLark recommends waiting about three seconds before
// calling GetInfo for the connection port and password.
func (c *Client) SetStatus(ctx context.Context, req SetStatusRequest) (*SetStatusResponse, error) {
	var resp SetStatusResponse
	if err := c.core.Do(ctx, "/open/v1/adb/setStatus", req, nil); err != nil {
		return nil, err
	}
	return &resp, nil
}

type GetInfoRequest struct {
	// Array of cloud phone IDs, maximum 200. Required.
	IDs []string `json:"ids"`
}

type GetInfoResponseItemsItem struct {
	// Error code: 0 indicates success; for other codes, refer to the error code table.
	Code int `json:"code"`
	// Cloud phone ID.
	ID string `json:"id"`
	// Connection IP.
	IP string `json:"ip"`
	// The password for glogin.
	Pwd  string `json:"pwd"`
	Port string `json:"port"`
}

type GetInfoResponse struct {
	// Per-phone ADB connection information.
	Items []GetInfoResponseItemsItem `json:"items"`
}

// GetInfo: https://open.geelark.com/api/adb-getData
func (c *Client) GetInfo(ctx context.Context, req GetInfoRequest) (*GetInfoResponse, error) {
	var resp GetInfoResponse
	if err := c.core.Do(ctx, "/open/v1/adb/getData", req, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
