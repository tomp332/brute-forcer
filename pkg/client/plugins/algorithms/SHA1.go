package algorithms

import (
	"crypto/sha1"
	"encoding/hex"
	"github.com/tomp332/gobrute/pkg/internalTypes"
)

type Sha1Plugin struct {
	internalTypes.Plugin
}

var Sha1PluginObj = &Sha1Plugin{
	internalTypes.Plugin{
		Name: "SHA1",
		Mode: 105,
	},
}

func (p Sha1Plugin) Execute(result *internalTypes.EncryptionTaskResult) error {
	h := sha1.New()
	h.Write([]byte(result.TargetPassword))
	result.EncryptedHash = hex.EncodeToString(h.Sum(nil))
	return nil
}
