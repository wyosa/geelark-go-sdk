package proxy

import (
	"context"

	"github.com/wyosa/geelark-go-sdk/internal/clientcore"
)

type Client struct {
	core *clientcore.Client
}

func New(core *clientcore.Client) *Client { return &Client{core: core} }

type AddRequest struct {
	List []AddItem `json:"list"`
}

type AddItem struct {
	Scheme            string `json:"scheme"`
	Server            string `json:"server"`
	Port              int    `json:"port"`
	Username          string `json:"username,omitempty"`
	Password          string `json:"password,omitempty"`
	ProxyQueryChannel *int   `json:"proxyQueryChannel,omitempty"`
	ProxyGroupID      string `json:"proxyGroupId,omitempty"`
}

type AddResponseFailDetailsItem struct {
	Index int    `json:"index"`
	Code  int    `json:"code"`
	Msg   string `json:"msg"`
}

type AddResponseSuccessDetailsItem struct {
	Index int    `json:"index"`
	ID    string `json:"id"`
}

type AddResponse struct {
	TotalAmount    int                             `json:"totalAmount"`
	SuccessAmount  int                             `json:"successAmount"`
	FailAmount     int                             `json:"failAmount"`
	FailDetails    []AddResponseFailDetailsItem    `json:"failDetails"`
	SuccessDetails []AddResponseSuccessDetailsItem `json:"successDetails"`
}

// Add calls /open/v1/proxy/add.
func (c *Client) Add(ctx context.Context, req AddRequest) (*AddResponse, error) {
	var resp AddResponse
	if err := c.core.Do(ctx, "/open/v1/proxy/add", req, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

type UpdateRequest struct {
	List []UpdateItem `json:"list"`
}

type UpdateItem struct {
	ID                string  `json:"id"`
	Scheme            string  `json:"scheme"`
	Server            string  `json:"server"`
	Port              int     `json:"port"`
	Username          *string `json:"username,omitempty"`
	Password          *string `json:"password,omitempty"`
	ProxyQueryChannel *int    `json:"proxyQueryChannel,omitempty"`
	ProxyGroupID      *string `json:"proxyGroupId,omitempty"`
}

type UpdateResponseFailDetailsItem struct {
	ID   string `json:"id"`
	Code int    `json:"code"`
	Msg  string `json:"msg"`
}

type UpdateResponse struct {
	TotalAmount   int                             `json:"totalAmount"`
	SuccessAmount int                             `json:"successAmount"`
	FailAmount    int                             `json:"failAmount"`
	FailDetails   []UpdateResponseFailDetailsItem `json:"failDetails"`
}

// Update calls /open/v1/proxy/update.
func (c *Client) Update(ctx context.Context, req UpdateRequest) (*UpdateResponse, error) {
	var resp UpdateResponse
	if err := c.core.Do(ctx, "/open/v1/proxy/update", req, &resp); err != nil {
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

// Delete calls /open/v1/proxy/delete.
func (c *Client) Delete(ctx context.Context, req DeleteRequest) (*DeleteResponse, error) {
	var resp DeleteResponse
	if err := c.core.Do(ctx, "/open/v1/proxy/delete", req, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

type ListRequest struct {
	Page         int      `json:"page"`
	PageSize     int      `json:"pageSize"`
	IDs          []string `json:"ids,omitempty"`
	ProxyGroupID string   `json:"proxyGroupId,omitempty"`
}

type ListResponseListItem struct {
	ID           string `json:"id"`
	SerialNo     int    `json:"serialNo"`
	Scheme       string `json:"scheme"`
	Server       string `json:"server"`
	Port         int    `json:"port"`
	Username     string `json:"username"`
	Password     string `json:"password"`
	ProxyGroupID string `json:"proxyGroupId,omitempty"`
}

type ListResponse struct {
	Total    int                    `json:"total"`
	Page     int                    `json:"page"`
	PageSize int                    `json:"pageSize"`
	List     []ListResponseListItem `json:"list"`
}

// List calls /open/v1/proxy/list.
func (c *Client) List(ctx context.Context, req ListRequest) (*ListResponse, error) {
	var resp ListResponse
	if err := c.core.Do(ctx, "/open/v1/proxy/list", req, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

type CheckRequest struct {
	ProxyQueryChannel string `json:"proxyQueryChannel"`
	ProxyType         string `json:"proxyType"`
	Server            string `json:"server"`
	Port              int    `json:"port"`
	Username          string `json:"username,omitempty"`
	Password          string `json:"password,omitempty"`
}

type CheckResponse struct {
	DetectStatus bool   `json:"detectStatus"`
	Message      string `json:"message"`
	OutboundIP   string `json:"outboundIP"`
	CountryCode  string `json:"countryCode"`
	CountryName  string `json:"countryName"`
	Subdivision  string `json:"subdivision"`
	City         string `json:"city"`
	Timezone     string `json:"timezone"`
	Isp          string `json:"isp"`
}

// Check calls /open/v1/proxy/check.
func (c *Client) Check(ctx context.Context, req CheckRequest) (*CheckResponse, error) {
	var resp CheckResponse
	if err := c.core.Do(ctx, "/open/v1/proxy/check", req, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
