package constant

import (
	"github.com/golang-module/carbon"
)

const (
	TestNet          string = "TestNet"
	MainNet          string = "MainNet"
	DefaultNamespace        = "0x00000000"
	BTCMainNet              = "mainnet"
	BTCTestNet              = "testnet"

	LevelError   = "Error"
	LevelWarn    = "Warn"
	LevelDebug   = "Debug"
	LevelVerbose = "Verbose"
)

// APPVersion ...
var APPVersion string = "v1.0"

// APPBuild ...
var APPBuild string = carbon.Now().ToShortDateString()

var LogLevel = LevelDebug

var NubitNet = TestNet

var BitcoinNet = BTCTestNet

var NubitRpc = "https://pre-alpha.indexer.nubit.org"

var ProxyRpc = "https://pre-alpha.indexer.nubit.org"

var NubitTestRpc = "https://pre-alpha.api.nubit.network"

var ProxyTestRpc = "https://pre-alpha.api.nubit.network"

var NubitLndProxy = "pre-alpha.lnd.nubit.org:50051"

const NubitHomeDir = ".nubit"

const NubitConfigFile = "config.json"

const DefaultWalletPwd = "NubitDAChain"

const DefaultWalletName = "nubitDAChain.key"

const (
	Text  = "text/plain"
	Video = "video/mp4"
	Image = "image/jpeg"
)

var ContextTypes = []string{
	Text, Video, Image,
}
