package encoding

import (
	"encoding/base64"
	"github.com/tomp332/gobrute/pkg/internalTypes"
)

type Base64Plugin struct {
	internalTypes.Plugin
}

var Base64PluginObj = &Base64Plugin{
	internalTypes.Plugin{
		Name: "BASE64",
	},
}

func (p Base64Plugin) Encode(data string) (string, error) {
	return base64.StdEncoding.EncodeToString([]byte(data)), nil
}

func (p Base64Plugin) Decode(data string) (string, error) {
	sDec, _ := base64.StdEncoding.DecodeString(data)
	return string(sDec), nil
}
