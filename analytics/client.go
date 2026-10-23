// Package analytics provides Data Assistant accounts, tags, and social account metrics.
package analytics

import (
	"context"

	"github.com/wyosa/geelark-go-sdk/common"
)

type Client struct {
	core common.Transport
}

func New(core common.Transport) *Client { return &Client{core: core} }

type DataRequest struct {
	// page. Required.
	Page int `json:"page"`
	// page size（1-100）. Required.
	PageSize int    `json:"pageSize"`
	Account  string `json:"account,omitempty"`
	// Search date as a Unix timestamp in seconds. Nil omits the filter.
	DataDate *int64 `json:"dataDate,omitempty"`
	// ID of the user who added the account.
	CreatedID string `json:"createdId,omitempty"`
	// account channel, 0: TikTok 1:YouTube 2:Instagram, 4: Reddit.
	Channel *int `json:"channel,omitempty"`
}

type DataResponseItemsItem struct {
	// account id.
	ID string `json:"id"`
	// account channel.
	Channel int    `json:"channel"`
	Account string `json:"account"`
	// play count , -1 indicates that the data has not been updated yet.
	PlayCount int64 `json:"playCount"`
	// follower count , -1 indicates that the data has not been updated yet.
	FollowerCount int64 `json:"followerCount"`
	// digg count , -1 indicates that the data has not been updated yet.
	DiggCount int64 `json:"diggCount"`
	// comment count , -1 indicates that the data has not been updated yet.
	CommentCount int64 `json:"commentCount"`
	// collect count , -1 indicates that the data has not been updated yet.
	CollectCount int64 `json:"collectCount"`
	// share count , -1 indicates that the data has not been updated yet.
	ShareCount int64 `json:"shareCount"`
	// Metrics date as a Unix timestamp in seconds; -1 indicates data not updated yet.
	DataDate int64 `json:"dataDate"`
	// Time the account was added as a Unix timestamp in seconds.
	AddAccDate int64  `json:"addAccDate"`
	Remark     string `json:"remark"`
	// ID of the user who added the account.
	CreatedID string `json:"createdId"`
	// Username of the user who added the account.
	Username string `json:"username"`
}

type DataResponse struct {
	Total int `json:"total"`
	// current page.
	Page int `json:"page"`
	// current page size.
	PageSize int `json:"pageSize"`
	// Account metrics for this page.
	Items []DataResponseItemsItem `json:"items"`
}

// Data: https://open.geelark.com/api/analytics-data
func (c *Client) Data(ctx context.Context, req DataRequest) (*DataResponse, error) {
	var resp DataResponse
	if err := c.core.Do(ctx, "/open/v1/analytics/data", req, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

type AddAccountsRequestAccountsDataItem struct {
	// Platform account, up to 64 characters. Required.
	Account string `json:"account,omitempty"`
	Remark  string `json:"remark,omitempty"`
	// Data Assistant tag ID list. If omitted or an empty array is provided, no tags will be
	// assigned; if a non-empty array is provided, the specified tags will be assigned to the
	// account. Duplicate IDs are automatically removed, and a maximum of 20 tags is supported
	// after deduplication. The tags must belong to the current team and must be Data Assistant
	// tags.
	TagIDs []string `json:"tagIds,omitempty"`
}

type AddAccountsRequest struct {
	// Platform; 0: TikTok 1: YouTube 2: Instagram. Required.
	Channel int `json:"channel"`
	// Account information list. A maximum of 200 accounts is supported per request. Required.
	AccountsData []AddAccountsRequestAccountsDataItem `json:"accountsData"`
}

type AddAccountsResponse struct {
	// Business status code: 0 indicates that all accounts were added successfully; 1 indicates
	// that the current account count exceeds the limit; 2 indicates partial success, with some
	// accounts failing because the account count exceeds the limit.
	BizCode int `json:"bizCode"`
	// Number of accounts successfully added.
	SuccessCount int `json:"successCount"`
	// Number of accounts that failed to be added.
	FailCount int `json:"failCount"`
	// Number of duplicate accounts.
	RepeatCount int `json:"repeatCount"`
}

// AddAccounts: https://open.geelark.com/api/analytics-accounts-add
func (c *Client) AddAccounts(ctx context.Context, req AddAccountsRequest) (*AddAccountsResponse, error) {
	var resp AddAccountsResponse
	if err := c.core.Do(ctx, "/open/v1/analytics/accounts/add", req, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

type DeleteAccountRequest struct {
	// Platform; 0: TikTok, 1: YouTube, 2: Instagram 4: Reddit. Required.
	Channel int `json:"channel"`
	// Account name, maximum length 64 characters. Required.
	Account string `json:"account"`
}

type DeleteAccountResponse struct{}

// DeleteAccount: https://open.geelark.com/api/analytics-accounts-delete
func (c *Client) DeleteAccount(ctx context.Context, req DeleteAccountRequest) (*DeleteAccountResponse, error) {
	var resp DeleteAccountResponse
	if err := c.core.Do(ctx, "/open/v1/analytics/accounts/delete", req, nil); err != nil {
		return nil, err
	}
	return &resp, nil
}

type ListAccountsRequest struct {
	// Page number. Required.
	Page int `json:"page"`
	// Number of items per page (1–100). Required.
	PageSize int `json:"pageSize"`
	// Account name.
	Account string `json:"account,omitempty"`
	// Platform, if not provided all platforms are included; 0: TikTok 1: YouTube 2: Instagram,
	// 4: Reddit.
	Channel *int `json:"channel,omitempty"`
	// Operator account.
	UserAccount string `json:"userAccount,omitempty"`
}

type ListAccountsResponseItemsItemTagsItem struct {
	// Data Assistant tag ID.
	ID string `json:"id"`
	// Tag name.
	Name string `json:"name"`
	// Tag color index.
	Color int `json:"color"`
}

type ListAccountsResponseItemsItem struct {
	// Data Assistant account ID.
	ID      string `json:"id"`
	Account string `json:"account"`
	// Platform: 0: TikTok 1: YouTube 2: Instagram.
	Channel int    `json:"channel"`
	Remark  string `json:"remark"`
	// Username of the last operator.
	Operator string `json:"operator"`
	// Creation time as a Unix timestamp in seconds.
	CreatedTime int64 `json:"created_time"`
	// Last modification time as a Unix timestamp in seconds.
	UpdatedTime int64 `json:"updated_time"`
	// Account tag list; empty array [] when no tags.
	Tags []ListAccountsResponseItemsItemTagsItem `json:"tags"`
}

type ListAccountsResponse struct {
	// Account data.
	Items []ListAccountsResponseItemsItem `json:"items"`
	// Total count.
	Total int `json:"total"`
	// Current page.
	Page int `json:"page"`
}

// ListAccounts: https://open.geelark.com/api/analytics-accounts-list
func (c *Client) ListAccounts(ctx context.Context, req ListAccountsRequest) (*ListAccountsResponse, error) {
	var resp ListAccountsResponse
	if err := c.core.Do(ctx, "/open/v1/analytics/accounts/list", req, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

type UpdateAccountRequest struct {
	// Account ID. Required.
	ID string `json:"id"`
	// Platform account, up to 64 characters.
	Account string `json:"account,omitempty"`
	// Platform; 0: TikTok 1: YouTube 2: Instagram 4: Reddit.
	Channel *int `json:"channel,omitempty"`
	// Account remark. An empty string omits remark; this type cannot send an explicit empty
	// string. The documentation does not specify how to clear an existing remark.
	Remark string `json:"remark,omitempty"`
	// Data Assistant tag IDs, at most 20 after deduplication. Nil omits the field; an empty
	// non-nil slice sends an empty array. The update page and changelog disagree about
	// omission, so do not assume it preserves existing tags.
	TagIDs []string `json:"tagIds,omitempty"`
}

type UpdateAccountResponse struct{}

// UpdateAccount: https://open.geelark.com/api/analytics-accounts-update
func (c *Client) UpdateAccount(ctx context.Context, req UpdateAccountRequest) (*UpdateAccountResponse, error) {
	var resp UpdateAccountResponse
	if err := c.core.Do(ctx, "/open/v1/analytics/accounts/update", req, nil); err != nil {
		return nil, err
	}
	return &resp, nil
}

type CreateTagRequest struct {
	// Tag name. Maximum 100 characters and must be unique within the current team. Required.
	Name string `json:"name"`
	// Tag color value. The default value is 0.
	Color *int `json:"color,omitempty"`
}

type CreateTagResponse struct{}

// CreateTag: https://open.geelark.com/api/analytics-tags-create
func (c *Client) CreateTag(ctx context.Context, req CreateTagRequest) (*CreateTagResponse, error) {
	var resp CreateTagResponse
	if err := c.core.Do(ctx, "/open/v1/analytics/tags/create", req, nil); err != nil {
		return nil, err
	}
	return &resp, nil
}

type ListTagsRequest struct {
	// Tag name. Fuzzy search is supported.
	Name string `json:"name,omitempty"`
	// Page number. The default value is 1.
	Page int `json:"page,omitempty"`
	// Number of records per page. The default and maximum value is 200.
	PageSize int `json:"pageSize,omitempty"`
}

type ListTagsResponseListItem struct {
	// Tag ID.
	ID string `json:"id"`
	// Tag name.
	Name string `json:"name"`
	// Tag color value.
	Color int `json:"color"`
}

type ListTagsResponse struct {
	// Total number of matching tags.
	Total int `json:"total"`
	// Current page number.
	Page int `json:"page"`
	// Number of records per page.
	PageSize int `json:"pageSize"`
	// Tag list.
	List []ListTagsResponseListItem `json:"list"`
}

// ListTags: https://open.geelark.com/api/analytics-tags-list
func (c *Client) ListTags(ctx context.Context, req ListTagsRequest) (*ListTagsResponse, error) {
	var resp ListTagsResponse
	if err := c.core.Do(ctx, "/open/v1/analytics/tags/list", req, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

type UpdateTagRequest struct {
	// Tag ID. Required.
	ID string `json:"id,omitempty"`
	// Tag name. Maximum 100 characters and must not duplicate another tag under the current
	// team. Required.
	Name string `json:"name,omitempty"`
	// Tag color value. The default value is 0.
	Color *int `json:"color,omitempty"`
}

type UpdateTagResponse struct{}

// UpdateTag: https://open.geelark.com/api/analytics-tags-update
func (c *Client) UpdateTag(ctx context.Context, req UpdateTagRequest) (*UpdateTagResponse, error) {
	var resp UpdateTagResponse
	if err := c.core.Do(ctx, "/open/v1/analytics/tags/update", req, nil); err != nil {
		return nil, err
	}
	return &resp, nil
}

type DeleteTagsRequest struct {
	// List of tag IDs to delete. At least one tag ID is required. Required.
	IDs []string `json:"ids,omitempty"`
}

type DeleteTagsResponse struct{}

// DeleteTags: https://open.geelark.com/api/analytics-tags-del
//
// Deleting a tag also removes its associations with Data Assistant accounts.
func (c *Client) DeleteTags(ctx context.Context, req DeleteTagsRequest) (*DeleteTagsResponse, error) {
	var resp DeleteTagsResponse
	if err := c.core.Do(ctx, "/open/v1/analytics/tags/del", req, nil); err != nil {
		return nil, err
	}
	return &resp, nil
}
