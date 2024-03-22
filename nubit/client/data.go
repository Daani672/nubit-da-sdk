package client

import (
	"context"
	"crypto/ecdsa"
	"encoding/json"
	"errors"
	"math"
	"net/url"
	"time"

	"github.com/RiemaLabs/nubit-da-sdk/constant"
	"github.com/RiemaLabs/nubit-da-sdk/log"
	types2 "github.com/RiemaLabs/nubit-da-sdk/types"
)

func (c *Client) Upload(ctx context.Context, data *types2.DataUploadReq, privateKey *ecdsa.PrivateKey, payParams *types2.PaymentParams) (target *types2.DataUploadRsp, err error) {
	if payParams != nil {
		if payParams.EstimateFeeMultiple > 0 {
			FeeMultiple := payParams.EstimateFeeMultiple * float64(data.StorageFee)
			data.StorageFee = uint64(int(math.Ceil(FeeMultiple)))
		}
		if int(data.StorageFee) < payParams.Approve || payParams.XAPIKEY != "" || payParams.Authorization != "" {
			payParams.AutoPay = true
		}
	}
	tx, err := c.newTx(ctx, data, privateKey, data.MethodName)
	if err != nil {
		log.Error("client.Namespace", "method", "TransferNamespace", "newTx", err)
		return nil, err
	}

	tx, err = c.getPreImage(ctx, tx, payParams)
	if err != nil {
		log.Error("client.data", "method", "Upload", "getPreImage.err", err)
		return nil, err
	}
	subTxData, err := c.SubmitTransaction(ctx, tx)
	if err != nil {
		log.Error("client.data", "method", "Upload", "submitTransaction", err)
		return nil, err
	}

	if !subTxData.Success {
		return nil, errors.New(subTxData.Msg)
	}

	return &types2.DataUploadRsp{
		TxID:      subTxData.TxID,
		Timestamp: time.Now().Unix(),
	}, nil
}

func (c *Client) GetData(ctx context.Context, req *types2.GetDataReq) (data *types2.GetDataRsp, err error) {
	path, err := url.JoinPath(c.c.Endpoint, c.u.GetUri(constant.GetData))
	if err != nil {
		log.Error("client.data", "method", "GetData", "JoinPath", err)
		return nil, err
	}
	post, err := c.c.Post(ctx, path, req, nil)
	if err != nil {
		log.Error("client.data", "method", "GetData", "GetData", err)
		return nil, err
	}

	err = json.Unmarshal(post, &data)
	if err != nil {
		log.Error("client.data", "method", "GetData", "Unmarshal", err)
		return nil, err
	}
	return data, err
}

func (c *Client) GetDataStatus(ctx context.Context, req *types2.GetDataStatusReq) (data *types2.GetDataStatusRsp, err error) {
	path, err := url.JoinPath(c.c.Endpoint, c.u.GetUri(constant.GetDataStatus))
	if err != nil {
		log.Error("client.data", "method", "GetData", "JoinPath", err)
		return nil, err
	}
	post, err := c.c.Post(ctx, path, req, nil)
	if err != nil {
		log.Error("client.data", "method", "GetData", "GetDataStatus", err)
		return nil, err
	}
	err = json.Unmarshal(post, &data)
	if err != nil {
		log.Error("client.data", "method", "GetData", "Unmarshal", err)
		return nil, err
	}
	return data, err
}

func (c *Client) GetDAIDByTxID(ctx context.Context, req *types2.QueryTxReq) (data *types2.GetDAIDRsp, err error) {
	path, err := url.JoinPath(c.c.Endpoint, c.u.GetUri(constant.GetDAIDByTxID))
	if err != nil {
		log.Error("client.data", "method", "GetDAIDByTxID", "JoinPath", err)
		return nil, err
	}
	post, err := c.c.Post(ctx, path, req, nil)
	if err != nil {
		log.Error("client.data", "method", "GetDAIDByTxID", "GetDAIDByTxID", err)
		return nil, err
	}
	err = json.Unmarshal(post, &data)
	if err != nil {
		log.Error("client.data", "method", "GetDAIDByTxID", "Unmarshal", err)
		return nil, err
	}
	return data, err
}

func (c *Client) GetDataByDAID(ctx context.Context, req *types2.GetDataReq) (data *types2.GetDataCallRsp, err error) {
	path, err := url.JoinPath(c.c.Endpoint, c.u.GetUri(constant.GetDataByDAID))
	if err != nil {
		log.Error("client.data", "method", "GetDataByDAID", "JoinPath", err)
		return nil, err
	}
	post, err := c.c.Post(ctx, path, req, nil)
	if err != nil {
		log.Error("client.data", "method", "GetDataByDAID", "GetDataByDAID", err)
		return nil, err
	}
	err = json.Unmarshal(post, &data)
	if err != nil {
		log.Error("client.data", "method", "GetDataByDAID", "Unmarshal", err)
		return nil, err
	}
	return data, err
}

func (c *Client) GetDataIDByUser(ctx context.Context, req *types2.GetDataByUserReq) (data *types2.GetDataByUserRsp, err error) {
	path, err := url.JoinPath(c.c.Endpoint, c.u.GetUri(constant.GetDataIDByUser))
	if err != nil {
		log.Error("client.data", "method", "GetDataIDByUser", "JoinPath", err)
		return nil, err
	}
	post, err := c.c.Post(ctx, path, req, nil)
	if err != nil {
		log.Error("client.data", "method", "GetDataIDByUser", "GetDataIDByUser", err)
		return nil, err
	}
	err = json.Unmarshal(post, &data)
	if err != nil {
		log.Error("client.data", "method", "GetDataIDByUser", "Unmarshal", err)
		return nil, err
	}
	return data, err
}

func (c *Client) GetEstimateFee(ctx context.Context, req any, privateKey *ecdsa.PrivateKey, method string, nid string) (data *types2.EstimateFee, err error) {
	body, err := c.newTransactionBody(ctx, req, privateKey, method)
	if err != nil {
		return nil, err
	}
	reqData, err := json.Marshal(body)
	if err != nil {
		log.Error("client.tx", "method", "GetEstimateFee", "Marshal", err)
		return nil, err
	}

	path, err := url.JoinPath(c.c.Endpoint, c.u.GetUri(constant.GetEstimateFee))
	if err != nil {
		log.Error("client.tx", "method", "GetEstimateFee", "JoinPath", err)
		return nil, err
	}

	datasize := len(reqData) + constant.BaseDataSize // because storagefee ,we try to make it bigger
	post, err := c.c.Post(ctx, path, &types2.EstimateFeeReq{
		DataSize:    datasize,
		MethodName:  method,
		NamespaceID: nid,
	}, nil)
	if err != nil {
		log.Error("client.tx", "method", "GetEstimateFee", "GetEstimateFee", err)
		return nil, err
	}

	err = json.Unmarshal(post, &data)
	if err != nil {
		log.Error("client.tx", "method", "GetEstimateFee", "Unmarshal", err)
		return nil, err
	}
	return data, err
}

// CountDatasize indexer-countDatasize
func (c *Client) CountDatasize(ctx context.Context) (data *types2.CountRsp, err error) {
	path, err := url.JoinPath(c.c.Endpoint, c.u.GetUri(constant.CountDatasize))
	if err != nil {
		log.Error("client.data", "method", "indexer-CountDatasize", "JoinPath", err)
		return nil, err
	}
	post, err := c.c.Post(ctx, path, nil, nil)
	if err != nil {
		log.Error("client.data", "method", "indexer-CountDatasize", "getNamespaceStatusReq", err)
		return nil, err
	}

	err = json.Unmarshal(post, &data)
	if err != nil {
		log.Error("client.data", "method", "indexer-CountDatasize", "Unmarshal", err)
		return nil, err
	}
	return data, err
}
