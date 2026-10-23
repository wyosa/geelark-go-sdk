package browser

import (
	"context"
	"encoding/json"
)

type ClearCacheRequest struct {
	// Browser ID. Required.
	IDs []string `json:"ids"`
	// Types of cache to clear. Supported types: local_storage, indexeddb, extension_cache,
	// cookie, history, image_file. Required.
	Type []string `json:"type"`
}

type ClearCacheResponse struct {
	// Number of successful operations.
	SuccessCount int `json:"success_count"`
	// Number of failed operations.
	ErrorCount int `json:"error_count"`
	// Per-profile cache clearing errors. Kept as raw JSON because the documentation does not
	// define a stable item schema.
	ErrorInfo []json.RawMessage `json:"error_info"`
}

// ClearCache: https://open.geelark.com/api/delete-browser-cache
//
// Close the browser profile before clearing its cache.
func (c *Client) ClearCache(ctx context.Context, req ClearCacheRequest) (*ClearCacheResponse, error) {
	var resp ClearCacheResponse
	if err := c.core.Do(ctx, "/api/v1/browser/deleteCache", req, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

type CookiesRequest struct {
	// Browser ID. Required.
	ID string `json:"id"`
}

type CookiesResponse struct {
	// Cookies exported by the browser profile, encoded as a string.
	Cookies string `json:"cookies"`
}

// Cookies: https://open.geelark.com/api/browser-get-cookies
func (c *Client) Cookies(ctx context.Context, req CookiesRequest) (*CookiesResponse, error) {
	var resp CookiesResponse
	if err := c.core.Do(ctx, "/api/v1/browser/getCookie", req, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

type SetBookmarksRequestBrowserBookmark struct {
	// Bookmark type, 0-Not set, 1-Upload file, 2-Manually create. Required.
	Type *int `json:"type,omitempty"`
	// Bookmark HTML file address.
	FileAddr *string `json:"fileAddr,omitempty"`
	// Manually create bookmark content, multiple bookmarks are separated by newline characters
	// \n, supported formats: Folder::Name::URL Name::URL URL.
	Text string `json:"text,omitempty"`
}

type SetBookmarksRequest struct {
	// Browser Bookmark. Required.
	BrowserBookmark SetBookmarksRequestBrowserBookmark `json:"browserBookmark"`
}

type SetBookmarksResponse struct{}

// SetBookmarks: https://open.geelark.com/api/browser-setBookmark
//
// Settings are applied automatically when a browser profile starts.
func (c *Client) SetBookmarks(ctx context.Context, req SetBookmarksRequest) (*SetBookmarksResponse, error) {
	var resp SetBookmarksResponse
	if err := c.core.Do(ctx, "/api/v1/browser/setBookmark", req, nil); err != nil {
		return nil, err
	}
	return &resp, nil
}

// GetBookmarksRequest contains no parameters for [Client.GetBookmarks].
type GetBookmarksRequest struct{}

type GetBookmarksResponseBrowserBookmark struct {
	// Bookmark type, 0 - No setting, 1 - Uploaded file, 2 - Manually created.
	Type int `json:"type"`
	// Bookmark HTML file address.
	FileAddr string `json:"fileAddr"`
	// Manually created bookmark content, multiple bookmarks separated by newline characters
	// \n, supported formats: Folder::Name::URL Name::URL URL.
	Text string `json:"text"`
}

type GetBookmarksResponse struct {
	// Browser Bookmarks.
	BrowserBookmark GetBookmarksResponseBrowserBookmark `json:"browserBookmark"`
}

// GetBookmarks: https://open.geelark.com/api/browser-getBookmark
func (c *Client) GetBookmarks(ctx context.Context, req GetBookmarksRequest) (*GetBookmarksResponse, error) {
	var resp GetBookmarksResponse
	if err := c.core.Do(ctx, "/api/v1/browser/getBookmark", req, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
