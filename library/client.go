package library

import (
	"context"

	"github.com/wyosa/geelark-go-sdk/internal/clientcore"
)

type Client struct {
	core *clientcore.Client
}

func New(core *clientcore.Client) *Client { return &Client{core: core} }

type GetUploadURLRequest struct {
	FileType string `json:"fileType"`
}

type GetUploadURLResponse struct {
	UploadURL   string `json:"uploadUrl"`
	ResourceURL string `json:"resourceUrl"`
}

// GetUploadURL calls /open/v1/material/getUploadUrl.
func (c *Client) GetUploadURL(ctx context.Context, req GetUploadURLRequest) (*GetUploadURLResponse, error) {
	var resp GetUploadURLResponse
	if err := c.core.Do(ctx, "/open/v1/material/getUploadUrl", req, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

type CreateRequest struct {
	URL      string   `json:"url"`
	TagsID   []string `json:"tagsId,omitempty"`
	FileName string   `json:"fileName,omitempty"`
}

type CreateResponseFailDetailsItem struct {
	ID   string `json:"id"`
	Code int    `json:"code"`
	Msg  string `json:"msg"`
}

type CreateResponse struct {
	ID          string                          `json:"id"`
	FailDetails []CreateResponseFailDetailsItem `json:"failDetails"`
}

// Create calls /open/v1/material/create.
func (c *Client) Create(ctx context.Context, req CreateRequest) (*CreateResponse, error) {
	var resp CreateResponse
	if err := c.core.Do(ctx, "/open/v1/material/create", req, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

type DeleteRequest struct {
	IDs []string `json:"ids"`
}

type DeleteResponseFailDetailsItem struct {
	ID   string `json:"id"`
	Code int    `json:"code"`
	Msg  string `json:"msg"`
}

type DeleteResponse struct {
	TotalAmount   int                             `json:"totalAmount"`
	SuccessAmount int                             `json:"successAmount"`
	FailAmount    int                             `json:"failAmount"`
	FailDetails   []DeleteResponseFailDetailsItem `json:"failDetails"`
}

// Delete calls /open/v1/material/del.
func (c *Client) Delete(ctx context.Context, req DeleteRequest) (*DeleteResponse, error) {
	var resp DeleteResponse
	if err := c.core.Do(ctx, "/open/v1/material/del", req, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

type SetTagRequest struct {
	MaterialsID []string `json:"materialsId"`
	TagsID      []string `json:"tagsId,omitempty"`
}

type SetTagResponseFailDetailsItem struct {
	ID   string `json:"id"`
	Code int    `json:"code"`
	Msg  string `json:"msg"`
}

type SetTagResponse struct {
	FailDetails []SetTagResponseFailDetailsItem `json:"failDetails"`
}

// SetTag calls /open/v1/material/tag/set.
func (c *Client) SetTag(ctx context.Context, req SetTagRequest) (*SetTagResponse, error) {
	var resp SetTagResponse
	if err := c.core.Do(ctx, "/open/v1/material/tag/set", req, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

type SearchRequest struct {
	Page     int      `json:"page,omitempty"`
	PageSize int      `json:"pageSize,omitempty"`
	FileName string   `json:"fileName,omitempty"`
	TagsID   []string `json:"tagsId,omitempty"`
	Source   *int     `json:"source,omitempty"`
	FileType []int    `json:"fileType,omitempty"`
	IDs      []string `json:"ids,omitempty"`
	TagIDs   []string `json:"tagIds,omitempty"`
}

type SearchResponseListItemTagsItem struct {
	ID    string `json:"id"`
	Name  string `json:"name"`
	Color int    `json:"color"`
}

type SearchResponseListItem struct {
	ID          string                           `json:"id"`
	CreatedTime int                              `json:"createdTime"`
	FileName    string                           `json:"fileName"`
	FileSize    int                              `json:"fileSize"`
	FileURL     string                           `json:"fileUrl"`
	FileType    int                              `json:"fileType"`
	Width       int                              `json:"width"`
	Height      int                              `json:"height"`
	Source      int                              `json:"source"`
	Tags        []SearchResponseListItemTagsItem `json:"tags"`
	UserName    string                           `json:"userName"`
}

type SearchResponse struct {
	Total    int                      `json:"total"`
	Page     int                      `json:"page"`
	PageSize int                      `json:"pageSize"`
	List     []SearchResponseListItem `json:"list"`
}

// Search calls /open/v1/material/search.
func (c *Client) Search(ctx context.Context, req SearchRequest) (*SearchResponse, error) {
	var resp SearchResponse
	if err := c.core.Do(ctx, "/open/v1/material/search", req, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

type CreateTagRequest struct {
	Name  string `json:"name"`
	Color *int   `json:"color,omitempty"`
}

type CreateTagResponse struct {
	ID string `json:"id"`
}

// CreateTag calls /open/v1/material/tag/create.
func (c *Client) CreateTag(ctx context.Context, req CreateTagRequest) (*CreateTagResponse, error) {
	var resp CreateTagResponse
	if err := c.core.Do(ctx, "/open/v1/material/tag/create", req, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

type DeleteTagRequest struct {
	IDs []string `json:"ids"`
}

type DeleteTagResponseFailDetailsItem struct {
	Code int    `json:"code"`
	ID   string `json:"id"`
	Msg  string `json:"msg"`
}

type DeleteTagResponse struct {
	TotalAmount   int                                `json:"totalAmount"`
	SuccessAmount int                                `json:"successAmount"`
	FailAmount    int                                `json:"failAmount"`
	FailDetails   []DeleteTagResponseFailDetailsItem `json:"failDetails"`
}

// DeleteTag calls /open/v1/material/tag/del.
func (c *Client) DeleteTag(ctx context.Context, req DeleteTagRequest) (*DeleteTagResponse, error) {
	var resp DeleteTagResponse
	if err := c.core.Do(ctx, "/open/v1/material/tag/del", req, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

type SearchTagsRequest struct {
	Page     int     `json:"page,omitempty"`
	PageSize int     `json:"pageSize,omitempty"`
	Name     *string `json:"name,omitempty"`
}

type SearchTagsResponseListItem struct {
	ID    string `json:"id"`
	Name  string `json:"name"`
	Color int    `json:"color"`
}

type SearchTagsResponse struct {
	Total    int                          `json:"total"`
	Page     int                          `json:"page"`
	PageSize int                          `json:"pageSize"`
	List     []SearchTagsResponseListItem `json:"list"`
}

// SearchTags calls /open/v1/material/tag/search.
func (c *Client) SearchTags(ctx context.Context, req SearchTagsRequest) (*SearchTagsResponse, error) {
	var resp SearchTagsResponse
	if err := c.core.Do(ctx, "/open/v1/material/tag/search", req, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
