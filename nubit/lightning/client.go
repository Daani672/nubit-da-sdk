package lightning

import (
	"context"
	"crypto/tls"
	"crypto/x509"
	"errors"
	"io"
	"io/ioutil"
	"strconv"
	"strings"

	"github.com/lightningnetwork/lnd/lnrpc"
	pb "github.com/lightningnetwork/lnd/lnrpc/routerrpc"
	"github.com/lightningnetwork/lnd/macaroons"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/metadata"
	macaroon "gopkg.in/macaroon.v2"

	"github.com/RiemaLabs/nubit-da-sdk/log"
	"github.com/RiemaLabs/nubit-da-sdk/types"
)

type Client struct {
	c   *grpc.ClientConn
	ctx context.Context
}

func NewClient(args *types.PaymentParams) *Client {
	var (
		credential macaroons.MacaroonCredential
		creds      credentials.TransportCredentials
		opts       []grpc.DialOption
	)
	ctx := context.TODO()
	creds = credentials.NewTLS(&tls.Config{})
	var target = args.Target
	switch true {
	case args.XAPIKEY != "":
		target = args.LndProxyTarget
		log.Debug("lightning", "NewClient", "x-api-key", "Target", target)
		ctx = metadata.AppendToOutgoingContext(ctx, "X-API-KEY", args.XAPIKEY)
		opts = append(opts, grpc.WithTransportCredentials(insecure.NewCredentials()))
	case args.Authorization != "":
		target = args.LndProxyTarget
		log.Debug("lightning", "NewClient", "authorization", "Target", target)
		ctx = metadata.AppendToOutgoingContext(ctx, "Authorization", args.Authorization)
		opts = append(opts, grpc.WithTransportCredentials(insecure.NewCredentials()))
	case args.MacaroonFile != "":
		log.Debug("lightning", "NewClient", "MacaroonFile")
		macaroonBytes, err := ioutil.ReadFile(args.MacaroonFile)
		if err != nil {
			log.Error("lightning", "ReadFile", err)
			return nil
		}
		mac := &macaroon.Macaroon{}
		if err = mac.UnmarshalBinary(macaroonBytes); err != nil {
			log.Error("lightning", "UnmarshalBinary", err)
			return nil
		}
		credential, err = macaroons.NewMacaroonCredential(mac)
		if err != nil {
			log.Error("lightning", "NewMacaroonCredential", err)
			return nil
		}
		if strings.TrimSpace(args.TlsCertFile) != "" {
			caCert, err := ioutil.ReadFile(args.TlsCertFile)
			if err != nil {
				panic(err)
			}
			caCertPool := x509.NewCertPool()
			caCertPool.AppendCertsFromPEM(caCert)
			creds = credentials.NewClientTLSFromCert(caCertPool, args.Target)
		}
		opts = append(opts, grpc.WithDefaultServiceConfig(`{"loadBalancingPolicy":"round_robin"}`), grpc.WithTransportCredentials(creds), grpc.WithPerRPCCredentials(credential))
	}

	_conn, err := grpc.DialContext(ctx, target, opts...)
	if err != nil {
		log.Error("lightning", "Dial.err", err, "uri", target)
		return nil
	}
	return &Client{_conn, ctx}
}

func (c *Client) Payment(ctx context.Context, invoice string, fee int64) (*types.PaymentStatus, error) {
	nCtx := c.ctx
	nCtx = metadata.AppendToOutgoingContext(nCtx, "invoice", invoice)
	if c == nil || c.c == nil {
		return nil, errors.New("lightning NewClient nil client")
	}
	defer c.c.Close()
	r_client := pb.NewRouterClient(c.c)
	stream, err := r_client.SendPaymentV2(nCtx, &pb.SendPaymentRequest{
		PaymentRequest: invoice,
		TimeoutSeconds: 60,
		FeeLimitSat:    fee,
	})
	if err != nil {
		return nil, err
	}
	var status *types.PaymentStatus
	loading := 0
	for {
		update, err := stream.Recv()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Error("lightning", "SendPaymentV2", err)
			return nil, err
		}
		preimage := update.GetPaymentPreimage()
		payment_status := lnrpc.Payment_PaymentStatus_name[int32(update.Status)]
		failure_reason := lnrpc.PaymentFailureReason_name[int32(update.FailureReason)]
		log.Debug("lightning", "lightning failure reason", failure_reason)
		log.Debug("lightning", "lightning status", payment_status)
		log.Debug("lightning", "lightning preimage", preimage)
		status = &types.PaymentStatus{
			Preimage:  preimage,
			Hash:      update.PaymentHash,
			Amount:    strconv.FormatInt(update.ValueMsat, 10),
			Success:   update.Status == lnrpc.Payment_SUCCEEDED,
			ErrorCode: int(update.FailureReason),
			ErrorMsg:  failure_reason,
		}
		if update.Status != lnrpc.Payment_SUCCEEDED {
			loading++
			if loading > 10 {
				return status, err
			}
			continue
		}
		break
	}
	return status, nil
}
