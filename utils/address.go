package utils

import (
	"crypto/ecdsa"
	"encoding/hex"
	"math/big"

	bitcoin "github.com/bitcoinschema/go-bitcoin/v2"
	"github.com/btcsuite/btcd/btcec/v2"
	"github.com/btcsuite/btcd/btcec/v2/schnorr"
	"github.com/btcsuite/btcd/btcutil"
	"github.com/btcsuite/btcd/chaincfg"
	"github.com/btcsuite/btcd/txscript"
	"golang.org/x/crypto/sha3"
)

const (
	// number of bits in a big.Word
	wordBits = 32 << (uint64(^big.Word(0)) >> 63)
	// number of bytes in a big.Word
	wordBytes = wordBits / 8
)

func PrivateStrToBtcAddress(private string) string {
	_, pub := btcec.PrivKeyFromBytes(PrivateStrToByte(private))
	taproot, err := btcutil.NewAddressTaproot(schnorr.SerializePubKey(txscript.ComputeTaprootKeyNoScript(pub)), &chaincfg.TestNet3Params)
	if err != nil {
		return ""
	}
	return taproot.EncodeAddress()
}

func PrivateStrToEcdsa(private string) *ecdsa.PrivateKey {
	privateKey, err := bitcoin.PrivateKeyFromString(private)
	if err != nil {
		return nil
	}
	return privateKey.ToECDSA()
}

func PrivateStrToByte(private string) []byte {
	ecd := PrivateStrToEcdsa(private)
	return FromECDSA(ecd)
}

func EcdsaToPrivateStr(ecd *ecdsa.PrivateKey) string {
	if ecd == nil {
		return ""
	}
	return hex.EncodeToString(FromECDSA(ecd))
}

// FromECDSA exports a private key into a binary dump.
func FromECDSA(priv *ecdsa.PrivateKey) []byte {
	if priv == nil {
		return nil
	}
	return PaddedBigBytes(priv.D, priv.Params().BitSize/8)
}

// PaddedBigBytes encodes a big integer as a big-endian byte slice. The length
// of the slice is at least n bytes.
func PaddedBigBytes(bigint *big.Int, n int) []byte {
	if bigint.BitLen()/8 >= n {
		return bigint.Bytes()
	}
	ret := make([]byte, n)
	ReadBits(bigint, ret)
	return ret
}

// ReadBits encodes the absolute value of bigint as big-endian bytes. Callers must ensure
// that buf has enough space. If buf is too short the result will be incomplete.
func ReadBits(bigint *big.Int, buf []byte) {
	i := len(buf)
	for _, d := range bigint.Bits() {
		for j := 0; j < wordBytes && i > 0; j++ {
			i--
			buf[i] = byte(d)
			d >>= 8
		}
	}
}

func BTCPRIKEYStrToHexAddr(privateKeyStr string) (string, string, error) {
	pubkeyStr, err := bitcoin.PubKeyFromPrivateKeyString(privateKeyStr, true)
	if err != nil {
		return "", "", nil
	}
	pubkey, err := bitcoin.PubKeyFromString(pubkeyStr)
	if err != nil {
		return "", "", nil
	}
	hash := sha3.NewLegacyKeccak256()
	hash.Write(pubkey.SerialiseUncompressed()[1:])
	addressByte := hash.Sum(nil)
	return "0x" + hex.EncodeToString(addressByte[12:]), pubkeyStr, nil
}
