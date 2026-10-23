package automation

import "context"

type FacebookLoginRequest = FaceBookLoginRequest

type FacebookLoginResponse = FaceBookLoginResponse

// FacebookLogin is an alias for [Client.FaceBookLogin].
func (c *Client) FacebookLogin(ctx context.Context, req FacebookLoginRequest) (*FacebookLoginResponse, error) {
	return c.FaceBookLogin(ctx, req)
}

type FacebookAutoCommentRequest = FaceBookAutoCommentRequest

type FacebookAutoCommentResponse = FaceBookAutoCommentResponse

// FacebookAutoComment is an alias for [Client.FaceBookAutoComment].
func (c *Client) FacebookAutoComment(ctx context.Context, req FacebookAutoCommentRequest) (*FacebookAutoCommentResponse, error) {
	return c.FaceBookAutoComment(ctx, req)
}

type FacebookPublishRequest = FaceBookPublishRequest

type FacebookPublishResponse = FaceBookPublishResponse

// FacebookPublish is an alias for [Client.FaceBookPublish].
func (c *Client) FacebookPublish(ctx context.Context, req FacebookPublishRequest) (*FacebookPublishResponse, error) {
	return c.FaceBookPublish(ctx, req)
}

type FacebookActiveAccountRequest = FaceBookActiveAccountRequest

type FacebookActiveAccountResponse = FaceBookActiveAccountResponse

// FacebookActiveAccount is an alias for [Client.FaceBookActiveAccount].
func (c *Client) FacebookActiveAccount(ctx context.Context, req FacebookActiveAccountRequest) (*FacebookActiveAccountResponse, error) {
	return c.FaceBookActiveAccount(ctx, req)
}

type FacebookPubReelsRequest = FaceBookPubReelsRequest

type FacebookPubReelsResponse = FaceBookPubReelsResponse

// FacebookPubReels is an alias for [Client.FaceBookPubReels].
func (c *Client) FacebookPubReels(ctx context.Context, req FacebookPubReelsRequest) (*FacebookPubReelsResponse, error) {
	return c.FaceBookPubReels(ctx, req)
}

type FacebookMessageRequest = FaceBookMessageRequest

type FacebookMessageResponse = FaceBookMessageResponse

// FacebookMessage is an alias for [Client.FaceBookMessage].
func (c *Client) FacebookMessage(ctx context.Context, req FacebookMessageRequest) (*FacebookMessageResponse, error) {
	return c.FaceBookMessage(ctx, req)
}
