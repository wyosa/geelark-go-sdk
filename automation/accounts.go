package automation

import (
	"context"
)

type GoogleLoginRequest struct {
	// Task name, up to 128 characters.
	Name string `json:"name,omitempty"`
	// Remarks, up to 200 characters.
	Remark string `json:"remark,omitempty"`
	// Scheduled execution time as a Unix timestamp in seconds. Required.
	ScheduleAt int64 `json:"scheduleAt"`
	// Cloud phone ID. Required.
	ID string `json:"id"`
	// Email, up to 64 characters. Required.
	Email string `json:"email"`
	// Password, up to 64 characters. Required.
	Password string `json:"password"`
	// 2fa code.
	Code2fa string `json:"code2fa,omitempty"`
}

type GoogleLoginResponse struct {
	// ID assigned to the asynchronous task.
	TaskID string `json:"taskId"`
}

// GoogleLogin: https://open.geelark.com/api/rpa-task-googleLogin
func (c *Client) GoogleLogin(ctx context.Context, req GoogleLoginRequest) (*GoogleLoginResponse, error) {
	var resp GoogleLoginResponse
	if err := c.core.Do(ctx, "/open/v1/rpa/task/googleLogin", req, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

type GoogleAppDownloadRequest struct {
	// Task name, up to 128 characters.
	Name string `json:"name,omitempty"`
	// Remarks, up to 200 characters.
	Remark string `json:"remark,omitempty"`
	// Scheduled execution time as a Unix timestamp in seconds. Required.
	ScheduleAt int64 `json:"scheduleAt"`
	// Cloud phone ID. Required.
	ID string `json:"id"`
	// app name. Required.
	AppName string `json:"appName"`
}

type GoogleAppDownloadResponse struct {
	// ID assigned to the asynchronous task.
	TaskID string `json:"taskId"`
}

// GoogleAppDownload: https://open.geelark.com/api/rpa-task-googleAppDownload
func (c *Client) GoogleAppDownload(ctx context.Context, req GoogleAppDownloadRequest) (*GoogleAppDownloadResponse, error) {
	var resp GoogleAppDownloadResponse
	if err := c.core.Do(ctx, "/open/v1/rpa/task/googleAppDownload", req, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

type GoogleAppBrowserRequest struct {
	// Task name, up to 128 characters.
	Name string `json:"name,omitempty"`
	// Remarks, up to 200 characters.
	Remark string `json:"remark,omitempty"`
	// Scheduled execution time as a Unix timestamp in seconds. Required.
	ScheduleAt int64 `json:"scheduleAt"`
	// Cloud phone ID. Required.
	ID string `json:"id"`
	// app name. Required.
	AppName string `json:"appName"`
	// describe your experience.
	Description string `json:"description,omitempty"`
}

type GoogleAppBrowserResponse struct {
	// ID assigned to the asynchronous task.
	TaskID string `json:"taskId"`
}

// GoogleAppBrowser: https://open.geelark.com/api/rpa-task-googleAppBrowser
func (c *Client) GoogleAppBrowser(ctx context.Context, req GoogleAppBrowserRequest) (*GoogleAppBrowserResponse, error) {
	var resp GoogleAppBrowserResponse
	if err := c.core.Do(ctx, "/open/v1/rpa/task/googleAppBrowser", req, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

type SheinLoginRequest struct {
	// Task name, up to 128 characters.
	Name string `json:"name,omitempty"`
	// Remarks, up to 200 characters.
	Remark string `json:"remark,omitempty"`
	// Scheduled execution time as a Unix timestamp in seconds. Required.
	ScheduleAt int64 `json:"scheduleAt"`
	// Cloud phone ID. Required.
	ID string `json:"id"`
	// Email, up to 64 characters. Required.
	Email string `json:"email"`
	// Password, up to 64 characters. Required.
	Password string `json:"password"`
}

type SheinLoginResponse struct {
	// ID assigned to the asynchronous task.
	TaskID string `json:"taskId"`
}

// SheinLogin: https://open.geelark.com/api/rpa-task-sheinLogin
func (c *Client) SheinLogin(ctx context.Context, req SheinLoginRequest) (*SheinLoginResponse, error) {
	var resp SheinLoginResponse
	if err := c.core.Do(ctx, "/open/v1/rpa/task/sheinLogin", req, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
