package types

import (
	"github.com/RiemaLabs/nubit-da-sdk/clients/grpc"
	"github.com/RiemaLabs/nubit-da-sdk/clients/http"
)

type (
	NubitClient struct {
		*NubitHttp
		Endpoint string
	}

	NubitRpc struct {
		c *grpc.Client
	}

	NubitHttp struct {
		*http.Client
	}
)

func NewHttp(c *http.Client) *NubitHttp {
	return &NubitHttp{c}
}
