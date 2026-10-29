package automation

import (
	"context"
)

type XPublishRequest struct {
	// Task name, up to 128 characters.
	Name string `json:"name,omitempty"`
	// Remarks, up to 200 characters.
	Remark string `json:"remark,omitempty"`
	// Scheduled execution time as a Unix timestamp in seconds. Required.
	ScheduleAt int64 `json:"scheduleAt"`
	// Cloud phone ID. Required.
	ID string `json:"id"`
	// Caption, up to 280 characters. Required.
	Description string `json:"description"`
	// Videos, up to 1, to upload videos, please refer to Upload Temporary Files to GeeLark
	// (https://open.geelark.com/api/upload-getUrl). Required.
	Video []string `json:"video"`
}

type XPublishResponse struct {
	// ID assigned to the asynchronous task.
	TaskID string `json:"taskId"`
}

// XPublish: https://open.geelark.com/api/rpa-task-xPublish
func (c *Client) XPublish(ctx context.Context, req XPublishRequest) (*XPublishResponse, error) {
	var resp XPublishResponse
	if err := c.core.Do(ctx, "/open/v1/rpa/task/xPublish", req, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

type MultiPlatformVideoDistributionRequest struct {
	// Task name, up to 128 characters.
	Name string `json:"name,omitempty"`
	// Remarks, up to 200 characters.
	Remark string `json:"remark,omitempty"`
	// Scheduled execution time as a Unix timestamp in seconds. Required.
	ScheduleAt int64 `json:"scheduleAt"`
	// Cloud phone ID. Required.
	ID string `json:"id"`
	// Title, maximum 100 characters, required when filling in all fields at once.
	Title string `json:"title,omitempty"`
	// TikTok title, maximum 4000 characters, required when filling in separately.
	TikTokTitle string `json:"tiktokTitle,omitempty"`
	// YouTube title, maximum 100 characters, required when filling in separately.
	YouTubeTitle string `json:"youtubeTitle,omitempty"`
	// Instagram title, maximum 2200 characters, optional when filling in separately.
	InstagramTitle string `json:"instagramTitle,omitempty"`
	// Videos, up to 10 videos, refer to the User Guide - File Upload for creating automation
	// tasks. Required.
	Video []string `json:"video"`
	// TikTok style link.
	TikTokRecreateLink string `json:"tiktokRecreateLink,omitempty"`
	// YouTube style link.
	YouTubeRecreateLink string `json:"youtubeRecreateLink,omitempty"`
	// Instagram style link.
	InstagramRecreateLink string `json:"instagramRecreateLink,omitempty"`
	// Same volume, range 0-100.
	SameStyleVoice *int `json:"sameStyleVoice,omitempty"`
	// Original volume, range 0-100.
	OriginalVoice *int `json:"originalVoice,omitempty"`
}

type MultiPlatformVideoDistributionResponse struct {
	// ID assigned to the asynchronous task.
	TaskID string `json:"taskId"`
}

// MultiPlatformVideoDistribution: https://open.geelark.com/api/rpa-task-multiPlatformVideoDistribution
func (c *Client) MultiPlatformVideoDistribution(ctx context.Context, req MultiPlatformVideoDistributionRequest) (*MultiPlatformVideoDistributionResponse, error) {
	var resp MultiPlatformVideoDistributionResponse
	if err := c.core.Do(ctx, "/open/v1/rpa/task/multiPlatformVideoDistribution", req, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

type RedditWarmupRequest struct {
	// Task name, up to 128 characters.
	Name string `json:"name,omitempty"`
	// Remarks, up to 200 characters.
	Remark string `json:"remark,omitempty"`
	// Scheduled execution time as a Unix timestamp in seconds. Required.
	ScheduleAt int64 `json:"scheduleAt"`
	// Cloud phone ID. Required.
	ID string `json:"id"`
	// Search keyword. Maximum length: 150 characters.
	//
	// Deprecated: GeeLark is retiring this field; use Keywords.
	Keyword string `json:"keyword,omitempty"`
	// Search keyword array. Supports up to 100 keywords, with each keyword having a maximum
	// length of 150 characters. If keywords is not empty, the keyword field is ignored.
	Keywords []string `json:"keywords,omitempty"`
	// Duration, in minutes.Valid range: [1, 9999]. Required.
	Duration int `json:"duration"`
}

type RedditWarmupResponse struct {
	// ID assigned to the asynchronous task.
	TaskID string `json:"taskId"`
}

// RedditWarmup: https://open.geelark.com/api/rpa-task-redditWarmup
func (c *Client) RedditWarmup(ctx context.Context, req RedditWarmupRequest) (*RedditWarmupResponse, error) {
	var resp RedditWarmupResponse
	if err := c.core.Do(ctx, "/open/v1/rpa/task/redditWarmup", req, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

type RedditVideoRequest struct {
	// Task name, up to 128 characters.
	Name string `json:"name,omitempty"`
	// Remarks, up to 200 characters.
	Remark string `json:"remark,omitempty"`
	// Scheduled execution time as a Unix timestamp in seconds. Required.
	ScheduleAt int64 `json:"scheduleAt"`
	// Cloud phone ID. Required.
	ID string `json:"id"`
	// Title. Required.
	Title       string `json:"title"`
	Description string `json:"description,omitempty"`
	// Videos, to upload videos, please refer to Upload Temporary Files to GeeLark
	// (https://open.geelark.com/api/upload-getUrl). Required.
	Video []string `json:"video"`
	// Community. Required.
	Community string `json:"community"`
	// flair, up to 100 characters.
	Flair string `json:"flair,omitempty"`
}

type RedditVideoResponse struct {
	// ID assigned to the asynchronous task.
	TaskID string `json:"taskId"`
}

// RedditVideo: https://open.geelark.com/api/rpa-task-redditVideo
func (c *Client) RedditVideo(ctx context.Context, req RedditVideoRequest) (*RedditVideoResponse, error) {
	var resp RedditVideoResponse
	if err := c.core.Do(ctx, "/open/v1/rpa/task/redditVideo", req, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

type RedditImageRequest struct {
	// Task name, up to 128 characters.
	Name string `json:"name,omitempty"`
	// Remarks, up to 200 characters.
	Remark string `json:"remark,omitempty"`
	// Scheduled execution time as a Unix timestamp in seconds. Required.
	ScheduleAt int64 `json:"scheduleAt"`
	// Cloud phone ID. Required.
	ID string `json:"id"`
	// Title. Required.
	Title       string `json:"title"`
	Description string `json:"description,omitempty"`
	// Images, to upload images, please refer to Upload Temporary Files to GeeLark
	// (https://open.geelark.com/api/upload-getUrl). Required.
	Images []string `json:"images"`
	// Community. Required.
	Community string `json:"community"`
	// flair, up to 100 characters.
	Flair string `json:"flair,omitempty"`
}

type RedditImageResponse struct {
	// ID assigned to the asynchronous task.
	TaskID string `json:"taskId"`
}

// RedditImage: https://open.geelark.com/api/rpa-task-redditImage
func (c *Client) RedditImage(ctx context.Context, req RedditImageRequest) (*RedditImageResponse, error) {
	var resp RedditImageResponse
	if err := c.core.Do(ctx, "/open/v1/rpa/task/redditImage", req, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

type PinterestVideoRequest struct {
	// Task name, up to 128 characters.
	Name string `json:"name,omitempty"`
	// Remarks, up to 200 characters.
	Remark string `json:"remark,omitempty"`
	// Scheduled execution time as a Unix timestamp in seconds. Required.
	ScheduleAt int64 `json:"scheduleAt"`
	// Cloud phone ID. Required.
	ID string `json:"id"`
	// Title. Maximum 100 characters. Required.
	Title string `json:"title"`
	// Description. Maximum 800 characters. Required.
	Description string `json:"description"`
	// Videos, to upload videos, please refer to Upload Temporary Files to GeeLark
	// (https://open.geelark.com/api/upload-getUrl). Required.
	Video []string `json:"video"`
	Link  string   `json:"link,omitempty"`
}

type PinterestVideoResponse struct {
	// ID assigned to the asynchronous task.
	TaskID string `json:"taskId"`
}

// PinterestVideo: https://open.geelark.com/api/rpa-task-pinterestVideo
func (c *Client) PinterestVideo(ctx context.Context, req PinterestVideoRequest) (*PinterestVideoResponse, error) {
	var resp PinterestVideoResponse
	if err := c.core.Do(ctx, "/open/v1/rpa/task/pinterestVideo", req, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

type PinterestImageRequest struct {
	// Task name, up to 128 characters.
	Name string `json:"name,omitempty"`
	// Remarks, up to 200 characters.
	Remark string `json:"remark,omitempty"`
	// Scheduled execution time as a Unix timestamp in seconds. Required.
	ScheduleAt int64 `json:"scheduleAt"`
	// Cloud phone ID. Required.
	ID string `json:"id"`
	// Title. Maximum 100 characters. Required.
	Title string `json:"title"`
	// Description. Maximum 800 characters. Required.
	Description string `json:"description"`
	// Images, to upload images, please refer to Upload Temporary Files to GeeLark
	// (https://open.geelark.com/api/upload-getUrl). Required.
	Images []string `json:"images"`
	Link   string   `json:"link,omitempty"`
}

type PinterestImageResponse struct {
	// ID assigned to the asynchronous task.
	TaskID string `json:"taskId"`
}

// PinterestImage: https://open.geelark.com/api/rpa-task-pinterestImage
func (c *Client) PinterestImage(ctx context.Context, req PinterestImageRequest) (*PinterestImageResponse, error) {
	var resp PinterestImageResponse
	if err := c.core.Do(ctx, "/open/v1/rpa/task/pinterestImage", req, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

type ThreadsVideoRequest struct {
	// Task name, up to 128 characters.
	Name string `json:"name,omitempty"`
	// Remarks, up to 200 characters.
	Remark string `json:"remark,omitempty"`
	// Scheduled execution time as a Unix timestamp in seconds. Required.
	ScheduleAt int64 `json:"scheduleAt"`
	// Cloud phone ID. Required.
	ID    string `json:"id"`
	Topic string `json:"topic,omitempty"`
	// Title. Maximum 500 characters. Required.
	Title string `json:"title"`
	// Videos, to upload videos, please refer to Upload Temporary Files to GeeLark
	// (https://open.geelark.com/api/upload-getUrl). Required.
	Video []string `json:"video"`
}

type ThreadsVideoResponse struct {
	// ID assigned to the asynchronous task.
	TaskID string `json:"taskId"`
}

// ThreadsVideo: https://open.geelark.com/api/rpa-task-threadsVideo
func (c *Client) ThreadsVideo(ctx context.Context, req ThreadsVideoRequest) (*ThreadsVideoResponse, error) {
	var resp ThreadsVideoResponse
	if err := c.core.Do(ctx, "/open/v1/rpa/task/threadsVideo", req, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

type ThreadsImageRequest struct {
	// Task name, up to 128 characters.
	Name string `json:"name,omitempty"`
	// Remarks, up to 200 characters.
	Remark string `json:"remark,omitempty"`
	// Scheduled execution time as a Unix timestamp in seconds. Required.
	ScheduleAt int64 `json:"scheduleAt"`
	// Cloud phone ID. Required.
	ID    string `json:"id"`
	Topic string `json:"topic,omitempty"`
	// Title. Maximum 500 characters. Required.
	Title string `json:"title"`
	// Images, to upload images, please refer to Upload Temporary Files to GeeLark
	// (https://open.geelark.com/api/upload-getUrl). Required.
	Images []string `json:"images"`
}

type ThreadsImageResponse struct {
	// ID assigned to the asynchronous task.
	TaskID string `json:"taskId"`
}

// ThreadsImage: https://open.geelark.com/api/rpa-task-threadsImage
func (c *Client) ThreadsImage(ctx context.Context, req ThreadsImageRequest) (*ThreadsImageResponse, error) {
	var resp ThreadsImageResponse
	if err := c.core.Do(ctx, "/open/v1/rpa/task/threadsImage", req, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
