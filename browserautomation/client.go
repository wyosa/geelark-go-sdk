package browserautomation

import (
	"context"

	"github.com/wyosa/geelark-go-sdk/internal/clientcore"
)

type Client struct {
	core *clientcore.Client
}

func New(core *clientcore.Client) *Client { return &Client{core: core} }

type SearchRequest struct {
	Page     int      `json:"page,omitempty"`
	PageSize int      `json:"pageSize,omitempty"`
	TaskIDs  []string `json:"taskIds,omitempty"`
}

type SearchResponseListItem struct {
	ID         string `json:"id"`
	Eid        string `json:"eid"`
	Name       string `json:"name"`
	Remark     string `json:"remark"`
	SerialName string `json:"serialName"`
	Status     int    `json:"status"`
	StartAt    int    `json:"startAt"`
	FinishAt   int    `json:"finishAt"`
	Cost       int    `json:"cost"`
	ResultCode int    `json:"resultCode"`
	ResultDesc string `json:"resultDesc"`
	ScheduleAt int    `json:"scheduleAt"`
}

type SearchResponse struct {
	Total    int                      `json:"total"`
	Page     int                      `json:"page"`
	PageSize int                      `json:"pageSize"`
	List     []SearchResponseListItem `json:"list"`
}

// Search calls /api/v1/browser/task/search.
func (c *Client) Search(ctx context.Context, req SearchRequest) (*SearchResponse, error) {
	var resp SearchResponse
	if err := c.core.Do(ctx, "/api/v1/browser/task/search", req, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

type CancelRequest struct {
	TaskID string `json:"taskId"`
}

type CancelResponse struct{}

// Cancel calls /api/v1/browser/task/cancel.
func (c *Client) Cancel(ctx context.Context, req CancelRequest) (*CancelResponse, error) {
	var resp CancelResponse
	if err := c.core.Do(ctx, "/api/v1/browser/task/cancel", req, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

type RestartRequest struct {
	TaskID string `json:"taskId"`
}

type RestartResponse struct{}

// Restart calls /api/v1/browser/task/restart.
func (c *Client) Restart(ctx context.Context, req RestartRequest) (*RestartResponse, error) {
	var resp RestartResponse
	if err := c.core.Do(ctx, "/api/v1/browser/task/restart", req, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

type DetailRequest struct {
	TaskID string `json:"taskId"`
}

type DetailResponse struct {
	ID         string   `json:"id"`
	Eid        string   `json:"eid"`
	Name       string   `json:"name"`
	Remark     string   `json:"remark"`
	SerialName string   `json:"serialName"`
	Status     int      `json:"status"`
	StartAt    int      `json:"startAt"`
	FinishAt   int      `json:"finishAt"`
	Cost       int      `json:"cost"`
	ResultCode int      `json:"resultCode"`
	ResultDesc string   `json:"resultDesc"`
	ScheduleAt int      `json:"scheduleAt"`
	Logs       []string `json:"logs"`
}

// Detail calls /api/v1/browser/task/detail.
func (c *Client) Detail(ctx context.Context, req DetailRequest) (*DetailResponse, error) {
	var resp DetailResponse
	if err := c.core.Do(ctx, "/api/v1/browser/task/detail", req, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

type FlowRequest struct {
	Page     int `json:"page,omitempty"`
	PageSize int `json:"pageSize,omitempty"`
}

type FlowResponseListItem struct {
	ID     string   `json:"id"`
	Title  string   `json:"title"`
	Desc   string   `json:"desc"`
	Params []string `json:"params"`
}

type FlowResponse struct {
	Total    int                    `json:"total"`
	Page     int                    `json:"page"`
	PageSize int                    `json:"pageSize"`
	List     []FlowResponseListItem `json:"list"`
}

// Flow calls /api/v1/browser/task/flow.
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
	Eid        string             `json:"eid"`
	Name       string             `json:"name,omitempty"`
	Remark     string             `json:"remark,omitempty"`
	ScheduleAt int                `json:"scheduleAt"`
	FlowID     string             `json:"flowId"`
	ParamMap   AddRequestParamMap `json:"paramMap,omitempty"`
}

type AddResponse struct {
	TaskID string `json:"taskId"`
}

// Add calls /api/v1/browser/task/add.
func (c *Client) Add(ctx context.Context, req AddRequest) (*AddResponse, error) {
	var resp AddResponse
	if err := c.core.Do(ctx, "/api/v1/browser/task/add", req, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

type TikTokSearchRequest struct {
	Eid           string `json:"eid"`
	Name          string `json:"name,omitempty"`
	Remark        string `json:"remark,omitempty"`
	ScheduleAt    int    `json:"scheduleAt"`
	SearchKeyword string `json:"searchKeyword"`
}

type TikTokSearchResponse struct {
	TaskID string `json:"taskId"`
}

// TikTokSearch calls /api/v1/browser/task/tiktokSearch.
func (c *Client) TikTokSearch(ctx context.Context, req TikTokSearchRequest) (*TikTokSearchResponse, error) {
	var resp TikTokSearchResponse
	if err := c.core.Do(ctx, "/api/v1/browser/task/tiktokSearch", req, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

type TikTokCommentRequest struct {
	Eid        string   `json:"eid"`
	Name       string   `json:"name,omitempty"`
	Remark     string   `json:"remark,omitempty"`
	ScheduleAt int      `json:"scheduleAt"`
	Comments   []string `json:"comments"`
}

type TikTokCommentResponse struct {
	TaskID string `json:"taskId"`
}

// TikTokComment calls /api/v1/browser/task/tiktokComment.
func (c *Client) TikTokComment(ctx context.Context, req TikTokCommentRequest) (*TikTokCommentResponse, error) {
	var resp TikTokCommentResponse
	if err := c.core.Do(ctx, "/api/v1/browser/task/tiktokComment", req, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

type TikTokLikeRequest struct {
	Eid        string `json:"eid"`
	Name       string `json:"name,omitempty"`
	Remark     string `json:"remark,omitempty"`
	ScheduleAt int    `json:"scheduleAt"`
	VideoLink  string `json:"videoLink"`
	Comment    string `json:"comment,omitempty"`
}

type TikTokLikeResponse struct {
	TaskID string `json:"taskId"`
}

// TikTokLike calls /api/v1/browser/task/tiktokLike.
func (c *Client) TikTokLike(ctx context.Context, req TikTokLikeRequest) (*TikTokLikeResponse, error) {
	var resp TikTokLikeResponse
	if err := c.core.Do(ctx, "/api/v1/browser/task/tiktokLike", req, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

type FacebookPostRequest struct {
	Eid        string `json:"eid"`
	Name       string `json:"name,omitempty"`
	Remark     string `json:"remark,omitempty"`
	ScheduleAt int    `json:"scheduleAt"`
	Content    string `json:"content"`
}

type FacebookPostResponse struct {
	TaskID string `json:"taskId"`
}

// FacebookPost calls /api/v1/browser/task/facebookPost.
func (c *Client) FacebookPost(ctx context.Context, req FacebookPostRequest) (*FacebookPostResponse, error) {
	var resp FacebookPostResponse
	if err := c.core.Do(ctx, "/api/v1/browser/task/facebookPost", req, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

type FacebookFriendsRequest struct {
	Eid        string `json:"eid"`
	Name       string `json:"name,omitempty"`
	Remark     string `json:"remark,omitempty"`
	ScheduleAt int    `json:"scheduleAt"`
}

type FacebookFriendsResponse struct {
	TaskID string `json:"taskId"`
}

// FacebookFriends calls /api/v1/browser/task/facebookFriends.
func (c *Client) FacebookFriends(ctx context.Context, req FacebookFriendsRequest) (*FacebookFriendsResponse, error) {
	var resp FacebookFriendsResponse
	if err := c.core.Do(ctx, "/api/v1/browser/task/facebookFriends", req, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

type FacebookLikeRequest struct {
	Eid        string `json:"eid"`
	Name       string `json:"name,omitempty"`
	Remark     string `json:"remark,omitempty"`
	ScheduleAt int    `json:"scheduleAt"`
}

type FacebookLikeResponse struct {
	TaskID string `json:"taskId"`
}

// FacebookLike calls /api/v1/browser/task/facebookLike.
func (c *Client) FacebookLike(ctx context.Context, req FacebookLikeRequest) (*FacebookLikeResponse, error) {
	var resp FacebookLikeResponse
	if err := c.core.Do(ctx, "/api/v1/browser/task/facebookLike", req, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

type FacebookHomepageRequest struct {
	Eid          string   `json:"eid"`
	Name         string   `json:"name,omitempty"`
	Remark       string   `json:"remark,omitempty"`
	ScheduleAt   int      `json:"scheduleAt"`
	MainPageName string   `json:"mainPageName"`
	Category     []string `json:"category"`
}

type FacebookHomepageResponse struct {
	TaskID string `json:"taskId"`
}

// FacebookHomepage calls /api/v1/browser/task/facebookHomepage.
func (c *Client) FacebookHomepage(ctx context.Context, req FacebookHomepageRequest) (*FacebookHomepageResponse, error) {
	var resp FacebookHomepageResponse
	if err := c.core.Do(ctx, "/api/v1/browser/task/facebookHomepage", req, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

type InstagramLikeRequest struct {
	Eid        string `json:"eid"`
	Name       string `json:"name,omitempty"`
	Remark     string `json:"remark,omitempty"`
	ScheduleAt int    `json:"scheduleAt"`
}

type InstagramLikeResponse struct {
	TaskID string `json:"taskId"`
}

// InstagramLike calls /api/v1/browser/task/instagramLike.
func (c *Client) InstagramLike(ctx context.Context, req InstagramLikeRequest) (*InstagramLikeResponse, error) {
	var resp InstagramLikeResponse
	if err := c.core.Do(ctx, "/api/v1/browser/task/instagramLike", req, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

type InstagramSearchRequest struct {
	Eid            string   `json:"eid"`
	Name           string   `json:"name,omitempty"`
	Remark         string   `json:"remark,omitempty"`
	ScheduleAt     int      `json:"scheduleAt"`
	SearchKeywords []string `json:"searchKeywords"`
}

type InstagramSearchResponse struct {
	TaskID string `json:"taskId"`
}

// InstagramSearch calls /api/v1/browser/task/instagramSearch.
func (c *Client) InstagramSearch(ctx context.Context, req InstagramSearchRequest) (*InstagramSearchResponse, error) {
	var resp InstagramSearchResponse
	if err := c.core.Do(ctx, "/api/v1/browser/task/instagramSearch", req, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

type YouTubeWatchRequest struct {
	Eid           string `json:"eid"`
	Name          string `json:"name,omitempty"`
	Remark        string `json:"remark,omitempty"`
	ScheduleAt    int    `json:"scheduleAt"`
	SearchKeyword string `json:"searchKeyword"`
	Title         string `json:"title"`
	Comment       string `json:"comment"`
}

type YouTubeWatchResponse struct {
	TaskID string `json:"taskId"`
}

// YouTubeWatch calls /api/v1/browser/task/youtubeWatch.
func (c *Client) YouTubeWatch(ctx context.Context, req YouTubeWatchRequest) (*YouTubeWatchResponse, error) {
	var resp YouTubeWatchResponse
	if err := c.core.Do(ctx, "/api/v1/browser/task/youtubeWatch", req, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

type XPostRequest struct {
	Eid        string `json:"eid"`
	Name       string `json:"name,omitempty"`
	Remark     string `json:"remark,omitempty"`
	ScheduleAt int    `json:"scheduleAt"`
	Content    string `json:"content"`
}

type XPostResponse struct {
	TaskID string `json:"taskId"`
}

// XPost calls /api/v1/browser/task/xPost.
func (c *Client) XPost(ctx context.Context, req XPostRequest) (*XPostResponse, error) {
	var resp XPostResponse
	if err := c.core.Do(ctx, "/api/v1/browser/task/xPost", req, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

type XLikeRequest struct {
	Eid        string `json:"eid"`
	Name       string `json:"name,omitempty"`
	Remark     string `json:"remark,omitempty"`
	ScheduleAt int    `json:"scheduleAt"`
}

type XLikeResponse struct {
	TaskID string `json:"taskId"`
}

// XLike calls /api/v1/browser/task/xLike.
func (c *Client) XLike(ctx context.Context, req XLikeRequest) (*XLikeResponse, error) {
	var resp XLikeResponse
	if err := c.core.Do(ctx, "/api/v1/browser/task/xLike", req, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

type RedditLikeRequest struct {
	Eid            string   `json:"eid"`
	Name           string   `json:"name,omitempty"`
	Remark         string   `json:"remark,omitempty"`
	ScheduleAt     int      `json:"scheduleAt"`
	SearchKeywords []string `json:"searchKeywords"`
}

type RedditLikeResponse struct {
	TaskID string `json:"taskId"`
}

// RedditLike calls /api/v1/browser/task/redditLike.
func (c *Client) RedditLike(ctx context.Context, req RedditLikeRequest) (*RedditLikeResponse, error) {
	var resp RedditLikeResponse
	if err := c.core.Do(ctx, "/api/v1/browser/task/redditLike", req, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

type CookieBotRequest struct {
	Eid        string   `json:"eid"`
	Name       string   `json:"name,omitempty"`
	Remark     string   `json:"remark,omitempty"`
	ScheduleAt int      `json:"scheduleAt"`
	Pages      []string `json:"pages"`
}

type CookieBotResponse struct {
	TaskID string `json:"taskId"`
}

// CookieBot calls /api/v1/browser/task/cookieBot.
func (c *Client) CookieBot(ctx context.Context, req CookieBotRequest) (*CookieBotResponse, error) {
	var resp CookieBotResponse
	if err := c.core.Do(ctx, "/api/v1/browser/task/cookieBot", req, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
