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
	"github.com/RiemaLabs/nubit-da-sdk/types"
)

func (c *Client) GetNIDByTxID(ctx context.Context, req *types.QueryTxReq) (data *types.GetNIDRsp, err error) {
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

// GetNamespaceByUser
// Deprecated: use GetTransactions instead.
func (c *Client) GetNamespaceByUser(ctx context.Context, req *types.GetNamespaceByUserReq) (data *types.GetNamespaceByUserRsp, err error) {
	path, err := url.JoinPath(c.c.Endpoint, c.u.GetUri(constant.GetNamespaceByUser))
	if err != nil {
		log.Error("client.data", "method", "GetNamespaceByUser", "JoinPath", err)
		return nil, err
	}
	post, err := c.c.Post(ctx, path, req, nil)
	if err != nil {
		log.Error("client.data", "method", "GetNamespaceByUser", "GetDataIDByUser", err)
		return nil, err
	}
	err = json.Unmarshal(post, &data)
	if err != nil {
		log.Error("client.data", "method", "GetNamespaceByUser", "Unmarshal", err)
		return nil, errors.Join(errors.New(string(post)), err)
	}
	return data, err
}

func (c *Client) GetNamespace(ctx context.Context, req *types.GetNamespaceReq) (data *types.GetNamespaceRsp, err error) {
	path, err := url.JoinPath(c.c.Endpoint, c.u.GetUri(constant.GetNamespace))
	if err != nil {
		log.Error("client.Namespace", "method", "GetNamespace", "JoinPath", err)
		return nil, err
	}
	post, err := c.c.Post(ctx, path, req, nil)
	if err != nil {
		log.Error("client.Namespace", "method", "GetNamespace", "getNamespaceReq", err)
		return nil, err
	}

	err = json.Unmarshal(post, &data)
	if err != nil {
		log.Error("client.Namespace", "method", "GetNamespace", "Unmarshal", err)
		return nil, err
	}
	return data, err
}

func (c *Client) GetNamespaceStatus(ctx context.Context, req *types.GetNamespaceStatusReq) (data *types.GetNamespaceStatusRsp, err error) {
	path, err := url.JoinPath(c.c.Endpoint, c.u.GetUri(constant.GetNamespaceStatus))
	if err != nil {
		log.Error("client.Namespace", "method", "GetNamespaceStatus", "JoinPath", err)
		return nil, err
	}
	post, err := c.c.Post(ctx, path, req, nil)
	if err != nil {
		log.Error("client.Namespace", "method", "GetNamespaceStatus", "getNamespaceStatusReq", err)
		return nil, err
	}

	err = json.Unmarshal(post, &data)
	if err != nil {
		log.Error("client.Namespace", "method", "GetNamespaceStatus", "Unmarshal", err)
		return nil, err
	}
	return data, err
}

func (c *Client) GetLastNamespaceID(ctx context.Context, req *types.GetLastNamespaceIDReq) (data *types.GetLastNamespaceIDRsp, err error) {
	path, err := url.JoinPath(c.c.Endpoint, c.u.GetUri(constant.GetLastNamespaceID))
	if err != nil {
		log.Error("client.Namespace", "method", "GetLastNamespaceID", "JoinPath", err)
		return nil, err
	}
	post, err := c.c.Post(ctx, path, req, nil)
	if err != nil {
		log.Error("client.Namespace", "method", "GetLastNamespaceID", "getLastNamespaceIDReq", err)
		return nil, err
	}

	err = json.Unmarshal(post, &data)
	if err != nil {
		log.Error("client.Namespace", "method", "GetLastNamespaceID", "Unmarshal", err)
		return nil, err
	}
	return data, err
}

func (c *Client) GetDataInNamespace(ctx context.Context, req *types.GetDataInNamespaceReq) (data *types.GetDataInNamespaceRsp, err error) {
	path, err := url.JoinPath(c.c.Endpoint, c.u.GetUri(constant.GetDataInNamespace))
	if err != nil {
		log.Error("client.Namespace", "method", "GetDataInNamespace", "JoinPath", err)
		return nil, err
	}
	post, err := c.c.Post(ctx, path, req, nil)
	if err != nil {
		log.Error("client.Namespace", "method", "GetDataInNamespace", "getDataInNamespace", err)
		return nil, err
	}

	if len(post) == 0 {
		return &types.GetDataInNamespaceRsp{}, nil
	}
	err = json.Unmarshal(post, &data)
	if err != nil {
		log.Error("client.Namespace", "method", "GetDataInNamespace", "Unmarshal", err)
		return nil, err
	}
	return data, err
}

func (c *Client) GetTotalDataIDsInNamesapce(ctx context.Context, req *types.GetTotalDataIDsInNamesapceReq) (data *types.GetTotalDataIDsInNamesapceRsp, err error) {
	path, err := url.JoinPath(c.c.Endpoint, c.u.GetUri(constant.GetTotalDataIDsInNamesapce))
	if err != nil {
		log.Error("client.Namespace", "method", "GetTotalDataIDsInNamesapce", "JoinPath", err)
		return nil, err
	}
	post, err := c.c.Post(ctx, path, req, nil)
	if err != nil {
		log.Error("client.Namespace", "method", "GetTotalDataIDsInNamesapce", "getDataInNamespace", err)
		return nil, err
	}

	if len(post) == 0 {
		return &types.GetTotalDataIDsInNamesapceRsp{}, nil
	}
	err = json.Unmarshal(post, &data)
	if err != nil {
		log.Error("client.Namespace", "method", "GetTotalDataIDsInNamesapce", "Unmarshal", err)
		return nil, err
	}
	return data, err
}

func (c *Client) CreateNamespace(ctx context.Context, req *types.CreateNameSpaceReq, privateKey *ecdsa.PrivateKey, payParams *types.PaymentParams) (data *types.DataUploadRsp, err error) {
	fee, err := c.GetEstimateFee(ctx, req, privateKey, req.MethodName, "")
	if err != nil {
		log.Error("client.data", "method", "Upload", "GetEstimateFee", err)
		return nil, err
	}
	if fee != nil {
		req.StorageFee = int64(fee.StorageFee)
	}
	if payParams != nil {
		if payParams.EstimateFeeMultiple > 0 {
			FeeMultiple := payParams.EstimateFeeMultiple * float64(fee.StorageFee)
			req.StorageFee = int64(int(math.Ceil(FeeMultiple)))
		}
		if int(req.StorageFee) < payParams.Approve || payParams.XAPIKEY != "" || payParams.Authorization != "" {
			payParams.AutoPay = true
		}
	}

	txReq, err := c.newTx(ctx, req, privateKey, req.MethodName)
	if err != nil {
		log.Error("client.Namespace", "method", "CreateNamespace", "getPreImage.err", err)
		return nil, err
	}
	txReq, err = c.getPreImage(ctx, txReq, payParams)
	if err != nil {
		log.Error("client.Namespace", "method", "CreateNamespace", "getPreImage.err", err)
		return nil, err
	}
	subTxData, err := c.SubmitTransaction(ctx, txReq)
	if err != nil {
		log.Error("client.Namespace", "method", "CreateNamespace", "submitTransaction", err)
		return nil, err
	}
	if !subTxData.Success {
		return nil, errors.New(subTxData.Msg)
	}

	return &types.DataUploadRsp{
		TxID:      subTxData.TxID,
		Timestamp: time.Now().Unix(),
	}, nil
}

func (c *Client) UpdateNamespace(ctx context.Context, req *types.UpdateNamespaceReq, privateKey *ecdsa.PrivateKey, payParams *types.PaymentParams) (data *types.DataUploadRsp, err error) {
	fee, err := c.GetEstimateFee(ctx, req, privateKey, req.MethodName, req.NID)
	if err != nil {
		log.Error("client.data", "method", "Upload", "GetEstimateFee", err)
		return nil, err
	}
	if fee != nil {
		req.StorageFee = uint64(fee.StorageFee)
	}
	if payParams != nil {
		if payParams.EstimateFeeMultiple > 0 {
			FeeMultiple := payParams.EstimateFeeMultiple * float64(fee.StorageFee)
			req.StorageFee = uint64(int(math.Ceil(FeeMultiple)))
		}
		if int(req.StorageFee) < payParams.Approve || payParams.XAPIKEY != "" || payParams.Authorization != "" {
			payParams.AutoPay = true
		}
	}
	txReq, err := c.newTx(ctx, req, privateKey, req.MethodName)
	if err != nil {
		log.Error("client.Namespace", "method", "UpdateNamespace", "newTx", err)
		return nil, err
	}
	txReq, err = c.getPreImage(ctx, txReq, payParams)
	if err != nil {
		log.Error("client.Namespace", "method", "UpdateNamespace", "getPreImage.err", err)
		return nil, err
	}
	subTxData, err := c.SubmitTransaction(ctx, txReq)
	if err != nil {
		log.Error("client.Namespace", "method", "UpdateNamespace", "submitTransaction", err)
		return nil, err
	}
	if !subTxData.Success {
		return nil, errors.New(subTxData.Msg)
	}

	return &types.DataUploadRsp{
		TxID:      subTxData.TxID,
		Timestamp: time.Now().Unix(),
	}, nil
}

func (c *Client) TransferNamespace(ctx context.Context, req *types.TransferNamespaceReq, privateKey *ecdsa.PrivateKey, payParams *types.PaymentParams) (data *types.DataUploadRsp, err error) {
	fee, err := c.GetEstimateFee(ctx, req, privateKey, req.MethodName, req.NID)
	if err != nil {
		log.Error("client.data", "method", "Upload", "GetEstimateFee", err)
		return nil, err
	}
	if fee != nil {
		req.StorageFee = uint64(fee.StorageFee)
	}
	if payParams != nil {
		if payParams.EstimateFeeMultiple > 0 {
			FeeMultiple := payParams.EstimateFeeMultiple * float64(fee.StorageFee)
			req.StorageFee = uint64(int(math.Ceil(FeeMultiple)))
		}
		if int(req.StorageFee) < payParams.Approve || payParams.XAPIKEY != "" || payParams.Authorization != "" {
			payParams.AutoPay = true
		}
	}
	txReq, err := c.newTx(ctx, req, privateKey, req.MethodName)
	if err != nil {
		log.Error("client.Namespace", "method", "TransferNamespace", "newTx", err)
		return nil, err
	}
	txReq, err = c.getPreImage(ctx, txReq, payParams)
	if err != nil {
		log.Error("client.Namespace", "method", "TransferNamespace", "getPreImage.err", err)
		return nil, err
	}
	subTxData, err := c.SubmitTransaction(ctx, txReq)
	if err != nil {
		log.Error("client.Namespace", "method", "TransferNamespace", "submitTransaction", err)
		return nil, err
	}

	if !subTxData.Success {
		return nil, errors.New(subTxData.Msg)
	}

	return &types.DataUploadRsp{
		TxID:      subTxData.TxID,
		Timestamp: time.Now().Unix(),
	}, nil
}

// GetNamespaces indexer-getNamespaces
func (c *Client) GetNamespaces(ctx context.Context, req *types.GetNamespacesReq) (data *types.GetNamespacesRsp, err error) {
	path, err := url.JoinPath(c.c.Endpoint, c.u.GetUri(constant.GetNamespaces))
	if err != nil {
		log.Error("client.Namespace", "method", "indexer-GetNamespaces", "JoinPath", err)
		return nil, err
	}
	post, err := c.c.Post(ctx, path, req, nil)
	if err != nil {
		log.Error("client.Namespace", "method", "indexer-GetNamespaces", "getNamespaceStatusReq", err)
		return nil, err
	}

	err = json.Unmarshal(post, &data)
	if err != nil {
		log.Error("client.Namespace", "method", "indexer-GetNamespaces", "Unmarshal", err)
		return nil, err
	}
	return data, err
}

// CountNamespace indexer-countNamespace
func (c *Client) CountNamespace(ctx context.Context) (data *types.CountRsp, err error) {
	path, err := url.JoinPath(c.c.Endpoint, c.u.GetUri(constant.CountNamespace))
	if err != nil {
		log.Error("client.Namespace", "method", "indexer-CountNamespace", "JoinPath", err)
		return nil, err
	}
	post, err := c.c.Post(ctx, path, nil, nil)
	if err != nil {
		log.Error("client.Namespace", "method", "indexer-CountNamespace", "getNamespaceStatusReq", err)
		return nil, err
	}

	err = json.Unmarshal(post, &data)
	if err != nil {
		log.Error("client.Namespace", "method", "indexer-CountNamespace", "Unmarshal", err)
		return nil, err
	}
	return data, err
}
