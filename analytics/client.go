package analytics

import (
	"context"

	"github.com/wyosa/geelark-go-sdk/internal/clientcore"
)

type Client struct {
	core *clientcore.Client
}

func New(core *clientcore.Client) *Client { return &Client{core: core} }

type DataRequest struct {
	Page      int    `json:"page"`
	PageSize  int    `json:"pageSize"`
	Account   string `json:"account,omitempty"`
	DataDate  *int   `json:"dataDate,omitempty"`
	CreatedID string `json:"createdId,omitempty"`
	Channel   *int   `json:"channel,omitempty"`
}

type DataResponseItemsItem struct {
	ID            string `json:"id"`
	Channel       int    `json:"channel"`
	Account       string `json:"account"`
	PlayCount     int    `json:"playCount"`
	FollowerCount int    `json:"followerCount"`
	DiggCount     int    `json:"diggCount"`
	CommentCount  int    `json:"commentCount"`
	CollectCount  int    `json:"collectCount"`
	ShareCount    int    `json:"shareCount"`
	DataDate      int    `json:"dataDate"`
	AddAccDate    int    `json:"addAccDate"`
	Remark        string `json:"remark"`
	CreatedID     string `json:"createdId"`
	Username      string `json:"username"`
}

type DataResponse struct {
	Total    int                     `json:"total"`
	Page     int                     `json:"page"`
	PageSize int                     `json:"pageSize"`
	Items    []DataResponseItemsItem `json:"items"`
}

// Data calls /open/v1/analytics/data.
func (c *Client) Data(ctx context.Context, req DataRequest) (*DataResponse, error) {
	var resp DataResponse
	if err := c.core.Do(ctx, "/open/v1/analytics/data", req, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

type AddAccountsRequestAccountsDataItem struct {
	Account string   `json:"account,omitempty"`
	Remark  string   `json:"remark,omitempty"`
	TagIDs  []string `json:"tagIds,omitempty"`
}

type AddAccountsRequest struct {
	Channel      int                                  `json:"channel"`
	AccountsData []AddAccountsRequestAccountsDataItem `json:"accountsData"`
}

type AddAccountsResponse struct {
	BizCode      int `json:"bizCode"`
	SuccessCount int `json:"successCount"`
	FailCount    int `json:"failCount"`
	RepeatCount  int `json:"repeatCount"`
}

// AddAccounts calls /open/v1/analytics/accounts/add.
func (c *Client) AddAccounts(ctx context.Context, req AddAccountsRequest) (*AddAccountsResponse, error) {
	var resp AddAccountsResponse
	if err := c.core.Do(ctx, "/open/v1/analytics/accounts/add", req, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

type DeleteAccountRequest struct {
	Channel int    `json:"channel"`
	Account string `json:"account"`
}

type DeleteAccountResponse struct{}

// DeleteAccount calls /open/v1/analytics/accounts/delete.
func (c *Client) DeleteAccount(ctx context.Context, req DeleteAccountRequest) (*DeleteAccountResponse, error) {
	var resp DeleteAccountResponse
	if err := c.core.Do(ctx, "/open/v1/analytics/accounts/delete", req, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

type ListAccountsRequest struct {
	Page        int    `json:"page"`
	PageSize    int    `json:"pageSize"`
	Account     string `json:"account,omitempty"`
	Channel     *int   `json:"channel,omitempty"`
	UserAccount string `json:"userAccount,omitempty"`
}

type ListAccountsResponseItemsItemTagsItem struct {
	ID    string `json:"id"`
	Name  string `json:"name"`
	Color int    `json:"color"`
}

type ListAccountsResponseItemsItem struct {
	ID          string                                  `json:"id"`
	Account     string                                  `json:"account"`
	Channel     int                                     `json:"channel"`
	Remark      string                                  `json:"remark"`
	Operator    string                                  `json:"operator"`
	CreatedTime int                                     `json:"created_time"`
	UpdatedTime int                                     `json:"updated_time"`
	Tags        []ListAccountsResponseItemsItemTagsItem `json:"tags"`
}

type ListAccountsResponse struct {
	Items []ListAccountsResponseItemsItem `json:"items"`
	Total int                             `json:"total"`
	Page  int                             `json:"page"`
}

// ListAccounts calls /open/v1/analytics/accounts/list.
func (c *Client) ListAccounts(ctx context.Context, req ListAccountsRequest) (*ListAccountsResponse, error) {
	var resp ListAccountsResponse
	if err := c.core.Do(ctx, "/open/v1/analytics/accounts/list", req, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

type UpdateAccountRequest struct {
	ID      string   `json:"id"`
	Account string   `json:"account,omitempty"`
	Channel *int     `json:"channel,omitempty"`
	Remark  string   `json:"remark,omitempty"`
	TagIDs  []string `json:"tagIds,omitempty"`
}

type UpdateAccountResponse struct{}

// UpdateAccount calls /open/v1/analytics/accounts/update.
func (c *Client) UpdateAccount(ctx context.Context, req UpdateAccountRequest) (*UpdateAccountResponse, error) {
	var resp UpdateAccountResponse
	if err := c.core.Do(ctx, "/open/v1/analytics/accounts/update", req, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

type CreateTagRequest struct {
	Name  string `json:"name"`
	Color *int   `json:"color,omitempty"`
}

type CreateTagResponse struct{}

// CreateTag calls /open/v1/analytics/tags/create.
func (c *Client) CreateTag(ctx context.Context, req CreateTagRequest) (*CreateTagResponse, error) {
	var resp CreateTagResponse
	if err := c.core.Do(ctx, "/open/v1/analytics/tags/create", req, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

type ListTagsRequest struct {
	Name     string `json:"name,omitempty"`
	Page     int    `json:"page,omitempty"`
	PageSize int    `json:"pageSize,omitempty"`
}

type ListTagsResponseListItem struct {
	ID    string `json:"id"`
	Name  string `json:"name"`
	Color int    `json:"color"`
}

type ListTagsResponse struct {
	Total    int                        `json:"total"`
	Page     int                        `json:"page"`
	PageSize int                        `json:"pageSize"`
	List     []ListTagsResponseListItem `json:"list"`
}

// ListTags calls /open/v1/analytics/tags/list.
func (c *Client) ListTags(ctx context.Context, req ListTagsRequest) (*ListTagsResponse, error) {
	var resp ListTagsResponse
	if err := c.core.Do(ctx, "/open/v1/analytics/tags/list", req, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

type UpdateTagRequest struct {
	ID    string `json:"id,omitempty"`
	Name  string `json:"name,omitempty"`
	Color *int   `json:"color,omitempty"`
}

type UpdateTagResponse struct{}

// UpdateTag calls /open/v1/analytics/tags/update.
func (c *Client) UpdateTag(ctx context.Context, req UpdateTagRequest) (*UpdateTagResponse, error) {
	var resp UpdateTagResponse
	if err := c.core.Do(ctx, "/open/v1/analytics/tags/update", req, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

type DeleteTagsRequest struct {
	IDs []string `json:"ids,omitempty"`
}

type DeleteTagsResponse struct{}

// DeleteTags calls /open/v1/analytics/tags/del.
func (c *Client) DeleteTags(ctx context.Context, req DeleteTagsRequest) (*DeleteTagsResponse, error) {
	var resp DeleteTagsResponse
	if err := c.core.Do(ctx, "/open/v1/analytics/tags/del", req, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
