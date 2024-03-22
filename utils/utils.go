package utils

import (
	"strings"

	"github.com/RiemaLabs/nubit-da-sdk/constant"
	"github.com/gabriel-vasile/mimetype"
)

func MIMEType(filePath string) map[string]interface{} {
	detectFile, err := mimetype.DetectFile(filePath)
	if err != nil {
		return nil
	}
	return map[string]interface{}{"contentType": detectFile.String()}
}

func CheckContentType(ty string) bool {
	if ty == "" {
		return false
	}
	for _, contextType := range constant.ContextTypes {
		if strings.HasPrefix(ty, contextType) {
			return true
		}
	}
	return false
}
