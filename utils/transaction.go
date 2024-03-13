package utils

import (
	"crypto/sha256"
	"encoding/hex"

	bitcoin "github.com/bitcoinschema/go-bitcoin/v2"

	"github.com/RiemaLabs/nubit-da-sdk/log"
)

func SignTransaction(privateKeyStr string, msg string) string {
	var (
		err       error
		signature string
	)
	if signature, err = bitcoin.SignMessage(privateKeyStr, msg, false); err != nil {
		log.Error("SignTransaction", "SignMessage", err)
		return signature
	}
	return signature
}

func Sha256hash(rawdata string) string {
	data := []byte(rawdata)
	// Create a new SHA-256 hash
	hash := sha256.New()
	// Write data to the hash
	hash.Write(data)
	// Get the final hash sum
	hashSum := hash.Sum(nil)
	// Convert the hash to a hexadecimal string
	hashString := hex.EncodeToString(hashSum)
	return hashString
}
