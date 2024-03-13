package client

import (
	"context"
	"encoding/json"
	"net/url"

	"github.com/RiemaLabs/nubit-da-sdk/constant"
	"github.com/RiemaLabs/nubit-da-sdk/log"
	"github.com/RiemaLabs/nubit-da-sdk/types"
)

func (c *Client) SubmitTransaction(ctx context.Context, req *types.TxReq) (data *types.TxRsp, err error) {
	path, err := url.JoinPath(c.c.Endpoint, c.u.GetUri(constant.SubmitTransaction))
	if err != nil {
		log.Error("client.Transaction", "method", "SubmitTransaction", "JoinPath", err)
		return nil, err
	}
	post, err := c.c.Post(ctx, path, req, nil)
	if err != nil {
		log.Error("client.Transaction", "method", "SubmitTransaction", "submitTransaction", err)
		return nil, err
	}
	err = json.Unmarshal(post, &data)
	if err != nil {
		log.Error("client.Transaction", "method", "SubmitTransaction", "Unmarshal", err)
		return nil, err
	}
	return data, err
}

func (c *Client) GetTransaction(ctx context.Context, req *types.GetTransactionReq) (data *types.GetTransactionRsp, err error) {
	path, err := url.JoinPath(c.c.Endpoint, c.u.GetUri(constant.GetTransaction))
	if err != nil {
		log.Error("client.Transaction", "method", "GetTransaction", "JoinPath", err)
		return nil, err
	}
	post, err := c.c.Post(ctx, path, req, nil)
	if err != nil {
		log.Error("client.Transaction", "method", "GetTransaction", "getTransaction", err)
		return nil, err
	}

	err = json.Unmarshal(post, &data)
	if err != nil {
		log.Error("client.Transaction", "method", "GetTransaction", "Unmarshal", err)
		return nil, err
	}
	return data, err
}

// CountTransactions indexer-countTransactions
func (c *Client) CountTransactions(ctx context.Context) (data *types.CountRsp, err error) {
	path, err := url.JoinPath(c.c.Endpoint, c.u.GetUri(constant.CountTransactions))
	if err != nil {
		log.Error("client.Transaction", "method", "indexer-CountTransactions", "JoinPath", err)
		return nil, err
	}
	post, err := c.c.Post(ctx, path, nil, nil)
	if err != nil {
		log.Error("client.Transaction", "method", "indexer-CountTransactions", "getNamespaceStatusReq", err)
		return nil, err
	}

	err = json.Unmarshal(post, &data)
	if err != nil {
		log.Error("client.Transaction", "method", "indexer-CountTransactions", "Unmarshal", err)
		return nil, err
	}
	return data, err
}

// GetTransactions indexer-getTransactions
func (c *Client) GetTransactions(ctx context.Context, req *types.GetTransactionsReq) (data *types.GetTransactionsRsp, err error) {
	path, err := url.JoinPath(c.c.Endpoint, c.u.GetUri(constant.GetTransactions))
	if err != nil {
		log.Error("client.Transaction", "method", "indexer-GetTransactions", "JoinPath", err)
		return nil, err
	}
	post, err := c.c.Post(ctx, path, req, nil)
	if err != nil {
		log.Error("client.Transaction", "method", "indexer-GetTransactions", "getNamespaceStatusReq", err)
		return nil, err
	}
	err = json.Unmarshal(post, &data)
	if err != nil {
		log.Error("client.Transaction", "method", "indexer-GetTransactions", "Unmarshal", err)
		return nil, err
	}
	return data, err
}

// GetTxIDByDataID indexer-getTxIDByDataID
func (c *Client) GetTxIDByDataID(ctx context.Context, req *types.GetTransactionsReq) (data *types.GetTransactionsRsp, err error) {
	path, err := url.JoinPath(c.c.Endpoint, c.u.GetUri(constant.GetTxIDByDataID))
	if err != nil {
		log.Error("client.Transaction", "method", "indexer-GetTxIDByDataID", "JoinPath", err)
		return nil, err
	}
	post, err := c.c.Post(ctx, path, req, nil)
	if err != nil {
		log.Error("client.Transaction", "method", "indexer-GetTxIDByDataID", "getNamespaceStatusReq", err)
		return nil, err
	}

	err = json.Unmarshal(post, &data)
	if err != nil {
		log.Error("client.Transaction", "method", "indexer-GetTxIDByDataID", "Unmarshal", err)
		return nil, err
	}
	return data, err
}
