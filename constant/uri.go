package constant

const (
	GetData                = "GetData"
	GetDataStatus          = "GetDataStatus"
	CountDatasize          = "CountDatasize"
	GetBlockHeaderByNumber = "GetBlockHeaderByNumber"
	GetLatestBlockNumber   = "GetLatestBlockNumber"
	GetNamespace           = "GetNamespace"
	GetNamespaceStatus     = "GetNamespaceStatus"
	GetLastNamespaceID     = "GetLastNamespaceID"
	GetDataInNamespace     = "GetDataInNamespace"
	SubmitTransaction      = "SubmitTransaction"
	GetTransaction         = "GetTransaction"
	GetTransactions        = "GetTransactions"
	GetTxIDByDataID        = "GetTxIDByDataID"
	CountTransactions      = "CountTransactions"
	CountNamespace         = "CountNamespace"
	GetInvoice             = "GetInvoice"
	GetPayee               = "GetPayee"
	GetEstimateFee         = "GetEstimateFee"
	GetDAIDByTxID          = "GetDAIDByTxID"
	GetNIDByTxID           = "GetNIDByTxID"
	GetDataByDAID          = "GetDataByDAID"
	GetDataIDByUser        = "GetDataIDByUser"
	GetNamespaceByUser     = "GetNamespaceByUser"
	CountUsers             = "CountUsers"
	GetNamespaces          = "GetNamespaces"

	GetNonce = "GetNonce"

	GetDatas = "GetDatas"
)

var MocApi = map[string]string{
	// data
	GetData:         "/v1/data",
	GetDataStatus:   "/v1/data/status",
	GetDAIDByTxID:   "/indexer/v1/getDataIDByTxID",
	GetDataByDAID:   "/v1/data/getDataByDAID",
	GetDataIDByUser: "/v1/data/getDataByUser",
	CountDatasize:   "/indexer/v1/countDatasize",
	GetDatas:        "/indexer/v1/getDatas",
	// block
	GetBlockHeaderByNumber: "/v1/block/number",
	GetLatestBlockNumber:   "/v1/block/latest",
	// Namespace
	GetNamespace:       "/v1/namespace",
	GetNamespaceStatus: "/v1/namespace/status",
	GetLastNamespaceID: "/v1/namespace/latest",
	GetDataInNamespace: "/v1/namespace/data",
	GetNIDByTxID:       "/v1/chain/getNIDByTxID",
	GetNamespaceByUser: "/v1/namespace/getNamespaceByUsers",
	CountNamespace:     "/indexer/v1/countNamespace",
	GetNamespaces:      "/indexer/v1/getNamespaces",

	// Transaction
	SubmitTransaction: "/v1/submitTransaction",
	GetTransaction:    "/v1/transaction",
	GetTransactions:   "/indexer/v1/getTransactions",
	CountTransactions: "/indexer/v1/countTransactions",
	GetTxIDByDataID:   "/indexer/v1/getTxIDByDataID",

	//Invoice
	GetInvoice:     "/v1/getInvoice",
	GetPayee:       "/v1/getPayee",
	GetEstimateFee: "/v1/getEstimateFee",

	// user
	CountUsers: "/indexer/v1/countUsers",
	// chain
	GetNonce: "/v1/getNonce",
}

var Release = map[string]string{
	// data
	GetData:         "/v1/data",
	GetDataStatus:   "/v1/data/status",
	GetDAIDByTxID:   "/indexer/v1/getDataIDByTxID",
	GetDataByDAID:   "/v1/data/getDataByDAID",
	GetDataIDByUser: "/v1/data/getDataByUser",
	CountDatasize:   "/indexer/v1/countDatasize",
	GetDatas:        "/indexer/v1/getDatas",
	// block
	GetBlockHeaderByNumber: "/v1/block/number",
	GetLatestBlockNumber:   "/v1/block/latest",
	// Namespace
	GetNamespace:       "/v1/namespace",
	GetNamespaceStatus: "/v1/namespace/status",
	GetLastNamespaceID: "/v1/namespace/latest",
	GetDataInNamespace: "/v1/namespace/data",
	GetNIDByTxID:       "/v1/chain/getNIDByTxID",
	GetNamespaceByUser: "/v1/namespace/getNamespaceByUsers",
	CountNamespace:     "/indexer/v1/countNamespace",
	GetNamespaces:      "/indexer/v1/getNamespaces",

	// Transaction
	SubmitTransaction: "/v1/submitTransaction",
	GetTransaction:    "/v1/transaction",
	GetTransactions:   "/indexer/v1/getTransactions",
	CountTransactions: "/indexer/v1/countTransactions",
	GetTxIDByDataID:   "/indexer/v1/getTxIDByDataID",

	//Invoice
	GetInvoice:     "/v1/getInvoice",
	GetPayee:       "/v1/getPayee",
	GetEstimateFee: "/v1/getEstimateFee",

	// user
	CountUsers: "/indexer/v1/countUsers",
	// chain
	GetNonce: "/v1/getNonce",
}
