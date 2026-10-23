package browser

import (
	"context"
)

// KernelsRequest contains no parameters for [Client.Kernels].
type KernelsRequest struct{}

type KernelsResponseItem struct {
	// Browser kernel version.
	Kernel string `json:"kernel"`
	// Whether this kernel is already downloaded in the desktop client.
	IsDownloaded bool `json:"isDownloaded"`
}

type KernelsResponse []KernelsResponseItem

// Kernels: https://open.geelark.com/api/browser-get-kernel
//
// Requires GeeLark desktop client version 4.3.8 or later.
func (c *Client) Kernels(ctx context.Context, req KernelsRequest) (*KernelsResponse, error) {
	var resp KernelsResponse
	if err := c.core.Do(ctx, "/api/v1/browser/getKernelsList", req, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

type UpdateKernelsRequest struct {
	// Browser kernel version. Required.
	KernelVersion string `json:"kernel_version"`
}

type UpdateKernelsResponse struct {
	// Kernel download/update status. The official response example uses "complete".
	Status string `json:"status"`
	// Progress value returned by the desktop client. Its range and units are not specified by
	// the documentation.
	Progress int `json:"progress"`
}

// UpdateKernels: https://open.geelark.com/api/browser-update-kernel
//
// Requires GeeLark desktop client version 4.3.8 or later.
func (c *Client) UpdateKernels(ctx context.Context, req UpdateKernelsRequest) (*UpdateKernelsResponse, error) {
	var resp UpdateKernelsResponse
	if err := c.core.Do(ctx, "/api/v1/browser/updateKernels", req, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
