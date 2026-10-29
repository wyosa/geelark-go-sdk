// Package geelark provides clients for the GeeLark cloud and local browser APIs.
package geelark

import (
	"errors"
	"net/http"
	"strings"
	"time"

	"github.com/wyosa/geelark-go-sdk/adb"
	"github.com/wyosa/geelark-go-sdk/analytics"
	"github.com/wyosa/geelark-go-sdk/app"
	"github.com/wyosa/geelark-go-sdk/automation"
	"github.com/wyosa/geelark-go-sdk/billing"
	"github.com/wyosa/geelark-go-sdk/browser"
	"github.com/wyosa/geelark-go-sdk/browserautomation"
	"github.com/wyosa/geelark-go-sdk/cloudphone"
	"github.com/wyosa/geelark-go-sdk/files"
	"github.com/wyosa/geelark-go-sdk/group"
	"github.com/wyosa/geelark-go-sdk/internal/clientcore"
	"github.com/wyosa/geelark-go-sdk/library"
	"github.com/wyosa/geelark-go-sdk/livestream"
	"github.com/wyosa/geelark-go-sdk/oem"
	"github.com/wyosa/geelark-go-sdk/patch"
	"github.com/wyosa/geelark-go-sdk/proxy"
	"github.com/wyosa/geelark-go-sdk/proxygroup"
	"github.com/wyosa/geelark-go-sdk/shell"
	"github.com/wyosa/geelark-go-sdk/tag"
	"github.com/wyosa/geelark-go-sdk/webhook"
)

const (
	// DefaultBaseURL is the remote GeeLark API origin.
	DefaultBaseURL = clientcore.DefaultBaseURL
	// DefaultBrowserBaseURL is the local GeeLark desktop API origin.
	DefaultBrowserBaseURL = clientcore.DefaultBrowserBaseURL
	// DefaultMinInterval spaces calls to the same path; Billing has higher minima.
	DefaultMinInterval = clientcore.DefaultMinInterval
)

// APIError describes a non-2xx HTTP response or a nonzero GeeLark result code.
// Use errors.As to inspect it and DecodeData to read partial-success results.
type APIError = clientcore.APIError

type config struct {
	baseURL       string
	remoteBaseURL string
	httpClient    *http.Client
	minInterval   time.Duration
}

// Option configures a cloud or browser client through the With functions.
type Option func(*config)

// Client provides the cloud and platform APIs through module services.
type Client struct {
	// CloudPhones manages cloud phone lifecycle and device configuration.
	CloudPhones *cloudphone.Client
	// ADB controls ADB access to cloud phones.
	ADB *adb.Client
	// Shell executes shell commands on cloud phones.
	Shell *shell.Client
	// Apps manages applications on cloud phones.
	Apps *app.Client
	// Files uploads temporary files and transfers files to cloud phones.
	Files *files.Client
	// Library manages retained Library files and their tags.
	Library *library.Client
	// Patches installs and removes patch modules on cloud phones.
	Patches *patch.Client
	// Automation creates and queries automation tasks.
	Automation *automation.Client
	// Analytics manages Data Assistant accounts, tags, and metrics.
	Analytics *analytics.Client
	// LiveStream controls cloud phone live streams.
	LiveStream *livestream.Client
	// Webhooks configures cloud callback delivery.
	Webhooks *webhook.Client
	// OEM customizes cloud phone views and branding.
	OEM *oem.Client
	// Proxies manages proxy connections.
	Proxies *proxy.Client
	// ProxyGroups manages proxy groups.
	ProxyGroups *proxygroup.Client
	// Groups manages cloud phone profile groups.
	Groups *group.Client
	// Tags manages cloud phone profile tags.
	Tags *tag.Client
	// Billing reads balances and manages subscriptions and time add-ons.
	Billing *billing.Client

	core *clientcore.Client
}

// BrowserClient provides APIs served by the local GeeLark desktop client.
type BrowserClient struct {
	// Browsers manages local browser profiles and browser settings.
	Browsers *browser.Client
	// Automation creates and queries automation tasks.
	Automation *browserautomation.Client

	core       *clientcore.Client
	remoteCore *clientcore.Client
}

// NewClient uses a GeeLark Bearer token for the cloud APIs.
func NewClient(token string, opts ...Option) (*Client, error) {
	if strings.TrimSpace(token) == "" {
		return nil, errors.New("geelark token is required")
	}
	cfg := applyOptions(DefaultBaseURL, opts)
	core, err := clientcore.New(clientcore.Config{
		BaseURL:     cfg.baseURL,
		Token:       token,
		HTTPClient:  cfg.httpClient,
		MinInterval: cfg.minInterval,
	})
	if err != nil {
		return nil, err
	}
	return newCloudClient(core), nil
}

// NewKeyClient uses the app ID and API key signature mode for cloud APIs.
func NewKeyClient(appID, apiKey string, opts ...Option) (*Client, error) {
	if strings.TrimSpace(appID) == "" || strings.TrimSpace(apiKey) == "" {
		return nil, errors.New("geelark app ID and API key are required")
	}
	cfg := applyOptions(DefaultBaseURL, opts)
	core, err := clientcore.New(clientcore.Config{
		BaseURL:     cfg.baseURL,
		AppID:       appID,
		APIKey:      apiKey,
		HTTPClient:  cfg.httpClient,
		MinInterval: cfg.minInterval,
	})
	if err != nil {
		return nil, err
	}
	return newCloudClient(core), nil
}

// NewBrowserClient uses a Bearer token and a running local GeeLark client.
// ExtensionGroups is a remote API and uses the same token with remoteBaseURL.
func NewBrowserClient(token string, opts ...Option) (*BrowserClient, error) {
	if strings.TrimSpace(token) == "" {
		return nil, errors.New("geelark token is required")
	}
	cfg := applyOptions(DefaultBrowserBaseURL, opts)
	local, err := clientcore.New(clientcore.Config{
		BaseURL:     cfg.baseURL,
		Token:       token,
		HTTPClient:  cfg.httpClient,
		MinInterval: cfg.minInterval,
	})
	if err != nil {
		return nil, err
	}
	remote, err := clientcore.New(clientcore.Config{
		BaseURL:     cfg.remoteBaseURL,
		Token:       token,
		HTTPClient:  cfg.httpClient,
		MinInterval: cfg.minInterval,
	})
	if err != nil {
		return nil, err
	}
	return &BrowserClient{
		Browsers:   browser.New(local, remote),
		Automation: browserautomation.New(local),
		core:       local,
		remoteCore: remote,
	}, nil
}

func newCloudClient(core *clientcore.Client) *Client {
	return &Client{
		CloudPhones: cloudphone.New(core),
		ADB:         adb.New(core),
		Shell:       shell.New(core),
		Apps:        app.New(core),
		Files:       files.New(core),
		Library:     library.New(core),
		Patches:     patch.New(core),
		Automation:  automation.New(core),
		Analytics:   analytics.New(core),
		LiveStream:  livestream.New(core),
		Webhooks:    webhook.New(core),
		OEM:         oem.New(core),
		Proxies:     proxy.New(core),
		ProxyGroups: proxygroup.New(core),
		Groups:      group.New(core),
		Tags:        tag.New(core),
		Billing:     billing.New(core),
		core:        core,
	}
}

func applyOptions(baseURL string, opts []Option) config {
	cfg := config{
		baseURL:       baseURL,
		remoteBaseURL: DefaultBaseURL,
		minInterval:   DefaultMinInterval,
	}
	for _, opt := range opts {
		if opt != nil {
			opt(&cfg)
		}
	}
	return cfg
}

// WithBaseURL changes the cloud or local API base URL for the chosen constructor.
// URLs require HTTPS except on loopback hosts. Credentials, query strings, and
// fragments are rejected. Do not append /open/v1 or /api/v1; methods add the API
// path. Empty values are ignored. NewBrowserClient's remote endpoint is configured
// separately with WithRemoteBaseURL.
func WithBaseURL(baseURL string) Option {
	return func(cfg *config) {
		if strings.TrimSpace(baseURL) != "" {
			cfg.baseURL = strings.TrimRight(baseURL, "/")
		}
	}
}

// WithRemoteBaseURL changes the host used by remote browser APIs.
func WithRemoteBaseURL(baseURL string) Option {
	return func(cfg *config) {
		if strings.TrimSpace(baseURL) != "" {
			cfg.remoteBaseURL = strings.TrimRight(baseURL, "/")
		}
	}
}

// WithHTTPClient copies the supplied client and refuses redirects.
// Its timeout and transport apply to API requests and storage uploads. The
// transport and cookie jar are shared with the supplied client and must be safe
// for concurrent use. A zero timeout remains unlimited; nil is ignored.
func WithHTTPClient(httpClient *http.Client) Option {
	return func(cfg *config) {
		if httpClient != nil {
			cfg.httpClient = httpClient
		}
	}
}

// WithMinInterval changes the minimum spacing between calls to the same path.
// Billing.PlanInfo always waits at least 62 seconds between calls and
// Billing.Balance at least 6.2 seconds, slightly above the documented 1 and 10
// requests per minute; a longer interval is honored.
// Nonpositive values are ignored. This setting applies to one client instance;
// callers coordinating several clients or processes need their own limiter.
func WithMinInterval(interval time.Duration) Option {
	return func(cfg *config) {
		if interval > 0 {
			cfg.minInterval = interval
		}
	}
}
