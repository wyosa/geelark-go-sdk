// Package proxy provides proxy registration, updates, lookup, and connection checks.
package proxy

import (
	"context"

	"github.com/wyosa/geelark-go-sdk/common"
)

type Client struct {
	core common.Transport
}

func New(core common.Transport) *Client { return &Client{core: core} }

type AddRequest struct {
	// The list of proxy information items can contain up to 100 entries. Required.
	List []AddItem `json:"list"`
}

type AddItem struct {
	// Proxy types，socks5，http，https. Required.
	Scheme string `json:"scheme"`
	// Proxy address. Required.
	Server string `json:"server"`
	// Proxy port. Required.
	Port int `json:"port"`
	// Proxy username.
	Username string `json:"username,omitempty"`
	// Proxy password.
	Password string `json:"password,omitempty"`
	// Detection channels: 1. IPApi; 2. IP2Location; default is IP2Location.
	ProxyQueryChannel *int `json:"proxyQueryChannel,omitempty"`
	// Proxy group ID. This field is described in the changelog but omitted from some endpoint
	// parameter tables.
	ProxyGroupID string `json:"proxyGroupId,omitempty"`
}

type AddResponseFailDetailsItem struct {
	// Index of the provided proxy information items.
	Index int `json:"index"`
	// Error code, refer to failure code and message.
	Code int `json:"code"`
	// Error message, refer to failure code and message.
	Msg string `json:"msg"`
}

type AddResponseSuccessDetailsItem struct {
	// Index of the provided proxy information items.
	Index int `json:"index"`
	// Proxy ID; if the provided proxy information items are the same, the proxy ID will be the
	// same.
	ID string `json:"id"`
}

type AddResponse struct {
	// Total processed, the same as the number of proxy information items provided.
	TotalAmount int `json:"totalAmount"`
	// Number of successful additions.
	SuccessAmount int `json:"successAmount"`
	// Number of failed additions.
	FailAmount int `json:"failAmount"`
	// Failed proxy information.
	FailDetails []AddResponseFailDetailsItem `json:"failDetails"`
	// Successful proxy information.
	SuccessDetails []AddResponseSuccessDetailsItem `json:"successDetails"`
}

// Add: https://open.geelark.com/api/proxy-add
func (c *Client) Add(ctx context.Context, req AddRequest) (*AddResponse, error) {
	var resp AddResponse
	if err := c.core.Do(ctx, "/open/v1/proxy/add", req, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

type UpdateRequest struct {
	// The list of proxy information items can contain up to 100 entries. Required.
	List []UpdateItem `json:"list"`
}

type UpdateItem struct {
	// Proxy ID. Required.
	ID string `json:"id"`
	// Proxy types，socks5，http，https. Required.
	Scheme string `json:"scheme"`
	// Proxy address. Required.
	Server string `json:"server"`
	// Proxy port. Required.
	Port int `json:"port"`
	// Proxy username.
	Username *string `json:"username,omitempty"`
	// Proxy password.
	Password *string `json:"password,omitempty"`
	// Detection channels: 1 IPApi; 2 IP2Location; default to original detection channel.
	ProxyQueryChannel *int `json:"proxyQueryChannel,omitempty"`
	// Proxy group ID. This field is described in the changelog but omitted from some endpoint
	// parameter tables.
	ProxyGroupID *string `json:"proxyGroupId,omitempty"`
}

type UpdateResponseFailDetailsItem struct {
	// Proxy ID.
	ID string `json:"id"`
	// Error code, refer to failure code and message.
	Code int `json:"code"`
	// Error message, refer to failure code and message.
	Msg string `json:"msg"`
}

type UpdateResponse struct {
	// Total processed, duplicate IDs provided will not be counted.
	TotalAmount int `json:"totalAmount"`
	// Number of successfully processed items.
	SuccessAmount int `json:"successAmount"`
	// Number of failed items.
	FailAmount int `json:"failAmount"`
	// Failed proxy information.
	FailDetails []UpdateResponseFailDetailsItem `json:"failDetails"`
}

// Update: https://open.geelark.com/api/proxy-update
func (c *Client) Update(ctx context.Context, req UpdateRequest) (*UpdateResponse, error) {
	var resp UpdateResponse
	if err := c.core.Do(ctx, "/open/v1/proxy/update", req, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

type DeleteRequest struct {
	// Proxy ID list, up to 100 IDs. Required.
	IDs []string `json:"ids"`
}

type DeleteResponseFailDetailsItem struct {
	// Proxy ID.
	ID string `json:"id"`
	// Error code, refer to failure code and message.
	Code int `json:"code"`
	// Error message, refer to failure code and message.
	Msg string `json:"msg"`
}

type DeleteResponse struct {
	// Total processed, duplicate IDs provided will not be counted.
	TotalAmount int `json:"totalAmount"`
	// Number of successfully processed items.
	SuccessAmount int `json:"successAmount"`
	// Number of failed items.
	FailAmount int `json:"failAmount"`
	// Failed proxy information.
	FailDetails []DeleteResponseFailDetailsItem `json:"failDetails"`
}

// Delete: https://open.geelark.com/api/proxy-delete
func (c *Client) Delete(ctx context.Context, req DeleteRequest) (*DeleteResponse, error) {
	var resp DeleteResponse
	if err := c.core.Do(ctx, "/open/v1/proxy/delete", req, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

type ListRequest struct {
	// Page number, minimum value is 1. Required.
	Page int `json:"page"`
	// Number of items per page, minimum is 1, maximum is 100. Required.
	PageSize int `json:"pageSize"`
	// Proxy ID list.
	IDs []string `json:"ids,omitempty"`
	// Proxy group ID. This field is described in the changelog but omitted from some endpoint
	// parameter tables.
	ProxyGroupID string `json:"proxyGroupId,omitempty"`
}

type ListResponseListItem struct {
	// Proxy ID.
	ID string `json:"id"`
	// Proxy serial number.
	SerialNo int `json:"serialNo"`
	// Proxy types，socks5，http，https.
	Scheme string `json:"scheme"`
	// Proxy address.
	Server string `json:"server"`
	// Proxy port.
	Port int `json:"port"`
	// Proxy username.
	Username string `json:"username"`
	// Proxy password.
	Password string `json:"password"`
	// Proxy group ID. This field is described in the changelog but omitted from some endpoint
	// parameter tables.
	ProxyGroupID string `json:"proxyGroupId,omitempty"`
}

type ListResponse struct {
	// Total number of items.
	Total int `json:"total"`
	// Page number.
	Page int `json:"page"`
	// Number of items per page.
	PageSize int `json:"pageSize"`
	// The list of proxy information items.
	List []ListResponseListItem `json:"list"`
}

// List: https://open.geelark.com/api/proxy-list
func (c *Client) List(ctx context.Context, req ListRequest) (*ListResponse, error) {
	var resp ListResponse
	if err := c.core.Do(ctx, "/open/v1/proxy/list", req, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

type CheckRequest struct {
	// IP lookup source, supports only IP-API or IP2Location. Required.
	ProxyQueryChannel string `json:"proxyQueryChannel"`
	// Proxy type, supports only socks5, http, or https. Required.
	ProxyType string `json:"proxyType"`
	// Host. Required.
	Server string `json:"server"`
	// Port number. Required.
	Port int `json:"port"`
	// Proxy username.
	Username string `json:"username,omitempty"`
	// Proxy password.
	Password string `json:"password,omitempty"`
}

type CheckResponse struct {
	// Whether the detection was successful.
	DetectStatus bool `json:"detectStatus"`
	// Reason for failure (if any).
	Message    string `json:"message"`
	OutboundIP string `json:"outboundIP"`
	// Country code of the outbound IP.
	CountryCode string `json:"countryCode"`
	// Country name of the outbound IP.
	CountryName string `json:"countryName"`
	// State/Province of the outbound IP.
	Subdivision string `json:"subdivision"`
	// City of the outbound IP.
	City string `json:"city"`
	// Time zone of the outbound IP.
	Timezone string `json:"timezone"`
	// ISP of the outbound IP.
	Isp string `json:"isp"`
}

// Check: https://open.geelark.com/api/proxy-check
func (c *Client) Check(ctx context.Context, req CheckRequest) (*CheckResponse, error) {
	var resp CheckResponse
	if err := c.core.Do(ctx, "/open/v1/proxy/check", req, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
