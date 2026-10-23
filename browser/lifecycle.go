package browser

import (
	"context"
)

type StartRequest struct {
	// browser id. Required.
	ID string `json:"id"`
	// Callback URL. Notification will be sent after the browser finishes starting.
	Webhook string `json:"webhook,omitempty"`
}

type StartResponse struct {
	// Browser debugging port returned by the desktop client.
	DebugPort int `json:"debugPort"`
}

// Start: https://open.geelark.com/api/launch-browser
func (c *Client) Start(ctx context.Context, req StartRequest) (*StartResponse, error) {
	var resp StartResponse
	if err := c.core.Do(ctx, "/api/v1/browser/start", req, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

type StopRequest struct {
	// browser id. Required.
	ID string `json:"id"`
}

type StopResponse struct{}

// Stop: https://open.geelark.com/api/close-browser
func (c *Client) Stop(ctx context.Context, req StopRequest) (*StopResponse, error) {
	var resp StopResponse
	if err := c.core.Do(ctx, "/api/v1/browser/stop", req, nil); err != nil {
		return nil, err
	}
	return &resp, nil
}

// APIStatusRequest contains no parameters for [Client.APIStatus].
type APIStatusRequest struct{}

type APIStatusResponse struct{}

// APIStatus: https://open.geelark.com/api/check-api-interface-status
func (c *Client) APIStatus(ctx context.Context, req APIStatusRequest) (*APIStatusResponse, error) {
	var resp APIStatusResponse
	if err := c.core.Do(ctx, "/api/v1/status", req, nil); err != nil {
		return nil, err
	}
	return &resp, nil
}

type StatusRequest struct {
	// Browser ID. Required.
	ID string `json:"id"`
}

type StatusResponse struct {
	// Browser status string returned by the desktop client. The documentation does not
	// enumerate all possible values.
	Status string `json:"status"`
	// Browser debugging port returned by the desktop client.
	DebugPort int `json:"debugPort"`
}

// Status: https://open.geelark.com/api/check-browser-status
func (c *Client) Status(ctx context.Context, req StatusRequest) (*StatusResponse, error) {
	var resp StatusResponse
	if err := c.core.Do(ctx, "/api/v1/browser/status", req, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
