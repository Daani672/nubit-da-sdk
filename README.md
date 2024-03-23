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

```go
go mod tidy
```

### 2. Initialize SDK
To start using the `nubit-da-sdk`, create a new instance and set it up with your network preferences, invitation code, and private key:

```go
// Initialize context and SDK settings
ctx := context.Background()
// Set network to mainnet
sdk.SetNet(constant.MainNet)
// Replace "your_invite_code" and "your_private_key" with actual values
client := sdk.NewNubit(sdk.WithCtx(ctx),
    sdk.WithInviteCode("your_invite_code"),
    sdk.WithPrivateKey("your_private_key"))
if client == nil {
    panic("client is nil") // Panic if the client creation fails
}
```

### 3. Create a Namespace
Namespaces are essential in nubit-da-sdk for organizing your data. Here's how to create one:

```go
// Replace "namespace_name" and "PrivacySetting" with actual values
// "PrivacySetting" should be either "Public" or "Private"
// "owners_address" should be the wallet address of the namespace owner
// "additional_admins" can be an array of addresses who can administer the namespace
ns, err := client.CreateNamespace("namespace_name", "PrivacySetting", "owners_address", []string{"additional_admins"})
if err != nil {
    panic(err) // Handle the error appropriately
}
fmt.Println("Created namespace:", ns)
```

### 4. Upload Data to Namespace
Once you have a namespace, you can start uploading data to it:
```go
// The path to the file you wish to upload
filePath := "/path/to/your/file"
// The namespace ID where you wish to upload the file
namespaceID := ns.ID
// Replace "0" with the storage fee if you wish to specify it
// Using "0" will automatically calculate the necessary fee
upload, err := client.Upload(filePath, namespaceID, 0)
if err != nil {
    panic(err) // Handle the error appropriately
}
fmt.Println("Uploaded data:", upload)
```

git