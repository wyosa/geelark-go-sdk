package adb

import (
	"context"

	"github.com/wyosa/geelark-go-sdk/internal/clientcore"
)

type Client struct {
	core *clientcore.Client
}

func New(core *clientcore.Client) *Client { return &Client{core: core} }

type SetStatusRequest struct {
	IDs  []string `json:"ids"`
	Open bool     `json:"open"`
}

type SetStatusResponse struct{}

// SetStatus calls /open/v1/adb/setStatus.
func (c *Client) SetStatus(ctx context.Context, req SetStatusRequest) (*SetStatusResponse, error) {
	var resp SetStatusResponse
	if err := c.core.Do(ctx, "/open/v1/adb/setStatus", req, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

type GetInfoRequest struct {
	IDs []string `json:"ids"`
}

type GetInfoResponseItemsItem struct {
	Code int    `json:"code"`
	ID   string `json:"id"`
	IP   string `json:"ip"`
	Pwd  string `json:"pwd"`
	Port string `json:"port"`
}

type GetInfoResponse struct {
	Items []GetInfoResponseItemsItem `json:"items"`
}

// GetInfo calls /open/v1/adb/getData.
func (c *Client) GetInfo(ctx context.Context, req GetInfoRequest) (*GetInfoResponse, error) {
	var resp GetInfoResponse
	if err := c.core.Do(ctx, "/open/v1/adb/getData", req, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
