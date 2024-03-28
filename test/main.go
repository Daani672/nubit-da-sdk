package main

import (
	"context"
	"encoding/json"
	"fmt"
	"time"

	sdk "github.com/RiemaLabs/nubit-da-sdk"
	"github.com/RiemaLabs/nubit-da-sdk/constant"
	"github.com/RiemaLabs/nubit-da-sdk/types"
)

func main() {
	ctx := context.Background()
	sdk.SetNet(constant.PreAlphaTestNet)
	client := sdk.NewNubit(sdk.WithCtx(ctx),
		sdk.WithGasCode("{gascode}"),
		sdk.WithPrivateKey("{PrivateKey}"))
	if client == nil {
		panic("client is nil")
	}

	ns, err := client.CreateNamespace("test123", "Private", "{BitCoinAddress}", []string{"{BitCoinAddress}"})
	if err != nil {
		panic(err)
	}
	fmt.Println("\n\n namespace---:", ns)

	time.Sleep(time.Second * 22)
	transaction, err := client.Client.GetTransaction(ctx, &types.GetTransactionReq{
		TxID: ns.TxID,
	})
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println("\n\n transaction:", transaction)

	upload, err := client.Upload("/path/file", transaction.NID, 0)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println("\n upload:", upload)
	time.Sleep(time.Second * 22)
	namespaces, err := client.Client.GetNamespaces(ctx, &types.GetNamespacesReq{Limit: 50, Offset: 0, Filter: struct {
		Owner string `json:"owner,omitempty"`
		Admin string `json:"admin,omitempty"`
	}{
		Owner: "{BitCoinAddress}",
	},
	})
	if err != nil {
		return
	}

	var Nss []string
	if len(namespaces.Namespaces) > 0 {
		for _, ns := range namespaces.Namespaces {
			fmt.Println("namespace:", ns.NamespaceID)
			Nss = append(Nss, ns.NamespaceID)
		}
	}
	fmt.Println("namespace:", Nss)
	datas, err := client.Client.GetDatas(ctx, &types.GetDatasReq{
		NID:         Nss,
		BlockNumber: 0,
	})
	if err != nil {
		return
	}
	marshal, err := json.Marshal(datas)
	if err != nil {
		return
	}
	fmt.Println("\n datas:", string(marshal))
}
