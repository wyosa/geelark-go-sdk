// Package webhook provides cloud callback configuration, callback models, and header
// verification.
package webhook

import (
	"context"

	"github.com/wyosa/geelark-go-sdk/common"
)

type Client struct {
	core common.Transport
}

func New(core common.Transport) *Client { return &Client{core: core} }

type SetURLRequest struct {
	// Callback interface address. Required.
	URL string `json:"url"`
}

type SetURLResponse struct{}

// SetURL: https://open.geelark.com/api/callback-set
func (c *Client) SetURL(ctx context.Context, req SetURLRequest) (*SetURLResponse, error) {
	var resp SetURLResponse
	if err := c.core.Do(ctx, "/open/v1/callback/set", req, nil); err != nil {
		return nil, err
	}
	return &resp, nil
}

// GetURLRequest contains no parameters for [Client.GetURL].
type GetURLRequest struct{}

type GetURLResponse struct {
	// The set callback URL.
	URL string `json:"url"`
}

// GetURL: https://open.geelark.com/api/callback-get
func (c *Client) GetURL(ctx context.Context, req GetURLRequest) (*GetURLResponse, error) {
	var resp GetURLResponse
	if err := c.core.Do(ctx, "/open/v1/callback/get", req, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
