package client

import (
	"fmt"
	"net/url"

	"github.com/RiemaLabs/nubit-da-sdk/clients/http"
	"github.com/RiemaLabs/nubit-da-sdk/types"
)

const (
	maxRequestContentLength = 1024 * 1024 * 5
	contentType             = "application/json"
)

// DialContext creates a new RPC client, just like Dial.
//
// The context is used to cancel or time out the initial connection establishment. It does
// not affect subsequent interactions with the client.
func DialContext(ctx Context) (*types.NubitClient, error) {
	u, err := url.Parse(ctx.ProxyRpc)
	if err != nil {
		return nil, err
	}
	switch u.Scheme {
	case "http", "https":
		return DialHTTP(ctx)
	//case "ws", "wss":
	//	return DialWebsocket(ctx, rawurl, "")
	//case "stdio":
	//	return DialStdIO(ctx)
	//case "":
	//	return DialIPC(ctx, rawurl)
	default:
		return nil, fmt.Errorf("no known transport for URL scheme %q", u.Scheme)
	}
}

// DialHTTP creates a new RPC client that connects to an RPC server over HTTP.
func DialHTTP(ctx Context) (*types.NubitClient, error) {
	return DialHTTPWithClient(ctx)
}

// DialHTTPWithClient creates a new RPC client that connects to an RPC server over HTTP
// using the provided HTTP Client.
func DialHTTPWithClient(ctx Context) (*types.NubitClient, error) {
	// Sanity check URL so we don't end up with a client that will fail every request.
	_, err := url.Parse(ctx.ProxyRpc)
	if err != nil {
		return nil, err
	}
	return &types.NubitClient{
		NubitHttp: types.NewHttp(http.NewClient()),
		Endpoint:  ctx.NubitRpc,
	}, err
}
