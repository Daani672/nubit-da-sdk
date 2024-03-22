package types

type Keys struct {
	PublicKey  string `json:"public_key,omitempty" yaml:"PublicKey,omitempty"`
	Mnemonic   string `json:"mnemonic,omitempty" yaml:"Mnemonic,omitempty"`
	PrivateKey string `json:"private_key,omitempty" yaml:"PrivateKey,omitempty"`
	FileName   string `json:"file_name,omitempty" yaml:"FileName,omitempty"`
}
