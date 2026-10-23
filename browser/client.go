// Package browser provides local GeeLark browser profiles, kernels, cookies, and bookmarks.
package browser

import (
	"github.com/wyosa/geelark-go-sdk/common"
)

type Client struct {
	core   common.Transport
	remote common.Transport
}

func New(core, remote common.Transport) *Client { return &Client{core: core, remote: remote} }
