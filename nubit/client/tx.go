package client

import (
	"context"
	"crypto/ecdsa"
	"encoding/json"

	"github.com/RiemaLabs/nubit-da-sdk/constant"
	"github.com/RiemaLabs/nubit-da-sdk/log"
	types2 "github.com/RiemaLabs/nubit-da-sdk/types"
	"github.com/RiemaLabs/nubit-da-sdk/utils"
)

func (c *Client) newTx(ctx context.Context, req any, privateKey *ecdsa.PrivateKey, method string) (*types2.TxReq, error) {
	body, err := c.newTransactionBody(ctx, req, privateKey, method)
	if err != nil {
		return nil, err
	}
	reqData, err := json.Marshal(body)
	if err != nil {
		return nil, err
	}
	hash := utils.Sha256hash(string(reqData))
	Signature := utils.SignTransaction(utils.EcdsaToPrivateStr(privateKey), hash)
	txReq := &types2.TxReq{
		TransactionBody: body,
		Invoice:         "",
		PreImage:        "",
		Signature:       Signature,
	}
	return txReq, nil
}

func (c *Client) newTransactionBody(ctx context.Context, req interface{}, privateKey *ecdsa.PrivateKey, method string) (*types2.TransactionBody, error) {
	_, PublicKey, err := utils.BTCPRIKEYStrToETHAddr(utils.EcdsaToPrivateStr(privateKey))
	if err != nil {
		return nil, err
	}
	var (
		body *types2.TransactionBody
	)
	switch method {
	case constant.DataUpload:
		reqD := req.(*types2.DataUploadReq)
		nonce, err := c.GetNonce(ctx, &types2.GetNonceReq{
			PublicKeyStr:   PublicKey,
			BitcoinAddress: reqD.From,
		})
		if err != nil {
			return nil, err
		}
		log.Debug("client.tx", "method", "newTransactionBody", "NonceAt", nonce)
		body = &types2.TransactionBody{
			BitcoinAddress: reqD.From,
			MethodName:     reqD.MethodName,
			Nonce:          nonce.Nonce,
			Payload: &types2.DataPayload{
				Data:        reqD.RawData,
				Labels:      reqD.Labels,
				NamespaceID: reqD.NID,
			},
			PublicKey:  PublicKey,
			StorageFee: int(reqD.StorageFee),
		}
	case constant.CreateNamespace:
		reqD := req.(*types2.CreateNameSpaceReq)
		nonce, err := c.GetNonce(ctx, &types2.GetNonceReq{
			PublicKeyStr:   PublicKey,
			BitcoinAddress: reqD.From,
		})
		if err != nil {
			return nil, err
		}
		body = &types2.TransactionBody{
			BitcoinAddress: reqD.From,
			MethodName:     reqD.MethodName,
			Nonce:          nonce.Nonce,
			Payload: &types2.CreateNameSpacePayload{
				Admins:     reqD.Admins,
				Name:       reqD.Name,
				Owner:      reqD.Owner,
				Permission: reqD.Permission,
			},
			PublicKey:  PublicKey,
			StorageFee: int(reqD.StorageFee),
		}
	case constant.UpdateNamespace:
		reqD := req.(*types2.UpdateNamespaceReq)
		nonce, err := c.GetNonce(ctx, &types2.GetNonceReq{
			PublicKeyStr:   PublicKey,
			BitcoinAddress: reqD.From,
		})
		if err != nil {
			return nil, err
		}
		body = &types2.TransactionBody{
			BitcoinAddress: reqD.From,
			MethodName:     reqD.MethodName,
			Nonce:          nonce.Nonce,
			Payload: &types2.UpdateNameSpacePayload{
				Admins:      reqD.Admins,
				Name:        reqD.Name,
				NamespaceID: reqD.NID,
				Permission:  reqD.Permission,
			},
			PublicKey:  PublicKey,
			StorageFee: int(reqD.StorageFee),
		}
	}
	return body, nil
}
