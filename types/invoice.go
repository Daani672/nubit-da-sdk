package types

type GetInvoiceReq struct {
	PublicKeyStr   string `json:"publicKey,omitempty" yaml:"publicKeyStr,omitempty"`
	BitcoinAddress string `json:"bitcoinAddress,omitempty" yaml:"bitcoinAddress,omitempty"`
	StorageFee     int64  `json:"storageFee,omitempty" yaml:"storageFee,omitempty"`
}
