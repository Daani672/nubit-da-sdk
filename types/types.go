package types

type CountRsp struct {
	Count int `json:"count,omitempty"`
}

type SysInfo struct {
	SysType string `json:"sysType,omitempty" yaml:"sysType,omitempty"`
	Arch    string `json:"arch,omitempty" yaml:"arch,omitempty"`
}

type CommandChain struct {
	Info string
	Err  string
}

type LightningInfo struct {
	MacaroonFile string `json:"macaroonFile,omitempty" yaml:"MacaroonFile,omitempty"`
	TlsFile      string `json:"tlsFile,omitempty" yaml:"TlsFile,omitempty"`
	GrpcTarget   string `json:"grpcTarget,omitempty" yaml:"GrpcTarget,omitempty"`
}

type Options func(*NubitClient)
