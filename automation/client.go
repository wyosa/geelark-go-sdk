package automation

import (
	"context"

	"github.com/wyosa/geelark-go-sdk/internal/clientcore"
)

type Client struct {
	core *clientcore.Client
}

func New(core *clientcore.Client) *Client { return &Client{core: core} }

type QueryTaskRequest struct {
	IDs []string `json:"ids"`
}

type QueryTaskResponseItemsItem struct {
	ID         string `json:"id"`
	PlanName   string `json:"planName"`
	TaskType   int    `json:"taskType"`
	SerialName string `json:"serialName"`
	EnvID      string `json:"envId"`
	ScheduleAt int    `json:"scheduleAt"`
	Status     int    `json:"status"`
	FailCode   int    `json:"failCode,omitempty"`
	FailDesc   string `json:"failDesc,omitempty"`
	Cost       int    `json:"cost,omitempty"`
	ShareLink  string `json:"shareLink,omitempty"`
}

type QueryTaskResponse struct {
	Total int                          `json:"total"`
	Items []QueryTaskResponseItemsItem `json:"items"`
}

// QueryTask calls /open/v1/task/query.
func (c *Client) QueryTask(ctx context.Context, req QueryTaskRequest) (*QueryTaskResponse, error) {
	var resp QueryTaskResponse
	if err := c.core.Do(ctx, "/open/v1/task/query", req, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

type ListTasksRequest struct {
	Size   *int     `json:"size,omitempty"`
	LastID string   `json:"lastId,omitempty"`
	IDs    []string `json:"ids,omitempty"`
}

type ListTasksResponseItemsItem struct {
	ID         string `json:"id"`
	PlanName   string `json:"planName"`
	TaskType   int    `json:"taskType"`
	SerialName string `json:"serialName"`
	EnvID      string `json:"envId"`
	ScheduleAt int    `json:"scheduleAt"`
	Status     int    `json:"status"`
	FailCode   int    `json:"failCode,omitempty"`
	FailDesc   string `json:"failDesc,omitempty"`
	Cost       int    `json:"cost,omitempty"`
}

type ListTasksResponse struct {
	Total int                          `json:"total"`
	Items []ListTasksResponseItemsItem `json:"items"`
}

// ListTasks calls /open/v1/task/historyRecords.
func (c *Client) ListTasks(ctx context.Context, req ListTasksRequest) (*ListTasksResponse, error) {
	var resp ListTasksResponse
	if err := c.core.Do(ctx, "/open/v1/task/historyRecords", req, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

type CancelTaskRequest struct {
	IDs []string `json:"ids"`
}

type CancelTaskResponse struct {
	TotalAmount   int                                 `json:"totalAmount"`
	SuccessAmount int                                 `json:"successAmount"`
	FailAmount    int                                 `json:"failAmount"`
	FailDetails   []CancelTaskResponseFailDetailsItem `json:"failDetails"`
}

type CancelTaskResponseFailDetailsItem struct {
	ID   string `json:"id"`
	Code int    `json:"code"`
	Msg  string `json:"msg"`
}

// CancelTask calls /open/v1/task/cancel.
func (c *Client) CancelTask(ctx context.Context, req CancelTaskRequest) (*CancelTaskResponse, error) {
	var resp CancelTaskResponse
	if err := c.core.Do(ctx, "/open/v1/task/cancel", req, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

type RetryTaskRequest struct {
	IDs []string `json:"ids"`
}

type RetryTaskResponse struct {
	TotalAmount   int                                `json:"totalAmount"`
	SuccessAmount int                                `json:"successAmount"`
	FailAmount    int                                `json:"failAmount"`
	FailDetails   []RetryTaskResponseFailDetailsItem `json:"failDetails"`
}

type RetryTaskResponseFailDetailsItem struct {
	ID   string `json:"id"`
	Code int    `json:"code"`
	Msg  string `json:"msg"`
}

// RetryTask calls /open/v1/task/restart.
func (c *Client) RetryTask(ctx context.Context, req RetryTaskRequest) (*RetryTaskResponse, error) {
	var resp RetryTaskResponse
	if err := c.core.Do(ctx, "/open/v1/task/restart", req, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

type TaskDetailRequest struct {
	ID          string  `json:"id"`
	SearchAfter []int64 `json:"searchAfter,omitempty"`
}

type TaskDetailResponse struct {
	ID           string   `json:"id"`
	PlanName     string   `json:"planName"`
	TaskType     int      `json:"taskType"`
	SerialName   string   `json:"serialName"`
	EnvID        string   `json:"envId"`
	ScheduleAt   int      `json:"scheduleAt"`
	Status       int      `json:"status"`
	Cost         int      `json:"cost"`
	ResultImages []string `json:"resultImages"`
	Logs         []string `json:"logs"`
	SearchAfter  []int64  `json:"searchAfter"`
	LogContinue  bool     `json:"logContinue"`
	FailCode     int      `json:"failCode,omitempty"`
	FailDesc     string   `json:"failDesc,omitempty"`
}

// TaskDetail calls /open/v1/task/detail.
func (c *Client) TaskDetail(ctx context.Context, req TaskDetailRequest) (*TaskDetailResponse, error) {
	var resp TaskDetailResponse
	if err := c.core.Do(ctx, "/open/v1/task/detail", req, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

type AddTaskRequest struct {
	PlanName string        `json:"planName,omitempty"`
	Remark   string        `json:"remark,omitempty"`
	TaskType int           `json:"taskType"`
	List     []AddTaskItem `json:"list"`
}

// AddTaskItem carries the fields for video, warmup, and image tasks.
// TaskType selects which fields GeeLark requires.
type AddTaskItem struct {
	ScheduleAt        int64               `json:"scheduleAt"`
	EnvID             string              `json:"envId"`
	Video             string              `json:"video,omitempty"`
	VideoDesc         string              `json:"videoDesc,omitempty"`
	ProductID         string              `json:"productId,omitempty"`
	ProductTitle      string              `json:"productTitle,omitempty"`
	RefVideoID        string              `json:"refVideoId,omitempty"`
	MaxTryTimes       *int                `json:"maxTryTimes,omitempty"`
	TimeoutMin        *int                `json:"timeoutMin,omitempty"`
	SameVideoVolume   *int                `json:"sameVideoVolume,omitempty"`
	SourceVideoVolume *int                `json:"sourceVideoVolume,omitempty"`
	MarkAI            *bool               `json:"markAI,omitempty"`
	Cover             string              `json:"cover,omitempty"`
	NeedShareLink     *bool               `json:"needShareLink,omitempty"`
	RandomUse         *bool               `json:"randomUse,omitempty"`
	BrowseTime        *int                `json:"browseTime,omitempty"`
	GameName          string              `json:"gameName,omitempty"`
	Location          string              `json:"location,omitempty"`
	Index             *int                `json:"index,omitempty"`
	QRCode            string              `json:"qrcode,omitempty"`
	Action            string              `json:"action,omitempty"`
	Keywords          []string            `json:"keywords,omitempty"`
	Duration          *int                `json:"duration,omitempty"`
	Probability       *AddTaskProbability `json:"probability,omitempty"`
	Images            []string            `json:"images,omitempty"`
	VideoID           string              `json:"videoId,omitempty"`
	VideoTitle        string              `json:"videoTitle,omitempty"`
	RandomBGM         *bool               `json:"randomBgm,omitempty"`
}

type AddTaskProbability struct {
	HasKey *AddTaskBehavior `json:"hasKey,omitempty"`
	NoKey  *AddTaskBehavior `json:"noKey,omitempty"`
}

type AddTaskBehavior struct {
	Like        *int `json:"like,omitempty"`
	Follow      *int `json:"follow,omitempty"`
	ReadComment *int `json:"readComment,omitempty"`
	Collect     *int `json:"collect,omitempty"`
}

type AddTaskResponse struct {
	TaskIDs []string `json:"taskIds"`
}

// AddTask calls /open/v1/task/add.
func (c *Client) AddTask(ctx context.Context, req AddTaskRequest) (*AddTaskResponse, error) {
	var resp AddTaskResponse
	if err := c.core.Do(ctx, "/open/v1/task/add", req, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

type TikTokLoginRequest struct {
	Name       string `json:"name,omitempty"`
	Remark     string `json:"remark,omitempty"`
	ScheduleAt int    `json:"scheduleAt"`
	ID         string `json:"id"`
	Account    string `json:"account"`
	Password   string `json:"password"`
	TwoFAKey   string `json:"twoFAKey,omitempty"`
}

type TikTokLoginResponse struct {
	TaskID string `json:"taskId"`
}

// TikTokLogin calls /open/v1/rpa/task/tiktokLogin.
func (c *Client) TikTokLogin(ctx context.Context, req TikTokLoginRequest) (*TikTokLoginResponse, error) {
	var resp TikTokLoginResponse
	if err := c.core.Do(ctx, "/open/v1/rpa/task/tiktokLogin", req, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

type TikTokEditRequest struct {
	Name       string `json:"name,omitempty"`
	Remark     string `json:"remark,omitempty"`
	ScheduleAt int    `json:"scheduleAt"`
	ID         string `json:"id"`
	Avatar     string `json:"avatar,omitempty"`
	NickName   string `json:"nickName,omitempty"`
	Bio        string `json:"bio,omitempty"`
	Site       string `json:"site,omitempty"`
	Username   string `json:"username,omitempty"`
}

type TikTokEditResponse struct {
	TaskID string `json:"taskId"`
}

// TikTokEdit calls /open/v1/rpa/task/tiktokEdit.
func (c *Client) TikTokEdit(ctx context.Context, req TikTokEditRequest) (*TikTokEditResponse, error) {
	var resp TikTokEditResponse
	if err := c.core.Do(ctx, "/open/v1/rpa/task/tiktokEdit", req, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

type TikTokRandomStarRequest struct {
	Name            string `json:"name,omitempty"`
	Remark          string `json:"remark,omitempty"`
	ScheduleAt      int    `json:"scheduleAt"`
	ID              string `json:"id"`
	LikeProbability *int   `json:"likeProbability,omitempty"`
}

type TikTokRandomStarResponse struct {
	TaskID string `json:"taskId"`
}

// TikTokRandomStar calls /open/v1/rpa/task/tiktokRandomStar.
func (c *Client) TikTokRandomStar(ctx context.Context, req TikTokRandomStarRequest) (*TikTokRandomStarResponse, error) {
	var resp TikTokRandomStarResponse
	if err := c.core.Do(ctx, "/open/v1/rpa/task/tiktokRandomStar", req, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

type TikTokRandomStarAsiaRequest struct {
	Name            string `json:"name,omitempty"`
	Remark          string `json:"remark,omitempty"`
	ScheduleAt      int    `json:"scheduleAt"`
	ID              string `json:"id"`
	LikeProbability *int   `json:"likeProbability,omitempty"`
}

type TikTokRandomStarAsiaResponse struct {
	TaskID string `json:"taskId"`
}

// TikTokRandomStarAsia calls /open/v1/rpa/task/tiktokRandomStarAsia.
func (c *Client) TikTokRandomStarAsia(ctx context.Context, req TikTokRandomStarAsiaRequest) (*TikTokRandomStarAsiaResponse, error) {
	var resp TikTokRandomStarAsiaResponse
	if err := c.core.Do(ctx, "/open/v1/rpa/task/tiktokRandomStarAsia", req, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

type TikTokRandomCommentRequest struct {
	Name               string   `json:"name,omitempty"`
	Remark             string   `json:"remark,omitempty"`
	ScheduleAt         int      `json:"scheduleAt"`
	ID                 string   `json:"id"`
	UseAi              int      `json:"useAi"`
	Comment            string   `json:"comment"`
	Links              []string `json:"links,omitempty"`
	CommentProbability *int     `json:"commentProbability,omitempty"`
	SearchKeywords     []string `json:"searchKeywords,omitempty"`
	LikeVideo          *bool    `json:"likeVideo,omitempty"`
	ImageURL           string   `json:"imageUrl,omitempty"`
}

type TikTokRandomCommentResponse struct {
	TaskID string `json:"taskId"`
}

// TikTokRandomComment calls /open/v1/rpa/task/tiktokRandomComment.
func (c *Client) TikTokRandomComment(ctx context.Context, req TikTokRandomCommentRequest) (*TikTokRandomCommentResponse, error) {
	var resp TikTokRandomCommentResponse
	if err := c.core.Do(ctx, "/open/v1/rpa/task/tiktokRandomComment", req, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

type TikTokRandomCommentAsiaRequest struct {
	Name               string   `json:"name,omitempty"`
	Remark             string   `json:"remark,omitempty"`
	ScheduleAt         int      `json:"scheduleAt"`
	ID                 string   `json:"id"`
	UseAi              int      `json:"useAi"`
	Comment            string   `json:"comment"`
	Links              []string `json:"links,omitempty"`
	CommentProbability *int     `json:"commentProbability,omitempty"`
	SearchKeywords     []string `json:"searchKeywords,omitempty"`
	LikeVideo          *bool    `json:"likeVideo,omitempty"`
	ImageURL           string   `json:"imageUrl,omitempty"`
}

type TikTokRandomCommentAsiaResponse struct {
	TaskID string `json:"taskId"`
}

// TikTokRandomCommentAsia calls /open/v1/rpa/task/tiktokRandomCommentAsia.
func (c *Client) TikTokRandomCommentAsia(ctx context.Context, req TikTokRandomCommentAsiaRequest) (*TikTokRandomCommentAsiaResponse, error) {
	var resp TikTokRandomCommentAsiaResponse
	if err := c.core.Do(ctx, "/open/v1/rpa/task/tiktokRandomCommentAsia", req, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

type TikTokMessageRequest struct {
	Name       string   `json:"name,omitempty"`
	Remark     string   `json:"remark,omitempty"`
	ScheduleAt int      `json:"scheduleAt"`
	ID         string   `json:"id"`
	Usernames  []string `json:"usernames"`
	Content    string   `json:"content"`
}

type TikTokMessageResponse struct {
	TaskID string `json:"taskId"`
}

// TikTokMessage calls /open/v1/rpa/task/tiktokMessage.
func (c *Client) TikTokMessage(ctx context.Context, req TikTokMessageRequest) (*TikTokMessageResponse, error) {
	var resp TikTokMessageResponse
	if err := c.core.Do(ctx, "/open/v1/rpa/task/tiktokMessage", req, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

type TikTokMessageAsiaRequest struct {
	Name       string   `json:"name,omitempty"`
	Remark     string   `json:"remark,omitempty"`
	ScheduleAt int      `json:"scheduleAt"`
	ID         string   `json:"id"`
	Usernames  []string `json:"usernames"`
	Content    string   `json:"content"`
}

type TikTokMessageAsiaResponse struct {
	TaskID string `json:"taskId"`
}

// TikTokMessageAsia calls /open/v1/rpa/task/tiktokMessageAsia.
func (c *Client) TikTokMessageAsia(ctx context.Context, req TikTokMessageAsiaRequest) (*TikTokMessageAsiaResponse, error) {
	var resp TikTokMessageAsiaResponse
	if err := c.core.Do(ctx, "/open/v1/rpa/task/tiktokMessageAsia", req, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

type TikTokRandomFollowRequest struct {
	Name              string `json:"name,omitempty"`
	Remark            string `json:"remark,omitempty"`
	ScheduleAt        int    `json:"scheduleAt"`
	ID                string `json:"id"`
	FollowProbability int    `json:"followProbability"`
}

type TikTokRandomFollowResponse struct {
	TaskID string `json:"taskId"`
}

// TikTokRandomFollow calls /open/v1/rpa/task/tiktokRandomFollow.
func (c *Client) TikTokRandomFollow(ctx context.Context, req TikTokRandomFollowRequest) (*TikTokRandomFollowResponse, error) {
	var resp TikTokRandomFollowResponse
	if err := c.core.Do(ctx, "/open/v1/rpa/task/tiktokRandomFollow", req, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

type TikTokRandomFollowAsiaRequest struct {
	Name              string `json:"name,omitempty"`
	Remark            string `json:"remark,omitempty"`
	ScheduleAt        int    `json:"scheduleAt"`
	ID                string `json:"id"`
	FollowProbability int    `json:"followProbability"`
}

type TikTokRandomFollowAsiaResponse struct {
	TaskID string `json:"taskId"`
}

// TikTokRandomFollowAsia calls /open/v1/rpa/task/tiktokRandomFollowAsia.
func (c *Client) TikTokRandomFollowAsia(ctx context.Context, req TikTokRandomFollowAsiaRequest) (*TikTokRandomFollowAsiaResponse, error) {
	var resp TikTokRandomFollowAsiaResponse
	if err := c.core.Do(ctx, "/open/v1/rpa/task/tiktokRandomFollowAsia", req, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

type TikTokHideRequest struct {
	Name       string `json:"name,omitempty"`
	Remark     string `json:"remark,omitempty"`
	ScheduleAt int    `json:"scheduleAt"`
	ID         string `json:"id"`
	Number     *int   `json:"number,omitempty"`
}

type TikTokHideResponse struct {
	TaskID string `json:"taskId"`
}

// TikTokHide calls /open/v1/rpa/task/tiktokHide.
func (c *Client) TikTokHide(ctx context.Context, req TikTokHideRequest) (*TikTokHideResponse, error) {
	var resp TikTokHideResponse
	if err := c.core.Do(ctx, "/open/v1/rpa/task/tiktokHide", req, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

type TikTokHideAsiaRequest struct {
	Name       string `json:"name,omitempty"`
	Remark     string `json:"remark,omitempty"`
	ScheduleAt int    `json:"scheduleAt"`
	ID         string `json:"id"`
	Number     *int   `json:"number,omitempty"`
}

type TikTokHideAsiaResponse struct {
	TaskID string `json:"taskId"`
}

// TikTokHideAsia calls /open/v1/rpa/task/tiktokHideAsia.
func (c *Client) TikTokHideAsia(ctx context.Context, req TikTokHideAsiaRequest) (*TikTokHideAsiaResponse, error) {
	var resp TikTokHideAsiaResponse
	if err := c.core.Do(ctx, "/open/v1/rpa/task/tiktokHideAsia", req, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

type TikTokDelRequest struct {
	Name       string `json:"name,omitempty"`
	Remark     string `json:"remark,omitempty"`
	ScheduleAt int    `json:"scheduleAt"`
	ID         string `json:"id"`
}

type TikTokDelResponse struct {
	TaskID string `json:"taskId"`
}

// TikTokDel calls /open/v1/rpa/task/tiktokDel.
func (c *Client) TikTokDel(ctx context.Context, req TikTokDelRequest) (*TikTokDelResponse, error) {
	var resp TikTokDelResponse
	if err := c.core.Do(ctx, "/open/v1/rpa/task/tiktokDel", req, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

type TikTokDelAsiaRequest struct {
	Name       string `json:"name,omitempty"`
	Remark     string `json:"remark,omitempty"`
	ScheduleAt int    `json:"scheduleAt"`
	ID         string `json:"id"`
}

type TikTokDelAsiaResponse struct {
	TaskID string `json:"taskId"`
}

// TikTokDelAsia calls /open/v1/rpa/task/tiktokDelAsia.
func (c *Client) TikTokDelAsia(ctx context.Context, req TikTokDelAsiaRequest) (*TikTokDelAsiaResponse, error) {
	var resp TikTokDelAsiaResponse
	if err := c.core.Do(ctx, "/open/v1/rpa/task/tiktokDelAsia", req, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

type TikTokCommentInterceptionRequest struct {
	Name            string   `json:"name,omitempty"`
	Remark          string   `json:"remark,omitempty"`
	ScheduleAt      int      `json:"scheduleAt"`
	ID              string   `json:"id"`
	ExecutionTime   int      `json:"executionTime"`
	VideoKeywords   []string `json:"videoKeywords"`
	CommentKeywords []string `json:"commentKeywords"`
	ReplyComment    []string `json:"replyComment"`
	TimeLimitDays   *int     `json:"timeLimitDays,omitempty"`
}

type TikTokCommentInterceptionResponse struct {
	TaskID string `json:"taskId"`
}

// TikTokCommentInterception calls /open/v1/rpa/task/tiktokCommentInterception.
func (c *Client) TikTokCommentInterception(ctx context.Context, req TikTokCommentInterceptionRequest) (*TikTokCommentInterceptionResponse, error) {
	var resp TikTokCommentInterceptionResponse
	if err := c.core.Do(ctx, "/open/v1/rpa/task/tiktokCommentInterception", req, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

type TikTokCommentInterceptionAsiaRequest struct {
	Name            string   `json:"name,omitempty"`
	Remark          string   `json:"remark,omitempty"`
	ScheduleAt      int      `json:"scheduleAt"`
	ID              string   `json:"id"`
	ExecutionTime   int      `json:"executionTime"`
	VideoKeywords   []string `json:"videoKeywords"`
	CommentKeywords []string `json:"commentKeywords"`
	ReplyComment    []string `json:"replyComment"`
	TimeLimitDays   *int     `json:"timeLimitDays,omitempty"`
}

type TikTokCommentInterceptionAsiaResponse struct {
	TaskID string `json:"taskId"`
}

// TikTokCommentInterceptionAsia calls /open/v1/rpa/task/tiktokCommentInterceptionAsia.
func (c *Client) TikTokCommentInterceptionAsia(ctx context.Context, req TikTokCommentInterceptionAsiaRequest) (*TikTokCommentInterceptionAsiaResponse, error) {
	var resp TikTokCommentInterceptionAsiaResponse
	if err := c.core.Do(ctx, "/open/v1/rpa/task/tiktokCommentInterceptionAsia", req, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

type TikTokDeleteCommentRequest struct {
	Name       string   `json:"name,omitempty"`
	Remark     string   `json:"remark,omitempty"`
	ScheduleAt int      `json:"scheduleAt"`
	ID         string   `json:"id"`
	Keywords   []string `json:"keywords"`
}

type TikTokDeleteCommentResponse struct {
	TaskID string `json:"taskId"`
}

// TikTokDeleteComment calls /open/v1/rpa/task/tiktokDeleteComment.
func (c *Client) TikTokDeleteComment(ctx context.Context, req TikTokDeleteCommentRequest) (*TikTokDeleteCommentResponse, error) {
	var resp TikTokDeleteCommentResponse
	if err := c.core.Do(ctx, "/open/v1/rpa/task/tiktokDeleteComment", req, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

type TikTokDeleteCommentAsiaRequest struct {
	Name       string   `json:"name,omitempty"`
	Remark     string   `json:"remark,omitempty"`
	ScheduleAt int      `json:"scheduleAt"`
	ID         string   `json:"id"`
	Keywords   []string `json:"keywords"`
}

type TikTokDeleteCommentAsiaResponse struct {
	TaskID string `json:"taskId"`
}

// TikTokDeleteCommentAsia calls /open/v1/rpa/task/tiktokDeleteCommentAsia.
func (c *Client) TikTokDeleteCommentAsia(ctx context.Context, req TikTokDeleteCommentAsiaRequest) (*TikTokDeleteCommentAsiaResponse, error) {
	var resp TikTokDeleteCommentAsiaResponse
	if err := c.core.Do(ctx, "/open/v1/rpa/task/tiktokDeleteCommentAsia", req, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

type FaceBookLoginRequest struct {
	Name       string `json:"name,omitempty"`
	Remark     string `json:"remark,omitempty"`
	ScheduleAt int    `json:"scheduleAt"`
	ID         string `json:"id"`
	Email      string `json:"email"`
	Password   string `json:"password"`
}

type FaceBookLoginResponse struct {
	TaskID string `json:"taskId"`
}

// FaceBookLogin calls /open/v1/rpa/task/faceBookLogin.
func (c *Client) FaceBookLogin(ctx context.Context, req FaceBookLoginRequest) (*FaceBookLoginResponse, error) {
	var resp FaceBookLoginResponse
	if err := c.core.Do(ctx, "/open/v1/rpa/task/faceBookLogin", req, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

type FaceBookAutoCommentRequest struct {
	Name        string   `json:"name,omitempty"`
	Remark      string   `json:"remark,omitempty"`
	ScheduleAt  int      `json:"scheduleAt"`
	ID          string   `json:"id"`
	PostAddress string   `json:"postAddress"`
	Comment     []string `json:"comment"`
	Keyword     []string `json:"keyword"`
}

type FaceBookAutoCommentResponse struct {
	TaskID string `json:"taskId"`
}

// FaceBookAutoComment calls /open/v1/rpa/task/faceBookAutoComment.
func (c *Client) FaceBookAutoComment(ctx context.Context, req FaceBookAutoCommentRequest) (*FaceBookAutoCommentResponse, error) {
	var resp FaceBookAutoCommentResponse
	if err := c.core.Do(ctx, "/open/v1/rpa/task/faceBookAutoComment", req, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

type FaceBookPublishRequest struct {
	Name          string   `json:"name,omitempty"`
	Remark        string   `json:"remark,omitempty"`
	ScheduleAt    int      `json:"scheduleAt"`
	ID            string   `json:"id"`
	Title         string   `json:"title"`
	Video         []string `json:"video"`
	NeedShareLink *bool    `json:"needShareLink,omitempty"`
}

type FaceBookPublishResponse struct {
	TaskID string `json:"taskId"`
}

// FaceBookPublish calls /open/v1/rpa/task/faceBookPublish.
func (c *Client) FaceBookPublish(ctx context.Context, req FaceBookPublishRequest) (*FaceBookPublishResponse, error) {
	var resp FaceBookPublishResponse
	if err := c.core.Do(ctx, "/open/v1/rpa/task/faceBookPublish", req, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

type FaceBookActiveAccountRequest struct {
	Name           string   `json:"name,omitempty"`
	Remark         string   `json:"remark,omitempty"`
	ScheduleAt     int      `json:"scheduleAt"`
	ID             string   `json:"id"`
	BrowsePostsNum int      `json:"browsePostsNum"`
	Keyword        []string `json:"keyword"`
	PostAddress    string   `json:"postAddress,omitempty"`
}

type FaceBookActiveAccountResponse struct {
	TaskID string `json:"taskId"`
}

// FaceBookActiveAccount calls /open/v1/rpa/task/faceBookActiveAccount.
func (c *Client) FaceBookActiveAccount(ctx context.Context, req FaceBookActiveAccountRequest) (*FaceBookActiveAccountResponse, error) {
	var resp FaceBookActiveAccountResponse
	if err := c.core.Do(ctx, "/open/v1/rpa/task/faceBookActiveAccount", req, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

type FaceBookPubReelsRequest struct {
	Name          string `json:"name,omitempty"`
	Remark        string `json:"remark,omitempty"`
	ScheduleAt    int    `json:"scheduleAt"`
	ID            string `json:"id"`
	Description   string `json:"description"`
	Video         string `json:"video"`
	Page          string `json:"page,omitempty"`
	NeedShareLink *bool  `json:"needShareLink,omitempty"`
}

type FaceBookPubReelsResponse struct {
	TaskID string `json:"taskId"`
}

// FaceBookPubReels calls /open/v1/rpa/task/faceBookPubReels.
func (c *Client) FaceBookPubReels(ctx context.Context, req FaceBookPubReelsRequest) (*FaceBookPubReelsResponse, error) {
	var resp FaceBookPubReelsResponse
	if err := c.core.Do(ctx, "/open/v1/rpa/task/faceBookPubReels", req, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

type FaceBookMessageRequest struct {
	Name       string   `json:"name,omitempty"`
	Remark     string   `json:"remark,omitempty"`
	ScheduleAt int      `json:"scheduleAt"`
	ID         string   `json:"id"`
	Usernames  []string `json:"usernames"`
	Content    string   `json:"content"`
}

type FaceBookMessageResponse struct {
	TaskID string `json:"taskId"`
}

// FaceBookMessage calls /open/v1/rpa/task/faceBookMessage.
func (c *Client) FaceBookMessage(ctx context.Context, req FaceBookMessageRequest) (*FaceBookMessageResponse, error) {
	var resp FaceBookMessageResponse
	if err := c.core.Do(ctx, "/open/v1/rpa/task/faceBookMessage", req, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

type FacebookReelsActiveRequest struct {
	Name           string   `json:"name,omitempty"`
	Remark         string   `json:"remark,omitempty"`
	ScheduleAt     int      `json:"scheduleAt"`
	ID             string   `json:"id"`
	VideoNumber    int      `json:"videoNumber"`
	SearchKeywords []string `json:"searchKeywords,omitempty"`
}

type FacebookReelsActiveResponse struct {
	TaskID string `json:"taskId"`
}

// FacebookReelsActive calls /open/v1/rpa/task/facebookReelsActive.
func (c *Client) FacebookReelsActive(ctx context.Context, req FacebookReelsActiveRequest) (*FacebookReelsActiveResponse, error) {
	var resp FacebookReelsActiveResponse
	if err := c.core.Do(ctx, "/open/v1/rpa/task/facebookReelsActive", req, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

type InstagramPubReelsRequest struct {
	Name           string   `json:"name,omitempty"`
	Remark         string   `json:"remark,omitempty"`
	ScheduleAt     int      `json:"scheduleAt"`
	ID             string   `json:"id"`
	Description    string   `json:"description,omitempty"`
	Video          []string `json:"video"`
	SameStyleURL   string   `json:"sameStyleUrl,omitempty"`
	SameStyleVoice *int     `json:"sameStyleVoice,omitempty"`
	OriginalVoice  *int     `json:"originalVoice,omitempty"`
	AiTag          *bool    `json:"aiTag,omitempty"`
	NeedShareLink  *bool    `json:"needShareLink,omitempty"`
	Cover          []string `json:"cover,omitempty"`
}

type InstagramPubReelsResponse struct {
	TaskID string `json:"taskId"`
}

// InstagramPubReels calls /open/v1/rpa/task/instagramPubReels.
func (c *Client) InstagramPubReels(ctx context.Context, req InstagramPubReelsRequest) (*InstagramPubReelsResponse, error) {
	var resp InstagramPubReelsResponse
	if err := c.core.Do(ctx, "/open/v1/rpa/task/instagramPubReels", req, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

type InstagramWarmupRequest struct {
	Name        string `json:"name,omitempty"`
	Remark      string `json:"remark,omitempty"`
	ScheduleAt  int    `json:"scheduleAt"`
	ID          string `json:"id"`
	BrowseVideo *int   `json:"browseVideo,omitempty"`
	Keyword     string `json:"keyword,omitempty"`
}

type InstagramWarmupResponse struct {
	TaskID string `json:"taskId"`
}

// InstagramWarmup calls /open/v1/rpa/task/instagramWarmup.
func (c *Client) InstagramWarmup(ctx context.Context, req InstagramWarmupRequest) (*InstagramWarmupResponse, error) {
	var resp InstagramWarmupResponse
	if err := c.core.Do(ctx, "/open/v1/rpa/task/instagramWarmup", req, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

type InstagramMessageRequest struct {
	Name       string   `json:"name,omitempty"`
	Remark     string   `json:"remark,omitempty"`
	ScheduleAt int      `json:"scheduleAt"`
	ID         string   `json:"id"`
	Usernames  []string `json:"usernames"`
	Content    string   `json:"content"`
}

type InstagramMessageResponse struct {
	TaskID string `json:"taskId"`
}

// InstagramMessage calls /open/v1/rpa/task/instagramMessage.
func (c *Client) InstagramMessage(ctx context.Context, req InstagramMessageRequest) (*InstagramMessageResponse, error) {
	var resp InstagramMessageResponse
	if err := c.core.Do(ctx, "/open/v1/rpa/task/instagramMessage", req, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

type InstagramPubReelsImagesRequest struct {
	Name          string   `json:"name,omitempty"`
	Remark        string   `json:"remark,omitempty"`
	ScheduleAt    int      `json:"scheduleAt"`
	ID            string   `json:"id"`
	Description   string   `json:"description,omitempty"`
	Image         []string `json:"image"`
	SameStyleURL  string   `json:"sameStyleUrl,omitempty"`
	AiTag         *bool    `json:"aiTag,omitempty"`
	PublishPost   *bool    `json:"publishPost,omitempty"`
	NeedShareLink *bool    `json:"needShareLink,omitempty"`
}

type InstagramPubReelsImagesResponse struct {
	TaskID string `json:"taskId"`
}

// InstagramPubReelsImages calls /open/v1/rpa/task/instagramPubReelsImages.
func (c *Client) InstagramPubReelsImages(ctx context.Context, req InstagramPubReelsImagesRequest) (*InstagramPubReelsImagesResponse, error) {
	var resp InstagramPubReelsImagesResponse
	if err := c.core.Do(ctx, "/open/v1/rpa/task/instagramPubReelsImages", req, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

type InstagramLoginRequest struct {
	Name       string `json:"name,omitempty"`
	Remark     string `json:"remark,omitempty"`
	ScheduleAt int    `json:"scheduleAt"`
	ID         string `json:"id"`
	Account    string `json:"account"`
	Password   string `json:"password"`
}

type InstagramLoginResponse struct {
	TaskID string `json:"taskId"`
}

// InstagramLogin calls /open/v1/rpa/task/instagramLogin.
func (c *Client) InstagramLogin(ctx context.Context, req InstagramLoginRequest) (*InstagramLoginResponse, error) {
	var resp InstagramLoginResponse
	if err := c.core.Do(ctx, "/open/v1/rpa/task/instagramLogin", req, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

type InstagramEditRequest struct {
	Name           string   `json:"name,omitempty"`
	Remark         string   `json:"remark,omitempty"`
	ScheduleAt     int      `json:"scheduleAt"`
	ID             string   `json:"id"`
	ProfilePicture []string `json:"profilePicture,omitempty"`
	Nickname       string   `json:"nickname,omitempty"`
	Username       string   `json:"username,omitempty"`
	Biography      string   `json:"biography,omitempty"`
	LinkURL        string   `json:"linkURL,omitempty"`
	LinkTitle      string   `json:"linkTitle,omitempty"`
}

type InstagramEditResponse struct {
	TaskID string `json:"taskId"`
}

// InstagramEdit calls /open/v1/rpa/task/instagramEdit.
func (c *Client) InstagramEdit(ctx context.Context, req InstagramEditRequest) (*InstagramEditResponse, error) {
	var resp InstagramEditResponse
	if err := c.core.Do(ctx, "/open/v1/rpa/task/instagramEdit", req, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

type InstagramAccountStatusRequest struct {
	Name       string `json:"name,omitempty"`
	Remark     string `json:"remark,omitempty"`
	ScheduleAt int    `json:"scheduleAt"`
	ID         string `json:"id"`
	Public     *bool  `json:"public,omitempty"`
}

type InstagramAccountStatusResponse struct {
	TaskID string `json:"taskId"`
}

// InstagramAccountStatus calls /open/v1/rpa/task/instagramAccountStatus.
func (c *Client) InstagramAccountStatus(ctx context.Context, req InstagramAccountStatusRequest) (*InstagramAccountStatusResponse, error) {
	var resp InstagramAccountStatusResponse
	if err := c.core.Do(ctx, "/open/v1/rpa/task/instagramAccountStatus", req, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

type InstagramFollowAccountRequest struct {
	Name       string   `json:"name,omitempty"`
	Remark     string   `json:"remark,omitempty"`
	ScheduleAt int      `json:"scheduleAt"`
	ID         string   `json:"id"`
	Username   []string `json:"username"`
}

type InstagramFollowAccountResponse struct {
	TaskID string `json:"taskId"`
}

// InstagramFollowAccount calls /open/v1/rpa/task/instagramFollowAccount.
func (c *Client) InstagramFollowAccount(ctx context.Context, req InstagramFollowAccountRequest) (*InstagramFollowAccountResponse, error) {
	var resp InstagramFollowAccountResponse
	if err := c.core.Do(ctx, "/open/v1/rpa/task/instagramFollowAccount", req, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

type InstagramAiCommentRequest struct {
	Name       string `json:"name,omitempty"`
	Remark     string `json:"remark,omitempty"`
	ScheduleAt int    `json:"scheduleAt"`
	ID         string `json:"id"`
	UseAi      *bool  `json:"useAi,omitempty"`
	RandomRate int    `json:"randomRate"`
}

type InstagramAiCommentResponse struct {
	TaskID string `json:"taskId"`
}

// InstagramAiComment calls /open/v1/rpa/task/instagramAiComment.
func (c *Client) InstagramAiComment(ctx context.Context, req InstagramAiCommentRequest) (*InstagramAiCommentResponse, error) {
	var resp InstagramAiCommentResponse
	if err := c.core.Do(ctx, "/open/v1/rpa/task/instagramAiComment", req, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

type YouTubePubShortRequest struct {
	Name                  string `json:"name,omitempty"`
	Remark                string `json:"remark,omitempty"`
	ScheduleAt            int    `json:"scheduleAt"`
	ID                    string `json:"id"`
	Title                 string `json:"title"`
	Video                 string `json:"video"`
	SameStyleURL          string `json:"sameStyleUrl,omitempty"`
	SameStyleVoice        int    `json:"sameStyleVoice"`
	OriginalVoice         int    `json:"originalVoice"`
	IsDisclosureMandatory *bool  `json:"isDisclosureMandatory,omitempty"`
	CoverFirstFrame       *bool  `json:"coverFirstFrame,omitempty"`
	NeedShareLink         *bool  `json:"needShareLink,omitempty"`
}

type YouTubePubShortResponse struct {
	TaskID string `json:"taskId"`
}

// YouTubePubShort calls /open/v1/rpa/task/youtubePubShort.
func (c *Client) YouTubePubShort(ctx context.Context, req YouTubePubShortRequest) (*YouTubePubShortResponse, error) {
	var resp YouTubePubShortResponse
	if err := c.core.Do(ctx, "/open/v1/rpa/task/youtubePubShort", req, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

type YouTubePubVideoRequest struct {
	Name                    string `json:"name,omitempty"`
	Remark                  string `json:"remark,omitempty"`
	ScheduleAt              int    `json:"scheduleAt"`
	ID                      string `json:"id"`
	Title                   string `json:"title"`
	Description             string `json:"description"`
	Video                   string `json:"video"`
	IsDisclosureMandatory   *bool  `json:"isDisclosureMandatory,omitempty"`
	WaitTimeAfterPublishing *int   `json:"waitTimeAfterPublishing,omitempty"`
}

type YouTubePubVideoResponse struct {
	TaskID string `json:"taskId"`
}

// YouTubePubVideo calls /open/v1/rpa/task/youtubePubVideo.
func (c *Client) YouTubePubVideo(ctx context.Context, req YouTubePubVideoRequest) (*YouTubePubVideoResponse, error) {
	var resp YouTubePubVideoResponse
	if err := c.core.Do(ctx, "/open/v1/rpa/task/youtubePubVideo", req, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

type YouTubeActiveAccountRequest struct {
	Name           string   `json:"name,omitempty"`
	Remark         string   `json:"remark,omitempty"`
	ScheduleAt     int      `json:"scheduleAt"`
	ID             string   `json:"id"`
	BrowseVideoNum int      `json:"browseVideoNum"`
	Keyword        []string `json:"keyword,omitempty"`
}

type YouTubeActiveAccountResponse struct {
	TaskID string `json:"taskId"`
}

// YouTubeActiveAccount calls /open/v1/rpa/task/youTubeActiveAccount.
func (c *Client) YouTubeActiveAccount(ctx context.Context, req YouTubeActiveAccountRequest) (*YouTubeActiveAccountResponse, error) {
	var resp YouTubeActiveAccountResponse
	if err := c.core.Do(ctx, "/open/v1/rpa/task/youTubeActiveAccount", req, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

type YouTubeEditRequest struct {
	Name        string `json:"name,omitempty"`
	Remark      string `json:"remark,omitempty"`
	ScheduleAt  int    `json:"scheduleAt"`
	ID          string `json:"id"`
	ProfileName string `json:"profileName,omitempty"`
	Handle      string `json:"handle,omitempty"`
	Description string `json:"description,omitempty"`
}

type YouTubeEditResponse struct {
	TaskID string `json:"taskId"`
}

// YouTubeEdit calls /open/v1/rpa/task/youtubeEdit.
func (c *Client) YouTubeEdit(ctx context.Context, req YouTubeEditRequest) (*YouTubeEditResponse, error) {
	var resp YouTubeEditResponse
	if err := c.core.Do(ctx, "/open/v1/rpa/task/youtubeEdit", req, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

type GoogleLoginRequest struct {
	Name       string `json:"name,omitempty"`
	Remark     string `json:"remark,omitempty"`
	ScheduleAt int    `json:"scheduleAt"`
	ID         string `json:"id"`
	Email      string `json:"email"`
	Password   string `json:"password"`
	Code2fa    string `json:"code2fa,omitempty"`
}

type GoogleLoginResponse struct {
	TaskID string `json:"taskId"`
}

// GoogleLogin calls /open/v1/rpa/task/googleLogin.
func (c *Client) GoogleLogin(ctx context.Context, req GoogleLoginRequest) (*GoogleLoginResponse, error) {
	var resp GoogleLoginResponse
	if err := c.core.Do(ctx, "/open/v1/rpa/task/googleLogin", req, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

type GoogleAppDownloadRequest struct {
	Name       string `json:"name,omitempty"`
	Remark     string `json:"remark,omitempty"`
	ScheduleAt int    `json:"scheduleAt"`
	ID         string `json:"id"`
	AppName    string `json:"appName"`
}

type GoogleAppDownloadResponse struct {
	TaskID string `json:"taskId"`
}

// GoogleAppDownload calls /open/v1/rpa/task/googleAppDownload.
func (c *Client) GoogleAppDownload(ctx context.Context, req GoogleAppDownloadRequest) (*GoogleAppDownloadResponse, error) {
	var resp GoogleAppDownloadResponse
	if err := c.core.Do(ctx, "/open/v1/rpa/task/googleAppDownload", req, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

type GoogleAppBrowserRequest struct {
	Name        string `json:"name,omitempty"`
	Remark      string `json:"remark,omitempty"`
	ScheduleAt  int    `json:"scheduleAt"`
	ID          string `json:"id"`
	AppName     string `json:"appName"`
	Description string `json:"description,omitempty"`
}

type GoogleAppBrowserResponse struct {
	TaskID string `json:"taskId"`
}

// GoogleAppBrowser calls /open/v1/rpa/task/googleAppBrowser.
func (c *Client) GoogleAppBrowser(ctx context.Context, req GoogleAppBrowserRequest) (*GoogleAppBrowserResponse, error) {
	var resp GoogleAppBrowserResponse
	if err := c.core.Do(ctx, "/open/v1/rpa/task/googleAppBrowser", req, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

type SheinLoginRequest struct {
	Name       string `json:"name,omitempty"`
	Remark     string `json:"remark,omitempty"`
	ScheduleAt int    `json:"scheduleAt"`
	ID         string `json:"id"`
	Email      string `json:"email"`
	Password   string `json:"password"`
}

type SheinLoginResponse struct {
	TaskID string `json:"taskId"`
}

// SheinLogin calls /open/v1/rpa/task/sheinLogin.
func (c *Client) SheinLogin(ctx context.Context, req SheinLoginRequest) (*SheinLoginResponse, error) {
	var resp SheinLoginResponse
	if err := c.core.Do(ctx, "/open/v1/rpa/task/sheinLogin", req, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

type XPublishRequest struct {
	Name        string   `json:"name,omitempty"`
	Remark      string   `json:"remark,omitempty"`
	ScheduleAt  int      `json:"scheduleAt"`
	ID          string   `json:"id"`
	Description string   `json:"description"`
	Video       []string `json:"video"`
}

type XPublishResponse struct {
	TaskID string `json:"taskId"`
}

// XPublish calls /open/v1/rpa/task/xPublish.
func (c *Client) XPublish(ctx context.Context, req XPublishRequest) (*XPublishResponse, error) {
	var resp XPublishResponse
	if err := c.core.Do(ctx, "/open/v1/rpa/task/xPublish", req, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

type ListFlowsRequest struct {
	Page     int `json:"page"`
	PageSize int `json:"pageSize"`
}

type ListFlowsResponseItemsItem struct {
	ID     string   `json:"id"`
	Title  string   `json:"title"`
	Desc   string   `json:"desc"`
	Params []string `json:"params"`
}

type ListFlowsResponse struct {
	Total    int                          `json:"total"`
	Page     int                          `json:"page"`
	PageSize int                          `json:"pageSize"`
	Items    []ListFlowsResponseItemsItem `json:"items"`
}

// ListFlows calls /open/v1/task/flow/list.
func (c *Client) ListFlows(ctx context.Context, req ListFlowsRequest) (*ListFlowsResponse, error) {
	var resp ListFlowsResponse
	if err := c.core.Do(ctx, "/open/v1/task/flow/list", req, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// CreateCustomTaskRequestParamMap contains the named inputs of a custom flow.
// Values may be strings, string slices, numbers, or booleans.
type CreateCustomTaskRequestParamMap map[string]any

type CreateCustomTaskRequest struct {
	Name       string                          `json:"name,omitempty"`
	Remark     string                          `json:"remark,omitempty"`
	ScheduleAt int                             `json:"scheduleAt"`
	ID         string                          `json:"id"`
	FlowID     string                          `json:"flowId"`
	ParamMap   CreateCustomTaskRequestParamMap `json:"paramMap,omitempty"`
}

type CreateCustomTaskResponse struct {
	TaskID string `json:"taskId"`
}

// CreateCustomTask calls /open/v1/task/rpa/add.
func (c *Client) CreateCustomTask(ctx context.Context, req CreateCustomTaskRequest) (*CreateCustomTaskResponse, error) {
	var resp CreateCustomTaskResponse
	if err := c.core.Do(ctx, "/open/v1/task/rpa/add", req, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

type ImportFlowRequest struct {
	ID  string `json:"id,omitempty"`
	Gal string `json:"gal,omitempty"`
}

type ImportFlowResponse struct {
	ID string `json:"id"`
}

// ImportFlow calls /open/v1/task/flow/import.
func (c *Client) ImportFlow(ctx context.Context, req ImportFlowRequest) (*ImportFlowResponse, error) {
	var resp ImportFlowResponse
	if err := c.core.Do(ctx, "/open/v1/task/flow/import", req, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

type ExportFlowRequest struct {
	ID string `json:"id"`
}

type ExportFlowResponse struct {
	Gal string `json:"gal"`
}

// ExportFlow calls /open/v1/task/flow/export.
func (c *Client) ExportFlow(ctx context.Context, req ExportFlowRequest) (*ExportFlowResponse, error) {
	var resp ExportFlowResponse
	if err := c.core.Do(ctx, "/open/v1/task/flow/export", req, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

type MultiPlatformVideoDistributionRequest struct {
	Name                  string   `json:"name,omitempty"`
	Remark                string   `json:"remark,omitempty"`
	ScheduleAt            int      `json:"scheduleAt"`
	ID                    string   `json:"id"`
	Title                 string   `json:"title,omitempty"`
	TikTokTitle           string   `json:"tiktokTitle,omitempty"`
	YouTubeTitle          string   `json:"youtubeTitle,omitempty"`
	InstagramTitle        string   `json:"instagramTitle,omitempty"`
	Video                 []string `json:"video"`
	TikTokRecreateLink    string   `json:"tiktokRecreateLink,omitempty"`
	YouTubeRecreateLink   string   `json:"youtubeRecreateLink,omitempty"`
	InstagramRecreateLink string   `json:"instagramRecreateLink,omitempty"`
	SameStyleVoice        *int     `json:"sameStyleVoice,omitempty"`
	OriginalVoice         *int     `json:"originalVoice,omitempty"`
}

type MultiPlatformVideoDistributionResponse struct {
	TaskID string `json:"taskId"`
}

// MultiPlatformVideoDistribution calls /open/v1/rpa/task/multiPlatformVideoDistribution.
func (c *Client) MultiPlatformVideoDistribution(ctx context.Context, req MultiPlatformVideoDistributionRequest) (*MultiPlatformVideoDistributionResponse, error) {
	var resp MultiPlatformVideoDistributionResponse
	if err := c.core.Do(ctx, "/open/v1/rpa/task/multiPlatformVideoDistribution", req, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

type FileUploadRequest struct {
	Name       string   `json:"name,omitempty"`
	Remark     string   `json:"remark,omitempty"`
	ScheduleAt int      `json:"scheduleAt"`
	ID         string   `json:"id"`
	Files      []string `json:"files"`
}

type FileUploadResponse struct {
	TaskID string `json:"taskId"`
}

// FileUpload calls /open/v1/rpa/task/fileUpload.
func (c *Client) FileUpload(ctx context.Context, req FileUploadRequest) (*FileUploadResponse, error) {
	var resp FileUploadResponse
	if err := c.core.Do(ctx, "/open/v1/rpa/task/fileUpload", req, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

type ImportContactsRequestContactsItem struct {
	Email1    string `json:"email1,omitempty"`
	Email2    string `json:"email2,omitempty"`
	Fax       string `json:"fax,omitempty"`
	FirstName string `json:"firstName,omitempty"`
	LastName  string `json:"lastName,omitempty"`
	Mobile    string `json:"mobile,omitempty"`
	Work      string `json:"work,omitempty"`
}

type ImportContactsRequest struct {
	Name       string                              `json:"name,omitempty"`
	Remark     string                              `json:"remark,omitempty"`
	ScheduleAt int                                 `json:"scheduleAt"`
	ID         string                              `json:"id"`
	Contacts   []ImportContactsRequestContactsItem `json:"contacts"`
}

type ImportContactsResponse struct {
	TaskID string `json:"taskId"`
}

// ImportContacts calls /open/v1/rpa/task/importContacts.
func (c *Client) ImportContacts(ctx context.Context, req ImportContactsRequest) (*ImportContactsResponse, error) {
	var resp ImportContactsResponse
	if err := c.core.Do(ctx, "/open/v1/rpa/task/importContacts", req, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

type KeyboxUploadRequest struct {
	Name       string   `json:"name,omitempty"`
	Remark     string   `json:"remark,omitempty"`
	ScheduleAt int      `json:"scheduleAt"`
	ID         string   `json:"id"`
	Files      []string `json:"files"`
}

type KeyboxUploadResponse struct {
	TaskID string `json:"taskId"`
}

// KeyboxUpload calls /open/v1/rpa/task/keyboxUpload.
func (c *Client) KeyboxUpload(ctx context.Context, req KeyboxUploadRequest) (*KeyboxUploadResponse, error) {
	var resp KeyboxUploadResponse
	if err := c.core.Do(ctx, "/open/v1/rpa/task/keyboxUpload", req, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

type RedditWarmupRequest struct {
	Name       string `json:"name,omitempty"`
	Remark     string `json:"remark,omitempty"`
	ScheduleAt int    `json:"scheduleAt"`
	ID         string `json:"id"`
	Keyword    string `json:"keyword,omitempty"`
	Keywords   string `json:"keywords,omitempty"`
	Duration   int    `json:"duration"`
}

type RedditWarmupResponse struct {
	TaskID string `json:"taskId"`
}

// RedditWarmup calls /open/v1/rpa/task/redditWarmup.
func (c *Client) RedditWarmup(ctx context.Context, req RedditWarmupRequest) (*RedditWarmupResponse, error) {
	var resp RedditWarmupResponse
	if err := c.core.Do(ctx, "/open/v1/rpa/task/redditWarmup", req, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

type RedditVideoRequest struct {
	Name        string   `json:"name,omitempty"`
	Remark      string   `json:"remark,omitempty"`
	ScheduleAt  int      `json:"scheduleAt"`
	ID          string   `json:"id"`
	Title       string   `json:"title"`
	Description string   `json:"description,omitempty"`
	Video       []string `json:"video"`
	Community   string   `json:"community"`
	Flair       string   `json:"flair,omitempty"`
}

type RedditVideoResponse struct {
	TaskID string `json:"taskId"`
}

// RedditVideo calls /open/v1/rpa/task/redditVideo.
func (c *Client) RedditVideo(ctx context.Context, req RedditVideoRequest) (*RedditVideoResponse, error) {
	var resp RedditVideoResponse
	if err := c.core.Do(ctx, "/open/v1/rpa/task/redditVideo", req, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

type RedditImageRequest struct {
	Name        string   `json:"name,omitempty"`
	Remark      string   `json:"remark,omitempty"`
	ScheduleAt  int      `json:"scheduleAt"`
	ID          string   `json:"id"`
	Title       string   `json:"title"`
	Description string   `json:"description,omitempty"`
	Images      []string `json:"images"`
	Community   string   `json:"community"`
	Flair       string   `json:"flair,omitempty"`
}

type RedditImageResponse struct {
	TaskID string `json:"taskId"`
}

// RedditImage calls /open/v1/rpa/task/redditImage.
func (c *Client) RedditImage(ctx context.Context, req RedditImageRequest) (*RedditImageResponse, error) {
	var resp RedditImageResponse
	if err := c.core.Do(ctx, "/open/v1/rpa/task/redditImage", req, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

type PinterestVideoRequest struct {
	Name        string   `json:"name,omitempty"`
	Remark      string   `json:"remark,omitempty"`
	ScheduleAt  int      `json:"scheduleAt"`
	ID          string   `json:"id"`
	Title       string   `json:"title"`
	Description string   `json:"description"`
	Video       []string `json:"video"`
	Link        string   `json:"link,omitempty"`
}

type PinterestVideoResponse struct {
	TaskID string `json:"taskId"`
}

// PinterestVideo calls /open/v1/rpa/task/pinterestVideo.
func (c *Client) PinterestVideo(ctx context.Context, req PinterestVideoRequest) (*PinterestVideoResponse, error) {
	var resp PinterestVideoResponse
	if err := c.core.Do(ctx, "/open/v1/rpa/task/pinterestVideo", req, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

type PinterestImageRequest struct {
	Name        string   `json:"name,omitempty"`
	Remark      string   `json:"remark,omitempty"`
	ScheduleAt  int      `json:"scheduleAt"`
	ID          string   `json:"id"`
	Title       string   `json:"title"`
	Description string   `json:"description"`
	Images      []string `json:"images"`
	Link        string   `json:"link,omitempty"`
}

type PinterestImageResponse struct {
	TaskID string `json:"taskId"`
}

// PinterestImage calls /open/v1/rpa/task/pinterestImage.
func (c *Client) PinterestImage(ctx context.Context, req PinterestImageRequest) (*PinterestImageResponse, error) {
	var resp PinterestImageResponse
	if err := c.core.Do(ctx, "/open/v1/rpa/task/pinterestImage", req, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

type ThreadsVideoRequest struct {
	Name       string   `json:"name,omitempty"`
	Remark     string   `json:"remark,omitempty"`
	ScheduleAt int      `json:"scheduleAt"`
	ID         string   `json:"id"`
	Topic      string   `json:"topic,omitempty"`
	Title      string   `json:"title"`
	Video      []string `json:"video"`
}

type ThreadsVideoResponse struct {
	TaskID string `json:"taskId"`
}

// ThreadsVideo calls /open/v1/rpa/task/threadsVideo.
func (c *Client) ThreadsVideo(ctx context.Context, req ThreadsVideoRequest) (*ThreadsVideoResponse, error) {
	var resp ThreadsVideoResponse
	if err := c.core.Do(ctx, "/open/v1/rpa/task/threadsVideo", req, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

type ThreadsImageRequest struct {
	Name       string   `json:"name,omitempty"`
	Remark     string   `json:"remark,omitempty"`
	ScheduleAt int      `json:"scheduleAt"`
	ID         string   `json:"id"`
	Topic      string   `json:"topic,omitempty"`
	Title      string   `json:"title"`
	Images     []string `json:"images"`
}

type ThreadsImageResponse struct {
	TaskID string `json:"taskId"`
}

// ThreadsImage calls /open/v1/rpa/task/threadsImage.
func (c *Client) ThreadsImage(ctx context.Context, req ThreadsImageRequest) (*ThreadsImageResponse, error) {
	var resp ThreadsImageResponse
	if err := c.core.Do(ctx, "/open/v1/rpa/task/threadsImage", req, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
