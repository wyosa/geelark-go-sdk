package common

import (
	"context"
	"io"
)

// Transport sends authenticated API requests and decodes the response envelope.
// Implementations must be safe for concurrent calls.
type Transport interface {
	// Do sends a JSON POST to path. Nonzero API codes and non-2xx HTTP statuses
	// must return an error. A non-nil out requires non-null response data;
	// nil out skips data decoding for endpoints without a response model.
	Do(ctx context.Context, path string, requestBody, out any) error
}

// OptionalDataTransport allows missing response data, as in CloudPhones.Update.
type OptionalDataTransport interface {
	Transport
	// DoOptionalData behaves like Do but allows missing or null data.
	DoOptionalData(ctx context.Context, path string, requestBody, out any) error
}

// UploadTransport additionally sends file bytes to signed storage URLs.
type UploadTransport interface {
	Transport
	// Upload sends exactly size bytes without adding API credentials. Zero
	// denotes an empty body; it does not mean that the size is unknown.
	Upload(ctx context.Context, method, signedURL string, headers map[string]string, body io.Reader, size int64) error
}
