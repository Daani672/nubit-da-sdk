package nubit_da_sdk

import (
	"context"
	"encoding/base64"
	"errors"
	"fmt"
	"io"
	"os"
	"strings"

	"github.com/RiemaLabs/nubit-da-sdk/constant"
	"github.com/RiemaLabs/nubit-da-sdk/log"
	"github.com/RiemaLabs/nubit-da-sdk/nubit/client"
	"github.com/RiemaLabs/nubit-da-sdk/types"
	"github.com/RiemaLabs/nubit-da-sdk/utils"
	"github.com/golang-module/carbon"
)

func init() {
	log.SetLevel(log.LevelDebug)
	log.SetVerion("v1.0", carbon.Now().ToShortDateString())
}

type Opt func(c *Option)

type Option struct {
	ctx           context.Context
	PaymentParams *types.PaymentParams
	privateKey    string
}
type NubitSDK struct {
	Client *client.Client
	Opts   *Option
}

func SetNet(netstr string) {
	constant.NubitNet = netstr
}

func WithPrivateKey(privateKey string) Opt {
	return func(c *Option) {
		c.privateKey = privateKey
	}
}

func WithCtx(ctx context.Context) Opt {
	return func(c *Option) {
		c.ctx = ctx
	}
}

func WithInviteCode(code string) Opt {
	return func(c *Option) {
		c.PaymentParams.XAPIKEY = code
		c.PaymentParams.LndProxyTarget = constant.NubitLndProxy
	}
}
func WithLndProxy(target string) Opt {
	return func(c *Option) {
		c.PaymentParams.LndProxyTarget = target
	}
}

func WithMacaroonFile(file string) Opt {
	return func(c *Option) {
		c.PaymentParams.MacaroonFile = file
	}
}

func WithLndTarget(target string) Opt {
	return func(c *Option) {
		c.PaymentParams.Target = target
	}
}

func WithApprove(Approve int) Opt {
	return func(c *Option) {
		c.PaymentParams.Approve = Approve
	}
}

func NewNubit(opts ...Opt) *NubitSDK {
	sdk := &NubitSDK{
		Opts: &Option{
			ctx:           client.Background(),
			PaymentParams: &types.PaymentParams{},
		},
	}
	var (
		ctx client.Context
		err error
	)
	for _, op := range opts {
		op(sdk.Opts)
	}
	ctx = client.Background()
	sdk.Client, err = client.Dial(ctx)
	if err != nil {
		return sdk
	}
	return sdk
}

func (sdk *NubitSDK) Upload(filePath string, nid string) (target *types.DataUploadRsp, err error) {
	file, err := os.OpenFile(filePath, os.O_RDONLY, os.ModePerm)
	if err != nil {
		return nil, err
	}
	defer file.Close()
	all, err := io.ReadAll(file)
	if err != nil {
		return nil, err
	}
	var Labels map[string]interface{}
	Labels = make(map[string]interface{})
	mty := utils.MIMEType(filePath)
	for k, v := range mty {
		Labels[k] = v
	}
	for k, v := range Labels {
		if !strings.EqualFold(k, "contentType") {
			continue
		}
		if !utils.CheckContentType(v.(string)) {
			return nil, fmt.Errorf("file type %s is not supported", v.(string))
		}
	}
	if nid == "" {
		nid = constant.DefaultNamespace
	}

	btc := utils.PrivateStrToBtcAddress(sdk.Opts.privateKey)

	data := &types.DataUploadReq{
		NID:        nid,
		From:       btc,
		RawData:    base64.StdEncoding.EncodeToString(all),
		Labels:     Labels,
		MethodName: constant.DataUpload,
	}
	fee, err := sdk.Client.GetEstimateFee(sdk.Opts.ctx, data, utils.PrivateStrToEcdsa(sdk.Opts.privateKey), constant.DataUpload, nid)
	if err != nil {
		return nil, err
	}
	data.StorageFee = uint64(fee.StorageFee)
	return sdk.Client.Upload(sdk.Opts.ctx, data, utils.PrivateStrToEcdsa(sdk.Opts.privateKey), sdk.Opts.PaymentParams)
}

func (sdk *NubitSDK) GetEstimateFee(req any, method string, nid string) (data *types.EstimateFee, err error) {
	if nid == "" {
		nid = constant.DefaultNamespace
	}
	return sdk.Client.GetEstimateFee(sdk.Opts.ctx, req, utils.PrivateStrToEcdsa(sdk.Opts.privateKey), method, nid)
}

func (sdk *NubitSDK) CreateNamespace(name, permission, owner string, admins []string) (data *types.DataUploadRsp, err error) {
	btc := utils.PrivateStrToBtcAddress(sdk.Opts.privateKey)
	if strings.EqualFold(owner, "") {
		owner = btc
	}
	if strings.EqualFold(name, "") {
		return nil, errors.New("The name cannot be empty")
	}
	if len(admins) == 0 {
		admins = append(admins, btc)
	}
	return sdk.Client.CreateNamespace(sdk.Opts.ctx, &types.CreateNameSpaceReq{
		From:            btc,
		Name:            name,
		Permission:      permission,
		MethodName:      constant.CreateNamespace,
		AvailableHeight: 0,
		Owner:           owner,
		Admins:          admins,
	}, utils.PrivateStrToEcdsa(sdk.Opts.privateKey), sdk.Opts.PaymentParams)
}

func (sdk *NubitSDK) UpdateNamespace(nid, name, permission, owner string, admins []string) (data *types.DataUploadRsp, err error) {
	btc := utils.PrivateStrToBtcAddress(sdk.Opts.privateKey)
	if strings.EqualFold(owner, "") {
		owner = btc
	}
	if strings.EqualFold(nid, "") {
		return nil, errors.New("namespace-id is empty")
	}
	namespace, err := sdk.Client.GetNamespace(sdk.Opts.ctx, &types.GetNamespaceReq{
		NID: nid,
	})
	if name == "" {
		name = namespace.Name
	}
	if permission == "" {
		permission = namespace.Permission
	}
	if len(admins) == 0 {
		admins = namespace.Admins
	}
	return sdk.Client.UpdateNamespace(sdk.Opts.ctx, &types.UpdateNamespaceReq{
		From:       btc,
		PrikeyStr:  sdk.Opts.privateKey,
		NID:        nid,
		Name:       name,
		Admins:     admins,
		Owner:      namespace.Owner,
		Permission: permission,
		StorageFee: 100,
		MethodName: constant.UpdateNamespace,
	}, utils.PrivateStrToEcdsa(sdk.Opts.privateKey), sdk.Opts.PaymentParams)
}