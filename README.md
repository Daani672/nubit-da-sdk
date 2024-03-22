# nubit-da-sdk
nubit-da-sdk is a software package for programmatically accessing Nubit Data Availability Layer. 

**Warning**
This release is specifically for the Pre-alpha Testnet and may include changes that are not backward compatible in the future.

## Usage
```go
package main

import (
	"context"
	"fmt"

	sdk "github.com/RiemaLabs/nubit-da-sdk"
	"github.com/RiemaLabs/nubit-da-sdk/constant"
	"github.com/RiemaLabs/nubit-da-sdk/types"
)

func main() {
	// Initialize context and SDK settings
	ctx := context.Background()
	// Set network to PreAlphaTestNet
	sdk.SetNet(constant.PreAlphaTestNet)
	client := sdk.NewNubit(sdk.WithCtx(ctx),
		sdk.WithInviteCode("YouInviteCode"),
		// We won't record your PK.
		sdk.WithPrivateKey("YourPrivateKey"))
	    // Set Nubit Node RPC address
		sdk.WithRpc("{rpc}")
	if client == nil {
		panic("client is nil") // Panic if the client creation fails
	}
	// Create a namespace
	ns, err := client.CreateNamespace("test", "Private", "OwnerAddress", []string{"AdminAddress", "AdminAddress"})
	if err != nil {
		panic(err) // Print created namespace information
	}
	fmt.Println("namespace:", ns)
	// Query transaction details
	transaction, err := client.Client.GetTransaction(ctx, &types.GetTransactionReq{
		TxID: ns.TxID,
	})
	if err != nil {
		fmt.Println(err)// Print error message if query fails
		return
	}

	fmt.Println("transaction:", transaction)
	// Upload file
	// If the fee is set to 0, it will be obtained automatically.
	upload, err := client.Upload("FilePath", transaction.NID,0)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("upload:", upload)

}


```