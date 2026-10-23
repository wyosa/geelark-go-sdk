package tag

import (
	"context"

	"github.com/wyosa/geelark-go-sdk/internal/clientcore"
)

type Client struct {
	core *clientcore.Client
}

func New(core *clientcore.Client) *Client { return &Client{core: core} }

type CreateRequestListItem struct {
	Name  string  `json:"name,omitempty"`
	Color *string `json:"color,omitempty"`
}

type CreateRequest struct {
	List []CreateRequestListItem `json:"list"`
}

type CreateResponseSuccessDetailsItem struct {
	ID    string `json:"id"`
	Name  string `json:"name"`
	Color string `json:"color"`
}

type CreateResponseFailDetailsItem struct {
	Code int    `json:"code"`
	Name string `json:"name"`
	Msg  string `json:"msg"`
}

type CreateResponse struct {
	TotalAmount    int                                `json:"totalAmount"`
	SuccessAmount  int                                `json:"successAmount"`
	FailAmount     int                                `json:"failAmount"`
	SuccessDetails []CreateResponseSuccessDetailsItem `json:"successDetails"`
	FailDetails    []CreateResponseFailDetailsItem    `json:"failDetails"`
}

// Create calls /open/v1/tag/add.
func (c *Client) Create(ctx context.Context, req CreateRequest) (*CreateResponse, error) {
	var resp CreateResponse
	if err := c.core.Do(ctx, "/open/v1/tag/add", req, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

type DeleteRequest struct {
	IDs []string `json:"ids"`
}

type DeleteResponseFailDetailsItem struct {
	Code int    `json:"code"`
	ID   string `json:"id"`
	Msg  string `json:"msg"`
}

type DeleteResponse struct {
	TotalAmount   int                             `json:"totalAmount"`
	SuccessAmount int                             `json:"successAmount"`
	FailAmount    int                             `json:"failAmount"`
	FailDetails   []DeleteResponseFailDetailsItem `json:"failDetails"`
}

// Delete calls /open/v1/tag/delete.
func (c *Client) Delete(ctx context.Context, req DeleteRequest) (*DeleteResponse, error) {
	var resp DeleteResponse
	if err := c.core.Do(ctx, "/open/v1/tag/delete", req, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

type ListRequest struct {
	Page     int      `json:"page"`
	PageSize int      `json:"pageSize"`
	IDs      []string `json:"ids,omitempty"`
	Names    []string `json:"names,omitempty"`
	Colors   []string `json:"colors,omitempty"`
}

type ListResponseListItem struct {
	ID    string `json:"id"`
	Name  string `json:"name"`
	Color string `json:"color"`
}

type ListResponse struct {
	Total       int                      `json:"total"`
	Page        int                      `json:"page"`
	PageSize    int                      `json:"pageSize"`
	List        []ListResponseListItem   `json:"list"`
	FailDetails []ListResponseFailDetail `json:"failDetails,omitempty"`
}

type ListResponseFailDetail struct {
	Code  int    `json:"code"`
	Type  int    `json:"type"`
	Param string `json:"param"`
	Msg   string `json:"msg"`
}

// List calls /open/v1/tag/list.
func (c *Client) List(ctx context.Context, req ListRequest) (*ListResponse, error) {
	var resp ListResponse
	if err := c.core.Do(ctx, "/open/v1/tag/list", req, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

type UpdateRequestListItem struct {
	ID    *string `json:"id,omitempty"`
	Name  *string `json:"name,omitempty"`
	Color *string `json:"color,omitempty"`
}

type UpdateRequest struct {
	List []UpdateRequestListItem `json:"list"`
}

type UpdateResponse struct {
	TotalAmount   int `json:"totalAmount"`
	SuccessAmount int `json:"successAmount"`
	FailAmount    int `json:"failAmount"`
}

// Update calls /open/v1/tag/update.
func (c *Client) Update(ctx context.Context, req UpdateRequest) (*UpdateResponse, error) {
	var resp UpdateResponse
	if err := c.core.Do(ctx, "/open/v1/tag/update", req, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
