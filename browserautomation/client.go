// Package browserautomation provides scheduled browser tasks and custom flows in the GeeLark
// desktop client.
package browserautomation

import (
	"context"

	"github.com/wyosa/geelark-go-sdk/common"
)

type Client struct {
	core common.Transport
}

func New(core common.Transport) *Client { return &Client{core: core} }

type SearchRequest struct {
	// Page number, defaults to 1 if not specified.
	Page int `json:"page,omitempty"`
	// Page size, maximum 100, defaults to 10 if not specified.
	PageSize int `json:"pageSize,omitempty"`
	// Task ID.
	TaskIDs []string `json:"taskIds,omitempty"`
}

type SearchResponseListItem struct {
	// Task ID.
	ID string `json:"id"`
	// Profile ID.
	Eid string `json:"eid"`
	// Task Name.
	Name string `json:"name"`
	// Task Remark.
	Remark string `json:"remark"`
	// Profile Name.
	SerialName string `json:"serialName"`
	// Task Status, 1-Waiting for Execution 2-Executing 3-Task Completed 4-Task Failed 7-Task
	// Cancelled.
	Status int `json:"status"`
	// Start Time, Second-level Timestamp.
	StartAt int64 `json:"startAt"`
	// End Time, Second-level Timestamp.
	FinishAt int64 `json:"finishAt"`
	// Time Elapsed in Seconds.
	Cost int `json:"cost"`
	// Task Result Code, Refer to Task Failure Code and Reason.
	ResultCode int `json:"resultCode"`
	// Task Result Description, Refer to Task Failure Code and Reason.
	ResultDesc string `json:"resultDesc"`
	// Scheduled execution time as a Unix timestamp in seconds.
	ScheduleAt int64 `json:"scheduleAt"`
}

type SearchResponse struct {
	// Total number of data.
	Total int `json:"total"`
	// Page number.
	Page     int `json:"page"`
	PageSize int `json:"pageSize"`
	// Data list.
	List []SearchResponseListItem `json:"list"`
}

// Search: https://open.geelark.com/api/browser-task-search
func (c *Client) Search(ctx context.Context, req SearchRequest) (*SearchResponse, error) {
	var resp SearchResponse
	if err := c.core.Do(ctx, "/api/v1/browser/task/search", req, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

type CancelRequest struct {
	// task id. Required.
	TaskID string `json:"taskId"`
}

type CancelResponse struct{}

// Cancel: https://open.geelark.com/api/browser-task-cancel
//
// Only waiting or running tasks can be cancelled.
func (c *Client) Cancel(ctx context.Context, req CancelRequest) (*CancelResponse, error) {
	var resp CancelResponse
	if err := c.core.Do(ctx, "/api/v1/browser/task/cancel", req, nil); err != nil {
		return nil, err
	}
	return &resp, nil
}

type RestartRequest struct {
	// task id. Required.
	TaskID string `json:"taskId"`
}

type RestartResponse struct{}

// Restart: https://open.geelark.com/api/browser-task-restart
//
// Completed, failed, or cancelled tasks can be retried.
func (c *Client) Restart(ctx context.Context, req RestartRequest) (*RestartResponse, error) {
	var resp RestartResponse
	if err := c.core.Do(ctx, "/api/v1/browser/task/restart", req, nil); err != nil {
		return nil, err
	}
	return &resp, nil
}

type DetailRequest struct {
	// Task ID. Required.
	TaskID string `json:"taskId"`
}

type DetailResponse struct {
	// Task ID.
	ID string `json:"id"`
	// Profile ID.
	Eid string `json:"eid"`
	// Task Name.
	Name string `json:"name"`
	// Task Remark.
	Remark string `json:"remark"`
	// Profile Name.
	SerialName string `json:"serialName"`
	// Task Status, 1-Waiting for Execution 2-Executing 3-Task Completed 4-Task Failed
	// 7-Canceled.
	Status int `json:"status"`
	// Start time, timestamp in seconds.
	StartAt int64 `json:"startAt"`
	// End time, timestamp in seconds.
	FinishAt int64 `json:"finishAt"`
	// Time elapsed in seconds.
	Cost int `json:"cost"`
	// Task result code, refer to task failure codes and reasons.
	ResultCode int `json:"resultCode"`
	// Task result description, refer to task failure codes and reasons.
	ResultDesc string `json:"resultDesc"`
	// Scheduled execution time as a Unix timestamp in seconds.
	ScheduleAt int64 `json:"scheduleAt"`
	// Task logs.
	Logs []string `json:"logs"`
}

// Detail: https://open.geelark.com/api/browser-task-detail
func (c *Client) Detail(ctx context.Context, req DetailRequest) (*DetailResponse, error) {
	var resp DetailResponse
	if err := c.core.Do(ctx, "/api/v1/browser/task/detail", req, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

type FlowRequest struct {
	// Page number, defaults to 1 if not specified.
	Page int `json:"page,omitempty"`
	// Page size, maximum 100, defaults to 10 if not specified.
	PageSize int `json:"pageSize,omitempty"`
}

type FlowResponseListItem struct {
	// Task Flow ID.
	ID string `json:"id"`
	// Task Flow Title.
	Title string `json:"title"`
	// Task Flow Remarks.
	Desc string `json:"desc"`
	// Task Flow Parameter Field Names.
	Params []string `json:"params"`
}

type FlowResponse struct {
	// Total number of records.
	Total int `json:"total"`
	// Page number.
	Page     int `json:"page"`
	PageSize int `json:"pageSize"`
	// List of records.
	List []FlowResponseListItem `json:"list"`
}

// Flow: https://open.geelark.com/api/browser-task-flow
func (c *Client) Flow(ctx context.Context, req FlowRequest) (*FlowResponse, error) {
	var resp FlowResponse
	if err := c.core.Do(ctx, "/api/v1/browser/task/flow", req, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// AddRequestParamMap contains the named inputs of a custom browser flow.
// Values may be strings, string slices, numbers, or booleans.
type AddRequestParamMap map[string]any

type AddRequest struct {
	// Profile ID. Required.
	Eid string `json:"eid"`
	// Task name, maximum 128 characters.
	Name string `json:"name,omitempty"`
	// Remark, maximum 200 characters.
	Remark string `json:"remark,omitempty"`
	// Scheduled execution time as a Unix timestamp in seconds. Required.
	ScheduleAt int64 `json:"scheduleAt"`
	// Task flow ID, the ID field returned by the task flow query interface. Required.
	FlowID string `json:"flowId"`
	// Task flow parameters, with corresponding parameter types as follows:; String: string;
	// Batch text: array[string]; Number: number; Boolean: bool; File: array[string].
	ParamMap AddRequestParamMap `json:"paramMap,omitempty"`
}

type AddResponse struct {
	// ID assigned to the asynchronous task.
	TaskID string `json:"taskId"`
}

// Add: https://open.geelark.com/api/browser-task-add
//
// Call Flow first to obtain a flow ID and the input names expected by ParamMap.
func (c *Client) Add(ctx context.Context, req AddRequest) (*AddResponse, error) {
	var resp AddResponse
	if err := c.core.Do(ctx, "/api/v1/browser/task/add", req, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

type TikTokSearchRequest struct {
	// Profile ID. Required.
	Eid string `json:"eid"`
	// Task name, maximum 128 characters.
	Name string `json:"name,omitempty"`
	// Remark, maximum 200 characters.
	Remark string `json:"remark,omitempty"`
	// Scheduled execution time as a Unix timestamp in seconds. Required.
	ScheduleAt int64 `json:"scheduleAt"`
	// search keyword. Required.
	SearchKeyword string `json:"searchKeyword"`
}

type TikTokSearchResponse struct {
	// ID assigned to the asynchronous task.
	TaskID string `json:"taskId"`
}

// TikTokSearch: https://open.geelark.com/api/browser-task-tiktokSearch
func (c *Client) TikTokSearch(ctx context.Context, req TikTokSearchRequest) (*TikTokSearchResponse, error) {
	var resp TikTokSearchResponse
	if err := c.core.Do(ctx, "/api/v1/browser/task/tiktokSearch", req, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

type TikTokCommentRequest struct {
	// Profile ID. Required.
	Eid string `json:"eid"`
	// Task name, maximum 128 characters.
	Name string `json:"name,omitempty"`
	// Remark, maximum 200 characters.
	Remark string `json:"remark,omitempty"`
	// Scheduled execution time as a Unix timestamp in seconds. Required.
	ScheduleAt int64 `json:"scheduleAt"`
	// comments. Required.
	Comments []string `json:"comments"`
}

type TikTokCommentResponse struct {
	// ID assigned to the asynchronous task.
	TaskID string `json:"taskId"`
}

// TikTokComment: https://open.geelark.com/api/browser-task-tiktokComment
func (c *Client) TikTokComment(ctx context.Context, req TikTokCommentRequest) (*TikTokCommentResponse, error) {
	var resp TikTokCommentResponse
	if err := c.core.Do(ctx, "/api/v1/browser/task/tiktokComment", req, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

type TikTokLikeRequest struct {
	// Profile ID. Required.
	Eid string `json:"eid"`
	// Task name, maximum 128 characters.
	Name string `json:"name,omitempty"`
	// Remark, maximum 200 characters.
	Remark string `json:"remark,omitempty"`
	// Scheduled execution time as a Unix timestamp in seconds. Required.
	ScheduleAt int64 `json:"scheduleAt"`
	// Video link. Required.
	VideoLink string `json:"videoLink"`
	Comment   string `json:"comment,omitempty"`
}

type TikTokLikeResponse struct {
	// ID assigned to the asynchronous task.
	TaskID string `json:"taskId"`
}

// TikTokLike: https://open.geelark.com/api/browser-task-tiktokLike
func (c *Client) TikTokLike(ctx context.Context, req TikTokLikeRequest) (*TikTokLikeResponse, error) {
	var resp TikTokLikeResponse
	if err := c.core.Do(ctx, "/api/v1/browser/task/tiktokLike", req, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

type FacebookPostRequest struct {
	// Profile ID. Required.
	Eid string `json:"eid"`
	// Task name, maximum 128 characters.
	Name string `json:"name,omitempty"`
	// Remark, maximum 200 characters.
	Remark string `json:"remark,omitempty"`
	// Scheduled execution time as a Unix timestamp in seconds. Required.
	ScheduleAt int64 `json:"scheduleAt"`
	// Content. Required.
	Content string `json:"content"`
}

type FacebookPostResponse struct {
	// ID assigned to the asynchronous task.
	TaskID string `json:"taskId"`
}

// FacebookPost: https://open.geelark.com/api/browser-task-facebookPost
func (c *Client) FacebookPost(ctx context.Context, req FacebookPostRequest) (*FacebookPostResponse, error) {
	var resp FacebookPostResponse
	if err := c.core.Do(ctx, "/api/v1/browser/task/facebookPost", req, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

type FacebookFriendsRequest struct {
	// Profile ID. Required.
	Eid string `json:"eid"`
	// Task name, maximum 128 characters.
	Name string `json:"name,omitempty"`
	// Remark, maximum 200 characters.
	Remark string `json:"remark,omitempty"`
	// Scheduled execution time as a Unix timestamp in seconds. Required.
	ScheduleAt int64 `json:"scheduleAt"`
}

type FacebookFriendsResponse struct {
	// ID assigned to the asynchronous task.
	TaskID string `json:"taskId"`
}

// FacebookFriends: https://open.geelark.com/api/browser-task-facebookFriends
func (c *Client) FacebookFriends(ctx context.Context, req FacebookFriendsRequest) (*FacebookFriendsResponse, error) {
	var resp FacebookFriendsResponse
	if err := c.core.Do(ctx, "/api/v1/browser/task/facebookFriends", req, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

type FacebookLikeRequest struct {
	// Profile ID. Required.
	Eid string `json:"eid"`
	// Task name, maximum 128 characters.
	Name string `json:"name,omitempty"`
	// Remark, maximum 200 characters.
	Remark string `json:"remark,omitempty"`
	// Scheduled execution time as a Unix timestamp in seconds. Required.
	ScheduleAt int64 `json:"scheduleAt"`
}

type FacebookLikeResponse struct {
	// ID assigned to the asynchronous task.
	TaskID string `json:"taskId"`
}

// FacebookLike: https://open.geelark.com/api/browser-task-facebookLike
func (c *Client) FacebookLike(ctx context.Context, req FacebookLikeRequest) (*FacebookLikeResponse, error) {
	var resp FacebookLikeResponse
	if err := c.core.Do(ctx, "/api/v1/browser/task/facebookLike", req, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

type FacebookHomepageRequest struct {
	// Profile ID. Required.
	Eid string `json:"eid"`
	// Task name, maximum 128 characters.
	Name string `json:"name,omitempty"`
	// Remark, maximum 200 characters.
	Remark string `json:"remark,omitempty"`
	// Scheduled execution time as a Unix timestamp in seconds. Required.
	ScheduleAt int64 `json:"scheduleAt"`
	// Homepage name data. Required.
	MainPageName string `json:"mainPageName"`
	// Homepage category data. Required.
	Category []string `json:"category"`
}

type FacebookHomepageResponse struct {
	// ID assigned to the asynchronous task.
	TaskID string `json:"taskId"`
}

// FacebookHomepage: https://open.geelark.com/api/browser-task-facebookHomepage
func (c *Client) FacebookHomepage(ctx context.Context, req FacebookHomepageRequest) (*FacebookHomepageResponse, error) {
	var resp FacebookHomepageResponse
	if err := c.core.Do(ctx, "/api/v1/browser/task/facebookHomepage", req, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

type InstagramLikeRequest struct {
	// Profile ID. Required.
	Eid string `json:"eid"`
	// Task name, maximum 128 characters.
	Name string `json:"name,omitempty"`
	// Remark, maximum 200 characters.
	Remark string `json:"remark,omitempty"`
	// Scheduled execution time as a Unix timestamp in seconds. Required.
	ScheduleAt int64 `json:"scheduleAt"`
}

type InstagramLikeResponse struct {
	// ID assigned to the asynchronous task.
	TaskID string `json:"taskId"`
}

// InstagramLike: https://open.geelark.com/api/browser-task-instagramLike
func (c *Client) InstagramLike(ctx context.Context, req InstagramLikeRequest) (*InstagramLikeResponse, error) {
	var resp InstagramLikeResponse
	if err := c.core.Do(ctx, "/api/v1/browser/task/instagramLike", req, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

type InstagramSearchRequest struct {
	// Profile ID. Required.
	Eid string `json:"eid"`
	// Task name, maximum 128 characters.
	Name string `json:"name,omitempty"`
	// Remark, maximum 200 characters.
	Remark string `json:"remark,omitempty"`
	// Scheduled execution time as a Unix timestamp in seconds. Required.
	ScheduleAt int64 `json:"scheduleAt"`
	// Search keywords. Required.
	SearchKeywords []string `json:"searchKeywords"`
}

type InstagramSearchResponse struct {
	// ID assigned to the asynchronous task.
	TaskID string `json:"taskId"`
}

// InstagramSearch: https://open.geelark.com/api/browser-task-instagramSearch
func (c *Client) InstagramSearch(ctx context.Context, req InstagramSearchRequest) (*InstagramSearchResponse, error) {
	var resp InstagramSearchResponse
	if err := c.core.Do(ctx, "/api/v1/browser/task/instagramSearch", req, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

type YouTubeWatchRequest struct {
	// Profile ID. Required.
	Eid string `json:"eid"`
	// Task name, maximum 128 characters.
	Name string `json:"name,omitempty"`
	// Remark, maximum 200 characters.
	Remark string `json:"remark,omitempty"`
	// Scheduled execution time as a Unix timestamp in seconds. Required.
	ScheduleAt int64 `json:"scheduleAt"`
	// search keyword. Required.
	SearchKeyword string `json:"searchKeyword"`
	// title. Required.
	Title string `json:"title"`
	// comment. Required.
	Comment string `json:"comment"`
}

type YouTubeWatchResponse struct {
	// ID assigned to the asynchronous task.
	TaskID string `json:"taskId"`
}

// YouTubeWatch: https://open.geelark.com/api/browser-task-youtubeWatch
func (c *Client) YouTubeWatch(ctx context.Context, req YouTubeWatchRequest) (*YouTubeWatchResponse, error) {
	var resp YouTubeWatchResponse
	if err := c.core.Do(ctx, "/api/v1/browser/task/youtubeWatch", req, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

type XPostRequest struct {
	// Profile ID. Required.
	Eid string `json:"eid"`
	// Task name, maximum 128 characters.
	Name string `json:"name,omitempty"`
	// Remark, maximum 200 characters.
	Remark string `json:"remark,omitempty"`
	// Scheduled execution time as a Unix timestamp in seconds. Required.
	ScheduleAt int64 `json:"scheduleAt"`
	// Content. Required.
	Content string `json:"content"`
}

type XPostResponse struct {
	// ID assigned to the asynchronous task.
	TaskID string `json:"taskId"`
}

// XPost: https://open.geelark.com/api/browser-task-xPost
func (c *Client) XPost(ctx context.Context, req XPostRequest) (*XPostResponse, error) {
	var resp XPostResponse
	if err := c.core.Do(ctx, "/api/v1/browser/task/xPost", req, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

type XLikeRequest struct {
	// Profile ID. Required.
	Eid string `json:"eid"`
	// Task name, maximum 128 characters.
	Name string `json:"name,omitempty"`
	// Remark, maximum 200 characters.
	Remark string `json:"remark,omitempty"`
	// Scheduled execution time as a Unix timestamp in seconds. Required.
	ScheduleAt int64 `json:"scheduleAt"`
}

type XLikeResponse struct {
	// ID assigned to the asynchronous task.
	TaskID string `json:"taskId"`
}

// XLike: https://open.geelark.com/api/browser-task-xLike
func (c *Client) XLike(ctx context.Context, req XLikeRequest) (*XLikeResponse, error) {
	var resp XLikeResponse
	if err := c.core.Do(ctx, "/api/v1/browser/task/xLike", req, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

type XAiCommentRequest struct {
	// Task name, maximum 128 characters.
	Name string `json:"name,omitempty"`
	// Remark, maximum 200 characters.
	Remark string `json:"remark,omitempty"`
	// Scheduled execution time as a Unix timestamp in seconds. Required.
	ScheduleAt int64 `json:"scheduleAt"`
	// Profile ID. Required.
	Eid string `json:"eid"`
	// Whether to use AI: UseAIComment (1, Pro users only) or UseOwnComment (2, post Comment).
	// Required; zero is not a valid value.
	UseAI int `json:"useAI"`
	// Comment content, up to 300 characters; required when UseAI is UseOwnComment. Always
	// sent because the parameter table marks it required.
	Comment string `json:"comment"`
	// Specified links, up to 100 items, each up to 100 characters.
	Links []string `json:"links,omitempty"`
	// Comment probability, 0-100, default is 30.
	CommentProbability *int `json:"commentProbability,omitempty"`
	// Search keywords, up to 100 items, each up to 100 characters.
	SearchKeywords []string `json:"searchKeywords,omitempty"`
	// Whether to like, default is no.
	LikeVideo *bool `json:"likeVideo,omitempty"`
	// Comment image as a local file path on the computer running the desktop client, maximum
	// 500 characters. Takes effect when UseAI is UseOwnComment.
	ImagePath string `json:"imagePath,omitempty"`
}

type XAiCommentResponse struct {
	// ID assigned to the asynchronous task.
	TaskID string `json:"taskId"`
}

// XAiComment: https://open.geelark.com/api/browser-task-xAIComment
//
// Comments on X (Twitter) posts with AI-generated or supplied comments.
func (c *Client) XAiComment(ctx context.Context, req XAiCommentRequest) (*XAiCommentResponse, error) {
	var resp XAiCommentResponse
	if err := c.core.Do(ctx, "/api/v1/browser/task/xAiComment", req, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

type RedditLikeRequest struct {
	// Profile ID. Required.
	Eid string `json:"eid"`
	// Task name, maximum 128 characters.
	Name string `json:"name,omitempty"`
	// Remark, maximum 200 characters.
	Remark string `json:"remark,omitempty"`
	// Scheduled execution time as a Unix timestamp in seconds. Required.
	ScheduleAt int64 `json:"scheduleAt"`
	// Search keywords. Required.
	SearchKeywords []string `json:"searchKeywords"`
}

type RedditLikeResponse struct {
	// ID assigned to the asynchronous task.
	TaskID string `json:"taskId"`
}

// RedditLike: https://open.geelark.com/api/browser-task-redditLike
func (c *Client) RedditLike(ctx context.Context, req RedditLikeRequest) (*RedditLikeResponse, error) {
	var resp RedditLikeResponse
	if err := c.core.Do(ctx, "/api/v1/browser/task/redditLike", req, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

type CookieBotRequest struct {
	// Profile ID. Required.
	Eid string `json:"eid"`
	// Task name, maximum 128 characters.
	Name string `json:"name,omitempty"`
	// Remark, maximum 200 characters.
	Remark string `json:"remark,omitempty"`
	// Scheduled execution time as a Unix timestamp in seconds. Required.
	ScheduleAt int64 `json:"scheduleAt"`
	// The webpage to be accessed. Required.
	Pages []string `json:"pages"`
}

type CookieBotResponse struct {
	// ID assigned to the asynchronous task.
	TaskID string `json:"taskId"`
}

// CookieBot: https://open.geelark.com/api/browser-task-cookieBot
func (c *Client) CookieBot(ctx context.Context, req CookieBotRequest) (*CookieBotResponse, error) {
	var resp CookieBotResponse
	if err := c.core.Do(ctx, "/api/v1/browser/task/cookieBot", req, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
