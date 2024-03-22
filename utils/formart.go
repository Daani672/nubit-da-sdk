package utils

import (
	"encoding/json"

	yaml "gopkg.in/yaml.v3"
)

func ToJson(data interface{}) string {
	if data == nil {
		return ""
	}
	str, _ := json.MarshalIndent(data, "", "    ")
	return string(str)
}

func ToText(data interface{}) string {
	if data == nil {
		return ""
	}
	marshal, err := yaml.Marshal(data)
	if err != nil {
		return ""
	}
	return string(marshal)
}
