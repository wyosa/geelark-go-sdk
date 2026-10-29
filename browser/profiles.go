package browser

import (
	"context"
)

type DeleteRequest struct {
	// Browser ID array, up to 100. Required.
	EnvIDs []string `json:"envIds"`
}

type DeleteResponse struct {
	// IDs of profiles currently in use.
	BusyIDs []string `json:"busyIds"`
	// IDs of profiles handling exceptions.
	ServerErrIDs []string `json:"serverErrIds"`
	// IDs of profiles that were successfully deleted.
	SuccessIDs []string `json:"successIds"`
}

// Delete: https://open.geelark.com/api/delete-browser
func (c *Client) Delete(ctx context.Context, req DeleteRequest) (*DeleteResponse, error) {
	var resp DeleteResponse
	if err := c.core.Do(ctx, "/api/v1/browser/delete", req, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

type TransferRequest struct {
	// Target user account. Required.
	Username string `json:"username"`
	// Environment IDs to transfer. The maximum length of the array is 200. Entries exceeding
	// 200 will be ignored. Required.
	EnvIDs []string `json:"envIds"`
	// Optional transfer option parameters: name: environment name, proxy: proxy, tag: tag,
	// remark: remark, files: files.
	TransferOption []string `json:"transferOption,omitempty"`
}

type TransferResponse struct {
	// Number of successful transfers.
	SuccessCount int `json:"successCount"`
	// Number of failed transfers.
	FailCount int `json:"failCount"`
	// IDs of environments where transfers failed, including those in use or non-existent
	// environments.
	FailEnvIDs []string `json:"failEnvIds"`
}

// Transfer: https://open.geelark.com/api/transfer-browsers
func (c *Client) Transfer(ctx context.Context, req TransferRequest) (*TransferResponse, error) {
	var resp TransferResponse
	if err := c.core.Do(ctx, "/api/v1/browser/transfer", req, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

type MoveGroupRequest struct {
	// Browser IDs, up to 100. Required.
	EnvIDs []string `json:"envIds"`
	// Target group ID. Use "0", not an empty string, to move browsers to ungrouped. Required.
	GroupID string `json:"groupId"`
}

type MoveGroupResponse struct{}

// MoveGroup: https://open.geelark.com/api/browser-group-move
func (c *Client) MoveGroup(ctx context.Context, req MoveGroupRequest) (*MoveGroupResponse, error) {
	var resp MoveGroupResponse
	if err := c.core.Do(ctx, "/api/v1/browser/moveGroup", req, nil); err != nil {
		return nil, err
	}
	return &resp, nil
}

type CloneRequest struct {
	// Browser ID to clone. Required.
	EnvID string `json:"envId"`
	// Number of clones, range 1-100. Required.
	Amount int `json:"amount"`
	// Target group ID, if not specified, it will be placed in an ungrouped group.
	GroupID string `json:"groupId,omitempty"`
	// Whether to clone the name.
	CloneName *bool `json:"cloneName,omitempty"`
	// Whether to clone the remark.
	CloneRemark *bool `json:"cloneRemark,omitempty"`
	// Whether to clone the tag.
	CloneTag *bool `json:"cloneTag,omitempty"`
	// Whether to clone the proxy.
	CloneProxy *bool `json:"cloneProxy,omitempty"`
	// Whether to clone the cookie.
	CloneCookie *bool `json:"cloneCookie,omitempty"`
	// Whether to clone the account information.
	CloneAccount *bool `json:"cloneAccount,omitempty"`
}

type CloneResponse struct {
	// cloned browser ID.
	IDs []string `json:"ids"`
}

// Clone: https://open.geelark.com/api/browser-clone
func (c *Client) Clone(ctx context.Context, req CloneRequest) (*CloneResponse, error) {
	var resp CloneResponse
	if err := c.core.Do(ctx, "/api/v1/browser/clone", req, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
