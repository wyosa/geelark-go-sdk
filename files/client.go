package files

import (
	"context"

	"github.com/wyosa/geelark-go-sdk/internal/clientcore"
)

type Client struct {
	core *clientcore.Client
}

func New(core *clientcore.Client) *Client { return &Client{core: core} }

type GetUploadURLRequest struct {
	FileType string `json:"fileType"`
}

type GetUploadURLResponse struct {
	UploadURL   string `json:"uploadUrl"`
	ResourceURL string `json:"resourceUrl"`
}

// GetUploadURL calls /open/v1/upload/getUrl.
func (c *Client) GetUploadURL(ctx context.Context, req GetUploadURLRequest) (*GetUploadURLResponse, error) {
	var resp GetUploadURLResponse
	if err := c.core.Do(ctx, "/open/v1/upload/getUrl", req, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

type UploadToPhoneRequest struct {
	ID      string `json:"id"`
	FileURL string `json:"fileUrl"`
}

type UploadToPhoneResponse struct {
	TaskID string `json:"taskId"`
}

// UploadToPhone calls /open/v1/phone/uploadFile.
func (c *Client) UploadToPhone(ctx context.Context, req UploadToPhoneRequest) (*UploadToPhoneResponse, error) {
	var resp UploadToPhoneResponse
	if err := c.core.Do(ctx, "/open/v1/phone/uploadFile", req, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

type UploadResultRequest struct {
	TaskID string `json:"taskId"`
}

type UploadResultResponse struct {
	Status int `json:"status"`
}

// UploadResult calls /open/v1/phone/uploadFile/result.
func (c *Client) UploadResult(ctx context.Context, req UploadResultRequest) (*UploadResultResponse, error) {
	var resp UploadResultResponse
	if err := c.core.Do(ctx, "/open/v1/phone/uploadFile/result", req, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

type UploadKeyboxRequest struct {
	ID      string `json:"id"`
	FileURL string `json:"fileUrl"`
}

type UploadKeyboxResponse struct {
	TaskID string `json:"taskId"`
}

// UploadKeybox calls /open/v1/phone/keyboxUpload.
func (c *Client) UploadKeybox(ctx context.Context, req UploadKeyboxRequest) (*UploadKeyboxResponse, error) {
	var resp UploadKeyboxResponse
	if err := c.core.Do(ctx, "/open/v1/phone/keyboxUpload", req, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

type KeyboxResultRequest struct {
	TaskID string `json:"taskId"`
}

type KeyboxResultResponse struct {
	Status int `json:"status"`
}

// KeyboxResult calls /open/v1/phone/keyboxUpload/result.
func (c *Client) KeyboxResult(ctx context.Context, req KeyboxResultRequest) (*KeyboxResultResponse, error) {
	var resp KeyboxResultResponse
	if err := c.core.Do(ctx, "/open/v1/phone/keyboxUpload/result", req, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
