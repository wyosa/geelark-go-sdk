// Package oem provides OEM and white-label customization of cloud phone views.
package oem

import (
	"context"

	"github.com/wyosa/geelark-go-sdk/common"
)

type Client struct {
	core common.Transport
}

func New(core common.Transport) *Client { return &Client{core: core} }

type CustomizationRequestToolBarSettingsItem struct {
	// Toolbar item name. Refer to the item descriptions below.
	ToolBar string `json:"toolBar,omitempty"`
	// Whether to display the item. If omitted or set to false, the item will not be displayed.
	Visible *bool `json:"visible,omitempty"`
	// Icon URL, maximum length 255 bytes; if not provided or left empty, the original value
	// will remain unchanged; supported formats: SVG, PNG, JPG (SVG is recommended).
	// Recommended size: 16×16.
	IconURL string `json:"iconUrl,omitempty"`
}

type CustomizationRequest struct {
	// Title, maximum length 64 bytes. If not provided or left empty, the original value
	// remains unchanged.
	Title string `json:"title,omitempty"`
	// Logo URL, maximum length 255 bytes. If not provided or left empty, the original value
	// remains unchanged.
	Logo string `json:"logo,omitempty"`
	// Whether to hide the header at the top of the cloud phone. Defaults to false if not
	// provided.
	HideHeader *bool `json:"hideHeader,omitempty"`
	// The QR code and the url opened on the phone's browser in the "Mirror" entrance, limited
	// to 255 characters. if set this value, GeeLark will display the url in the "Mirror"
	// entrance (as shown picture blow).!demo
	// (https://singapore-upgrade.geelark.cn/en_mirror_url_demo.png) When open this url, please
	// redirect to https://mobile.geelark.com/mobile.html with all parameters (it is
	// recommended to use a iframe) ).
	MirrorURL string `json:"mirrorUrl,omitempty"`
	// Control whether the toolbar entrance on the side of the cloud phone should displayed. If
	// not set, all will be displayed by default.The display order of the menu matches the
	// order in which the parameters are passed. The positions of Network quality and App Store
	// indicator are fixed by default and cannot be modified.
	ToolBarSettings []CustomizationRequestToolBarSettingsItem `json:"toolBarSettings,omitempty"`
	// Whether to hide the sidebar. Defaults to false if not specified.
	HideSidebar *bool `json:"hideSidebar,omitempty"`
	// Custom domain name. Maximum length: 255 characters. Do not include the http:// or
	// https:// prefix. Example: phone.example.com. For instructions on configuring a custom
	// domain, please refer to OEM Custom Domain
	// (https://open.geelark.com/api/OEM-custom-domain).
	CustomDomain string `json:"customDomain,omitempty"`
}

type CustomizationResponse struct{}

// Customization: https://open.geelark.com/api/phone-customization
//
// Requires a GeeLark plan with at least 50 profiles. Set customization before creating and
// starting the cloud phones that use it.
func (c *Client) Customization(ctx context.Context, req CustomizationRequest) (*CustomizationResponse, error) {
	var resp CustomizationResponse
	if err := c.core.Do(ctx, "/open/v1/phone/customization", req, nil); err != nil {
		return nil, err
	}
	return &resp, nil
}
