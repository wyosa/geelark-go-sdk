package oem

import (
	"context"

	"github.com/wyosa/geelark-go-sdk/internal/clientcore"
)

type Client struct {
	core *clientcore.Client
}

func New(core *clientcore.Client) *Client { return &Client{core: core} }

type CustomizationRequestToolBarSettingsItem struct {
	ToolBar string `json:"toolBar,omitempty"`
	Visible *bool  `json:"visible,omitempty"`
	IconURL string `json:"iconUrl,omitempty"`
}

type CustomizationRequest struct {
	Title           string                                    `json:"title,omitempty"`
	Logo            string                                    `json:"logo,omitempty"`
	HideHeader      *bool                                     `json:"hideHeader,omitempty"`
	MirrorURL       string                                    `json:"mirrorUrl,omitempty"`
	ToolBarSettings []CustomizationRequestToolBarSettingsItem `json:"toolBarSettings,omitempty"`
	HideSidebar     *bool                                     `json:"hideSidebar,omitempty"`
	CustomDomain    string                                    `json:"customDomain,omitempty"`
}

type CustomizationResponse struct{}

// Customization calls /open/v1/phone/customization.
func (c *Client) Customization(ctx context.Context, req CustomizationRequest) (*CustomizationResponse, error) {
	var resp CustomizationResponse
	if err := c.core.Do(ctx, "/open/v1/phone/customization", req, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
