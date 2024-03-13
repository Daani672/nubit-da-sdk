package main

import (
	"context"
	"fmt"

	sdk "github.com/RiemaLabs/nubit-da-sdk"
	"github.com/RiemaLabs/nubit-da-sdk/constant"
	"github.com/RiemaLabs/nubit-da-sdk/types"
)

func main() {
	ctx := context.Background()
	sdk.SetNet(constant.MainNet)
	client := sdk.NewNubit(sdk.WithCtx(ctx),
		sdk.WithInviteCode("7mkEPWPBBrMr12WKNsL2UALvqYfbox"),
		sdk.WithPrivateKey("9541ea760acc451684d28033566379a95bfe5a1b4da4a56a7df6055e4fa93eac"))
	if client == nil {
		panic("client is nil")
	}
	ns, err := client.CreateNamespace("test", "Private", "1JqocHHUBgLKZxzQpCqrrzMnV6QV4XrUJr", []string{"18JTw53V9MMtGax7es3GMPQHwjpjNFyPj1", "1JqocHHUBgLKZxzQpCqrrzMnV6QV4XrUJr"})
	if err != nil {
		panic(err)
	}
	fmt.Println("namespace:", ns)

	transaction, err := client.Client.GetTransaction(ctx, &types.GetTransactionReq{
		TxID: ns.TxID,
	})
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println("transaction:", transaction)

	upload, err := client.Upload("/Users/{USER}/Documents/RiemaLabs/nubit-da-sdk/test/main.go", transaction.NID)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println("upload:", upload)

}
