package library

import (
	"context"
	"io"
	"net/http"
	"os"
)

// Upload obtains a signed URL and uploads a Library file with PUT.
// Call Create with the returned ResourceURL to save it in the Library.
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
