package algorithms

import (
	"github.com/tomp332/gobrute/pkg/internalTypes"
	"golang.org/x/crypto/md4"
)

type Md4Plugin struct {
	internalTypes.Plugin
}

var Md4PluginObj = &Md4Plugin{
	internalTypes.Plugin{
		Name: "MD4",
		Mode: 102,
	},
}

func (p Md4Plugin) Execute(result *internalTypes.EncryptionTaskResult) error {
	h := md4.New()
	h.Write([]byte(result.TargetPassword))
	result.EncryptedHash = string(h.Sum(nil))
	return nil
}
