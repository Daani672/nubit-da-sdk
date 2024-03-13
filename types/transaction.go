package types

type QueryTxReq struct {
	TxID string `json:"TxID,omitempty"`
}

type TransactionIDs struct {
	Data      []string `json:"data,omitempty"`
	Namespace []string `json:"namespace,omitempty"`
}

type TxReq struct {
	TransactionBody *TransactionBody `json:"transactionBody,omitempty"`
	Invoice         string           `json:"invoice,omitempty"`
	PreImage        string           `json:"preImage,omitempty"`
	Signature       string           `json:"signature,omitempty"`
}

type TransactionBody struct {
	BitcoinAddress string `json:"bitcoinAddress,omitempty"`
	MethodName     string `json:"methodName,omitempty"`
	Nonce          int    `json:"nonce"`
	Payload        any    `json:"payload"`
	PublicKey      string `json:"publicKey"`
	StorageFee     int    `json:"storageFee"`
}

type TxRsp struct {
	TxID    string `json:"transactionID,omitempty"`
	Msg     string `json:"msg,omitempty"`
	Success bool   `json:"success,omitempty"`
}

type GetTransactionReq struct {
	TxID string `json:"transactionID,omitempty"`
}

type GetTransactionRsp struct {
	From          string `json:"from"`
	Nonce         int64  `json:"nonce"`
	BlockHash     string `json:"blockHash"`
	BlockNumber   uint64 `json:"blockNumber"`
	StorageFee    uint64 `json:"storageFee"`
	DAID          string `json:"dataID"`
	Datasize      int    `json:"datasize"`
	NID           string `json:"namespaceID"`
	TransactionID string `json:"transactionID"`
	Timestamp     uint64 `json:"timestamp"`
	MethodName    string `json:"methodName"`
	Status        string `json:"status"`
	ContentType   string `json:"contentType"`

	RawData string `json:"data"`
}

// /indexer/v1/getTransactions

type GetTransactionsReq struct {
	Limit   int                        `json:"limit,omitempty"`
	Offset  int                        `json:"offset,omitempty"`
	Filters *GetTransactionsReqFilters `json:"filters,omitempty"`
}

type GetTransactionsReqFilters struct {
	BlockNumber int    `json:"blockNumber,omitempty"`
	From        string `json:"from,omitempty"`
}

type GetTransactionsRsp struct {
	Transactions []struct {
		From          string `json:"from,omitempty"`
		Nonce         int    `json:"nonce,omitempty"`
		BlockHash     string `json:"blockHash,omitempty"`
		BlockNumber   int    `json:"blockNumber,omitempty"`
		StorageFee    int    `json:"storageFee,omitempty"`
		DataID        string `json:"dataID,omitempty"`
		NamespaceID   string `json:"namespaceID,omitempty"`
		Timestamp     int    `json:"timestamp,omitempty"`
		MethodName    string `json:"methodName,omitempty"`
		ContentType   string `json:"contentType,omitempty"`
		Status        string `json:"status,omitempty"`
		TransactionID string `json:"transactionID,omitempty"`
	} `json:"transactions,omitempty"`
	LastOffset int `json:"lastOffset,omitempty"`
}
