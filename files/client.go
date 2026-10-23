// Package files provides temporary file uploads and file transfer to cloud phones.
package files

import (
	"context"

	"github.com/wyosa/geelark-go-sdk/common"
)

type Client struct {
	core common.UploadTransport
}

func New(core common.UploadTransport) *Client { return &Client{core: core} }

type GetUploadURLRequest struct {
	// File type, Currently supported: "jpg", "jpeg", "png", "gif", "bmp", "webp","heif",
	// "heic", "mp4", "webm","xml", "apk", "xapk". Required.
	FileType string `json:"fileType"`
}

type GetUploadURLResponse struct {
	// URL for uploading the file (valid for 30 minutes).
	UploadURL string `json:"uploadUrl"`
	// URL to access the resource.
	ResourceURL string `json:"resourceUrl"`
}

// GetUploadURL: https://open.geelark.com/api/upload-getUrl
//
// The upload URL expires after 30 minutes; the uploaded temporary file expires after three
// days. Use Library uploads for the documented longer retention period.
func (c *Client) GetUploadURL(ctx context.Context, req GetUploadURLRequest) (*GetUploadURLResponse, error) {
	var resp GetUploadURLResponse
	if err := c.core.Do(ctx, "/open/v1/upload/getUrl", req, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

type UploadToPhoneRequest struct {
	// Cloud phone ID. Required.
	ID string `json:"id"`
	// File URL. Required.
	FileURL string `json:"fileUrl"`
}

type UploadToPhoneResponse struct {
	TaskID string `json:"taskId"`
}

// UploadToPhone: https://open.geelark.com/api/phone-uploadFile
//
// Start the cloud phone first. The file is placed in its Downloads folder. The returned
// task ID represents an asynchronous upload; query UploadResult or process its callback.
func (c *Client) UploadToPhone(ctx context.Context, req UploadToPhoneRequest) (*UploadToPhoneResponse, error) {
	var resp UploadToPhoneResponse
	if err := c.core.Do(ctx, "/open/v1/phone/uploadFile", req, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

type UploadResultRequest struct {
	// Task ID. Required.
	TaskID string `json:"taskId"`
}

type UploadResultResponse struct {
	// 0: Failed to retrieve; 1: Uploading; 2: Upload successful; 3: Upload failed.
	Status int `json:"status"`
}

// UploadResult: https://open.geelark.com/api/phone-uploadFile-result
//
// Retrieve the result within one hour of starting the upload task; retrieval fails after
// the result expires.
func (c *Client) UploadResult(ctx context.Context, req UploadResultRequest) (*UploadResultResponse, error) {
	var resp UploadResultResponse
	if err := c.core.Do(ctx, "/open/v1/phone/uploadFile/result", req, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

type UploadKeyboxRequest struct {
	// Cloud phone ID. Required.
	ID string `json:"id"`
	// File URL. Required.
	FileURL string `json:"fileUrl"`
}

type UploadKeyboxResponse struct {
	TaskID string `json:"taskId"`
}

// UploadKeybox: https://open.geelark.com/api/phone-keyboxUpload
//
// Only Android 12, 13, and 15 are supported. On Android 15, first update Google Play and
// GMS, sign in to a valid Google account, and install Play Integrity API Checker from
// Google Play.
func (c *Client) UploadKeybox(ctx context.Context, req UploadKeyboxRequest) (*UploadKeyboxResponse, error) {
	var resp UploadKeyboxResponse
	if err := c.core.Do(ctx, "/open/v1/phone/keyboxUpload", req, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

type KeyboxResultRequest struct {
	// Task ID. Required.
	TaskID string `json:"taskId"`
}

type KeyboxResultResponse struct {
	// 0: Uploading; 1: Upload successful; 2: Upload failed.
	Status int `json:"status"`
}

// KeyboxResult: https://open.geelark.com/api/phone-keyboxUpload-result
func (c *Client) KeyboxResult(ctx context.Context, req KeyboxResultRequest) (*KeyboxResultResponse, error) {
	var resp KeyboxResultResponse
	if err := c.core.Do(ctx, "/open/v1/phone/keyboxUpload/result", req, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
