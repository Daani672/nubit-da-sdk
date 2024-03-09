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

	PaymentParams struct {
		Target        string `json:"target,omitempty" yaml:"Target,omitempty"`
		MacaroonFile  string `json:"macaroonFile,omitempty" yaml:"MacaroonFile,omitempty"`
		TlsCertFile   string `json:"tlsCertFile,omitempty" yaml:"TlsCertFile,omitempty"`
		Authorization string `json:"authorization,omitempty" yaml:"Authorization,omitempty"`
		XAPIKEY       string `json:"x-api-key,omitempty" yaml:"x-api-key,omitempty"`
	}
)
