package automation

import (
	"context"
)

type ListFlowsRequest struct {
	// Page number, minimum value is 1. Required.
	Page int `json:"page"`
	// Number of items per page, minimum is 1, maximum is 100. Required.
	PageSize int `json:"pageSize"`
}

type ListFlowsResponseItemsItem struct {
	// Task flow id.
	ID string `json:"id"`
	// Task flow title.
	Title string `json:"title"`
	// Task flow description.
	Desc string `json:"desc"`
	// Task flow parameter field name.
	Params []string `json:"params"`
}

type ListFlowsResponse struct {
	// Total number of items.
	Total int `json:"total"`
	// Page number.
	Page int `json:"page"`
	// Number of items per page.
	PageSize int `json:"pageSize"`
	// Task flow array.
	Items []ListFlowsResponseItemsItem `json:"items"`
}

// ListFlows: https://open.geelark.com/api/task-flow-list
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
	// Task name, up to 128 characters.
	Name string `json:"name,omitempty"`
	// Remarks, up to 200 characters.
	Remark string `json:"remark,omitempty"`
	// Scheduled execution time as a Unix timestamp in seconds. Required.
	ScheduleAt int64 `json:"scheduleAt"`
	// Cloud phone ID. Required.
	ID string `json:"id"`
	// Task flow id(The ID field of the Task flow query response). Required.
	FlowID string `json:"flowId"`
	// Task flow parameters, with corresponding parameter types as follows:; String: string;
	// Batch text: array[string]; Number: number; Boolean: bool; File: array[string].
	ParamMap CreateCustomTaskRequestParamMap `json:"paramMap,omitempty"`
}

type CreateCustomTaskResponse struct {
	// ID assigned to the asynchronous task.
	TaskID string `json:"taskId"`
}

// CreateCustomTask: https://open.geelark.com/api/task-rpa-add
//
// Use ListFlows first to obtain the FlowID and the parameter names expected by ParamMap.
func (c *Client) CreateCustomTask(ctx context.Context, req CreateCustomTaskRequest) (*CreateCustomTaskResponse, error) {
	var resp CreateCustomTaskResponse
	if err := c.core.Do(ctx, "/open/v1/task/rpa/add", req, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

type ImportFlowRequest struct {
	// custom task flow id, If the corresponding ID is passed, it will be updated; if not, a
	// new one will be created.
	ID string `json:"id,omitempty"`
	// custom task flow data.
	Gal string `json:"gal,omitempty"`
}

type ImportFlowResponse struct {
	// ID of the imported custom flow.
	ID string `json:"id"`
}

// ImportFlow: https://open.geelark.com/api/task-flow-import
func (c *Client) ImportFlow(ctx context.Context, req ImportFlowRequest) (*ImportFlowResponse, error) {
	var resp ImportFlowResponse
	if err := c.core.Do(ctx, "/open/v1/task/flow/import", req, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

type ExportFlowRequest struct {
	// custom task flow id. Required.
	ID string `json:"id"`
}

type ExportFlowResponse struct {
	// Exported flow data in GeeLark GAL format.
	Gal string `json:"gal"`
}

// ExportFlow: https://open.geelark.com/api/task-flow-export
func (c *Client) ExportFlow(ctx context.Context, req ExportFlowRequest) (*ExportFlowResponse, error) {
	var resp ExportFlowResponse
	if err := c.core.Do(ctx, "/open/v1/task/flow/export", req, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
