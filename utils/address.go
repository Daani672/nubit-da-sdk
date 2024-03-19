package utils

import (
	"crypto/ecdsa"
	"encoding/hex"
	"fmt"
	"os"

	"github.com/RiemaLabs/nubit-da-sdk/constant"
	bitcoin "github.com/bitcoinschema/go-bitcoin/v2"
	"github.com/btcsuite/btcd/btcec/v2"
	"github.com/btcsuite/btcd/btcutil"
	"github.com/btcsuite/btcd/btcutil/hdkeychain"
	"github.com/btcsuite/btcd/chaincfg"
	"github.com/ethereum/go-ethereum/crypto"
	"golang.org/x/crypto/sha3"
)

func KeyTo0xAddress(key *hdkeychain.ExtendedKey) string {
	privateKey, err := key.ECPrivKey()
	if err != nil {
		return ""
	}
	publicKey := privateKey.ToECDSA().Public()
	publicKeyECDSA := publicKey.(*ecdsa.PublicKey)
	return crypto.PubkeyToAddress(*publicKeyECDSA).String()
}

func KeyToBtcAddress(key *hdkeychain.ExtendedKey) string {
	switch constant.NubitNet {
	case constant.MainNet:
		privateKey, err := key.ECPrivKey()
		if err != nil {
			return ""
		}
		btcaddr, err := bitcoin.GetAddressFromPrivateKeyString(EcdsaToPrivateStr(privateKey.ToECDSA()), false)
		if err != nil {
			return ""
		}
		return btcaddr
	case constant.TestNet:
		privateKey, err := key.ECPrivKey()
		if err != nil {
			return ""
		}
		_, pub := btcec.PrivKeyFromBytes(PrivateStrToByte(EcdsaToPrivateStr(privateKey.ToECDSA())))
		publicKeyHash := btcutil.Hash160(pub.SerializeUncompressed())
		p2pkhAddr, err := btcutil.NewAddressPubKeyHash(publicKeyHash, &chaincfg.TestNet3Params)
		if err != nil {
			return ""
		}
		return p2pkhAddr.EncodeAddress()
	}
	return ""
}

func PrivateStrToBtcAddress(private string) string {
	switch constant.NubitNet {
	case constant.MainNet:
		address, err := bitcoin.GetAddressFromPrivateKeyString(private, true)
		if err != nil {
			return ""
		}

		return address
	case constant.TestNet:
		_, pub := btcec.PrivKeyFromBytes(PrivateStrToByte(private))
		publicKeyHash := btcutil.Hash160(pub.SerializeUncompressed())
		p2pkhAddr, err := btcutil.NewAddressPubKeyHash(publicKeyHash, &chaincfg.TestNet3Params)
		if err != nil {
			return ""
		}
		return p2pkhAddr.EncodeAddress()
	}
	return ""
}

func PrivateStrToEcdsa(private string) *ecdsa.PrivateKey {
	toECDSA, err := crypto.HexToECDSA(private)
	if err != nil {
		fmt.Println(err)
		os.Exit(0)
		return nil
	}
	return toECDSA
}

func PrivateStrToByte(private string) []byte {
	ecd := PrivateStrToEcdsa(private)
	ecd.Public()
	return crypto.FromECDSA(ecd)
}

func ByteToEcdsa(b []byte) *ecdsa.PrivateKey {
	toECDSA, err := crypto.ToECDSA(b)
	if err != nil {
		return nil
	}
	return toECDSA
}

func EcdsaToPrivateStr(ecd *ecdsa.PrivateKey) string {
	PirKeyByte := crypto.FromECDSA(ecd)
	return hex.EncodeToString(PirKeyByte)
}

func PrivateByteToStr(b []byte) string {
	return EcdsaToPrivateStr(ByteToEcdsa(b))
}

func BtcAddrTo0x(pubkeyStr string) string {
	pubkey, _ := bitcoin.PubKeyFromString(pubkeyStr)
	return crypto.PubkeyToAddress(*pubkey.ToECDSA()).String()
}

func BTCPRIKEYStrToETHAddr(privateKeyStr string) (string, string, error) {
	pubkeyStr, err := bitcoin.PubKeyFromPrivateKeyString(privateKeyStr, true)
	pubkey, err := bitcoin.PubKeyFromString(pubkeyStr)
	if err != nil {
		return "", "", nil
	}
	hash := sha3.NewLegacyKeccak256()
	hash.Write(pubkey.SerialiseUncompressed()[1:])
	addressByte := hash.Sum(nil)
	return "0x" + hex.EncodeToString(addressByte[12:]), pubkeyStr, nil
}
