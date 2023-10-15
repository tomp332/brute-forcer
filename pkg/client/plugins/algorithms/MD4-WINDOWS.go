package algorithms

import (
	"github.com/tomp332/gobrute/pkg/internalTypes"
	"golang.org/x/crypto/md4"
)

type Md4WindowsPlugin struct {
	internalTypes.Plugin
}

var Md4WindowsPluginObj = &Md4WindowsPlugin{
	internalTypes.Plugin{
		Name: "MD4-WINDOWS",
		Mode: 103,
	},
}

func (p Md4WindowsPlugin) Execute(result *internalTypes.EncryptionTaskResult) error {
	u := ""
	for _, c := range result.TargetPassword {
		u = u + string(c) + "\x00"
	}
	h := md4.New()
	h.Write([]byte(u))
	result.EncryptedHash = string(h.Sum(nil))
	return nil
}
