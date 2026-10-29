// Package patch installs Xposed-style patch modules built with the GeeLark patch SDK
// (https://open.geelark.com/api/patch-sdk) on cloud phones running Android 12, 13, or 15.
package patch

import (
	"context"

	"github.com/wyosa/geelark-go-sdk/common"
)

const (
	// QueryByPatchID makes InstallStatus look up PatchID on the cloud phone ID.
	QueryByPatchID = 1
	// QueryByTaskID makes InstallStatus look up the TaskID returned by Install.
	QueryByTaskID = 2
)

const (
	// InstallStatusInstalling means the patch is still being installed.
	InstallStatusInstalling = 0
	// InstallStatusSucceeded means the patch was installed.
	InstallStatusSucceeded = 1
	// InstallStatusFailed means the patch installation failed.
	InstallStatusFailed = 2
)

const (
	// UninstallStatusFailed means the patch was not uninstalled.
	UninstallStatusFailed = 0
	// UninstallStatusSucceeded means the patch was uninstalled.
	UninstallStatusSucceeded = 1
)

type Client struct {
	core common.Transport
}

func New(core common.Transport) *Client { return &Client{core: core} }

type InstallRequest struct {
	// Link to the patch installation file; only apk and xapk are supported. Upload the file
	// with files.Client.Upload first. Required.
	ApkURL string `json:"apkURL"`
	// Cloud phone ID. Required.
	ID string `json:"id"`
}

type InstallResponse struct {
	// Task ID, used with QueryByTaskID to query the installation status.
	TaskID string `json:"taskId"`
}

// Install: https://open.geelark.com/api/phone-patch-install
//
// A patch takes effect immediately, without rebooting the device or restarting apps.
func (c *Client) Install(ctx context.Context, req InstallRequest) (*InstallResponse, error) {
	var resp InstallResponse
	if err := c.core.Do(ctx, "/open/v1/phone/patch/install", req, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

type InstallStatusRequest struct {
	// Query type: QueryByPatchID (1) or QueryByTaskID (2). Required.
	Type int `json:"type"`
	// Patch ID set in the patch project's build.gradle. Required when Type is QueryByPatchID.
	PatchID string `json:"patchId,omitempty"`
	// Cloud phone ID. Required when Type is QueryByPatchID.
	ID string `json:"id,omitempty"`
	// Task ID returned by Install. Required when Type is QueryByTaskID.
	TaskID string `json:"taskId,omitempty"`
}

type InstallStatusResponse struct {
	// Status: InstallStatusInstalling (0), InstallStatusSucceeded (1), or InstallStatusFailed (2).
	Status int `json:"status"`
}

// InstallStatus: https://open.geelark.com/api/phone-patch-install-status
//
// Error code 42007 means the queried task does not exist.
func (c *Client) InstallStatus(ctx context.Context, req InstallStatusRequest) (*InstallStatusResponse, error) {
	var resp InstallStatusResponse
	if err := c.core.Do(ctx, "/open/v1/phone/patch/install/status", req, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

type UninstallRequest struct {
	// Patch ID set in the patch project's build.gradle.
	PatchID string `json:"patchId,omitempty"`
	// Cloud phone ID.
	ID string `json:"id,omitempty"`
}

type UninstallResponse struct {
	// Status: UninstallStatusFailed (0) or UninstallStatusSucceeded (1).
	Status int `json:"status"`
}

// Uninstall: https://open.geelark.com/api/phone-patch-uninstall
func (c *Client) Uninstall(ctx context.Context, req UninstallRequest) (*UninstallResponse, error) {
	var resp UninstallResponse
	if err := c.core.Do(ctx, "/open/v1/phone/patch/uninstall", req, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
