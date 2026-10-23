package shell

import (
	"context"

	"github.com/wyosa/geelark-go-sdk/internal/clientcore"
)

type Client struct {
	core *clientcore.Client
}

func New(core *clientcore.Client) *Client { return &Client{core: core} }

type ExecuteRequest struct {
	ID  string `json:"id"`
	Cmd string `json:"cmd"`
}

type ExecuteResponse struct {
	Status bool   `json:"status"`
	Output string `json:"output"`
}

// Execute calls /open/v1/shell/execute.
func (c *Client) Execute(ctx context.Context, req ExecuteRequest) (*ExecuteResponse, error) {
	var resp ExecuteResponse
	if err := c.core.Do(ctx, "/open/v1/shell/execute", req, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
