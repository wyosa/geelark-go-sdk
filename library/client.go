// Package library provides Library files and tags.
package library

import (
	"context"

	"github.com/wyosa/geelark-go-sdk/common"
)

type Client struct {
	core common.UploadTransport
}

func New(core common.UploadTransport) *Client { return &Client{core: core} }

type GetUploadURLRequest struct {
	// File type, Currently supported: "jpg", "jpeg", "png", "gif", "bmp", "webp","heif",
	// "heic", "mp4", "webm","xml", "apk", "xapk". Required.
	FileType string `json:"fileType"`
}

type GetUploadURLResponse struct {
	// URL for uploading the file (valid for 30 minutes).
	UploadURL string `json:"uploadUrl"`
	// URL to access the resource.
	ResourceURL string `json:"resourceUrl"`
}

// GetUploadURL: https://open.geelark.com/api/get-material-upload-url
//
// GeeLark documents a 30-day retention period for Library files. Call Create immediately
// after the storage upload; unregistered uploads are deleted. The signed upload URL itself
// expires after 30 minutes.
func (c *Client) GetUploadURL(ctx context.Context, req GetUploadURLRequest) (*GetUploadURLResponse, error) {
	var resp GetUploadURLResponse
	if err := c.core.Do(ctx, "/open/v1/material/getUploadUrl", req, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

type CreateRequest struct {
	// please refer to Upload a file to the Library
	// (https://open.geelark.com/api/get-material-upload-url). Required.
	URL    string   `json:"url"`
	TagsID []string `json:"tagsId,omitempty"`
	// file name, up to 200 characters.
	FileName string `json:"fileName,omitempty"`
}

type CreateResponseFailDetailsItem struct {
	// Tag ID.
	ID string `json:"id"`
	// Error code.
	Code int `json:"code"`
	// Error msg.
	Msg string `json:"msg"`
}

type CreateResponse struct {
	// ID of the created Library material.
	ID string `json:"id"`
	// Failure details.
	FailDetails []CreateResponseFailDetailsItem `json:"failDetails"`
}

// Create: https://open.geelark.com/api/material-create
func (c *Client) Create(ctx context.Context, req CreateRequest) (*CreateResponse, error) {
	var resp CreateResponse
	if err := c.core.Do(ctx, "/open/v1/material/create", req, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

type DeleteRequest struct {
	// material id. Required.
	IDs []string `json:"ids"`
}

type DeleteResponseFailDetailsItem struct {
	// Tag ID.
	ID string `json:"id"`
	// Error code.
	Code int `json:"code"`
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

// Delete: https://open.geelark.com/api/material-del
func (c *Client) Delete(ctx context.Context, req DeleteRequest) (*DeleteResponse, error) {
	var resp DeleteResponse
	if err := c.core.Do(ctx, "/open/v1/material/del", req, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

type SetTagRequest struct {
	// material id. Required.
	MaterialsID []string `json:"materialsId"`
	// New material tag IDs; each operation replaces all existing tags. Nil and empty slices
	// both omit tagsId in this SDK. The server behavior for an omitted field versus an empty
	// array is not specified.
	TagsID []string `json:"tagsId,omitempty"`
}

type SetTagResponseFailDetailsItem struct {
	// Tag ID.
	ID string `json:"id"`
	// Error code.
	Code int `json:"code"`
	// Error msg.
	Msg string `json:"msg"`
}

type SetTagResponse struct {
	// Details of individual failures, including the affected ID and error code.
	FailDetails []SetTagResponseFailDetailsItem `json:"failDetails"`
}

// SetTag: https://open.geelark.com/api/material-tag-set
func (c *Client) SetTag(ctx context.Context, req SetTagRequest) (*SetTagResponse, error) {
	var resp SetTagResponse
	if err := c.core.Do(ctx, "/open/v1/material/tag/set", req, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

type SearchRequest struct {
	Page int `json:"page,omitempty"`
	// page size （max: 200）.
	PageSize int `json:"pageSize,omitempty"`
	// search file name.
	FileName string `json:"fileName,omitempty"`
	// search tag id.
	TagsID []string `json:"tagsId,omitempty"`
	// source: 0-upload，1-AI Edit 2Baidu Cloud Drive 3GhostCut 4GoogleDrive 5Image to video.
	Source *int `json:"source,omitempty"`
	// file type 1-image，2-video.
	FileType []int `json:"fileType,omitempty"`
	// ID array，The maximum length of the array is 100.
	IDs []string `json:"ids,omitempty"`
	// Additional tag ID filter shown in the JSON example. The parameter table lists TagsID
	// instead; server precedence is not documented.
	TagIDs []string `json:"tagIds,omitempty"`
}

type SearchResponseListItemTagsItem struct {
	// Library tag ID.
	ID string `json:"id"`
	// Library tag name. The example uses a string despite the table describing it as an
	// integer.
	Name string `json:"name"`
	// tag color: 0 White 1 Red 2 Blue 3 Green 4 Yellow 5 Purple.
	Color int `json:"color"`
}

type SearchResponseListItem struct {
	// Library material ID.
	ID string `json:"id"`
	// Creation time, second level timestamp.
	CreatedTime int64 `json:"createdTime"`
	// material name.
	FileName string `json:"fileName"`
	// material file size, byte.
	FileSize int64 `json:"fileSize"`
	// material url.
	FileURL string `json:"fileUrl"`
	// material file type 1-image，2-video.
	FileType int `json:"fileType"`
	Width    int `json:"width"`
	Height   int `json:"height"`
	// source: 0-upload，1-AI Edit 2Baidu Cloud Drive 3GhostCut 4GoogleDrive 5Image to video.
	Source int `json:"source"`
	// tag data.
	Tags []SearchResponseListItemTagsItem `json:"tags"`
	// upload user name.
	UserName string `json:"userName"`
}

type SearchResponse struct {
	// total data.
	Total int `json:"total"`
	// current page.
	Page int `json:"page"`
	// current page size.
	PageSize int `json:"pageSize"`
	// material data.
	List []SearchResponseListItem `json:"list"`
}

// Search: https://open.geelark.com/api/material-search
func (c *Client) Search(ctx context.Context, req SearchRequest) (*SearchResponse, error) {
	var resp SearchResponse
	if err := c.core.Do(ctx, "/open/v1/material/search", req, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

type CreateTagRequest struct {
	// tag name（Up to 30 characters）. Required.
	Name string `json:"name"`
	// tag color: 0 White 1 Red 2 Blue 3 Green 4 Yellow 5 Purple.
	Color *int `json:"color,omitempty"`
}

type CreateTagResponse struct {
	// tag id.
	ID string `json:"id"`
}

// CreateTag: https://open.geelark.com/api/material-tag-create
func (c *Client) CreateTag(ctx context.Context, req CreateTagRequest) (*CreateTagResponse, error) {
	var resp CreateTagResponse
	if err := c.core.Do(ctx, "/open/v1/material/tag/create", req, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

type DeleteTagRequest struct {
	// tag id. Required.
	IDs []string `json:"ids"`
}

type DeleteTagResponseFailDetailsItem struct {
	// Error code.
	Code int `json:"code"`
	// Tag ID.
	ID string `json:"id"`
	// Error msg.
	Msg string `json:"msg"`
}

type DeleteTagResponse struct {
	// Total delete requests.
	TotalAmount int `json:"totalAmount"`
	// Total successes.
	SuccessAmount int `json:"successAmount"`
	// Total failures.
	FailAmount int `json:"failAmount"`
	// Failure details.
	FailDetails []DeleteTagResponseFailDetailsItem `json:"failDetails"`
}

// DeleteTag: https://open.geelark.com/api/material-tag-del
func (c *Client) DeleteTag(ctx context.Context, req DeleteTagRequest) (*DeleteTagResponse, error) {
	var resp DeleteTagResponse
	if err := c.core.Do(ctx, "/open/v1/material/tag/del", req, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

type SearchTagsRequest struct {
	Page int `json:"page,omitempty"`
	// page size （max: 200）.
	PageSize int `json:"pageSize,omitempty"`
	// search tag name.
	Name *string `json:"name,omitempty"`
}

type SearchTagsResponseListItem struct {
	// tag id.
	ID string `json:"id"`
	// tag name.
	Name string `json:"name"`
	// tag color: 0 White 1 Red 2 Blue 3 Green 4 Yellow 5 Purple.
	Color int `json:"color"`
}

type SearchTagsResponse struct {
	// total data.
	Total int `json:"total"`
	// current page.
	Page int `json:"page"`
	// current page size.
	PageSize int `json:"pageSize"`
	// tag data.
	List []SearchTagsResponseListItem `json:"list"`
}

// SearchTags: https://open.geelark.com/api/material-tag-search
func (c *Client) SearchTags(ctx context.Context, req SearchTagsRequest) (*SearchTagsResponse, error) {
	var resp SearchTagsResponse
	if err := c.core.Do(ctx, "/open/v1/material/tag/search", req, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
