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

func (p Base64Plugin) Execute(result *internalTypes.TaskResult) error {
	result.Hash = base64.StdEncoding.EncodeToString([]byte(result.Password))
	return nil
}
