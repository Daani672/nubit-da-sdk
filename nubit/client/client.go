package client

import (
	"github.com/RiemaLabs/nubit-da-sdk/types"
)

type Client struct {
	c *types.NubitClient
	u *Uri
}

func Dial(ctx Context) (*Client, error) {
	c, err := DialContext(ctx)
	if err != nil {
		return nil, err
	}
	return NewClient(ctx, c), nil
}

func NewClient(ctx Context, c *types.NubitClient) *Client {
	return &Client{c: c, u: NewUri(ctx)}
}
