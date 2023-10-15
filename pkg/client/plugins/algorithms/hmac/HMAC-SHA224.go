package hmac

import (
	"crypto/hmac"
	"encoding/hex"
	"github.com/tomp332/gobrute/pkg/internalTypes"
	"golang.org/x/crypto/sha3"
)

type HmacSha224 struct {
	internalTypes.Plugin
}

var HmacSha224PluginObj = &HmacSha224{
	internalTypes.Plugin{
		Name: "HMAC-SHA224",
		Mode: 203,
	},
}

func (p HmacSha224) Execute(result *internalTypes.EncryptionTaskResult) error {
	h := hmac.New(sha3.New224, []byte(result.TargetPassword))
	result.EncryptedHash = hex.EncodeToString(h.Sum(nil))
	return nil
}
