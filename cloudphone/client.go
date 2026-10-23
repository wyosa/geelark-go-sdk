// Package cloudphone provides cloud phone lifecycle, configuration, device information, and
// asynchronous operations.
package cloudphone

import (
	"github.com/wyosa/geelark-go-sdk/common"
)

type Client struct {
	core common.OptionalDataTransport
}

func New(core common.OptionalDataTransport) *Client { return &Client{core: core} }
