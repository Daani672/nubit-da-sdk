package types

type (
	PaymentStatus struct {
		Preimage  string `json:"preimage,omitempty" yaml:"Preimage,omitempty"`
		Route     string `json:"route,omitempty" yaml:"Route,omitempty"`
		Hash      string `json:"hash,omitempty" yaml:"Hash,omitempty"`
		Amount    string `json:"amount,omitempty" yaml:"Amount,omitempty"`
		Success   bool   `json:"success,omitempty" yaml:"Success,omitempty"`
		ErrorCode int    `json:"error_code,omitempty" yaml:"ErrorCode,omitempty"`
		ErrorMsg  string `json:"error_msg,omitempty" yaml:"ErrorMsg,omitempty"`
	}
	PaymentRequest struct {
		Invoice string `json:"invoice,omitempty" yaml:"Invoice,omitempty"`
		Amount  string `json:"amount,omitempty" yaml:"Amount,omitempty"`
	}
	PaymentParams struct {
		Target              string  `json:"target,omitempty" yaml:"Target,omitempty"`
		AutoPay             bool    `json:"auto_pay,omitempty" yaml:"AutoPay,omitempty"`
		Approve             int     `json:"approve,omitempty"`     //Approve
		EstimateFeeMultiple float64 `json:"estimateFee,omitempty"` // EstimateFeeMultiple Pay at this multiple
		MacaroonFile        string  `json:"macaroonFile,omitempty" yaml:"MacaroonFile,omitempty"`
		FeeLimitSat         int     `json:"feeLimitSat,omitempty"`
		TlsCertFile         string  `json:"tlsCertFile,omitempty" yaml:"TlsCertFile,omitempty"`
		Authorization       string  `json:"authorization,omitempty" yaml:"Authorization,omitempty"`
		XAPIKEY             string  `json:"x-api-key,omitempty" yaml:"x-api-key,omitempty"`
		LndProxyTarget      string  `json:"lndProxyTarget,omitempty" yaml:"lndProxyTarget,omitempty"`
	}

	Invoice struct {
		Fee     int    `json:"fee,omitempty"`
		Invoice string `json:"invoice,omitempty"`
	}

	Payee struct {
		PublicKey string `json:"publicKey,omitempty"`
		Address   string `json:"address,omitempty"`
	}

	EstimateFee struct {
		StorageFee int `json:"storageFee,omitempty"`
	}

	EstimateFeeReq struct {
		DataSize    int    `json:"dataSize,omitempty"`
		MethodName  string `json:"methodName,omitempty"`
		NamespaceID string `json:"namespaceID,omitempty"`
	}
)
