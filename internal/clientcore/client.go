package clientcore

import (
	"bytes"
	"context"
	"crypto/rand"
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net"
	"net/http"
	"net/url"
	"strings"
	"sync"
	"time"
)

const (
	DefaultBaseURL        = "https://openapi.geelark.com"
	DefaultBrowserBaseURL = "http://localhost:40185"
	DefaultMinInterval    = 310 * time.Millisecond
	maxResponseBytes      = 10 << 20
	maxErrorBodyBytes     = 4 << 10
)

type Config struct {
	BaseURL     string
	Token       string
	AppID       string
	APIKey      string
	HTTPClient  *http.Client
	MinInterval time.Duration
}

type Client struct {
	baseURL      string
	token        string
	appID        string
	apiKey       string
	httpClient   *http.Client
	uploadClient *http.Client
	minInterval  time.Duration
	mu           sync.Mutex
	nextRequest  map[string]time.Time
}

type APIError struct {
	HTTPStatus int
	Code       int
	Message    string
	TraceID    string
	Data       json.RawMessage
	Body       string
}

// DecodeData decodes data returned with a nonzero code, including partial success.
func (e *APIError) DecodeData(out any) error {
	if len(e.Data) == 0 || string(e.Data) == "null" {
		return errors.New("geelark api error has no data")
	}
	return json.Unmarshal(e.Data, out)
}

func (e *APIError) Error() string {
	parts := make([]string, 0, 4)
	if e.HTTPStatus > 0 {
		parts = append(parts, fmt.Sprintf("http_status=%d", e.HTTPStatus))
	}
	if e.Code != 0 {
		parts = append(parts, fmt.Sprintf("code=%d", e.Code))
	}
	if e.TraceID != "" {
		parts = append(parts, "trace_id="+e.TraceID)
	}
	if e.Message != "" {
		parts = append(parts, e.Message)
	} else if len(parts) == 0 && e.Body != "" {
		parts = append(parts, e.Body)
	}
	if len(parts) == 0 {
		return "geelark api error"
	}
	return "geelark api error: " + strings.Join(parts, ", ")
}

type envelope struct {
	TraceID string          `json:"traceId"`
	Code    *int            `json:"code"`
	Message string          `json:"msg"`
	Data    json.RawMessage `json:"data"`
}

func New(cfg Config) (*Client, error) {
	baseURL, err := normalizeBaseURL(cfg.BaseURL)
	if err != nil {
		return nil, err
	}
	token := strings.TrimSpace(cfg.Token)
	appID := strings.TrimSpace(cfg.AppID)
	apiKey := strings.TrimSpace(cfg.APIKey)
	if (token == "" && (appID == "" || apiKey == "")) || (token != "" && (appID != "" || apiKey != "")) {
		return nil, errors.New("provide either a GeeLark token or app ID and API key")
	}
	interval := cfg.MinInterval
	if interval <= 0 {
		interval = DefaultMinInterval
	}
	httpClient := withoutRedirects(cfg.HTTPClient)
	uploadClient := *httpClient
	if cfg.HTTPClient == nil {
		// Large uploads use the caller's context rather than the API's 30s timeout.
		uploadClient.Timeout = 0
	}
	return &Client{
		baseURL:      baseURL,
		token:        token,
		appID:        appID,
		apiKey:       apiKey,
		httpClient:   httpClient,
		uploadClient: &uploadClient,
		minInterval:  interval,
		nextRequest:  make(map[string]time.Time),
	}, nil
}

func (c *Client) Do(ctx context.Context, path string, requestBody, out any) error {
	return c.do(ctx, path, requestBody, out, false)
}

// DoOptionalData decodes data when present for endpoints whose documented
// success response may omit it. Ordinary data-bearing endpoints must use Do.
func (c *Client) DoOptionalData(ctx context.Context, path string, requestBody, out any) error {
	return c.do(ctx, path, requestBody, out, true)
}

func (c *Client) do(ctx context.Context, path string, requestBody, out any, allowMissingData bool) error {
	if ctx == nil {
		ctx = context.Background()
	}
	if !strings.HasPrefix(path, "/open/v1/") && !strings.HasPrefix(path, "/open/v2/") && !strings.HasPrefix(path, "/api/v1/") {
		return fmt.Errorf("invalid API path %q", path)
	}
	if err := ctx.Err(); err != nil {
		return err
	}
	payload := []byte("{}")
	if requestBody != nil {
		var err error
		payload, err = json.Marshal(requestBody)
		if err != nil {
			return fmt.Errorf("marshal request: %w", err)
		}
	}
	req, err := http.NewRequestWithContext(ctx, http.MethodPost, c.baseURL+path, bytes.NewReader(payload))
	if err != nil {
		return fmt.Errorf("build request: %w", err)
	}
	if err := c.waitTurn(ctx, path); err != nil {
		return err
	}
	traceID, err := newTraceID()
	if err != nil {
		return fmt.Errorf("generate trace ID: %w", err)
	}
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("traceId", traceID)
	if c.token != "" {
		req.Header.Set("Authorization", "Bearer "+c.token)
	} else {
		ts := fmt.Sprintf("%d", time.Now().UnixMilli())
		nonce := traceID[:6]
		req.Header.Set("appId", c.appID)
		req.Header.Set("ts", ts)
		req.Header.Set("nonce", nonce)
		req.Header.Set("sign", signature(c.appID, traceID, ts, nonce, c.apiKey))
	}
	resp, err := c.httpClient.Do(req)
	if err != nil {
		return fmt.Errorf("send request: %w", err)
	}
	defer resp.Body.Close()
	body, err := io.ReadAll(io.LimitReader(resp.Body, maxResponseBytes+1))
	if err != nil {
		return fmt.Errorf("read response: %w", err)
	}
	if len(body) > maxResponseBytes {
		return fmt.Errorf("read response: body exceeds %d bytes", maxResponseBytes)
	}
	var apiResp envelope
	decodeErr := json.Unmarshal(body, &apiResp)
	code := 0
	if apiResp.Code != nil {
		code = *apiResp.Code
	}
	if resp.StatusCode < 200 || resp.StatusCode > 299 || (decodeErr == nil && code != 0) {
		return &APIError{
			HTTPStatus: resp.StatusCode,
			Code:       code,
			Message:    apiResp.Message,
			TraceID:    apiResp.TraceID,
			Data:       apiResp.Data,
			Body:       truncate(body),
		}
	}
	if decodeErr != nil {
		return fmt.Errorf("decode response: %w", decodeErr)
	}
	if apiResp.Code == nil {
		return errors.New("decode response: missing or null code")
	}
	if out == nil {
		return nil
	}
	if len(apiResp.Data) == 0 || bytes.Equal(bytes.TrimSpace(apiResp.Data), []byte("null")) {
		if allowMissingData {
			return nil
		}
		return errors.New("decode response: missing or null data")
	}
	if err := json.Unmarshal(apiResp.Data, out); err != nil {
		return fmt.Errorf("decode data: %w", err)
	}
	return nil
}

// Upload sends data to a signed storage URL without GeeLark credentials.
func (c *Client) Upload(ctx context.Context, method, signedURL string, headers map[string]string, body io.Reader, size int64) error {
	if ctx == nil {
		ctx = context.Background()
	}
	if method == "" {
		method = http.MethodPut
	}
	if method != http.MethodPut && method != http.MethodPost {
		return fmt.Errorf("upload: unsupported method %q", method)
	}
	u, err := url.Parse(signedURL)
	if err != nil || u.Host == "" || (u.Scheme != "https" && u.Scheme != "http") || u.User != nil {
		return errors.New("upload: invalid signed URL")
	}
	if size < 0 {
		return errors.New("upload: size must be nonnegative")
	}
	if size > 0 && body == nil {
		return errors.New("upload: body is required for nonzero size")
	}
	if size == 0 {
		// Zero is an exact length, not a request to infer the stream length.
		// Read at most one byte to distinguish a valid empty stream from data
		// that would otherwise be silently discarded. Do not read after cancellation.
		if err := ctx.Err(); err != nil {
			return err
		}
		if body != nil {
			var probe [1]byte
			n, err := io.ReadFull(body, probe[:])
			if n != 0 {
				return errors.New("upload: body exceeds declared size 0")
			}
			if err != io.EOF {
				return fmt.Errorf("upload: read empty body: %w", err)
			}
		}
		body = http.NoBody
	}
	req, err := http.NewRequestWithContext(ctx, method, signedURL, body)
	if err != nil {
		return errors.New("upload: build request failed")
	}
	req.ContentLength = size
	for k, v := range headers {
		if strings.EqualFold(k, "Authorization") || strings.EqualFold(k, "appId") || strings.EqualFold(k, "sign") || strings.EqualFold(k, "traceId") {
			return fmt.Errorf("upload: credential header %q is not allowed", k)
		}
		req.Header.Set(k, v)
	}
	resp, err := c.uploadClient.Do(req)
	if err != nil {
		var urlErr *url.Error
		if errors.As(err, &urlErr) {
			err = urlErr.Err
		}
		return fmt.Errorf("upload: %w", err)
	}
	defer resp.Body.Close()
	respBody, err := io.ReadAll(io.LimitReader(resp.Body, maxResponseBytes+1))
	if err != nil {
		return fmt.Errorf("upload: read response: %w", err)
	}
	if len(respBody) > maxResponseBytes {
		return fmt.Errorf("upload: response body exceeds %d bytes", maxResponseBytes)
	}
	if resp.StatusCode < 200 || resp.StatusCode > 299 {
		return &APIError{HTTPStatus: resp.StatusCode, Body: truncate(respBody)}
	}
	var apiResp envelope
	if json.Unmarshal(respBody, &apiResp) == nil && apiResp.Code != nil && *apiResp.Code != 0 {
		return &APIError{HTTPStatus: resp.StatusCode, Code: *apiResp.Code, Message: apiResp.Message, TraceID: apiResp.TraceID, Data: apiResp.Data, Body: truncate(respBody)}
	}
	return nil
}

func signature(appID, traceID, ts, nonce, apiKey string) string {
	sum := sha256.Sum256([]byte(appID + traceID + ts + nonce + apiKey))
	return strings.ToUpper(hex.EncodeToString(sum[:]))
}

func newTraceID() (string, error) {
	var id [16]byte
	if _, err := rand.Read(id[:]); err != nil {
		return "", err
	}
	id[6] = id[6]&0x0f | 0x40
	id[8] = id[8]&0x3f | 0x80
	return fmt.Sprintf("%08x-%04x-%04x-%04x-%012x", id[:4], id[4:6], id[6:8], id[8:10], id[10:]), nil
}

func (c *Client) waitTurn(ctx context.Context, path string) error {
	interval := c.minInterval
	// Billing endpoints have stricter limits than the API-wide default.
	switch path {
	case "/open/v1/pay/plan/info":
		interval = max(interval, time.Minute)
	case "/open/v1/pay/wallet":
		interval = max(interval, 6*time.Second)
	}
	for {
		if err := ctx.Err(); err != nil {
			return err
		}
		c.mu.Lock()
		now := time.Now()
		delay := c.nextRequest[path].Sub(now)
		if delay <= 0 {
			c.nextRequest[path] = now.Add(interval)
			c.mu.Unlock()
			return nil
		}
		c.mu.Unlock()
		// Reserve a slot only when it is used. Cancelled callers must not
		// leave holes in the queue or delay unrelated future requests.
		timer := time.NewTimer(delay)
		select {
		case <-ctx.Done():
			timer.Stop()
			return ctx.Err()
		case <-timer.C:
		}
	}
}

func normalizeBaseURL(raw string) (string, error) {
	raw = strings.TrimRight(strings.TrimSpace(raw), "/")
	if raw == "" {
		return "", errors.New("base URL is required")
	}
	u, err := url.Parse(raw)
	if err != nil {
		return "", fmt.Errorf("invalid base URL: %w", err)
	}
	if u.User != nil || u.Hostname() == "" || strings.ContainsAny(raw, "?#") {
		return "", errors.New("invalid base URL: expected scheme://host[/path] without credentials, query or fragment")
	}
	if u.Scheme != "https" && !(u.Scheme == "http" && isLoopback(u.Hostname())) {
		return "", errors.New("invalid base URL: HTTPS is required except for loopback hosts")
	}
	return raw, nil
}

func isLoopback(host string) bool {
	if host == "localhost" {
		return true
	}
	ip := net.ParseIP(host)
	return ip != nil && ip.IsLoopback()
}

func withoutRedirects(hc *http.Client) *http.Client {
	var c http.Client
	if hc != nil {
		c = *hc
	} else {
		c.Timeout = 30 * time.Second
	}
	c.CheckRedirect = func(_ *http.Request, _ []*http.Request) error {
		return http.ErrUseLastResponse
	}
	return &c
}

func truncate(body []byte) string {
	if len(body) > maxErrorBodyBytes {
		body = body[:maxErrorBodyBytes]
	}
	return string(body)
}
