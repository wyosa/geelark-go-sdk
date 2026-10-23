package automation

import (
	"context"
)

type YouTubePubShortRequest struct {
	// Task name, up to 128 characters.
	Name string `json:"name,omitempty"`
	// Remarks, up to 200 characters.
	Remark string `json:"remark,omitempty"`
	// Scheduled execution time as a Unix timestamp in seconds. Required.
	ScheduleAt int64 `json:"scheduleAt"`
	// Cloud phone ID. Required.
	ID string `json:"id"`
	// Title, up to 100 characters. Required.
	Title string `json:"title"`
	// Video, to upload videos, please refer to Upload Temporary Files to GeeLark
	// (https://open.geelark.com/api/upload-getUrl). Required.
	Video string `json:"video"`
	// Same style URL, up to 500 characters.
	SameStyleURL string `json:"sameStyleUrl,omitempty"`
	// Same style volume, 0-100. If you do not want to send the same URL, just send 0.
	// Required.
	SameStyleVoice int `json:"sameStyleVoice"`
	// Original voice volume, 0-100. If you do not want to send the same URL, just send 0.
	// Required.
	OriginalVoice int `json:"originalVoice"`
	// Whether to force disclosure, defaults to false.
	IsDisclosureMandatory *bool `json:"isDisclosureMandatory,omitempty"`
	// The first frame of the cover is false by default.
	CoverFirstFrame *bool `json:"coverFirstFrame,omitempty"`
	// Whether to retrieve the sharing link, the default is false.
	NeedShareLink *bool `json:"needShareLink,omitempty"`
}

type YouTubePubShortResponse struct {
	// ID assigned to the asynchronous task.
	TaskID string `json:"taskId"`
}

// YouTubePubShort: https://open.geelark.com/api/rpa-task-youtubePubShort
func (c *Client) YouTubePubShort(ctx context.Context, req YouTubePubShortRequest) (*YouTubePubShortResponse, error) {
	var resp YouTubePubShortResponse
	if err := c.core.Do(ctx, "/open/v1/rpa/task/youtubePubShort", req, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

type YouTubePubVideoRequest struct {
	// Task name, up to 128 characters.
	Name string `json:"name,omitempty"`
	// Remarks, up to 200 characters.
	Remark string `json:"remark,omitempty"`
	// Scheduled execution time as a Unix timestamp in seconds. Required.
	ScheduleAt int64 `json:"scheduleAt"`
	// Cloud phone ID. Required.
	ID string `json:"id"`
	// Title, up to 100 characters. Required.
	Title string `json:"title"`
	// Description, up to 5000 characters. Required.
	Description string `json:"description"`
	// Video, to upload videos, please refer to Upload Temporary Files to GeeLark
	// (https://open.geelark.com/api/upload-getUrl). Required.
	Video string `json:"video"`
	// Whether to force disclosure, defaults to false.
	IsDisclosureMandatory *bool `json:"isDisclosureMandatory,omitempty"`
	// Waiting time after publishing, in minutes. Valid values: [0, 100].
	WaitTimeAfterPublishing *int `json:"waitTimeAfterPublishing,omitempty"`
}

type YouTubePubVideoResponse struct {
	// ID assigned to the asynchronous task.
	TaskID string `json:"taskId"`
}

// YouTubePubVideo: https://open.geelark.com/api/rpa-task-youtubePubVideo
func (c *Client) YouTubePubVideo(ctx context.Context, req YouTubePubVideoRequest) (*YouTubePubVideoResponse, error) {
	var resp YouTubePubVideoResponse
	if err := c.core.Do(ctx, "/open/v1/rpa/task/youtubePubVideo", req, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

type YouTubeActiveAccountRequest struct {
	// Task name, up to 128 characters.
	Name string `json:"name,omitempty"`
	// Remarks, up to 200 characters.
	Remark string `json:"remark,omitempty"`
	// Scheduled execution time as a Unix timestamp in seconds. Required.
	ScheduleAt int64 `json:"scheduleAt"`
	// Cloud phone ID. Required.
	ID string `json:"id"`
	// Estimated number of videos to browse, minimum 1, maximum 100. Required.
	BrowseVideoNum int `json:"browseVideoNum"`
	// Keywords, up to 10, each keyword up to 150 characters.
	Keyword []string `json:"keyword,omitempty"`
}

type YouTubeActiveAccountResponse struct {
	// ID assigned to the asynchronous task.
	TaskID string `json:"taskId"`
}

// YouTubeActiveAccount: https://open.geelark.com/api/rpa-task-youTubeActiveAccount
func (c *Client) YouTubeActiveAccount(ctx context.Context, req YouTubeActiveAccountRequest) (*YouTubeActiveAccountResponse, error) {
	var resp YouTubeActiveAccountResponse
	if err := c.core.Do(ctx, "/open/v1/rpa/task/youTubeActiveAccount", req, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

type YouTubeEditRequest struct {
	// Task name, up to 128 characters.
	Name string `json:"name,omitempty"`
	// Remarks, up to 200 characters.
	Remark string `json:"remark,omitempty"`
	// Scheduled execution time as a Unix timestamp in seconds. Required.
	ScheduleAt int64 `json:"scheduleAt"`
	// Cloud phone ID. Required.
	ID string `json:"id"`
	// Name, maximum 50 characters.
	ProfileName string `json:"profileName,omitempty"`
	// Identifier name, maximum 100 characters.
	Handle string `json:"handle,omitempty"`
	// Description, maximum 1000 characters.
	Description string `json:"description,omitempty"`
}

type YouTubeEditResponse struct {
	// ID assigned to the asynchronous task.
	TaskID string `json:"taskId"`
}

// YouTubeEdit: https://open.geelark.com/api/rpa-task-youtubeEdit
func (c *Client) YouTubeEdit(ctx context.Context, req YouTubeEditRequest) (*YouTubeEditResponse, error) {
	var resp YouTubeEditResponse
	if err := c.core.Do(ctx, "/open/v1/rpa/task/youtubeEdit", req, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
