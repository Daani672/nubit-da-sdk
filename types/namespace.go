package types

type GetNamespaceByUserReq struct {
	Address string `json:"address,omitempty"`
}

type GetNamespaceByUserRsp struct {
	NamespaceIDs []string `json:"namespaceIDs,omitempty"`
}

type CreateNameSpaceReq struct {
	From            string   `json:"from,omitempty"`
	Name            string   `json:"name,omitempty"`
	Permission      string   `json:"permission,omitempty"`
	MethodName      string   `json:"methodName,omitempty"`
	AvailableHeight int      `json:"AvailableHeight,omitempty"`
	Owner           string   `json:"owner,omitempty"`
	Admins          []string `json:"admins,omitempty"`
	StorageFee      int64    `json:"storageFee,omitempty"`
}

type CreateNameSpacePayload struct {
	Admins     []string `json:"admins,omitempty"`
	Name       string   `json:"name,omitempty"`
	Owner      string   `json:"owner,omitempty"`
	Permission string   `json:"permission,omitempty"`
}

type UpdateNameSpaceReq struct {
	NID             int      `json:"NID,omitempty"`
	Name            string   `json:"name,omitempty"`
	Permission      string   `json:"permission,omitempty"`
	MethodName      string   `json:"methodName,omitempty"`
	AvailableHeight int      `json:"AvailableHeight,omitempty"`
	Owner           string   `json:"owner,omitempty"`
	Admins          []string `json:"admins,omitempty"`
}

type UpdateNameSpacePayload struct {
	Admins      []string `json:"admins,omitempty"`
	Name        string   `json:"name,omitempty"`
	NamespaceID string   `json:"namespaceID,omitempty"`
	Permission  string   `json:"permission,omitempty"`
}

type UpdateNamespaceReq struct {
	PrikeyStr string `json:"prikeyStr,omitempty"`

	From            string   `json:"from,omitempty"`
	NID             string   `json:"namespaceID,omitempty"`
	Name            string   `json:"name,omitempty"`
	Admins          []string `json:"admins,omitempty"`
	Owner           string   `json:"owner,omitempty"`
	Permission      string   `json:"permission,omitempty"`
	StorageFee      uint64   `json:"storageFee,omitempty"`
	AvailableHeight int64    `json:"availableHeight,omitempty"`

	BlockNumber      int64 `json:"blockNumber,omitempty"`
	TransactionIndex uint  `json:"transactionIndex,omitempty"`

	MethodName string `json:"methodName,omitempty"`
}

type TransferNamespaceReq struct {
	PrikeyStr string `json:"prikeyStr,omitempty"`

	NID        string `json:"namespaceID,omitempty"`
	StorageFee uint64 `json:"storageFee,omitempty"`
	Owner      string `json:"owner,omitempty"`

	MethodName string `json:"methodName,omitempty"`
}

type TransferNamespacePayload struct {
	NamespaceID string `json:"namespaceID,omitempty"`
	Owner       string `json:"owner,omitempty"`
}

type GetNamespaceReq struct {
	NID string `json:"namespaceID,omitempty"`
}

type GetNamespaceStatusReq struct {
	NID string `json:"namespaceID,omitempty"`
}

type GetLastNamespaceIDReq struct {
}

type GetDataInNamespaceReq struct {
	NID    string `json:"namespaceID,omitempty"`
	Limit  int    `json:"limit,omitempty"`
	Offset int    `json:"offset,omitempty"`
}

type GetNamespaceRsp struct {
	NID             string   `json:"namespaceID,omitempty" yaml:"namespaceID,omitempty"`
	Name            string   `json:"name,omitempty" yaml:"name,omitempty"`
	Admins          []string `json:"admins,omitempty" yaml:"admins,omitempty"`
	Owner           string   `json:"owner,omitempty" yaml:"owner,omitempty"`
	Permission      string   `json:"permission,omitempty" yaml:"permission,omitempty"`
	StorageFee      uint64   `json:"storageFee,omitempty" yaml:"storageFee,omitempty"`
	AvailableHeight int64    `json:"availableHeight,omitempty" yaml:"availableHeight,omitempty"`
	BlockNumber     int64    `json:"blockNumber,omitempty" yaml:"blockNumber,omitempty"`
	MethodName      string   `json:"methodName,omitempty" yaml:"methodName,omitempty"`
}

type GetNamespaceStatusRsp struct {
	Status string `json:"status,omitempty"`
}

type GetLastNamespaceIDRsp struct {
	NID string `json:"namespaceID,omitempty"`
}

type GetDataInNamespaceRsp struct {
	DataIDs    []string `json:"dataIDs,omitempty"`
	LastOffset int64    `json:"lastOffset,omitempty"`
}

// /indexer/v1/getNamespaces

type GetNamespacesReq struct {
	Limit  int `json:"limit,omitempty"`
	Offset int `json:"offset,omitempty"`
	Filter struct {
		Owner string `json:"owner,omitempty"`
		Admin string `json:"admin,omitempty"`
	} `json:"filter,omitempty"`
}

type GetNamespacesRsp struct {
	Namespaces []struct {
		Name            string   `json:"name,omitempty"`
		Permission      string   `json:"permission,omitempty"`
		AvailableHeight int      `json:"availableHeight,omitempty"`
		Owner           string   `json:"owner,omitempty"`
		Admins          []string `json:"admins,omitempty"`
		NamespaceID     string   `json:"namespaceID,omitempty"`
	} `json:"namespaces,omitempty"`
	LastOffset int `json:"lastOffset,omitempty"`
}

// /indexer/v1/countNamespace

type GetTotalDataIDsInNamesapceReq struct {
	NID string `json:"namespaceID"`
}

type GetTotalDataIDsInNamesapceRsp struct {
	Count int64 `json:"count"`
}
