package automation

import (
	"context"
)

type FaceBookLoginRequest struct {
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

type FaceBookLoginResponse struct {
	// ID assigned to the asynchronous task.
	TaskID string `json:"taskId"`
}

// FaceBookLogin: https://open.geelark.com/api/rpa-task-faceBookLogin
//
// Deprecated: use FacebookLogin.
func (c *Client) FaceBookLogin(ctx context.Context, req FaceBookLoginRequest) (*FaceBookLoginResponse, error) {
	var resp FaceBookLoginResponse
	if err := c.core.Do(ctx, "/open/v1/rpa/task/faceBookLogin", req, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

type FaceBookAutoCommentRequest struct {
	// Task name, up to 128 characters.
	Name string `json:"name,omitempty"`
	// Remarks, up to 200 characters.
	Remark string `json:"remark,omitempty"`
	// Scheduled execution time as a Unix timestamp in seconds. Required.
	ScheduleAt int64 `json:"scheduleAt"`
	// Cloud phone ID. Required.
	ID string `json:"id"`
	// Post address, up to 128 characters. Required.
	PostAddress string `json:"postAddress"`
	// Comments, up to 10, each comment up to 8000 characters. Required.
	Comment []string `json:"comment"`
	// Keywords, up to 10, each keyword up to 100 characters. Required.
	Keyword []string `json:"keyword"`
}

type FaceBookAutoCommentResponse struct {
	// ID assigned to the asynchronous task.
	TaskID string `json:"taskId"`
}

// FaceBookAutoComment: https://open.geelark.com/api/rpa-task-faceBookAutoComment
//
// Deprecated: use FacebookAutoComment.
func (c *Client) FaceBookAutoComment(ctx context.Context, req FaceBookAutoCommentRequest) (*FaceBookAutoCommentResponse, error) {
	var resp FaceBookAutoCommentResponse
	if err := c.core.Do(ctx, "/open/v1/rpa/task/faceBookAutoComment", req, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

type FaceBookPublishRequest struct {
	// Task name, up to 128 characters.
	Name string `json:"name,omitempty"`
	// Remarks, up to 200 characters.
	Remark string `json:"remark,omitempty"`
	// Scheduled execution time as a Unix timestamp in seconds. Required.
	ScheduleAt int64 `json:"scheduleAt"`
	// Cloud phone ID. Required.
	ID string `json:"id"`
	// Title, up to 200 characters. Required.
	Title string `json:"title"`
	// Videos, up to 10 videos, to upload videos, please refer to Upload Temporary Files to
	// GeeLark (https://open.geelark.com/api/upload-getUrl). Required.
	Video []string `json:"video"`
	// Whether to retrieve the sharing link, the default is false.
	NeedShareLink *bool `json:"needShareLink,omitempty"`
}

type FaceBookPublishResponse struct {
	// ID assigned to the asynchronous task.
	TaskID string `json:"taskId"`
}

// FaceBookPublish: https://open.geelark.com/api/rpa-task-faceBookPublish
//
// Deprecated: use FacebookPublish.
func (c *Client) FaceBookPublish(ctx context.Context, req FaceBookPublishRequest) (*FaceBookPublishResponse, error) {
	var resp FaceBookPublishResponse
	if err := c.core.Do(ctx, "/open/v1/rpa/task/faceBookPublish", req, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

type FaceBookActiveAccountRequest struct {
	// Task name, up to 128 characters.
	Name string `json:"name,omitempty"`
	// Remarks, up to 200 characters.
	Remark string `json:"remark,omitempty"`
	// Scheduled execution time as a Unix timestamp in seconds. Required.
	ScheduleAt int64 `json:"scheduleAt"`
	// Cloud phone ID. Required.
	ID string `json:"id"`
	// Estimated number of posts to browse, minimum 1, maximum 100. Required.
	BrowsePostsNum int `json:"browsePostsNum"`
	// Keywords, up to 10, each keyword up to 150 characters. Required.
	Keyword []string `json:"keyword"`
	// Post URL shown in the official request example but omitted from the parameter table.
	PostAddress string `json:"postAddress,omitempty"`
}

type FaceBookActiveAccountResponse struct {
	// ID assigned to the asynchronous task.
	TaskID string `json:"taskId"`
}

// FaceBookActiveAccount: https://open.geelark.com/api/rpa-task-faceBookActiveAccount
//
// Deprecated: use FacebookActiveAccount.
func (c *Client) FaceBookActiveAccount(ctx context.Context, req FaceBookActiveAccountRequest) (*FaceBookActiveAccountResponse, error) {
	var resp FaceBookActiveAccountResponse
	if err := c.core.Do(ctx, "/open/v1/rpa/task/faceBookActiveAccount", req, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

type FaceBookPubReelsRequest struct {
	// Task name, up to 128 characters.
	Name string `json:"name,omitempty"`
	// Remarks, up to 200 characters.
	Remark string `json:"remark,omitempty"`
	// Scheduled execution time as a Unix timestamp in seconds. Required.
	ScheduleAt int64 `json:"scheduleAt"`
	// Cloud phone ID. Required.
	ID string `json:"id"`
	// Caption, up to 500 characters. Required.
	Description string `json:"description"`
	// Video, to upload videos, please refer to Upload Temporary Files to GeeLark
	// (https://open.geelark.com/api/upload-getUrl). Required.
	Video string `json:"video"`
	Page  string `json:"page,omitempty"`
	// Whether to retrieve the sharing link, the default is false.
	NeedShareLink *bool `json:"needShareLink,omitempty"`
}

type FaceBookPubReelsResponse struct {
	// ID assigned to the asynchronous task.
	TaskID string `json:"taskId"`
}

// FaceBookPubReels: https://open.geelark.com/api/rpa-task-faceBookPubReels
//
// Deprecated: use FacebookPubReels.
func (c *Client) FaceBookPubReels(ctx context.Context, req FaceBookPubReelsRequest) (*FaceBookPubReelsResponse, error) {
	var resp FaceBookPubReelsResponse
	if err := c.core.Do(ctx, "/open/v1/rpa/task/faceBookPubReels", req, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

type FaceBookMessageRequest struct {
	// Task name, up to 128 characters.
	Name string `json:"name,omitempty"`
	// Remarks, up to 200 characters.
	Remark string `json:"remark,omitempty"`
	// Scheduled execution time as a Unix timestamp in seconds. Required.
	ScheduleAt int64 `json:"scheduleAt"`
	// Cloud phone ID. Required.
	ID string `json:"id"`
	// usernames. Required.
	Usernames []string `json:"usernames"`
	// content. Maximum 20000 characters. Required.
	Content string `json:"content"`
}

type FaceBookMessageResponse struct {
	// ID assigned to the asynchronous task.
	TaskID string `json:"taskId"`
}

// FaceBookMessage: https://open.geelark.com/api/rpa-task-faceBookMessage
//
// Deprecated: use FacebookMessage.
func (c *Client) FaceBookMessage(ctx context.Context, req FaceBookMessageRequest) (*FaceBookMessageResponse, error) {
	var resp FaceBookMessageResponse
	if err := c.core.Do(ctx, "/open/v1/rpa/task/faceBookMessage", req, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

type FacebookReelsActiveRequest struct {
	// Task name, up to 128 characters.
	Name string `json:"name,omitempty"`
	// Remarks, up to 200 characters.
	Remark string `json:"remark,omitempty"`
	// Scheduled execution time as a Unix timestamp in seconds. Required.
	ScheduleAt int64 `json:"scheduleAt"`
	// Cloud phone ID. Required.
	ID string `json:"id"`
	// Estimated number of videos viewed. Required.
	VideoNumber    int      `json:"videoNumber"`
	SearchKeywords []string `json:"searchKeywords,omitempty"`
}

type FacebookReelsActiveResponse struct {
	// ID assigned to the asynchronous task.
	TaskID string `json:"taskId"`
}

// FacebookReelsActive: https://open.geelark.com/api/rpa-task-facebookReelsActive
func (c *Client) FacebookReelsActive(ctx context.Context, req FacebookReelsActiveRequest) (*FacebookReelsActiveResponse, error) {
	var resp FacebookReelsActiveResponse
	if err := c.core.Do(ctx, "/open/v1/rpa/task/facebookReelsActive", req, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
