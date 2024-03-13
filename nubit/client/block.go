package client

import (
	"context"
	"encoding/json"
	"net/url"

	"github.com/RiemaLabs/nubit-da-sdk/constant"
	"github.com/RiemaLabs/nubit-da-sdk/log"
	"github.com/RiemaLabs/nubit-da-sdk/types"
)

func (c *Client) GetBlockHeaderByNumber(ctx context.Context, req *types.GetBlockHeaderByNumberReq) (data *types.GetBlockHeaderByNumberRsp, err error) {
	path, err := url.JoinPath(c.c.Endpoint, c.u.GetUri(constant.GetBlockHeaderByNumber))
	if err != nil {
		log.Error("client.Block", "method", "GetBlockHeaderByNumber", "JoinPath", err)
		return nil, err
	}
	post, err := c.c.Post(ctx, path, req, nil)
	if err != nil {
		log.Error("client.Block", "method", "GetBlockHeaderByNumber", "getBlockHeaderByNumber", err)
		return nil, err
	}

	err = json.Unmarshal(post, &data)
	if err != nil {
		log.Error("client.Block", "method", "GetBlockHeaderByNumber", "Unmarshal", err)
		return nil, err
	}
	return data, err
}

func (c *Client) GetLatestBlockNumber(ctx context.Context) (data *types.GetLatestBlockNumberRsp, err error) {
	path, err := url.JoinPath(c.c.Endpoint, c.u.GetUri(constant.GetLatestBlockNumber))
	if err != nil {
		log.Error("client.Block", "method", "GetLatestBlockNumber", "JoinPath", err)
		return nil, err
	}
	post, err := c.c.Post(ctx, path, nil, nil)
	if err != nil {
		log.Error("client.Block", "method", "GetLatestBlockNumber", "getLatestBlockNumber", err)
		return nil, err
	}

	err = json.Unmarshal(post, &data)
	if err != nil {
		log.Error("client.Block", "method", "GetLatestBlockNumber", "Unmarshal", err)
		return nil, err
	}
	return data, err
}

func (c *Client) GetNonce(ctx context.Context, req *types.GetNonceReq) (data *types.GetNonceRsp, err error) {
	path, err := url.JoinPath(c.c.Endpoint, c.u.GetUri(constant.GetNonce))
	if err != nil {
		log.Error("client.Block", "method", "GetNonce", "JoinPath", err)
		return nil, err
	}
	post, err := c.c.Post(ctx, path, req, nil)
	if err != nil {
		log.Error("client.Block", "method", "GetNonce", "getLatestBlockNumber", err)
		return nil, err
	}

	err = json.Unmarshal(post, &data)
	if err != nil {
		log.Error("client.Block", "method", "GetNonce", "Unmarshal", err)
		return nil, err
	}
	return data, err
}
