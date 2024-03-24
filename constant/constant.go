package constant

import (
	"github.com/golang-module/carbon"
)

const (
	TestNet         string = "TestNet"
	PreAlphaTestNet string = "Pre-Alpha Testnet"
)

// APPBuild ...
var APPBuild string = carbon.Now().ToShortDateString()

var NubitNet = TestNet

var NubitRpc = "https://pre-alpha.indexer.nubit.org"

var ProxyRpc = "https://pre-alpha.indexer.nubit.org"

var NubitTestRpc = "https://test.indexer.nubit.network"

var ProxyTestRpc = "https://test.indexer.nubit.network"

var NubitLndProxy = "pre-alpha.lnd.nubit.org:50051"

const (
	Text  = "text/plain"
	Video = "video/mp4"
	Image = "image/jpeg"
	JSON  = "application/json"
)

var ContextTypes = []string{
	Text, Video, Image, JSON,
}
