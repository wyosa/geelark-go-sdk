package livestream

import (
	"context"

	"github.com/wyosa/geelark-go-sdk/internal/clientcore"
)

type Client struct {
	core *clientcore.Client
}

func New(core *clientcore.Client) *Client { return &Client{core: core} }

type StartRequest struct {
	ID      string `json:"id"`
	FileURL string `json:"fileUrl"`
	IsLoop  int    `json:"isLoop"`
}

type StartResponse struct {
	TaskID string `json:"taskId"`
}

// Start calls /open/v1/phone/videoPush/upload.
func (c *Client) Start(ctx context.Context, req StartRequest) (*StartResponse, error) {
	var resp StartResponse
	if err := c.core.Do(ctx, "/open/v1/phone/videoPush/upload", req, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

type ResultRequest struct {
	ID string `json:"id"`
}

type ResultResponse struct {
	Status int `json:"status"`
}

// Result calls /open/v1/phone/videoPush/upload/result.
func (c *Client) Result(ctx context.Context, req ResultRequest) (*ResultResponse, error) {
	var resp ResultResponse
	if err := c.core.Do(ctx, "/open/v1/phone/videoPush/upload/result", req, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

type StopRequest struct {
	ID string `json:"id"`
}

type StopResponse struct{}

// Stop calls /open/v1/phone/videoPush/stop.
func (c *Client) Stop(ctx context.Context, req StopRequest) (*StopResponse, error) {
	var resp StopResponse
	if err := c.core.Do(ctx, "/open/v1/phone/videoPush/stop", req, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
