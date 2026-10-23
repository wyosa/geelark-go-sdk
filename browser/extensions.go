package browser

import (
	"context"
)

type ExtensionGroupsRequest struct {
	// Page number. The minimum value is 1. Defaults to 1 if not provided.
	Page int `json:"page,omitempty"`
	// Number of records per page. The minimum value is 1 and the maximum value is 100.
	// Defaults to 10 if not provided.
	PageSize int `json:"pageSize,omitempty"`
}

type ExtensionGroupsResponseItemsItem struct {
	// Browser extension category ID; use it as ExtGroup in profile requests.
	ID string `json:"id"`
	// Browser extension category name.
	Name string `json:"name"`
}

type ExtensionGroupsResponse struct {
	// Total number of matching records.
	Total int `json:"total"`
	// Current page number.
	Page int `json:"page"`
	// Number of records per page.
	PageSize int `json:"pageSize"`
	// Browser extension categories for this page.
	Items []ExtensionGroupsResponseItemsItem `json:"items"`
}

// ExtensionGroups: https://open.geelark.com/api/browser-extgroup-list
//
// This method calls the remote cloud API using the browser client token.
func (c *Client) ExtensionGroups(ctx context.Context, req ExtensionGroupsRequest) (*ExtensionGroupsResponse, error) {
	var resp ExtensionGroupsResponse
	if err := c.remote.Do(ctx, "/open/v1/browser/extGroup/list", req, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
