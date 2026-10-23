// Package shell provides shell command execution on cloud phones.
package shell

import (
	"context"

	"github.com/wyosa/geelark-go-sdk/common"
)

type Client struct {
	core common.Transport
}

func New(core common.Transport) *Client { return &Client{core: core} }

type ExecuteRequest struct {
	// Cloud phone ID. Required.
	ID string `json:"id"`
	// Command to execute. Required.
	Cmd string `json:"cmd"`
}

type ExecuteResponse struct {
	// true: Execution successful, false: Execution failed.
	Status bool `json:"status"`
	// Execution result.
	Output string `json:"output"`
}

// Execute: https://open.geelark.com/api/shell-execute
func (c *Client) Execute(ctx context.Context, req ExecuteRequest) (*ExecuteResponse, error) {
	var resp ExecuteResponse
	if err := c.core.Do(ctx, "/open/v1/shell/execute", req, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
