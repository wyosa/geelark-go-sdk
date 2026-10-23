package webhook

import (
	"context"

	"github.com/wyosa/geelark-go-sdk/internal/clientcore"
)

type Client struct {
	core *clientcore.Client
}

func New(core *clientcore.Client) *Client { return &Client{core: core} }

type SetURLRequest struct {
	URL string `json:"url"`
}

type SetURLResponse struct{}

// SetURL calls /open/v1/callback/set.
func (c *Client) SetURL(ctx context.Context, req SetURLRequest) (*SetURLResponse, error) {
	var resp SetURLResponse
	if err := c.core.Do(ctx, "/open/v1/callback/set", req, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

type GetURLRequest struct{}

type GetURLResponse struct {
	URL string `json:"url"`
}

// GetURL calls /open/v1/callback/get.
func (c *Client) GetURL(ctx context.Context, req GetURLRequest) (*GetURLResponse, error) {
	var resp GetURLResponse
	if err := c.core.Do(ctx, "/open/v1/callback/get", req, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
