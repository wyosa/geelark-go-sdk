// Package proxygroup provides proxy groups and their membership categories.
package proxygroup

import (
	"context"

	"github.com/wyosa/geelark-go-sdk/common"
)

type Client struct {
	core common.Transport
}

func New(core common.Transport) *Client { return &Client{core: core} }

type CreateRequest struct {
	// Proxy group name, up to 50 characters. Required.
	Name string `json:"name"`
}

type CreateResponse struct {
	// Proxy group ID.
	ID string `json:"id"`
}

// Create: https://open.geelark.com/api/proxy-group-add
func (c *Client) Create(ctx context.Context, req CreateRequest) (*CreateResponse, error) {
	var resp CreateResponse
	if err := c.core.Do(ctx, "/open/v1/proxyGroup/add", req, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

type ListRequest struct {
	// Page number, default is 1.
	Page int `json:"page,omitempty"`
	// Number of items per page. Required.
	PageSize int `json:"pageSize"`
	// Proxy group name.
	Name string `json:"name,omitempty"`
}

type ListResponseListItem struct {
	// Proxy group ID. 0 represents ungrouped proxies.
	ID string `json:"id"`
	// Proxy group name.
	Name string `json:"name"`
}

type ListResponse struct {
	// Total number of records, including the virtual ungrouped category.
	Total int `json:"total"`
	// Page number.
	Page int `json:"page"`
	// Number of items per page.
	PageSize int `json:"pageSize"`
	// Proxy group list.
	List []ListResponseListItem `json:"list"`
}

// List: https://open.geelark.com/api/proxy-group-search
//
// ID "0" denotes the virtual Ungrouped Proxies category. It appears on the first page and
// is included in Total.
func (c *Client) List(ctx context.Context, req ListRequest) (*ListResponse, error) {
	var resp ListResponse
	if err := c.core.Do(ctx, "/open/v1/proxyGroup/search", req, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

type UpdateRequest struct {
	// Proxy group ID. Required.
	ID string `json:"id"`
	// New proxy group name, up to 50 characters. Required.
	Name string `json:"name"`
}

type UpdateResponse struct{}

// Update: https://open.geelark.com/api/proxy-group-update
func (c *Client) Update(ctx context.Context, req UpdateRequest) (*UpdateResponse, error) {
	var resp UpdateResponse
	if err := c.core.Do(ctx, "/open/v1/proxyGroup/update", req, nil); err != nil {
		return nil, err
	}
	return &resp, nil
}

type DeleteRequest struct {
	// Proxy group ID list. Duplicate IDs are counted only once. Required.
	IDs []string `json:"ids"`
}

type DeleteResponse struct {
	// Total number of items processed. Duplicate IDs are counted only once.
	TotalAmount int `json:"totalAmount"`
	// Number of successfully deleted proxy groups.
	SuccessAmount int `json:"successAmount"`
	// Number of failed deletions.
	FailAmount int `json:"failAmount"`
	// Failed proxy group details.
	FailDetails []DeleteResponseFailDetailsItem `json:"failDetails"`
}

type DeleteResponseFailDetailsItem struct {
	// Proxy group ID.
	ID string `json:"id"`
	// Error code.
	Code int `json:"code"`
	// Error message.
	Msg string `json:"msg"`
}

// Delete: https://open.geelark.com/api/proxy-group-delete
//
// Proxies in deleted groups are moved to the ungrouped category. Nonexistent groups,
// already deleted groups, and group ID "0" are treated as successfully deleted.
func (c *Client) Delete(ctx context.Context, req DeleteRequest) (*DeleteResponse, error) {
	var resp DeleteResponse
	if err := c.core.Do(ctx, "/open/v1/proxyGroup/del", req, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
