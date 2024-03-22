package types

type GetBlockHeaderByNumberReq struct {
	BlockNumber int64 `json:"blockNumber,omitempty"`
}

type GetBlockHeaderByNumberRsp struct {
	BlockHash      string         `json:"blockHash,omitempty"`
	ParentHash     string         `json:"parentHash,omitempty"`
	Timestamp      int64          `json:"timestamp,omitempty"`
	Size           int64          `json:"size,omitempty"`
	TransactionIDs TransactionIDs `json:"transactionIDs,omitempty"`
}

type GetLatestBlockNumberRsp struct {
	BlockNumber int64 `json:"blockNumber,omitempty"`
}

type GetNonceReq struct {
	PublicKeyStr   string `json:"publicKey,omitempty"`
	BitcoinAddress string `json:"bitcoinAddress,omitempty"`
}

type GetNonceRsp struct {
	Nonce     int    `json:"nonce,omitempty"`
	ErrorInfo string `json:"errorInfo,omitempty"`
	IsSuccess bool   `json:"isSuccess,omitempty"`
}
