package app

import (
	"context"

	"github.com/wyosa/geelark-go-sdk/internal/clientcore"
)

type Client struct {
	core *clientcore.Client
}

func New(core *clientcore.Client) *Client { return &Client{core: core} }

type ListInstalledRequest struct {
	EnvID    string `json:"envId"`
	Page     int    `json:"page"`
	PageSize int    `json:"pageSize"`
}

type ListInstalledResponseItemsItem struct {
	AppIcon       string `json:"appIcon"`
	AppID         string `json:"appId"`
	AppName       string `json:"appName"`
	AppVersionID  string `json:"appVersionId"`
	InstallStatus int    `json:"installStatus"`
	InstallTime   string `json:"installTime"`
	PackageName   string `json:"packageName"`
	VersionCode   string `json:"versionCode"`
	VersionName   string `json:"versionName"`
}

type ListInstalledResponse struct {
	Items    []ListInstalledResponseItemsItem `json:"items"`
	Total    int                              `json:"total"`
	Page     int                              `json:"page"`
	PageSize int                              `json:"pageSize"`
}

// ListInstalled calls /open/v1/app/list.
func (c *Client) ListInstalled(ctx context.Context, req ListInstalledRequest) (*ListInstalledResponse, error) {
	var resp ListInstalledResponse
	if err := c.core.Do(ctx, "/open/v1/app/list", req, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

type ListAvailableRequest struct {
	Name         string `json:"name,omitempty"`
	EnvID        string `json:"envId"`
	GetUploadApp *bool  `json:"getUploadApp,omitempty"`
	Page         int    `json:"page"`
	PageSize     int    `json:"pageSize"`
}

type ListAvailableResponseItemsItemAppVersionInfoListItem struct {
	ID            string `json:"id"`
	InstallStatus int    `json:"installStatus"`
	VersionCode   int    `json:"versionCode"`
	VersionName   string `json:"versionName"`
}

type ListAvailableResponseItemsItem struct {
	AppIcon            string                                                 `json:"appIcon"`
	AppName            string                                                 `json:"appName"`
	AppVersionInfoList []ListAvailableResponseItemsItemAppVersionInfoListItem `json:"appVersionInfoList"`
	ID                 string                                                 `json:"id"`
	InstallStatus      int                                                    `json:"installStatus"`
	PackageName        string                                                 `json:"packageName"`
}

type ListAvailableResponse struct {
	Items    []ListAvailableResponseItemsItem `json:"items"`
	Total    int                              `json:"total"`
	Page     int                              `json:"page"`
	PageSize int                              `json:"pageSize"`
}

// ListAvailable calls /open/v1/app/installable/list.
func (c *Client) ListAvailable(ctx context.Context, req ListAvailableRequest) (*ListAvailableResponse, error) {
	var resp ListAvailableResponse
	if err := c.core.Do(ctx, "/open/v1/app/installable/list", req, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

type InstallRequest struct {
	EnvID        string `json:"envId"`
	AppVersionID string `json:"appVersionId"`
}

type InstallResponse struct{}

// Install calls /open/v1/app/install.
func (c *Client) Install(ctx context.Context, req InstallRequest) (*InstallResponse, error) {
	var resp InstallResponse
	if err := c.core.Do(ctx, "/open/v1/app/install", req, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

type StartRequest struct {
	EnvID        string `json:"envId"`
	AppVersionID string `json:"appVersionId,omitempty"`
	PackageName  string `json:"packageName,omitempty"`
}

type StartResponse struct{}

// Start calls /open/v1/app/start.
func (c *Client) Start(ctx context.Context, req StartRequest) (*StartResponse, error) {
	var resp StartResponse
	if err := c.core.Do(ctx, "/open/v1/app/start", req, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

type StopRequest struct {
	EnvID        string `json:"envId"`
	AppVersionID string `json:"appVersionId,omitempty"`
	PackageName  string `json:"packageName,omitempty"`
}

type StopResponse struct{}

// Stop calls /open/v1/app/stop.
func (c *Client) Stop(ctx context.Context, req StopRequest) (*StopResponse, error) {
	var resp StopResponse
	if err := c.core.Do(ctx, "/open/v1/app/stop", req, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

type UninstallRequest struct {
	EnvID       string `json:"envId"`
	PackageName string `json:"packageName"`
}

type UninstallResponse struct{}

// Uninstall calls /open/v1/app/uninstall.
func (c *Client) Uninstall(ctx context.Context, req UninstallRequest) (*UninstallResponse, error) {
	var resp UninstallResponse
	if err := c.core.Do(ctx, "/open/v1/app/uninstall", req, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

type UploadRequest struct {
	FileURL string `json:"fileUrl"`
	Desc    string `json:"desc,omitempty"`
}

type UploadResponse struct {
	TaskID string `json:"taskId"`
}

// Upload calls /open/v1/app/upload.
func (c *Client) Upload(ctx context.Context, req UploadRequest) (*UploadResponse, error) {
	var resp UploadResponse
	if err := c.core.Do(ctx, "/open/v1/app/upload", req, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

type UploadStatusRequest struct {
	TaskID string `json:"taskId"`
}

type UploadStatusResponse struct {
	Status    int    `json:"status"`
	AppName   string `json:"appName"`
	AppIcon   string `json:"appIcon"`
	AppID     string `json:"appId"`
	VersionID string `json:"versionId"`
}

// UploadStatus calls /open/v1/app/upload/status.
func (c *Client) UploadStatus(ctx context.Context, req UploadStatusRequest) (*UploadStatusResponse, error) {
	var resp UploadStatusResponse
	if err := c.core.Do(ctx, "/open/v1/app/upload/status", req, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

type ListShopRequest struct {
	Key          string `json:"key,omitempty"`
	GetUploadApp *bool  `json:"getUploadApp,omitempty"`
	Page         int    `json:"page"`
	PageSize     int    `json:"pageSize"`
}

type ListShopResponseItemsItemAppVersionListItem struct {
	ID          string `json:"id"`
	VersionCode int    `json:"versionCode"`
	VersionName string `json:"versionName"`
}

type ListShopResponseItemsItem struct {
	AppIcon        string                                        `json:"appIcon"`
	AppName        string                                        `json:"appName"`
	AppVersionList []ListShopResponseItemsItemAppVersionListItem `json:"appVersionList"`
	ID             string                                        `json:"id"`
}

type ListShopResponse struct {
	Items    []ListShopResponseItemsItem `json:"items"`
	Total    int                         `json:"total"`
	Page     int                         `json:"page"`
	PageSize int                         `json:"pageSize"`
}

// ListShop calls /open/v1/app/shop/list.
func (c *Client) ListShop(ctx context.Context, req ListShopRequest) (*ListShopResponse, error) {
	var resp ListShopResponse
	if err := c.core.Do(ctx, "/open/v1/app/shop/list", req, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

type AddToTeamRequest struct {
	ID              string   `json:"id"`
	VersionID       string   `json:"versionId"`
	InstallGroupIDs []string `json:"installGroupIds,omitempty"`
}

type AddToTeamResponse struct{}

// AddToTeam calls /open/v1/app/add.
func (c *Client) AddToTeam(ctx context.Context, req AddToTeamRequest) (*AddToTeamResponse, error) {
	var resp AddToTeamResponse
	if err := c.core.Do(ctx, "/open/v1/app/add", req, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

type ListTeamRequest struct {
	Page     int `json:"page"`
	PageSize int `json:"pageSize"`
}

type ListTeamResponseItemsItem struct {
	ID           string   `json:"id"`
	AppName      string   `json:"appName"`
	AppIcon      string   `json:"appIcon"`
	VersionID    string   `json:"versionId"`
	VersionCode  int      `json:"versionCode"`
	VersionName  string   `json:"versionName"`
	Status       int      `json:"status"`
	IsUpload     bool     `json:"isUpload"`
	UploadStatus int      `json:"uploadStatus"`
	AppAuth      int      `json:"appAuth"`
	AppRoot      int      `json:"appRoot"`
	HideApp      bool     `json:"hideApp"`
	EnvGroups    []string `json:"envGroups"`
}

type ListTeamResponse struct {
	Total    int                         `json:"total"`
	Page     int                         `json:"page"`
	PageSize int                         `json:"pageSize"`
	Items    []ListTeamResponseItemsItem `json:"items"`
}

// ListTeam calls /open/v1/app/teamApp/list.
func (c *Client) ListTeam(ctx context.Context, req ListTeamRequest) (*ListTeamResponse, error) {
	var resp ListTeamResponse
	if err := c.core.Do(ctx, "/open/v1/app/teamApp/list", req, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

type RemoveTeamRequest struct {
	ID string `json:"id"`
}

type RemoveTeamResponse struct{}

// RemoveTeam calls /open/v1/app/remove.
func (c *Client) RemoveTeam(ctx context.Context, req RemoveTeamRequest) (*RemoveTeamResponse, error) {
	var resp RemoveTeamResponse
	if err := c.core.Do(ctx, "/open/v1/app/remove", req, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

type SetAutoInstallRequest struct {
	ID              string   `json:"id"`
	Status          int      `json:"status"`
	InstallGroupIDs []string `json:"installGroupIds,omitempty"`
}

type SetAutoInstallResponse struct{}

// SetAutoInstall calls /open/v1/app/setStatus.
func (c *Client) SetAutoInstall(ctx context.Context, req SetAutoInstallRequest) (*SetAutoInstallResponse, error) {
	var resp SetAutoInstallResponse
	if err := c.core.Do(ctx, "/open/v1/app/setStatus", req, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

type SetAuthorizationRequest struct {
	ID      string `json:"id"`
	AppAuth int    `json:"appAuth"`
}

type SetAuthorizationResponse struct{}

// SetAuthorization calls /open/v1/app/auth/status.
func (c *Client) SetAuthorization(ctx context.Context, req SetAuthorizationRequest) (*SetAuthorizationResponse, error) {
	var resp SetAuthorizationResponse
	if err := c.core.Do(ctx, "/open/v1/app/auth/status", req, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

type SetRootRequest struct {
	ID    string `json:"id"`
	Opera int    `json:"opera"`
}

type SetRootResponse struct{}

// SetRoot calls /open/v1/app/root.
func (c *Client) SetRoot(ctx context.Context, req SetRootRequest) (*SetRootResponse, error) {
	var resp SetRootResponse
	if err := c.core.Do(ctx, "/open/v1/app/root", req, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

type BatchOperateRequest struct {
	Action      int      `json:"action"`
	GroupIDs    []string `json:"groupIds,omitempty"`
	PackageName string   `json:"packageName,omitempty"`
	VersionID   string   `json:"versionId,omitempty"`
}

type BatchOperateResponseItemsItem struct {
	ID      string `json:"id"`
	ErrCode int    `json:"errCode"`
}

type BatchOperateResponse struct {
	Items []BatchOperateResponseItemsItem `json:"items"`
}

// BatchOperate calls /open/v1/app/operation/batch.
func (c *Client) BatchOperate(ctx context.Context, req BatchOperateRequest) (*BatchOperateResponse, error) {
	var resp BatchOperateResponse
	if err := c.core.Do(ctx, "/open/v1/app/operation/batch", req, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

type SetKeepAliveRequest struct {
	ID    string `json:"id"`
	Opera int    `json:"opera"`
}

type SetKeepAliveResponse struct{}

// SetKeepAlive calls /open/v1/app/setKeepAlive.
func (c *Client) SetKeepAlive(ctx context.Context, req SetKeepAliveRequest) (*SetKeepAliveResponse, error) {
	var resp SetKeepAliveResponse
	if err := c.core.Do(ctx, "/open/v1/app/setKeepAlive", req, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

type SetAutoStartRequest struct {
	ID    string `json:"id"`
	Opera int    `json:"opera"`
}

type SetAutoStartResponse struct{}

// SetAutoStart calls /open/v1/app/setAutoStart.
func (c *Client) SetAutoStart(ctx context.Context, req SetAutoStartRequest) (*SetAutoStartResponse, error) {
	var resp SetAutoStartResponse
	if err := c.core.Do(ctx, "/open/v1/app/setAutoStart", req, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

type SetHiddenRequest struct {
	ID    string `json:"id"`
	Opera int    `json:"opera"`
}

type SetHiddenResponse struct{}

// SetHidden calls /open/v1/app/setHide.
func (c *Client) SetHidden(ctx context.Context, req SetHiddenRequest) (*SetHiddenResponse, error) {
	var resp SetHiddenResponse
	if err := c.core.Do(ctx, "/open/v1/app/setHide", req, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
