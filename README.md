# Nubit-da-sdk [![Join Nubit Discord Community](https://img.shields.io/discord/916984413944967180?logo=discord&style=flat)](https://discord.gg/5sVBzYa4Sg) [![Follow Nubit On X](https://img.shields.io/twitter/follow/nubit_org)](https://twitter.com/Nubit_org)

<img src="assets/logo.svg" width="600px" alt="Nubit Logo" />

## Background
`nubit-da-sdk` equips developers with the necessary tools and libraries to efficiently interact with the Nubit Data Availability (DA) Chain. It streamlines complex blockchain functionalities into easy-to-use operations.

## What is Nubit-da-sdk?
`nubit-da-sdk` is an all-encompassing Golang SDK for the Nubit Data Availability (DA)  Chain. It simplifies the complex processes of blockchain interactions into user-friendly API calls. 
The SDK covers a broad range of blockchain functionalities, from generating wallets to managing data in namespaces. It's designed to streamline blockchain operations for developers, fostering easy development and integration within the Nubit ecosystem. 
This SDK is essential for anyone looking to build applications on the Nubit platform, providing a robust set of tools to efficiently handle various blockchain-related tasks.


## Getting Started
TTo use `nubit-da-sdk`, you will need Golang installed on your system. Follow the steps below to set up and start using the SDK. `Go` version `1.22.0` or later is required for running repository. Please visit [Golang download Page](https://go.dev/doc/install) to get latest Golang installed.

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
Expected Outcome: This script initializes your SDK client. If successful, you'll see "SDK client successfully initialized".

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
Expected Outcome: The script creates a new namespace and returns its ID if successful.

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
Expected Outcome: The script uploads a file to the specified namespace and provides a transaction ID upon successful upload.


## FAQ
- **Q: How do I integrate nubit-da-sdk into my project?**
    - A: Start by ensuring your system has Golang installed. Follow our "Getting Started" guide to integrate nubit-da-sdk into your project. This involves installing dependencies, initializing the SDK, setting up network preferences, and creating namespaces for data organization.

- **Q: How does nubit-da-sdk ensure data security and privacy?**
    - A: nubit-da-sdk employs encryption techniques and offers both public and private namespaces for data handling. Users can choose the appropriate privacy settings for their data, managed securely with private keys.

- **Q: Where can I get help if I encounter issues with nubit-da-sdk?**
    - A: First, ensure you're following the documentation correctly. If the issue persists, seek help through our community forums, GitHub issues, or contact our technical support. We're here to assist with any SDK-related queries or difficulties.

- **Q: How does nubit-da-sdk handle large-scale data uploads and storage?**
    - A: nubit-da-sdk supports efficient data uploading and storage management. You can upload and manage data in the created namespaces, with the storage fees calculated automatically or specified manually based on your needs.

- **Q: Can I use nubit-da-sdk for Lightning Network payments?**
    - A: Yes, nubit-da-sdk supports Lightning Network transactions, enabling fast and efficient micropayments on the blockchain. You can integrate Lightning payment functionalities into your application using the SDK, providing a seamless user experience for transactions.

- **Q: What kind of blockchain operations can I perform with nubit-da-sdk?**
    - A: The SDK is designed for a wide range of blockchain operations, including creating and managing wallets, executing and tracking transactions, managing namespaces for organizing data, and interacting with the Nubit blockchain for various decentralized applications (dApps).