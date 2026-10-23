package proxygroup

import (
	"context"

	"github.com/wyosa/geelark-go-sdk/internal/clientcore"
)

type Client struct {
	core *clientcore.Client
}

func New(core *clientcore.Client) *Client { return &Client{core: core} }

type CreateRequest struct {
	Name string `json:"name"`
}

type CreateResponse struct {
	ID string `json:"id"`
}

// Create calls /open/v1/proxyGroup/add.
func (c *Client) Create(ctx context.Context, req CreateRequest) (*CreateResponse, error) {
	var resp CreateResponse
	if err := c.core.Do(ctx, "/open/v1/proxyGroup/add", req, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

type ListRequest struct {
	Page     int    `json:"page,omitempty"`
	PageSize int    `json:"pageSize"`
	Name     string `json:"name,omitempty"`
}

type ListResponseListItem struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

type ListResponse struct {
	Total    int                    `json:"total"`
	Page     int                    `json:"page"`
	PageSize int                    `json:"pageSize"`
	List     []ListResponseListItem `json:"list"`
}

// List calls /open/v1/proxyGroup/search.
func (c *Client) List(ctx context.Context, req ListRequest) (*ListResponse, error) {
	var resp ListResponse
	if err := c.core.Do(ctx, "/open/v1/proxyGroup/search", req, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

type UpdateRequest struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

type UpdateResponse struct{}

// Update calls /open/v1/proxyGroup/update.
func (c *Client) Update(ctx context.Context, req UpdateRequest) (*UpdateResponse, error) {
	var resp UpdateResponse
	if err := c.core.Do(ctx, "/open/v1/proxyGroup/update", req, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

type DeleteRequest struct {
	IDs []string `json:"ids"`
}

type DeleteResponse struct {
	TotalAmount   int                             `json:"totalAmount"`
	SuccessAmount int                             `json:"successAmount"`
	FailAmount    int                             `json:"failAmount"`
	FailDetails   []DeleteResponseFailDetailsItem `json:"failDetails"`
}

type DeleteResponseFailDetailsItem struct {
	ID   string `json:"id"`
	Code int    `json:"code"`
	Msg  string `json:"msg"`
}

// Delete calls /open/v1/proxyGroup/del.
func (c *Client) Delete(ctx context.Context, req DeleteRequest) (*DeleteResponse, error) {
	var resp DeleteResponse
	if err := c.core.Do(ctx, "/open/v1/proxyGroup/del", req, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
