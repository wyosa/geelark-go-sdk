// Package tag provides cloud phone profile tags.
package tag

import (
	"context"

	"github.com/wyosa/geelark-go-sdk/common"
)

type Client struct {
	core common.Transport
}

func New(core common.Transport) *Client { return &Client{core: core} }

type CreateRequestListItem struct {
	// Tag name, up to 30 characters. Required.
	Name string `json:"name,omitempty"`
	// Tag color.
	Color *string `json:"color,omitempty"`
}

type CreateRequest struct {
	// Tag data list. Required.
	List []CreateRequestListItem `json:"list"`
}

type CreateResponseSuccessDetailsItem struct {
	// Tag ID.
	ID string `json:"id"`
	// Tag name.
	Name string `json:"name"`
	// Tag color.
	Color string `json:"color"`
}

type CreateResponseFailDetailsItem struct {
	// Error code.
	Code int `json:"code"`
	// Tag name.
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

// Create: https://open.geelark.com/api/tag-add
func (c *Client) Create(ctx context.Context, req CreateRequest) (*CreateResponse, error) {
	var resp CreateResponse
	if err := c.core.Do(ctx, "/open/v1/tag/add", req, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

type DeleteRequest struct {
	// List of tag IDs to delete. Required.
	IDs []string `json:"ids"`
}

type DeleteResponseFailDetailsItem struct {
	// Error code.
	Code int `json:"code"`
	// Tag ID.
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

// Delete: https://open.geelark.com/api/tag-delete
func (c *Client) Delete(ctx context.Context, req DeleteRequest) (*DeleteResponse, error) {
	var resp DeleteResponse
	if err := c.core.Do(ctx, "/open/v1/tag/delete", req, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

type ListRequest struct {
	// Page number, minimum is 1. Required.
	Page int `json:"page"`
	// Number of items per page, minimum is 1, maximum is 100. Required.
	PageSize int `json:"pageSize"`
	// List of tag IDs.
	IDs []string `json:"ids,omitempty"`
	// List of tag names.
	Names []string `json:"names,omitempty"`
	// List of tag colors.
	Colors []string `json:"colors,omitempty"`
}

type ListResponseListItem struct {
	// Tag ID.
	ID string `json:"id"`
	// Tag name.
	Name string `json:"name"`
	// Tag color.
	Color string `json:"color"`
}

type ListResponse struct {
	// Total count.
	Total int `json:"total"`
	// Page number.
	Page     int `json:"page"`
	PageSize int `json:"pageSize"`
	// Tag list.
	List []ListResponseListItem `json:"list"`
	// Failure details.
	FailDetails []ListResponseFailDetail `json:"failDetails,omitempty"`
}

type ListResponseFailDetail struct {
	// Error code.
	Code int `json:"code"`
	// Failure type: 1-Tag ID, 2-Tag name, 3-Tag color.
	Type int `json:"type"`
	// Failed parameter.
	Param string `json:"param"`
	// Failure message.
	Msg string `json:"msg"`
}

// List: https://open.geelark.com/api/tag-list
func (c *Client) List(ctx context.Context, req ListRequest) (*ListResponse, error) {
	var resp ListResponse
	if err := c.core.Do(ctx, "/open/v1/tag/list", req, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

type UpdateRequestListItem struct {
	// Tag ID. Required.
	ID *string `json:"id,omitempty"`
	// New tag name, up to 30 characters.
	Name *string `json:"name,omitempty"`
	// New tag color.
	Color *string `json:"color,omitempty"`
}

type UpdateRequest struct {
	// Array of tags to update. Required.
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

// Update: https://open.geelark.com/api/tag-update
func (c *Client) Update(ctx context.Context, req UpdateRequest) (*UpdateResponse, error) {
	var resp UpdateResponse
	if err := c.core.Do(ctx, "/open/v1/tag/update", req, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
