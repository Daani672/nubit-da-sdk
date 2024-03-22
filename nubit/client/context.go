package client

import (
	"context"

	"github.com/RiemaLabs/nubit-da-sdk/constant"
)

type Context struct {
	context.Context
	NubitRpc string
	ProxyRpc string
	UriMap   map[string]string
}

func Background() Context {
	ctx := Context{
		Context: context.Background(),
		UriMap:  nil,
	}
	switch constant.NubitNet {
	case constant.PreAlphaTestNet:
		ctx.NubitRpc = constant.NubitRpc
		ctx.ProxyRpc = constant.ProxyRpc
		ctx.UriMap = constant.Release
	case constant.TestNet:
		ctx.NubitRpc = constant.NubitTestRpc
		ctx.ProxyRpc = constant.ProxyTestRpc
		ctx.UriMap = constant.MockApi
	}
	return ctx
}

func WithMoc(ctx Context) Context {
	ctx.UriMap = constant.MockApi
	return ctx
}

func WithRelease(ctx Context) Context {
	ctx.UriMap = constant.Release
	return ctx
}

func UrlCtx(url string) Context {
	return Context{
		Context:  context.Background(),
		NubitRpc: url,
		UriMap:   nil,
	}
}
