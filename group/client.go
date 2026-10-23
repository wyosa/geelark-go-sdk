// Package group provides cloud phone profile groups.
package group

import (
	"context"

	"github.com/wyosa/geelark-go-sdk/common"
)

type Client struct {
	core common.Transport
}

func New(core common.Transport) *Client { return &Client{core: core} }

type CreateRequestListItem struct {
	// Group name, up to 50 characters. Required.
	Name string `json:"name,omitempty"`
	// Group remark, up to 500 characters.
	Remark string `json:"remark,omitempty"`
}

type CreateRequest struct {
	// Group data list. Required.
	List []CreateRequestListItem `json:"list"`
}

type CreateResponseSuccessDetailsItem struct {
	// Group ID.
	ID string `json:"id"`
	// Group name.
	Name string `json:"name"`
	// Remark (not shown if empty).
	Remark string `json:"remark,omitempty"`
}

type CreateResponseFailDetailsItem struct {
	// Error code.
	Code int `json:"code"`
	// Group name.
	Name string `json:"name"`
	// Error msg.
	Msg string `json:"msg"`
}

type CreateResponse struct {
	// Total requests.
	TotalAmount int `json:"totalAmount"`
	// Total successes.
	SuccessAmount int `json:"successAmount"`
	// Total failures.
	FailAmount     int                                `json:"failAmount"`
	SuccessDetails []CreateResponseSuccessDetailsItem `json:"successDetails"`
	// Failure details.
	FailDetails []CreateResponseFailDetailsItem `json:"failDetails"`
}

// Create: https://open.geelark.com/api/group-add
func (c *Client) Create(ctx context.Context, req CreateRequest) (*CreateResponse, error) {
	var resp CreateResponse
	if err := c.core.Do(ctx, "/open/v1/group/add", req, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

type DeleteRequest struct {
	// List of group IDs to delete. Required.
	IDs []string `json:"ids"`
}

type DeleteResponseFailDetailsItem struct {
	// Error code.
	Code int `json:"code"`
	// Group ID.
	ID string `json:"id"`
	// Error msg.
	Msg string `json:"msg"`
}

type DeleteResponse struct {
	// Total delete requests.
	TotalAmount int `json:"totalAmount"`
	// Total successes.
	SuccessAmount int `json:"successAmount"`
	// Total failures.
	FailAmount int `json:"failAmount"`
	// Failure details.
	FailDetails []DeleteResponseFailDetailsItem `json:"failDetails"`
}

// Delete: https://open.geelark.com/api/group-delete
func (c *Client) Delete(ctx context.Context, req DeleteRequest) (*DeleteResponse, error) {
	var resp DeleteResponse
	if err := c.core.Do(ctx, "/open/v1/group/delete", req, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

type ListRequest struct {
	// Page number, minimum is 1. Required.
	Page int `json:"page"`
	// Number of items per page, minimum is 1, maximum is 100. Required.
	PageSize int `json:"pageSize"`
	// List of group IDs.
	IDs []string `json:"ids,omitempty"`
	// List of group names.
	Names []string `json:"names,omitempty"`
	// List of group remarks.
	Remarks []string `json:"remarks,omitempty"`
}

type ListResponseListItem struct {
	// Group ID.
	ID string `json:"id"`
	// Group name.
	Name string `json:"name"`
	// Group remark.
	Remark string `json:"remark"`
}

type ListResponse struct {
	// Total count.
	Total int `json:"total"`
	// Page number.
	Page     int `json:"page"`
	PageSize int `json:"pageSize"`
	// Group list.
	List []ListResponseListItem `json:"list"`
	// Failure details.
	FailDetails []ListResponseFailDetail `json:"failDetails,omitempty"`
}

type ListResponseFailDetail struct {
	// Error code.
	Code int `json:"code"`
	// Failure type: 1-Group ID, 2-Group name, 3-Group remark.
	Type int `json:"type"`
	// Failed parameter.
	Param string `json:"param"`
	// Failure message.
	Msg string `json:"msg"`
}

// List: https://open.geelark.com/api/group-list
func (c *Client) List(ctx context.Context, req ListRequest) (*ListResponse, error) {
	var resp ListResponse
	if err := c.core.Do(ctx, "/open/v1/group/list", req, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

type UpdateRequestListItem struct {
	// Group ID. Required.
	ID *string `json:"id,omitempty"`
	// New group name, up to 50 characters.
	Name *string `json:"name,omitempty"`
	// New group remark, up to 500 characters.
	Remark *string `json:"remark,omitempty"`
}

type UpdateRequest struct {
	// Group data. Required.
	List []UpdateRequestListItem `json:"list"`
}

type UpdateResponse struct {
	// Total number of processed items.
	TotalAmount int `json:"totalAmount"`
	// Number of successful items.
	SuccessAmount int `json:"successAmount"`
	// Number of failed items.
	FailAmount int `json:"failAmount"`
}

// Update: https://open.geelark.com/api/group-update
func (c *Client) Update(ctx context.Context, req UpdateRequest) (*UpdateResponse, error) {
	var resp UpdateResponse
	if err := c.core.Do(ctx, "/open/v1/group/update", req, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
