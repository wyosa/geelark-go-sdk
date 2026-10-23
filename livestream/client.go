// Package livestream provides cloud phone live streaming controls.
package livestream

import (
	"context"

	"github.com/wyosa/geelark-go-sdk/common"
)

type Client struct {
	core common.Transport
}

func New(core common.Transport) *Client { return &Client{core: core} }

type StartRequest struct {
	// Cloud phone ID. Required.
	ID string `json:"id"`
	// File URL. Required.
	FileURL string `json:"fileUrl"`
	// Whether to stream the video in a loop (1 - Yes, 0 - No). Required.
	IsLoop int `json:"isLoop"`
}

type StartResponse struct {
	TaskID string `json:"taskId"`
}

// Start: https://open.geelark.com/api/phone-videoPush
//
// Supports Android 12, 13, 14, 15, and 16. MP4 video files are recommended.
func (c *Client) Start(ctx context.Context, req StartRequest) (*StartResponse, error) {
	var resp StartResponse
	if err := c.core.Do(ctx, "/open/v1/phone/videoPush/upload", req, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

type ResultRequest struct {
	// Task ID (the taskId returned by the video streaming API). Required.
	ID string `json:"id"`
}

type ResultResponse struct {
	// Streaming status: 0 - In Progress; 1 - Success; 2 - Failed.
	Status int `json:"status"`
}

// Result: https://open.geelark.com/api/phone-videoPush-result
func (c *Client) Result(ctx context.Context, req ResultRequest) (*ResultResponse, error) {
	var resp ResultResponse
	if err := c.core.Do(ctx, "/open/v1/phone/videoPush/upload/result", req, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

type StopRequest struct {
	// Cloud phone ID. Required.
	ID string `json:"id"`
}

type StopResponse struct{}

// Stop: https://open.geelark.com/api/phone-videoPush-stop
func (c *Client) Stop(ctx context.Context, req StopRequest) (*StopResponse, error) {
	var resp StopResponse
	if err := c.core.Do(ctx, "/open/v1/phone/videoPush/stop", req, nil); err != nil {
		return nil, err
	}
	return &resp, nil
}
