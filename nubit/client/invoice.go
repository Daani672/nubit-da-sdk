package client

import (
	"context"
	"encoding/json"
	"errors"
	"net/url"

	"github.com/RiemaLabs/nubit-da-sdk/constant"
	"github.com/RiemaLabs/nubit-da-sdk/log"
	"github.com/RiemaLabs/nubit-da-sdk/nubit/lightning"
	"github.com/RiemaLabs/nubit-da-sdk/types"
	"golang.org/x/sync/errgroup"
)

func (c *Client) GetInvoice(ctx context.Context, req *types.GetInvoiceReq) (data *types.Invoice, err error) {
	path, err := url.JoinPath(c.c.Endpoint, c.u.GetUri(constant.GetInvoice))
	if err != nil {
		log.Error("client.lightning", "method", "GetInvoice", "JoinPath", err)
		return nil, err
	}
	post, err := c.c.Post(ctx, path, req, nil)
	if err != nil {
		log.Error("client.lightning", "method", "GetInvoice", "GetInvoice", err)
		return nil, err
	}

	err = json.Unmarshal(post, &data)
	if err != nil {
		log.Error("client.lightning", "method", "GetInvoice", "Unmarshal", err)
		return nil, err
	}
	return data, err
}

func (c *Client) GetPayee(ctx context.Context) (data *types.Payee, err error) {
	path, err := url.JoinPath(c.c.Endpoint, c.u.GetUri(constant.GetPayee))
	if err != nil {
		log.Error("client.data", "method", "GetPayee", "JoinPath", err)
		return nil, err
	}
	post, err := c.c.Post(ctx, path, nil, nil)
	if err != nil {
		log.Error("client.data", "method", "GetPayee", "GetPayee", err)
		return nil, err
	}

	err = json.Unmarshal(post, &data)
	if err != nil {
		log.Error("client.data", "method", "GetPayee", "Unmarshal", err)
		return nil, err
	}
	return data, err
}

func (c *Client) getPreImage(ctx context.Context, req *types.TxReq, payParams *types.PaymentParams) (data *types.TxReq, err error) {
	var (
		invoice  *types.Invoice
		preimage string
	)
	g, ctx := errgroup.WithContext(ctx)
	g.Go(func() error {
		invoice, err = c.GetInvoice(ctx, &types.GetInvoiceReq{
			PublicKeyStr:   req.TransactionBody.PublicKey,
			BitcoinAddress: req.TransactionBody.BitcoinAddress,
			StorageFee:     int64(req.TransactionBody.StorageFee),
		})
		if err != nil {
			return err
		}
		return nil
	})
	if err = g.Wait(); err != nil {
		return nil, err
	}
	if invoice == nil {
		return nil, errors.New("get invoice failed")
	}
	switch true {
	case payParams != nil && payParams.AutoPay:
		payment, err := lightning.NewClient(payParams).Payment(ctx, invoice.Invoice, int64(payParams.FeeLimitSat))
		if err != nil {
			log.Error("client.lightning", "method", "getPreImage", "Payment", err)
			return nil, err
		}
		if payment == nil {
			return nil, errors.New("payment failed")
		}
		preimage = payment.Preimage
	default:
	}

	req.Invoice = invoice.Invoice
	req.PreImage = preimage
	return req, nil
}
