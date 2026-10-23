// Package app provides cloud phone application installation, configuration, and catalogs.
package app

import (
	"context"

	"github.com/wyosa/geelark-go-sdk/common"
)

type Client struct {
	core common.Transport
}

func New(core common.Transport) *Client { return &Client{core: core} }

type ListInstalledRequest struct {
	// Cloud phone profile ID. Required.
	EnvID string `json:"envId"`
	// Page number, minimum is 1. Required.
	Page int `json:"page"`
	// Number of items per page, minimum is 1, maximum is 100. Required.
	PageSize int `json:"pageSize"`
}

type ListInstalledResponseItemsItem struct {
	// Application icon URL.
	AppIcon string `json:"appIcon"`
	// Application ID.
	AppID string `json:"appId"`
	// Application name.
	AppName string `json:"appName"`
	// Application version ID.
	AppVersionID string `json:"appVersionId"`
	// Installation status: 0-Installing, 1-Installed, 2-Failed, 3-Uninstalling, 4-Uninstalled,
	// 5-Uninstall Failed, others-Not Installed.
	InstallStatus int `json:"installStatus"`
	// Installation time.
	InstallTime string `json:"installTime"`
	// Application package name.
	PackageName string `json:"packageName"`
	// Application version code.
	VersionCode string `json:"versionCode"`
	// Application version name.
	VersionName string `json:"versionName"`
}

type ListInstalledResponse struct {
	// Array of application data.
	Items []ListInstalledResponseItemsItem `json:"items"`
	// Total number of items.
	Total int `json:"total"`
	// Page number.
	Page     int `json:"page"`
	PageSize int `json:"pageSize"`
}

// ListInstalled: https://open.geelark.com/api/app-list
func (c *Client) ListInstalled(ctx context.Context, req ListInstalledRequest) (*ListInstalledResponse, error) {
	var resp ListInstalledResponse
	if err := c.core.Do(ctx, "/open/v1/app/list", req, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

type ListAvailableRequest struct {
	// Search keyword.
	Name string `json:"name,omitempty"`
	// Cloud phone profile ID. Required.
	EnvID        string `json:"envId"`
	GetUploadApp *bool  `json:"getUploadApp,omitempty"`
	// Page number, minimum is 1. Required.
	Page int `json:"page"`
	// Number of items per page, minimum is 1, maximum is 100. Required.
	PageSize int `json:"pageSize"`
}

type ListAvailableResponseItemsItemAppVersionInfoListItem struct {
	// Application version ID.
	ID string `json:"id"`
	// Installation status: 0 - Installing, 1 - Installed, 2 - Installation Failed, 3 -
	// Uninstalling, 4 - Uninstalled, 5 - Uninstallation Failed, others - Not Installed.
	InstallStatus int `json:"installStatus"`
	// App version code.
	VersionCode int `json:"versionCode"`
	// App version name.
	VersionName string `json:"versionName"`
}

type ListAvailableResponseItemsItem struct {
	// App icon URL.
	AppIcon string `json:"appIcon"`
	AppName string `json:"appName"`
	// List of app versions.
	AppVersionInfoList []ListAvailableResponseItemsItemAppVersionInfoListItem `json:"appVersionInfoList"`
	// Application ID.
	ID string `json:"id"`
	// Installation status: 0 - Installing, 1 - Installed, 2 - Installation Failed, 3 -
	// Uninstalling, 4 - Uninstalled, 5 - Uninstallation Failed, others - Not Installed.
	InstallStatus int `json:"installStatus"`
	// App package name.
	PackageName string `json:"packageName"`
}

type ListAvailableResponse struct {
	// List of data items.
	Items []ListAvailableResponseItemsItem `json:"items"`
	// Total number of tasks.
	Total int `json:"total"`
	// Page number.
	Page     int `json:"page"`
	PageSize int `json:"pageSize"`
}

// ListAvailable: https://open.geelark.com/api/app-installable-list
func (c *Client) ListAvailable(ctx context.Context, req ListAvailableRequest) (*ListAvailableResponse, error) {
	var resp ListAvailableResponse
	if err := c.core.Do(ctx, "/open/v1/app/installable/list", req, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

type InstallRequest struct {
	// Cloud phone profile ID. Required.
	EnvID string `json:"envId"`
	// App version ID. Required.
	AppVersionID string `json:"appVersionId"`
}

type InstallResponse struct{}

// Install: https://open.geelark.com/api/app-install
func (c *Client) Install(ctx context.Context, req InstallRequest) (*InstallResponse, error) {
	var resp InstallResponse
	if err := c.core.Do(ctx, "/open/v1/app/install", req, nil); err != nil {
		return nil, err
	}
	return &resp, nil
}

type StartRequest struct {
	// Cloud phone profile ID. Required.
	EnvID string `json:"envId"`
	// App version ID (either appVersionId or packageName must be provided).
	AppVersionID string `json:"appVersionId,omitempty"`
	// Application package name (either appVersionId or packageName must be provided).
	PackageName string `json:"packageName,omitempty"`
}

type StartResponse struct{}

// Start: https://open.geelark.com/api/app-start
func (c *Client) Start(ctx context.Context, req StartRequest) (*StartResponse, error) {
	var resp StartResponse
	if err := c.core.Do(ctx, "/open/v1/app/start", req, nil); err != nil {
		return nil, err
	}
	return &resp, nil
}

type StopRequest struct {
	// Cloud phone profile ID. Required.
	EnvID string `json:"envId"`
	// App version ID (either appVersionId or packageName must be provided).
	AppVersionID string `json:"appVersionId,omitempty"`
	// Application package name (either appVersionId or packageName must be provided).
	PackageName string `json:"packageName,omitempty"`
}

type StopResponse struct{}

// Stop: https://open.geelark.com/api/app-stop
func (c *Client) Stop(ctx context.Context, req StopRequest) (*StopResponse, error) {
	var resp StopResponse
	if err := c.core.Do(ctx, "/open/v1/app/stop", req, nil); err != nil {
		return nil, err
	}
	return &resp, nil
}

type UninstallRequest struct {
	// Cloud phone profile ID. Required.
	EnvID string `json:"envId"`
	// Application package name. Required.
	PackageName string `json:"packageName"`
}

type UninstallResponse struct{}

// Uninstall: https://open.geelark.com/api/app-uninstall
func (c *Client) Uninstall(ctx context.Context, req UninstallRequest) (*UninstallResponse, error) {
	var resp UninstallResponse
	if err := c.core.Do(ctx, "/open/v1/app/uninstall", req, nil); err != nil {
		return nil, err
	}
	return &resp, nil
}

type UploadRequest struct {
	// Application installation file link (supports apk, xapk only). Required.
	FileURL string `json:"fileUrl"`
	// Remarks.
	Desc string `json:"desc,omitempty"`
}

type UploadResponse struct {
	// Task ID for querying upload status.
	TaskID string `json:"taskId"`
}

// Upload: https://open.geelark.com/api/app-upload
func (c *Client) Upload(ctx context.Context, req UploadRequest) (*UploadResponse, error) {
	var resp UploadResponse
	if err := c.core.Do(ctx, "/open/v1/app/upload", req, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

type UploadStatusRequest struct {
	// Task ID. Required.
	TaskID string `json:"taskId"`
}

type UploadStatusResponse struct {
	// Status (0: in the process of being uploaded; 1: uploaded successfully; 2: upload failed;
	// 3: not approved in review).
	Status int `json:"status"`
	// Application name.
	AppName string `json:"appName"`
	// Application icon.
	AppIcon string `json:"appIcon"`
	// Application ID.
	AppID string `json:"appId"`
	// Application version ID.
	VersionID string `json:"versionId"`
}

// UploadStatus: https://open.geelark.com/api/app-upload-status
func (c *Client) UploadStatus(ctx context.Context, req UploadStatusRequest) (*UploadStatusResponse, error) {
	var resp UploadStatusResponse
	if err := c.core.Do(ctx, "/open/v1/app/upload/status", req, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

type ListShopRequest struct {
	// Search keyword.
	Key string `json:"key,omitempty"`
	// Get uploaded apps.
	GetUploadApp *bool `json:"getUploadApp,omitempty"`
	// Page number, minimum is 1. Required.
	Page int `json:"page"`
	// Number of data items per page, minimum is 1, maximum is 200. Required.
	PageSize int `json:"pageSize"`
}

type ListShopResponseItemsItemAppVersionListItem struct {
	// Application version ID.
	ID string `json:"id"`
	// Application version code.
	VersionCode int `json:"versionCode"`
	// Application version name.
	VersionName string `json:"versionName"`
}

type ListShopResponseItemsItem struct {
	// Application icon.
	AppIcon string `json:"appIcon"`
	// Application name.
	AppName string `json:"appName"`
	// Application version information.
	AppVersionList []ListShopResponseItemsItemAppVersionListItem `json:"appVersionList"`
	// Application ID.
	ID string `json:"id"`
}

type ListShopResponse struct {
	// Data array.
	Items []ListShopResponseItemsItem `json:"items"`
	// Total number.
	Total int `json:"total"`
	// Page number.
	Page     int `json:"page"`
	PageSize int `json:"pageSize"`
}

// ListShop: https://open.geelark.com/api/app-shop-list
func (c *Client) ListShop(ctx context.Context, req ListShopRequest) (*ListShopResponse, error) {
	var resp ListShopResponse
	if err := c.core.Do(ctx, "/open/v1/app/shop/list", req, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

type AddToTeamRequest struct {
	// Application id; Please call the “Get Application List” API to obtain it. Required.
	ID string `json:"id"`
	// Version id; Please call the “Get Application List” API to obtain it. Required.
	VersionID string `json:"versionId"`
	// Profile groups allowed for installation. Defaults to all profile groups. "0" indicates
	// no grouping;Please call the query group API first to obtain the group ID.
	InstallGroupIDs []string `json:"installGroupIds,omitempty"`
}

type AddToTeamResponse struct{}

// AddToTeam: https://open.geelark.com/api/app-add
//
// The application is added to team applications and will be installed automatically when
// the cloud phone starts.
func (c *Client) AddToTeam(ctx context.Context, req AddToTeamRequest) (*AddToTeamResponse, error) {
	var resp AddToTeamResponse
	if err := c.core.Do(ctx, "/open/v1/app/add", req, nil); err != nil {
		return nil, err
	}
	return &resp, nil
}

type ListTeamRequest struct {
	// Page number. The minimum value is 1. Required.
	Page int `json:"page"`
	// Number of records per page. The minimum value is 1 and the maximum value is 200.
	// Required.
	PageSize int `json:"pageSize"`
}

type ListTeamResponseItemsItem struct {
	// Team app ID.
	ID          string `json:"id"`
	AppName     string `json:"appName"`
	AppIcon     string `json:"appIcon"`
	VersionID   string `json:"versionId"`
	VersionCode int    `json:"versionCode"`
	VersionName string `json:"versionName"`
	// Whether to install the app automatically. 0: No, 1: Yes.
	Status int `json:"status"`
	// Whether the app was uploaded by the user.
	IsUpload bool `json:"isUpload"`
	// Upload status. 0: Uploading, 1: Upload successful, 2: Upload failed, 3: Review rejected.
	UploadStatus int `json:"uploadStatus"`
	// App authorization. 0: Disabled, 1: Enabled.
	AppAuth int `json:"appAuth"`
	// App root permission. 0: Disabled, 1: Enabled.
	AppRoot int `json:"appRoot"`
	// Whether the app is hidden. false: Not hidden, true: Hidden.
	HideApp bool `json:"hideApp"`
	// IDs of the profile groups where the app can be installed. An empty array means all
	// profile groups, and 0 means ungrouped profiles.
	EnvGroups []string `json:"envGroups"`
}

type ListTeamResponse struct {
	// Total number of records.
	Total int `json:"total"`
	// Page number.
	Page int `json:"page"`
	// Number of records per page.
	PageSize int `json:"pageSize"`
	// App data array.
	Items []ListTeamResponseItemsItem `json:"items"`
}

// ListTeam: https://open.geelark.com/api/app-teamApp-list
func (c *Client) ListTeam(ctx context.Context, req ListTeamRequest) (*ListTeamResponse, error) {
	var resp ListTeamResponse
	if err := c.core.Do(ctx, "/open/v1/app/teamApp/list", req, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

type RemoveTeamRequest struct {
	// Team application ID. Required.
	ID string `json:"id"`
}

type RemoveTeamResponse struct{}

// RemoveTeam: https://open.geelark.com/api/app-remove
func (c *Client) RemoveTeam(ctx context.Context, req RemoveTeamRequest) (*RemoveTeamResponse, error) {
	var resp RemoveTeamResponse
	if err := c.core.Do(ctx, "/open/v1/app/remove", req, nil); err != nil {
		return nil, err
	}
	return &resp, nil
}

type SetAutoInstallRequest struct {
	// Team application ID. Required.
	ID string `json:"id"`
	// Whether to install automatically, 0 for no, 1 for yes. Required.
	Status int `json:"status"`
	// Allowed profile groups for installation, defaults to all profile groups, 0 represents no
	// group.
	InstallGroupIDs []string `json:"installGroupIds,omitempty"`
}

type SetAutoInstallResponse struct{}

// SetAutoInstall: https://open.geelark.com/api/app-setStatus
//
// Enabled applications are installed automatically when cloud phones start.
func (c *Client) SetAutoInstall(ctx context.Context, req SetAutoInstallRequest) (*SetAutoInstallResponse, error) {
	var resp SetAutoInstallResponse
	if err := c.core.Do(ctx, "/open/v1/app/setStatus", req, nil); err != nil {
		return nil, err
	}
	return &resp, nil
}

type SetAuthorizationRequest struct {
	// Team application ID. Required.
	ID string `json:"id"`
	// App authorization 0 disables authorization; 1 enables authorization. Required.
	AppAuth int `json:"appAuth"`
}

type SetAuthorizationResponse struct{}

// SetAuthorization: https://open.geelark.com/api/app-auth-status
//
// The setting grants all app permissions, including location, and only affects newly
// installed applications.
func (c *Client) SetAuthorization(ctx context.Context, req SetAuthorizationRequest) (*SetAuthorizationResponse, error) {
	var resp SetAuthorizationResponse
	if err := c.core.Do(ctx, "/open/v1/app/auth/status", req, nil); err != nil {
		return nil, err
	}
	return &resp, nil
}

type SetRootRequest struct {
	// Team application ID. Required.
	ID string `json:"id"`
	// Operation, 0 off, 1 on. Required.
	Opera int `json:"opera"`
}

type SetRootResponse struct{}

// SetRoot: https://open.geelark.com/api/app-root
func (c *Client) SetRoot(ctx context.Context, req SetRootRequest) (*SetRootResponse, error) {
	var resp SetRootResponse
	if err := c.core.Do(ctx, "/open/v1/app/root", req, nil); err != nil {
		return nil, err
	}
	return &resp, nil
}

type BatchOperateRequest struct {
	// Operation type 1 Startup; 2. Close; 3. Restart; 4 Installation; 5 Uninstalling.
	// Required.
	Action int `json:"action"`
	// Group ID, operate on the opened cloud phone under the specified group, default to all.
	GroupIDs []string `json:"groupIds,omitempty"`
	// Application package name, startup/close/restart/uninstall, choose between
	// packageName/versionId , recommended to use package name.
	PackageName string `json:"packageName,omitempty"`
	// Application version ID, required for installation, can be obtained through
	// the/app/teamApp/list interface.
	VersionID string `json:"versionId,omitempty"`
}

type BatchOperateResponseItemsItem struct {
	// cloud phone id.
	ID string `json:"id"`
	// Error code 1: The application is currently being installed/uninstalled; 2 corresponding
	// apps not installed.
	ErrCode int `json:"errCode"`
}

type BatchOperateResponse struct {
	// Per-phone application operation results.
	Items []BatchOperateResponseItemsItem `json:"items"`
}

// BatchOperate: https://open.geelark.com/api/app-operation-batch
//
// Target cloud phones must already be running.
func (c *Client) BatchOperate(ctx context.Context, req BatchOperateRequest) (*BatchOperateResponse, error) {
	var resp BatchOperateResponse
	if err := c.core.Do(ctx, "/open/v1/app/operation/batch", req, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

type SetKeepAliveRequest struct {
	// Team application ID. Required.
	ID string `json:"id"`
	// Operation, 0 off, 1 on. Required.
	Opera int `json:"opera"`
}

type SetKeepAliveResponse struct{}

// SetKeepAlive: https://open.geelark.com/api/app-setKeepAlive
//
// Requires the Pro plan and Android 12, 13, 14, or 15. At most one application can be kept
// alive. On Android 12, 13, and 15, restart an already running app after enabling this
// setting. Android 14 applies the setting immediately.
func (c *Client) SetKeepAlive(ctx context.Context, req SetKeepAliveRequest) (*SetKeepAliveResponse, error) {
	var resp SetKeepAliveResponse
	if err := c.core.Do(ctx, "/open/v1/app/setKeepAlive", req, nil); err != nil {
		return nil, err
	}
	return &resp, nil
}

type SetAutoStartRequest struct {
	// Team application ID. Required.
	ID string `json:"id"`
	// Operation, 0 off, 1 on. Required.
	Opera int `json:"opera"`
}

type SetAutoStartResponse struct{}

// SetAutoStart: https://open.geelark.com/api/app-setAutoStart
func (c *Client) SetAutoStart(ctx context.Context, req SetAutoStartRequest) (*SetAutoStartResponse, error) {
	var resp SetAutoStartResponse
	if err := c.core.Do(ctx, "/open/v1/app/setAutoStart", req, nil); err != nil {
		return nil, err
	}
	return &resp, nil
}

type SetHiddenRequest struct {
	// Team app ID; call the "Get Team App List" API to obtain it. Required.
	ID string `json:"id"`
	// Operation type; 1: hide the app, 0: stop hiding the app. Required.
	Opera int `json:"opera"`
}

type SetHiddenResponse struct{}

// SetHidden: https://open.geelark.com/api/app-set-hide
//
// App hiding is supported on Android 12, 13, 14, 15, and 16.
func (c *Client) SetHidden(ctx context.Context, req SetHiddenRequest) (*SetHiddenResponse, error) {
	var resp SetHiddenResponse
	if err := c.core.Do(ctx, "/open/v1/app/setHide", req, nil); err != nil {
		return nil, err
	}
	return &resp, nil
}
