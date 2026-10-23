package automation

import (
	"context"
)

type QueryTaskRequest struct {
	// Array of task IDs, up to 100. Required.
	IDs []string `json:"ids"`
}

type QueryTaskResponseItemsItem struct {
	// Task ID.
	ID string `json:"id"`
	// Task plan name.
	PlanName string `json:"planName"`
	// Task type; 1 TikTok video posting; 2 TikTok AI account warmup; 3 TikTok carousel
	// posting; 4 TikTok account login; 6 TikTok profile editing; 42 Custom(Including Facebook,
	// YouTube and other platforms).
	TaskType int `json:"taskType"`
	// Cloud phone name.
	SerialName string `json:"serialName"`
	// Cloud phone ID.
	EnvID string `json:"envId"`
	// Scheduled execution time as a Unix timestamp in seconds.
	ScheduleAt int64 `json:"scheduleAt"`
	// Task status; 1 Waiting; 2 In progress; 3 Completed; 4 Failed; 7 Cancelled.
	Status int `json:"status"`
	// Failure code, refer to task failure codes and reasons.
	FailCode int `json:"failCode,omitempty"`
	// Failure reason, refer to task failure codes and reasons.
	FailDesc string `json:"failDesc,omitempty"`
	// The time a task takes to complete or fail (in seconds).
	Cost      int    `json:"cost,omitempty"`
	ShareLink string `json:"shareLink,omitempty"`
}

type QueryTaskResponse struct {
	// Total number of tasks.
	Total int `json:"total"`
	// Array of tasks.
	Items []QueryTaskResponseItemsItem `json:"items"`
}

// QueryTask: https://open.geelark.com/api/task-query
func (c *Client) QueryTask(ctx context.Context, req QueryTaskRequest) (*QueryTaskResponse, error) {
	var resp QueryTaskResponse
	if err := c.core.Do(ctx, "/open/v1/task/query", req, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

type ListTasksRequest struct {
	// Number of records per page, maximum is 100.
	Size *int `json:"size,omitempty"`
	// The id of the last item from the previous page’s data.items array.
	LastID string `json:"lastId,omitempty"`
	// Task IDs, maximum 100.
	IDs []string `json:"ids,omitempty"`
}

type ListTasksResponseItemsItem struct {
	// Task ID.
	ID string `json:"id"`
	// Task plan name.
	PlanName string `json:"planName"`
	// Task type; 1 TikTok video posting; 2 TikTok AI account warmup; 3 TikTok carousel
	// posting; 4 TikTok account login; 6 TikTok profile editing; 42 Custom(Including Facebook,
	// YouTube and other platforms).
	TaskType int `json:"taskType"`
	// Cloud phone name.
	SerialName string `json:"serialName"`
	// Cloud phone ID.
	EnvID string `json:"envId"`
	// Scheduled execution time as a Unix timestamp in seconds.
	ScheduleAt int64 `json:"scheduleAt"`
	// Task status; 1 Waiting; 2 In progress; 3 Completed; 4 Failed; 7 Cancelled.
	Status int `json:"status"`
	// Failure code, refer to task failure codes and reasons.
	FailCode int `json:"failCode,omitempty"`
	// Failure reason, refer to task failure codes and reasons.
	FailDesc string `json:"failDesc,omitempty"`
	// The time a task takes to complete or fail (in seconds).
	Cost int `json:"cost,omitempty"`
}

type ListTasksResponse struct {
	// Total number of tasks.
	Total int `json:"total"`
	// Array of tasks.
	Items []ListTasksResponseItemsItem `json:"items"`
}

// ListTasks: https://open.geelark.com/api/task-historyRecords
//
// Only tasks scheduled within the past seven days are returned. Set LastID to the final
// item ID from the preceding page to continue pagination.
func (c *Client) ListTasks(ctx context.Context, req ListTasksRequest) (*ListTasksResponse, error) {
	var resp ListTasksResponse
	if err := c.core.Do(ctx, "/open/v1/task/historyRecords", req, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

type CancelTaskRequest struct {
	// A task ID array, with a maximum of 100 entries. Required.
	IDs []string `json:"ids"`
}

type CancelTaskResponse struct {
	// Total number processed.
	TotalAmount int `json:"totalAmount"`
	// Number of successfully processed tasks.
	SuccessAmount int `json:"successAmount"`
	// Number of failed tasks.
	FailAmount int `json:"failAmount"`
	// Details of failed tasks.
	FailDetails []CancelTaskResponseFailDetailsItem `json:"failDetails"`
}

type CancelTaskResponseFailDetailsItem struct {
	// Task ID.
	ID string `json:"id"`
	// Error code.
	Code int `json:"code"`
	// Error message.
	Msg string `json:"msg"`
}

// CancelTask: https://open.geelark.com/api/task-cancel
//
// Only tasks that are waiting for execution or currently running can be cancelled.
func (c *Client) CancelTask(ctx context.Context, req CancelTaskRequest) (*CancelTaskResponse, error) {
	var resp CancelTaskResponse
	if err := c.core.Do(ctx, "/open/v1/task/cancel", req, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

type RetryTaskRequest struct {
	// Array of task IDs. Required.
	IDs []string `json:"ids"`
}

type RetryTaskResponse struct {
	// Total number of tasks processed.
	TotalAmount int `json:"totalAmount"`
	// Number of tasks processed successfully.
	SuccessAmount int `json:"successAmount"`
	// Number of tasks failed to process.
	FailAmount int `json:"failAmount"`
	// Details of failed tasks.
	FailDetails []RetryTaskResponseFailDetailsItem `json:"failDetails"`
}

type RetryTaskResponseFailDetailsItem struct {
	// Task ID.
	ID string `json:"id"`
	// Error code.
	Code int `json:"code"`
	// Error message.
	Msg string `json:"msg"`
}

// RetryTask: https://open.geelark.com/api/task-restart
//
// Only failed or cancelled tasks can be retried, at most five times per task. This method
// submits a retry; the SDK never retries the HTTP request automatically.
func (c *Client) RetryTask(ctx context.Context, req RetryTaskRequest) (*RetryTaskResponse, error) {
	var resp RetryTaskResponse
	if err := c.core.Do(ctx, "/open/v1/task/restart", req, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

type TaskDetailRequest struct {
	// Task ID. Required.
	ID string `json:"id"`
	// Log pagination cursor returned by the preceding page. Pass it back unchanged; it may
	// contain millisecond timestamps that require int64.
	SearchAfter []int64 `json:"searchAfter,omitempty"`
}

type TaskDetailResponse struct {
	// Task ID.
	ID string `json:"id"`
	// Task plan name.
	PlanName string `json:"planName"`
	// Task type; 1 TikTok video posting; 2 TikTok AI account warmup; 3 TikTok carousel
	// posting; 4 TikTok account login; 6 TikTok profile editing; 42 Custom(Including Facebook,
	// YouTube and other platforms).
	TaskType int `json:"taskType"`
	// Cloud phone name.
	SerialName string `json:"serialName"`
	// Cloud phone ID.
	EnvID string `json:"envId"`
	// Scheduled execution time as a Unix timestamp in seconds.
	ScheduleAt int64 `json:"scheduleAt"`
	// Task status; 1 Waiting; 2 In progress; 3 Completed; 4 Failed; 7 Cancelled.
	Status int `json:"status"`
	// The time a task takes to complete or fail (in seconds).
	Cost int `json:"cost"`
	// When the task is completed or fails, take a screenshot of the information.
	ResultImages []string `json:"resultImages"`
	// Task log, with time in Zone 0, returns when the task is completed or fails.
	Logs []string `json:"logs"`
	// Log pagination cursor returned by the preceding page. Pass it back unchanged; it may
	// contain millisecond timestamps that require int64.
	SearchAfter []int64 `json:"searchAfter"`
	// Is there still a log.
	LogContinue bool `json:"logContinue"`
	// Failure code, refer to task failure codes and reasons.
	FailCode int `json:"failCode,omitempty"`
	// Failure reason, refer to task failure codes and reasons.
	FailDesc string `json:"failDesc,omitempty"`
}

// TaskDetail: https://open.geelark.com/api/task-detail
//
// Omit SearchAfter on the first call. While LogContinue is true, pass the returned
// SearchAfter unchanged to obtain the next page of logs.
func (c *Client) TaskDetail(ctx context.Context, req TaskDetailRequest) (*TaskDetailResponse, error) {
	var resp TaskDetailResponse
	if err := c.core.Do(ctx, "/open/v1/task/detail", req, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

type AddTaskRequest struct {
	// Task plan name, auto-generated if not provided.
	PlanName string `json:"planName,omitempty"`
	// Remarks, up to 200 characters.
	Remark string `json:"remark,omitempty"`
	// Task type; 1 Publish video; 2 Warmup; 3 Publish image set. Required.
	TaskType int `json:"taskType"`
	// Task parameter array, create a maximum of 100 tasks at a time. Required.
	List []AddTaskItem `json:"list"`
}

// AddTaskItem carries the fields for video, warmup, and image tasks.
// TaskType selects which fields GeeLark requires.
type AddTaskItem struct {
	// Scheduled execution time as a Unix timestamp in seconds. Applicability depends on
	// AddTaskRequest.TaskType.
	ScheduleAt int64 `json:"scheduleAt"`
	// Cloud phone ID. Applicability depends on AddTaskRequest.TaskType.
	EnvID string `json:"envId"`
	// Video URL. Applicability depends on AddTaskRequest.TaskType.
	Video string `json:"video,omitempty"`
	// Video description. Maximum 4000 characters. Applicability depends on
	// AddTaskRequest.TaskType.
	VideoDesc string `json:"videoDesc,omitempty"`
	// product id. Applicability depends on AddTaskRequest.TaskType.
	ProductID string `json:"productId,omitempty"`
	// Product display title. Applicability depends on AddTaskRequest.TaskType.
	ProductTitle string `json:"productTitle,omitempty"`
	// Similar video ID. Applicability depends on AddTaskRequest.TaskType.
	RefVideoID string `json:"refVideoId,omitempty"`
	// Maximum number of automatic retries. The value ranges from 0 to 3. The default value is
	// 3. Applicability depends on AddTaskRequest.TaskType.
	MaxTryTimes *int `json:"maxTryTimes,omitempty"`
	// Time-out period. The value ranges from 30 to 80 (unit minute). The default value is 80.
	// Applicability depends on AddTaskRequest.TaskType.
	TimeoutMin *int `json:"timeoutMin,omitempty"`
	// Same video volume, 0-100. Applicability depends on AddTaskRequest.TaskType.
	SameVideoVolume *int `json:"sameVideoVolume,omitempty"`
	// Original video volume, 0-100. Applicability depends on AddTaskRequest.TaskType.
	SourceVideoVolume *int `json:"sourceVideoVolume,omitempty"`
	// Whether to label AI generated content, default is false. Applicability depends on
	// AddTaskRequest.TaskType.
	MarkAI *bool `json:"markAI,omitempty"`
	// Video cover. Applicability depends on AddTaskRequest.TaskType.
	Cover string `json:"cover,omitempty"`
	// Whether to obtain the sharing link, the default value is false. Applicability depends on
	// AddTaskRequest.TaskType.
	NeedShareLink *bool `json:"needShareLink,omitempty"`
	// Whether to browse randomly, default is false. Applicability depends on
	// AddTaskRequest.TaskType.
	RandomUse *bool `json:"randomUse,omitempty"`
	// Random browsing time in minutes, required when randomUse is true, minimum is 1, maximum
	// is 20. Applicability depends on AddTaskRequest.TaskType.
	BrowseTime *int `json:"browseTime,omitempty"`
	// Game name, maximum 256 characters. Applicability depends on AddTaskRequest.TaskType.
	GameName string `json:"gameName,omitempty"`
	// Location, maximum 500 characters. Applicability depends on AddTaskRequest.TaskType.
	Location string `json:"location,omitempty"`
	// The product number. If left blank or set to 0, the first product will be selected. The
	// value must be between 1 and 100. Applicability depends on AddTaskRequest.TaskType.
	Index *int `json:"index,omitempty"`
	// QR code. Used when publishing a project via QR code. Applicability depends on
	// AddTaskRequest.TaskType.
	QRCode string `json:"qrcode,omitempty"`
	// Warmup action; search profile - Search personal profile; search video - Search short
	// videos; browse video - Randomly browse videos. Applicability depends on
	// AddTaskRequest.TaskType.
	Action string `json:"action,omitempty"`
	// Search keyword, required when searching behavior, optional when browsing behavior.
	// Applicability depends on AddTaskRequest.TaskType.
	Keywords []string `json:"keywords,omitempty"`
	// Browsing duration, in minutes. Applicability depends on AddTaskRequest.TaskType.
	Duration *int `json:"duration,omitempty"`
	// Warmup behavior probability configuration. If not provided, it is null (default behavior
	// is used). See the description below for the field structure. Applicability depends on
	// AddTaskRequest.TaskType.
	Probability *AddTaskProbability `json:"probability,omitempty"`
	// Image URLs. Applicability depends on AddTaskRequest.TaskType.
	Images []string `json:"images,omitempty"`
	// Same video ID. Applicability depends on AddTaskRequest.TaskType.
	VideoID string `json:"videoId,omitempty"`
	// Gallery Title. Maximum 90 characters. Applicability depends on AddTaskRequest.TaskType.
	VideoTitle string `json:"videoTitle,omitempty"`
	// Randomly match background music, default is false. Applicability depends on
	// AddTaskRequest.TaskType.
	RandomBGM *bool `json:"randomBgm,omitempty"`
}

type AddTaskProbability struct {
	// Behavior probabilities when keywords are found. See the behavior probability object
	// below. Applicability depends on AddTaskRequest.TaskType.
	HasKey *AddTaskBehavior `json:"hasKey,omitempty"`
	// Behavior probabilities when no keywords are found. See the behavior probability object
	// below. Applicability depends on AddTaskRequest.TaskType.
	NoKey *AddTaskBehavior `json:"noKey,omitempty"`
}

type AddTaskBehavior struct {
	// Like probability, 0-100. Applicability depends on AddTaskRequest.TaskType.
	Like *int `json:"like,omitempty"`
	// Follow probability, 0-100. Applicability depends on AddTaskRequest.TaskType.
	Follow *int `json:"follow,omitempty"`
	// Read comment probability, 0-100. Applicability depends on AddTaskRequest.TaskType.
	ReadComment *int `json:"readComment,omitempty"`
	// Collect probability, 0-100. Applicability depends on AddTaskRequest.TaskType.
	Collect *int `json:"collect,omitempty"`
}

type AddTaskResponse struct {
	// Array of task IDs.
	TaskIDs []string `json:"taskIds"`
}

// AddTask: https://open.geelark.com/api/task-add
//
// Every item requires ScheduleAt and EnvID. Video tasks additionally require Video; warmup
// tasks require Action and Duration (and Keywords for search actions); image tasks require
// Images. Upload media before scheduling publishing tasks. Warmup tasks created through
// this endpoint are not automatically retried.
func (c *Client) AddTask(ctx context.Context, req AddTaskRequest) (*AddTaskResponse, error) {
	var resp AddTaskResponse
	if err := c.core.Do(ctx, "/open/v1/task/add", req, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
