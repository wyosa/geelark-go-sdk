package files

import (
	"context"
	"io"
	"net/http"
	"os"
)

// Upload obtains a temporary signed URL and uploads body with PUT, over HTTPS
// even when GeeLark returns a plain-HTTP Alibaba Cloud OSS URL.
// The returned resource URL is valid for three days according to GeeLark.
// Size must be the exact number of bytes; zero requires an empty reader or nil.
// A zero-size reader is probed before upload so nonempty content is not discarded.
// Readers must return promptly or support cancellation themselves. The HTTP
// transport may close body if it implements io.Closer; do not reuse it afterward.
func (c *Client) Upload(ctx context.Context, req GetUploadURLRequest, body io.Reader, size int64) (*GetUploadURLResponse, error) {
	signed, err := c.GetUploadURL(ctx, req)
	if err != nil {
		return nil, err
	}
	if err := c.core.Upload(ctx, http.MethodPut, signed.UploadURL, nil, body, size); err != nil {
		return nil, err
	}
	return signed, nil
}

// UploadFile opens path, uploads the complete file, and closes it on return.
// Set req.FileType to the API's file extension and give ctx a deadline to bound
// storage upload time. The returned resource URL is temporary, as with Upload.
func (c *Client) UploadFile(ctx context.Context, path string, req GetUploadURLRequest) (*GetUploadURLResponse, error) {
	f, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer f.Close()
	info, err := f.Stat()
	if err != nil {
		return nil, err
	}
	return c.Upload(ctx, req, f, info.Size())
}
