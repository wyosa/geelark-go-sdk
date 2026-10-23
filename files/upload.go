package files

import (
	"context"
	"io"
	"net/http"
	"os"
)

// Upload obtains a temporary signed URL and uploads body with PUT.
// The returned resource URL is valid for three days according to GeeLark.
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
