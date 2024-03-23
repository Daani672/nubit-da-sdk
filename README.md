# Nubit-da-sdk [![Join Nubit Discord Community](https://img.shields.io/discord/916984413944967180?logo=discord&style=flat)](https://discord.gg/5sVBzYa4Sg) [![Follow Nubit On X](https://img.shields.io/twitter/follow/nubit_org)](https://twitter.com/Nubit_org)

<img src="assets/logo.svg" width="600px" alt="Nubit Logo" />

## Background
`nubit-da-sdk` offers developers the tools and library support needed to interact with the Nubit Decentralized Autonomous (DA) Chain. It encapsulates a variety of functionalities, from wallet creation to namespace operations, making blockchain operations seamless and efficient.

## What is Nubit-da-sdk?
The `nubit-da-sdk` is a comprehensive Golang SDK designed for ease of use when working with the Nubit DA Chain. It abstracts complex blockchain interactions into simple API calls, enabling rapid development and integration with the Nubit ecosystem.

## Getting Started
To use `nubit-da-sdk`, you will need Golang installed on your system. You can run your own modular Indexer by following the procedure below. `Go` version `1.22.0` is required for running repository. Please visit [Golang download Page](https://go.dev/doc/install) to get latest Golang installed.

### 1. Install Dependencies
Dependencies are managed through Go Modules. To install all required dependencies, navigate to your project directory and run:

```Bash
go mod tidy
```

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
	// Set network to mainnet
	sdk.SetNet(constant.MainNet)
	client := sdk.NewNubit(sdk.WithCtx(ctx),
		sdk.WithInviteCode("7mkEPWPBBrMr12WKNsL2UALvqYfbox"),// Set invite code
		sdk.WithPrivateKey("9541ea760acc451684d28033566379a95bfe5a1b4da4a56a7df6055e4fa93eac")) // Set private key
	if client == nil {
		panic("client is nil") // Panic if the client creation fails
	}
	// Create a namespace
	ns, err := client.CreateNamespace("test", "Private", "1JqocHHUBgLKZxzQpCqrrzMnV6QV4XrUJr", []string{"18JTw53V9MMtGax7es3GMPQHwjpjNFyPj1", "1JqocHHUBgLKZxzQpCqrrzMnV6QV4XrUJr"})
	if err != nil {
		panic(err) // Print created namespace information
	}
	fmt.Println("namespace:", ns)
	// Query transaction details
	transaction, err := client.Client.GetTransaction(ctx, &types.GetTransactionReq{
		TxID: ns.TxID, // Query transaction info by namespace's transaction ID
	})
	if err != nil {
		fmt.Println(err)// Print error message if query fails
		return
	}

	fmt.Println("transaction:", transaction)
	// Upload file
	upload, err := client.Upload("/Users/{USER}/Documents/RiemaLabs/nubit-da-sdk/test/main.go", transaction.NID,0) // If the fee is 0, it will be obtained automatically, but of course it can also be obtained through the client. GetEstimateFee, 0 is recommended
	if err != nil {
		fmt.Println(err)// Print error message if file upload fails
		return
	}
	// Print upload result
	fmt.Println("upload:", upload)

}


```