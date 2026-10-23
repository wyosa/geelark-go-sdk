package cloudphone

import (
	"context"
)

type ScreenshotRequest struct {
	// clound phone id. Required.
	ID string `json:"id"`
}

type ScreenshotResponse struct {
	TaskID string `json:"taskId"`
}

// Screenshot: https://open.geelark.com/api/phone-screenShot
func (c *Client) Screenshot(ctx context.Context, req ScreenshotRequest) (*ScreenshotResponse, error) {
	var resp ScreenshotResponse
	if err := c.core.Do(ctx, "/open/v1/phone/screenShot", req, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

type ScreenshotResultRequest struct {
	// task id. Required.
	TaskID string `json:"taskId"`
}

type ScreenshotResultResponse struct {
	// 0 Acquisition failed；1 In progress；2 Execution succeeded；3 Execution failed.
	Status int `json:"status"`
	// screen shot download link.
	DownloadLink string `json:"downloadLink"`
}

// ScreenshotResult: https://open.geelark.com/api/phone-screenShot-result
//
// Retrieve the result within 30 minutes after requesting the screenshot; retrieval fails
// after the result expires.
func (c *Client) ScreenshotResult(ctx context.Context, req ScreenshotResultRequest) (*ScreenshotResultResponse, error) {
	var resp ScreenshotResultResponse
	if err := c.core.Do(ctx, "/open/v1/phone/screenShot/result", req, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

type UIDumpRequest struct {
	// cloud phone id. Required.
	ID string `json:"id"`
}

type UIDumpResponse struct {
	TaskID string `json:"taskId"`
}

// UIDump: https://open.geelark.com/api/get-ui-dump
func (c *Client) UIDump(ctx context.Context, req UIDumpRequest) (*UIDumpResponse, error) {
	var resp UIDumpResponse
	if err := c.core.Do(ctx, "/open/v1/phone/uiDump", req, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

type UIDumpResultRequest struct {
	// task id. Required.
	TaskID string `json:"taskId"`
}

type UIDumpResultResponse struct {
	// 0 Acquisition failed；1 In progress；2 Execution succeeded；3 Execution failed.
	Status int `json:"status"`
	// UI layout xml file download link.
	DownloadLink string `json:"downloadLink"`
}

// UIDumpResult: https://open.geelark.com/api/get-ui-dump-result
//
// Retrieve the result within 30 minutes after requesting the UI dump; retrieval fails
// after the result expires.
func (c *Client) UIDumpResult(ctx context.Context, req UIDumpResultRequest) (*UIDumpResultResponse, error) {
	var resp UIDumpResultResponse
	if err := c.core.Do(ctx, "/open/v1/phone/uiDump/result", req, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
