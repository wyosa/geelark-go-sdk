package automation

import (
	"context"
)

type TikTokLoginRequest struct {
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
	// 2FA Key.
	TwoFAKey string `json:"twoFAKey,omitempty"`
}

type TikTokLoginResponse struct {
	// ID assigned to the asynchronous task.
	TaskID string `json:"taskId"`
}

// TikTokLogin: https://open.geelark.com/api/rpa-task-tiktokLogin
func (c *Client) TikTokLogin(ctx context.Context, req TikTokLoginRequest) (*TikTokLoginResponse, error) {
	var resp TikTokLoginResponse
	if err := c.core.Do(ctx, "/open/v1/rpa/task/tiktokLogin", req, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

type TikTokEditRequest struct {
	// Task name, up to 128 characters.
	Name string `json:"name,omitempty"`
	// Remarks, up to 200 characters.
	Remark string `json:"remark,omitempty"`
	// Scheduled execution time as a Unix timestamp in seconds. Required.
	ScheduleAt int64 `json:"scheduleAt"`
	// Cloud phone ID. Required.
	ID string `json:"id"`
	// Avatar URL, refer to the User Guide - File Upload for creating automation tasks; the
	// uploaded image should have a 1:1 aspect ratio, otherwise the edit will fail.
	Avatar string `json:"avatar,omitempty"`
	// Nickname, up to 30 characters.
	NickName string `json:"nickName,omitempty"`
	// Bio, up to 160 characters.
	Bio string `json:"bio,omitempty"`
	// Website, please provide a URL starting with http/https.
	Site     string `json:"site,omitempty"`
	Username string `json:"username,omitempty"`
}

type TikTokEditResponse struct {
	// ID assigned to the asynchronous task.
	TaskID string `json:"taskId"`
}

// TikTokEdit: https://open.geelark.com/api/rpa-task-tiktokEdit
func (c *Client) TikTokEdit(ctx context.Context, req TikTokEditRequest) (*TikTokEditResponse, error) {
	var resp TikTokEditResponse
	if err := c.core.Do(ctx, "/open/v1/rpa/task/tiktokEdit", req, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

type TikTokRandomStarRequest struct {
	// Task name, up to 128 characters.
	Name string `json:"name,omitempty"`
	// Remarks, up to 200 characters.
	Remark string `json:"remark,omitempty"`
	// Scheduled execution time as a Unix timestamp in seconds. Required.
	ScheduleAt int64 `json:"scheduleAt"`
	// Cloud phone ID. Required.
	ID string `json:"id"`
	// Probability of liking, 0-100, default is 30.
	LikeProbability *int `json:"likeProbability,omitempty"`
}

type TikTokRandomStarResponse struct {
	// ID assigned to the asynchronous task.
	TaskID string `json:"taskId"`
}

// TikTokRandomStar: https://open.geelark.com/api/rpa-task-tiktokRandomStar
func (c *Client) TikTokRandomStar(ctx context.Context, req TikTokRandomStarRequest) (*TikTokRandomStarResponse, error) {
	var resp TikTokRandomStarResponse
	if err := c.core.Do(ctx, "/open/v1/rpa/task/tiktokRandomStar", req, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

type TikTokRandomStarAsiaRequest struct {
	// Task name, up to 128 characters.
	Name string `json:"name,omitempty"`
	// Remarks, up to 200 characters.
	Remark string `json:"remark,omitempty"`
	// Scheduled execution time as a Unix timestamp in seconds. Required.
	ScheduleAt int64 `json:"scheduleAt"`
	// Cloud phone ID. Required.
	ID string `json:"id"`
	// Probability of liking, 0-100, default is 30.
	LikeProbability *int `json:"likeProbability,omitempty"`
}

type TikTokRandomStarAsiaResponse struct {
	// ID assigned to the asynchronous task.
	TaskID string `json:"taskId"`
}

// TikTokRandomStarAsia: https://open.geelark.com/api/rpa-task-tiktokRandomStarAsia
func (c *Client) TikTokRandomStarAsia(ctx context.Context, req TikTokRandomStarAsiaRequest) (*TikTokRandomStarAsiaResponse, error) {
	var resp TikTokRandomStarAsiaResponse
	if err := c.core.Do(ctx, "/open/v1/rpa/task/tiktokRandomStarAsia", req, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

type TikTokRandomCommentRequest struct {
	// Task name, up to 128 characters.
	Name string `json:"name,omitempty"`
	// Remarks, up to 200 characters.
	Remark string `json:"remark,omitempty"`
	// Scheduled execution time as a Unix timestamp in seconds. Required.
	ScheduleAt int64 `json:"scheduleAt"`
	// Cloud phone ID. Required.
	ID string `json:"id"`
	// Whether to use AI: UseAIComment (1, Pro users only) or UseOwnComment (2, post Comment).
	// Required; zero is not a valid value.
	UseAi int `json:"useAi"`
	// Comment content, up to 500 characters; required when useAi is 2. Required.
	Comment string `json:"comment"`
	// Specified link.
	Links []string `json:"links,omitempty"`
	// Comment probability, 0-100, default is 30.
	CommentProbability *int     `json:"commentProbability,omitempty"`
	SearchKeywords     []string `json:"searchKeywords,omitempty"`
	// Whether to like, default is no.
	LikeVideo *bool `json:"likeVideo,omitempty"`
	// Comment image url, maximum 500 characters. Takes effect when useAi is 2.
	ImageURL string `json:"imageUrl,omitempty"`
}

type TikTokRandomCommentResponse struct {
	// ID assigned to the asynchronous task.
	TaskID string `json:"taskId"`
}

// TikTokRandomComment: https://open.geelark.com/api/rpa-task-tiktokRandomComment
func (c *Client) TikTokRandomComment(ctx context.Context, req TikTokRandomCommentRequest) (*TikTokRandomCommentResponse, error) {
	var resp TikTokRandomCommentResponse
	if err := c.core.Do(ctx, "/open/v1/rpa/task/tiktokRandomComment", req, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

type TikTokRandomCommentAsiaRequest struct {
	// Task name, up to 128 characters.
	Name string `json:"name,omitempty"`
	// Remarks, up to 200 characters.
	Remark string `json:"remark,omitempty"`
	// Scheduled execution time as a Unix timestamp in seconds. Required.
	ScheduleAt int64 `json:"scheduleAt"`
	// Cloud phone ID. Required.
	ID string `json:"id"`
	// Whether to use AI: UseAIComment (1, Pro users only) or UseOwnComment (2, post Comment).
	// Required; zero is not a valid value.
	UseAi int `json:"useAi"`
	// Comment content, up to 500 characters; required when useAi is 2. Required.
	Comment string `json:"comment"`
	// Specified link.
	Links []string `json:"links,omitempty"`
	// Comment probability, 0-100, default is 30.
	CommentProbability *int     `json:"commentProbability,omitempty"`
	SearchKeywords     []string `json:"searchKeywords,omitempty"`
	// Whether to like, default is no.
	LikeVideo *bool `json:"likeVideo,omitempty"`
	// Comment image url, maximum 500 characters. Takes effect when useAi is 2.
	ImageURL string `json:"imageUrl,omitempty"`
}

type TikTokRandomCommentAsiaResponse struct {
	// ID assigned to the asynchronous task.
	TaskID string `json:"taskId"`
}

// TikTokRandomCommentAsia: https://open.geelark.com/api/rpa-task-tiktokRandomCommentAsia
func (c *Client) TikTokRandomCommentAsia(ctx context.Context, req TikTokRandomCommentAsiaRequest) (*TikTokRandomCommentAsiaResponse, error) {
	var resp TikTokRandomCommentAsiaResponse
	if err := c.core.Do(ctx, "/open/v1/rpa/task/tiktokRandomCommentAsia", req, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

type TikTokMessageRequest struct {
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
	// content. Maximum 6000 characters. Required.
	Content string `json:"content"`
}

type TikTokMessageResponse struct {
	// ID assigned to the asynchronous task.
	TaskID string `json:"taskId"`
}

// TikTokMessage: https://open.geelark.com/api/rpa-task-tiktokMessage
func (c *Client) TikTokMessage(ctx context.Context, req TikTokMessageRequest) (*TikTokMessageResponse, error) {
	var resp TikTokMessageResponse
	if err := c.core.Do(ctx, "/open/v1/rpa/task/tiktokMessage", req, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

type TikTokMessageAsiaRequest struct {
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
	// content. Maximum 6000 characters. Required.
	Content string `json:"content"`
}

type TikTokMessageAsiaResponse struct {
	// ID assigned to the asynchronous task.
	TaskID string `json:"taskId"`
}

// TikTokMessageAsia: https://open.geelark.com/api/rpa-task-tiktokMessageAsia
func (c *Client) TikTokMessageAsia(ctx context.Context, req TikTokMessageAsiaRequest) (*TikTokMessageAsiaResponse, error) {
	var resp TikTokMessageAsiaResponse
	if err := c.core.Do(ctx, "/open/v1/rpa/task/tiktokMessageAsia", req, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

type TikTokRandomFollowRequest struct {
	// Task name, up to 128 characters.
	Name string `json:"name,omitempty"`
	// Remarks, up to 200 characters.
	Remark string `json:"remark,omitempty"`
	// Scheduled execution time as a Unix timestamp in seconds. Required.
	ScheduleAt int64 `json:"scheduleAt"`
	// Cloud phone ID. Required.
	ID string `json:"id"`
	// Follow probability, 0-100. Required.
	FollowProbability int `json:"followProbability"`
}

type TikTokRandomFollowResponse struct {
	// ID assigned to the asynchronous task.
	TaskID string `json:"taskId"`
}

// TikTokRandomFollow: https://open.geelark.com/api/rpa-task-tiktokRandomFollow
func (c *Client) TikTokRandomFollow(ctx context.Context, req TikTokRandomFollowRequest) (*TikTokRandomFollowResponse, error) {
	var resp TikTokRandomFollowResponse
	if err := c.core.Do(ctx, "/open/v1/rpa/task/tiktokRandomFollow", req, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

type TikTokRandomFollowAsiaRequest struct {
	// Task name, up to 128 characters.
	Name string `json:"name,omitempty"`
	// Remarks, up to 200 characters.
	Remark string `json:"remark,omitempty"`
	// Scheduled execution time as a Unix timestamp in seconds. Required.
	ScheduleAt int64 `json:"scheduleAt"`
	// Cloud phone ID. Required.
	ID string `json:"id"`
	// Follow probability, 0-100. Required.
	FollowProbability int `json:"followProbability"`
}

type TikTokRandomFollowAsiaResponse struct {
	// ID assigned to the asynchronous task.
	TaskID string `json:"taskId"`
}

// TikTokRandomFollowAsia: https://open.geelark.com/api/rpa-task-tiktokRandomFollowAsia
func (c *Client) TikTokRandomFollowAsia(ctx context.Context, req TikTokRandomFollowAsiaRequest) (*TikTokRandomFollowAsiaResponse, error) {
	var resp TikTokRandomFollowAsiaResponse
	if err := c.core.Do(ctx, "/open/v1/rpa/task/tiktokRandomFollowAsia", req, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

type TikTokHideRequest struct {
	// Task name, up to 128 characters.
	Name string `json:"name,omitempty"`
	// Remarks, up to 200 characters.
	Remark string `json:"remark,omitempty"`
	// Scheduled execution time as a Unix timestamp in seconds. Required.
	ScheduleAt int64 `json:"scheduleAt"`
	// Cloud phone ID. Required.
	ID string `json:"id"`
	// Number value, range: [0, 999].If not specified or set to 0, all items will be hidden.
	Number *int `json:"number,omitempty"`
}

type TikTokHideResponse struct {
	// ID assigned to the asynchronous task.
	TaskID string `json:"taskId"`
}

// TikTokHide: https://open.geelark.com/api/rpa-task-tiktokHide
func (c *Client) TikTokHide(ctx context.Context, req TikTokHideRequest) (*TikTokHideResponse, error) {
	var resp TikTokHideResponse
	if err := c.core.Do(ctx, "/open/v1/rpa/task/tiktokHide", req, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

type TikTokHideAsiaRequest struct {
	// Task name, up to 128 characters.
	Name string `json:"name,omitempty"`
	// Remarks, up to 200 characters.
	Remark string `json:"remark,omitempty"`
	// Scheduled execution time as a Unix timestamp in seconds. Required.
	ScheduleAt int64 `json:"scheduleAt"`
	// Cloud phone ID. Required.
	ID string `json:"id"`
	// Number value, range: [0, 999].If not specified or set to 0, all items will be hidden.
	Number *int `json:"number,omitempty"`
}

type TikTokHideAsiaResponse struct {
	// ID assigned to the asynchronous task.
	TaskID string `json:"taskId"`
}

// TikTokHideAsia: https://open.geelark.com/api/rpa-task-tiktokHideAsia
func (c *Client) TikTokHideAsia(ctx context.Context, req TikTokHideAsiaRequest) (*TikTokHideAsiaResponse, error) {
	var resp TikTokHideAsiaResponse
	if err := c.core.Do(ctx, "/open/v1/rpa/task/tiktokHideAsia", req, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

type TikTokDelRequest struct {
	// Task name, up to 128 characters.
	Name string `json:"name,omitempty"`
	// Remarks, up to 200 characters.
	Remark string `json:"remark,omitempty"`
	// Scheduled execution time as a Unix timestamp in seconds. Required.
	ScheduleAt int64 `json:"scheduleAt"`
	// Cloud phone ID. Required.
	ID string `json:"id"`
}

type TikTokDelResponse struct {
	// ID assigned to the asynchronous task.
	TaskID string `json:"taskId"`
}

// TikTokDel: https://open.geelark.com/api/rpa-task-tiktokDel
func (c *Client) TikTokDel(ctx context.Context, req TikTokDelRequest) (*TikTokDelResponse, error) {
	var resp TikTokDelResponse
	if err := c.core.Do(ctx, "/open/v1/rpa/task/tiktokDel", req, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

type TikTokDelAsiaRequest struct {
	// Task name, up to 128 characters.
	Name string `json:"name,omitempty"`
	// Remarks, up to 200 characters.
	Remark string `json:"remark,omitempty"`
	// Scheduled execution time as a Unix timestamp in seconds. Required.
	ScheduleAt int64 `json:"scheduleAt"`
	// Cloud phone ID. Required.
	ID string `json:"id"`
}

type TikTokDelAsiaResponse struct {
	// ID assigned to the asynchronous task.
	TaskID string `json:"taskId"`
}

// TikTokDelAsia: https://open.geelark.com/api/rpa-task-tiktokDelAsia
func (c *Client) TikTokDelAsia(ctx context.Context, req TikTokDelAsiaRequest) (*TikTokDelAsiaResponse, error) {
	var resp TikTokDelAsiaResponse
	if err := c.core.Do(ctx, "/open/v1/rpa/task/tiktokDelAsia", req, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

type TikTokCommentInterceptionRequest struct {
	// Task name, up to 128 characters.
	Name string `json:"name,omitempty"`
	// Remarks, up to 200 characters.
	Remark string `json:"remark,omitempty"`
	// Scheduled execution time as a Unix timestamp in seconds. Required.
	ScheduleAt int64 `json:"scheduleAt"`
	// Cloud phone ID. Required.
	ID string `json:"id"`
	// Execution time, in minutes, minimum 1. Required.
	ExecutionTime int `json:"executionTime"`
	// Video keywords. Required.
	VideoKeywords []string `json:"videoKeywords"`
	// Comment keywords. Required.
	CommentKeywords []string `json:"commentKeywords"`
	// Comment. Required.
	ReplyComment []string `json:"replyComment"`
	// Date limit, enter the number of days, data before the corresponding day is not
	// commented.
	TimeLimitDays *int `json:"timeLimitDays,omitempty"`
}

type TikTokCommentInterceptionResponse struct {
	// ID assigned to the asynchronous task.
	TaskID string `json:"taskId"`
}

// TikTokCommentInterception: https://open.geelark.com/api/rpa-task-tiktokCommentInterception
func (c *Client) TikTokCommentInterception(ctx context.Context, req TikTokCommentInterceptionRequest) (*TikTokCommentInterceptionResponse, error) {
	var resp TikTokCommentInterceptionResponse
	if err := c.core.Do(ctx, "/open/v1/rpa/task/tiktokCommentInterception", req, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

type TikTokCommentInterceptionAsiaRequest struct {
	// Task name, up to 128 characters.
	Name string `json:"name,omitempty"`
	// Remarks, up to 200 characters.
	Remark string `json:"remark,omitempty"`
	// Scheduled execution time as a Unix timestamp in seconds. Required.
	ScheduleAt int64 `json:"scheduleAt"`
	// Cloud phone ID. Required.
	ID string `json:"id"`
	// Execution time, in minutes, minimum 1. Required.
	ExecutionTime int `json:"executionTime"`
	// Video keywords. Required.
	VideoKeywords []string `json:"videoKeywords"`
	// Comment keywords. Required.
	CommentKeywords []string `json:"commentKeywords"`
	// Comment. Required.
	ReplyComment []string `json:"replyComment"`
	// Date limit, enter the number of days, data before the corresponding day is not
	// commented.
	TimeLimitDays *int `json:"timeLimitDays,omitempty"`
}

type TikTokCommentInterceptionAsiaResponse struct {
	// ID assigned to the asynchronous task.
	TaskID string `json:"taskId"`
}

// TikTokCommentInterceptionAsia: https://open.geelark.com/api/rpa-task-tiktokCommentInterceptionAsia
func (c *Client) TikTokCommentInterceptionAsia(ctx context.Context, req TikTokCommentInterceptionAsiaRequest) (*TikTokCommentInterceptionAsiaResponse, error) {
	var resp TikTokCommentInterceptionAsiaResponse
	if err := c.core.Do(ctx, "/open/v1/rpa/task/tiktokCommentInterceptionAsia", req, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

type TikTokDeleteCommentRequest struct {
	// Task name, up to 32 characters.
	Name string `json:"name,omitempty"`
	// Remark, up to 200 characters.
	Remark string `json:"remark,omitempty"`
	// Scheduled execution time as a Unix timestamp in seconds. Required.
	ScheduleAt int64 `json:"scheduleAt"`
	// Cloud phone ID. Required.
	ID string `json:"id"`
	// Keywords, up to 100 keywords, with a maximum of 100 characters per keyword. Required.
	Keywords []string `json:"keywords"`
}

type TikTokDeleteCommentResponse struct {
	// ID assigned to the asynchronous task.
	TaskID string `json:"taskId"`
}

// TikTokDeleteComment: https://open.geelark.com/api/rpa-task-tiktokDeleteComment
func (c *Client) TikTokDeleteComment(ctx context.Context, req TikTokDeleteCommentRequest) (*TikTokDeleteCommentResponse, error) {
	var resp TikTokDeleteCommentResponse
	if err := c.core.Do(ctx, "/open/v1/rpa/task/tiktokDeleteComment", req, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

type TikTokDeleteCommentAsiaRequest struct {
	// Task name, up to 32 characters.
	Name string `json:"name,omitempty"`
	// Remark, up to 200 characters.
	Remark string `json:"remark,omitempty"`
	// Scheduled execution time as a Unix timestamp in seconds. Required.
	ScheduleAt int64 `json:"scheduleAt"`
	// Cloud phone ID. Required.
	ID string `json:"id"`
	// Keywords, up to 100 keywords, with a maximum of 100 characters per keyword. Required.
	Keywords []string `json:"keywords"`
}

type TikTokDeleteCommentAsiaResponse struct {
	// ID assigned to the asynchronous task.
	TaskID string `json:"taskId"`
}

// TikTokDeleteCommentAsia: https://open.geelark.com/api/rpa-task-tiktokDeleteCommentAsia
func (c *Client) TikTokDeleteCommentAsia(ctx context.Context, req TikTokDeleteCommentAsiaRequest) (*TikTokDeleteCommentAsiaResponse, error) {
	var resp TikTokDeleteCommentAsiaResponse
	if err := c.core.Do(ctx, "/open/v1/rpa/task/tiktokDeleteCommentAsia", req, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
