package types

type (
	DataUploadReq struct {
		DataID     string                 `json:"dataID,omitempty"`
		NID        string                 `json:"namespaceID,omitempty"`
		From       string                 `json:"from,omitempty"`
		RawData    string                 `json:"data,omitempty"`
		Labels     map[string]interface{} `json:"labels,omitempty"`
		StorageFee uint64                 `json:"storageFee,omitempty"`
		TxHash     string                 `gorm:"column:txhash" json:"txHash,omitempty"`

		MethodName string `json:"methodName,omitempty"`

		Timestamp uint64 `gorm:"column:timestamp" json:"timestamp,omitempty"`
	}

	DataPayload struct {
		Data        string                 `json:"data"`
		Labels      map[string]interface{} `json:"labels,omitempty"`
		NamespaceID string                 `json:"namespaceID"`
	}

	DataUploadMockReq struct {
		PrikeyStr string `json:"prikeyStr,omitempty"`

		DataID        string                 `json:"dataID,omitempty"`
		NID           string                 `json:"namespaceID,omitempty"`
		From          string                 `json:"from,omitempty"`
		Data          []byte                 `json:"data,omitempty"`
		Labels        map[string]interface{} `json:"labels,omitempty"`
		StorageFee    uint64                 `json:"storageFee,omitempty"`
		LightingProof []byte                 `json:"lightingProof,omitempty"`
		TxHash        string                 `gorm:"column:txhash" json:"txHash,omitempty"`

		MethodName string `json:"methodName,omitempty"`
	}

	Data struct {
		ID            string                 `json:"ID,omitempty"`
		NID           uint8                  `json:"NID,omitempty"`
		Data          []byte                 `json:"data,omitempty"`
		Labels        map[string]interface{} `json:"labels,omitempty"`
		StorageFee    uint64                 `json:"storageFee,omitempty"`
		LightingProof []byte                 `json:"lightingProof,omitempty"`
		TxHash        string                 `gorm:"column:txhash" json:"txHash,omitempty"`

		MethodName string `json:"methodName,omitempty"`
	}

	DataUploadRsp struct {
		TxID      string `json:"TxID,omitempty"`
		Timestamp int64  `json:"timestamp,omitempty"`
	}

	GetDataReq struct {
		DAID string `json:"dataID,omitempty"`
	}

	GetDataStatusReq struct {
		DAID string `json:"dataID,omitempty"`
	}

	GetDataByUserReq struct {
		Address string `json:"address,omitempty"`
		Limit   int    `json:"limit,omitempty"`
		Offset  int    `json:"offset,omitempty"`
	}

	GetDataByUserRsp struct {
		DataIDs   []string `json:"dataIDs,omitempty"`
		LastIndex int      `json:"lastIndex,omitempty"`
	}

	GetDataCallRsp struct {
		CallData string `json:"calldata,omitempty"`
	}

	GetDAIDRsp struct {
		DAID string `json:"DAID,omitempty"`
	}

	GetNIDRsp struct {
		NID string `json:"NID,omitempty"`
	}

	GetDataRsp struct {
		DataID     string                 `json:"dataID,omitempty" yaml:"dataID,omitempty"`
		NID        string                 `json:"namespaceID,omitempty" yaml:"namespaceID,omitempty"`
		From       string                 `json:"from,omitempty" yaml:"from,omitempty"`
		RawData    string                 `json:"data,omitempty" yaml:"rawData,omitempty"`
		Labels     map[string]interface{} `json:"labels,omitempty" yaml:"labels,omitempty"`
		StorageFee uint64                 `json:"storageFee,omitempty" yaml:"storageFee,omitempty"`
		TxHash     string                 `json:"txHash,omitempty" yaml:"txHash,omitempty"`
		MethodName string                 `json:"methodName,omitempty" yaml:"methodName,omitempty"`
		Timestamp  uint64                 `json:"timestamp,omitempty" yaml:"timestamp,omitempty"`
	}

	GetDataStatusRsp struct {
		Status string `json:"status,omitempty"`
	}
)
