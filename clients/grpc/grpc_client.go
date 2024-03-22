package grpc

import (
	"context"
	"time"

	"google.golang.org/grpc"

	"github.com/RiemaLabs/nubit-da-sdk/log"
)

type Client struct {
	Conn *grpc.ClientConn
	Opts *Option
}

func NewClient(opts ...Opt) (*Client, error) {
	client := &Client{
		Opts: DefaultOptions(),
	}
	for _, op := range opts {
		op(client.Opts)
	}
	ctx := context.TODO()
	ctx, _ = context.WithTimeout(ctx, time.Minute)
	_conn, err := grpc.DialContext(ctx, client.Opts.Uri, client.Opts.GrpcOpt...)
	if err != nil {
		log.Error("grpc_client", "Dial.err", err, "uri", client.Opts.Uri)
		return nil, err
	}
	client.Conn = _conn
	return client, nil
}
