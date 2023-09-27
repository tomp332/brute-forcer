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
		Mode: 101,
	},
}

func (p Base64Plugin) Execute(t *internalTypes.Task) (string, error) {
	return base64.StdEncoding.EncodeToString([]byte(t.PlaintText)), nil
}
