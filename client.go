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
	"github.com/wyosa/geelark-go-sdk/proxy"
	"github.com/wyosa/geelark-go-sdk/proxygroup"
	"github.com/wyosa/geelark-go-sdk/shell"
	"github.com/wyosa/geelark-go-sdk/tag"
	"github.com/wyosa/geelark-go-sdk/webhook"
)

const (
	DefaultBaseURL        = clientcore.DefaultBaseURL
	DefaultBrowserBaseURL = clientcore.DefaultBrowserBaseURL
	DefaultMinInterval    = clientcore.DefaultMinInterval
)

type APIError = clientcore.APIError

type config struct {
	baseURL       string
	remoteBaseURL string
	httpClient    *http.Client
	minInterval   time.Duration
}

type Option func(*config)

// Client provides the cloud and platform APIs through module services.
type Client struct {
	CloudPhones *cloudphone.Client
	ADB         *adb.Client
	Shell       *shell.Client
	Apps        *app.Client
	Files       *files.Client
	Library     *library.Client
	Automation  *automation.Client
	Analytics   *analytics.Client
	LiveStream  *livestream.Client
	Webhooks    *webhook.Client
	OEM         *oem.Client
	Proxies     *proxy.Client
	ProxyGroups *proxygroup.Client
	Groups      *group.Client
	Tags        *tag.Client
	Billing     *billing.Client

	core *clientcore.Client
}

// BrowserClient provides APIs served by the local GeeLark desktop client.
type BrowserClient struct {
	Browsers   *browser.Client
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
func WithHTTPClient(httpClient *http.Client) Option {
	return func(cfg *config) {
		if httpClient != nil {
			cfg.httpClient = httpClient
		}
	}
}

// WithMinInterval changes the minimum spacing between calls to the same path.
func WithMinInterval(interval time.Duration) Option {
	return func(cfg *config) {
		if interval > 0 {
			cfg.minInterval = interval
		}
	}
}
