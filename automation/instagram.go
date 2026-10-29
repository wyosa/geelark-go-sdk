package automation

import (
	"context"
)

type InstagramPubReelsRequest struct {
	// Task name, up to 128 characters.
	Name string `json:"name,omitempty"`
	// Remarks, up to 200 characters.
	Remark string `json:"remark,omitempty"`
	// Scheduled execution time as a Unix timestamp in seconds. Required.
	ScheduleAt int64 `json:"scheduleAt"`
	// Cloud phone ID. Required.
	ID string `json:"id"`
	// Caption, up to 2200 characters.
	Description string `json:"description,omitempty"`
	// Videos, up to 10, to upload videos, please refer to Upload Temporary Files to GeeLark
	// (https://open.geelark.com/api/upload-getUrl). Required.
	Video []string `json:"video"`
	// Same URL.
	SameStyleURL string `json:"sameStyleUrl,omitempty"`
	// Same volume, range 0-100.
	SameStyleVoice *int `json:"sameStyleVoice,omitempty"`
	// Original volume, range 0-100.
	OriginalVoice *int `json:"originalVoice,omitempty"`
	// AI tag, defaults to false.
	AiTag *bool `json:"aiTag,omitempty"`
	// Whether to retrieve the sharing link, the default is false.
	NeedShareLink *bool    `json:"needShareLink,omitempty"`
	Cover         []string `json:"cover,omitempty"`
}

type InstagramPubReelsResponse struct {
	// ID assigned to the asynchronous task.
	TaskID string `json:"taskId"`
}

// InstagramPubReels: https://open.geelark.com/api/rpa-task-instagramPubReels
func (c *Client) InstagramPubReels(ctx context.Context, req InstagramPubReelsRequest) (*InstagramPubReelsResponse, error) {
	var resp InstagramPubReelsResponse
	if err := c.core.Do(ctx, "/open/v1/rpa/task/instagramPubReels", req, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

type InstagramWarmupRequest struct {
	// Task name, up to 128 characters.
	Name string `json:"name,omitempty"`
	// Remarks, up to 200 characters.
	Remark string `json:"remark,omitempty"`
	// Scheduled execution time as a Unix timestamp in seconds. Required.
	ScheduleAt int64 `json:"scheduleAt"`
	// Cloud phone ID. Required.
	ID string `json:"id"`
	// Number of videos viewed, 1-100.
	BrowseVideo *int `json:"browseVideo,omitempty"`
	// Search keyword.
	Keyword string `json:"keyword,omitempty"`
}

type InstagramWarmupResponse struct {
	// ID assigned to the asynchronous task.
	TaskID string `json:"taskId"`
}

// InstagramWarmup: https://open.geelark.com/api/rpa-task-instagramWarmup
func (c *Client) InstagramWarmup(ctx context.Context, req InstagramWarmupRequest) (*InstagramWarmupResponse, error) {
	var resp InstagramWarmupResponse
	if err := c.core.Do(ctx, "/open/v1/rpa/task/instagramWarmup", req, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

type InstagramMessageRequest struct {
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
	// content. Maximum 1000 characters. Required.
	Content string `json:"content"`
}

type InstagramMessageResponse struct {
	// ID assigned to the asynchronous task.
	TaskID string `json:"taskId"`
}

// InstagramMessage: https://open.geelark.com/api/rpa-task-instagramMessage
func (c *Client) InstagramMessage(ctx context.Context, req InstagramMessageRequest) (*InstagramMessageResponse, error) {
	var resp InstagramMessageResponse
	if err := c.core.Do(ctx, "/open/v1/rpa/task/instagramMessage", req, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

type InstagramPubReelsImagesRequest struct {
	// Task name, up to 128 characters.
	Name string `json:"name,omitempty"`
	// Remarks, up to 200 characters.
	Remark string `json:"remark,omitempty"`
	// Scheduled execution time as a Unix timestamp in seconds. Required.
	ScheduleAt int64 `json:"scheduleAt"`
	// Cloud phone ID. Required.
	ID string `json:"id"`
	// Caption, up to 2200 characters.
	Description string `json:"description,omitempty"`
	// Images, up to 10, to upload images, please refer to Upload Temporary Files to GeeLark
	// (https://open.geelark.com/api/upload-getUrl). Required.
	Image []string `json:"image"`
	// Same URL.
	SameStyleURL string `json:"sameStyleUrl,omitempty"`
	// AI tag, defaults to false.
	AiTag *bool `json:"aiTag,omitempty"`
	// Whether to publish the images as a feed post instead of Reels; defaults to false.
	PublishPost *bool `json:"publishPost,omitempty"`
	// Whether to retrieve the sharing link, the default is false.
	NeedShareLink *bool `json:"needShareLink,omitempty"`
}

type InstagramPubReelsImagesResponse struct {
	// ID assigned to the asynchronous task.
	TaskID string `json:"taskId"`
}

// InstagramPubReelsImages: https://open.geelark.com/api/rpa-task-instagramPubReelsImages
func (c *Client) InstagramPubReelsImages(ctx context.Context, req InstagramPubReelsImagesRequest) (*InstagramPubReelsImagesResponse, error) {
	var resp InstagramPubReelsImagesResponse
	if err := c.core.Do(ctx, "/open/v1/rpa/task/instagramPubReelsImages", req, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

type InstagramLoginRequest struct {
	// Task name, up to 128 characters.
	Name string `json:"name,omitempty"`
	// Remarks, up to 200 characters.
	Remark string `json:"remark,omitempty"`
	// Scheduled execution time as a Unix timestamp in seconds. Required.
	ScheduleAt int64 `json:"scheduleAt"`
	// Cloud phone ID. Required.
	ID string `json:"id"`
	// Account, up to 64 characters. Required.
	Account string `json:"account"`
	// Password, up to 64 characters. Required.
	Password string `json:"password"`
}

type InstagramLoginResponse struct {
	// ID assigned to the asynchronous task.
	TaskID string `json:"taskId"`
}

// InstagramLogin: https://open.geelark.com/api/rpa-task-instagramLogin
func (c *Client) InstagramLogin(ctx context.Context, req InstagramLoginRequest) (*InstagramLoginResponse, error) {
	var resp InstagramLoginResponse
	if err := c.core.Do(ctx, "/open/v1/rpa/task/instagramLogin", req, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

type InstagramEditRequest struct {
	// Task name, up to 128 characters.
	Name string `json:"name,omitempty"`
	// Remarks, up to 200 characters.
	Remark string `json:"remark,omitempty"`
	// Scheduled execution time as a Unix timestamp in seconds. Required.
	ScheduleAt int64 `json:"scheduleAt"`
	// Cloud phone ID. Required.
	ID string `json:"id"`
	// Avatar.
	ProfilePicture []string `json:"profilePicture,omitempty"`
	Nickname       string   `json:"nickname,omitempty"`
	Username       string   `json:"username,omitempty"`
	Biography      string   `json:"biography,omitempty"`
	LinkURL        string   `json:"linkURL,omitempty"`
	LinkTitle      string   `json:"linkTitle,omitempty"`
}

type InstagramEditResponse struct {
	// ID assigned to the asynchronous task.
	TaskID string `json:"taskId"`
}

// InstagramEdit: https://open.geelark.com/api/rpa-task-instagramEdit
func (c *Client) InstagramEdit(ctx context.Context, req InstagramEditRequest) (*InstagramEditResponse, error) {
	var resp InstagramEditResponse
	if err := c.core.Do(ctx, "/open/v1/rpa/task/instagramEdit", req, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

type InstagramAccountStatusRequest struct {
	// Task name, up to 128 characters.
	Name string `json:"name,omitempty"`
	// Remarks, up to 200 characters.
	Remark string `json:"remark,omitempty"`
	// Scheduled execution time as a Unix timestamp in seconds. Required.
	ScheduleAt int64 `json:"scheduleAt"`
	// Cloud phone ID. Required.
	ID string `json:"id"`
	// Public, defaults to false.
	Public *bool `json:"public,omitempty"`
}

type InstagramAccountStatusResponse struct {
	// ID assigned to the asynchronous task.
	TaskID string `json:"taskId"`
}

// InstagramAccountStatus: https://open.geelark.com/api/rpa-task-instagramAccountStatus
func (c *Client) InstagramAccountStatus(ctx context.Context, req InstagramAccountStatusRequest) (*InstagramAccountStatusResponse, error) {
	var resp InstagramAccountStatusResponse
	if err := c.core.Do(ctx, "/open/v1/rpa/task/instagramAccountStatus", req, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

type InstagramFollowAccountRequest struct {
	// Task name, up to 128 characters.
	Name string `json:"name,omitempty"`
	// Remarks, up to 200 characters.
	Remark string `json:"remark,omitempty"`
	// Scheduled execution time as a Unix timestamp in seconds. Required.
	ScheduleAt int64 `json:"scheduleAt"`
	// Cloud phone ID. Required.
	ID string `json:"id"`
	// List of usernames. Supports up to 100 elements, with each username containing a maximum
	// of 30 characters. Required.
	Username []string `json:"username"`
}

type InstagramFollowAccountResponse struct {
	// ID assigned to the asynchronous task.
	TaskID string `json:"taskId"`
}

// InstagramFollowAccount: https://open.geelark.com/api/rpa-task-instagramFollowAccount
func (c *Client) InstagramFollowAccount(ctx context.Context, req InstagramFollowAccountRequest) (*InstagramFollowAccountResponse, error) {
	var resp InstagramFollowAccountResponse
	if err := c.core.Do(ctx, "/open/v1/rpa/task/instagramFollowAccount", req, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

type InstagramAiCommentRequest struct {
	// Task name, up to 128 characters.
	Name string `json:"name,omitempty"`
	// Remarks, up to 200 characters.
	Remark string `json:"remark,omitempty"`
	// Scheduled execution time as a Unix timestamp in seconds. Required.
	ScheduleAt int64 `json:"scheduleAt"`
	// Cloud phone ID. Required.
	ID string `json:"id"`
	// Whether to use AI for comments. Defaults to false.
	UseAi *bool `json:"useAi,omitempty"`
	// Random probability, ranging from 0 to 100. Required.
	RandomRate int `json:"randomRate"`
}

type InstagramAiCommentResponse struct {
	// ID assigned to the asynchronous task.
	TaskID string `json:"taskId"`
}

// InstagramAiComment: https://open.geelark.com/api/rpa-task-InstagramAIComments
func (c *Client) InstagramAiComment(ctx context.Context, req InstagramAiCommentRequest) (*InstagramAiCommentResponse, error) {
	var resp InstagramAiCommentResponse
	if err := c.core.Do(ctx, "/open/v1/rpa/task/instagramAiComment", req, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

type InstagramLikeRequest struct {
	// Task name, up to 128 characters.
	Name string `json:"name,omitempty"`
	// Remarks, up to 200 characters.
	Remark string `json:"remark,omitempty"`
	// Scheduled execution time as a Unix timestamp in seconds. Required.
	ScheduleAt int64 `json:"scheduleAt"`
	// Cloud phone ID. Required.
	ID string `json:"id"`
	// Links of the posts to like. Supports up to 100 elements, with each link containing a
	// maximum of 100 characters. Required.
	Links []string `json:"links"`
	// Comments to post. Supports up to 100 elements, with each comment containing a maximum of
	// 100 characters.
	Comments []string `json:"comments,omitempty"`
}

type InstagramLikeResponse struct {
	// ID assigned to the asynchronous task.
	TaskID string `json:"taskId"`
}

// InstagramLike: https://open.geelark.com/api/rpa-task-InstagramLike
//
// Likes the specified Instagram posts and optionally comments on them.
func (c *Client) InstagramLike(ctx context.Context, req InstagramLikeRequest) (*InstagramLikeResponse, error) {
	var resp InstagramLikeResponse
	if err := c.core.Do(ctx, "/open/v1/rpa/task/instagramLike", req, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
