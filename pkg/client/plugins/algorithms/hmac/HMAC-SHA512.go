package hmac

import (
	"crypto/hmac"
	"crypto/sha512"
	"encoding/hex"
	"github.com/tomp332/gobrute/pkg/internalTypes"
)

type HmacSha512 struct {
	internalTypes.Plugin
}

var HmacSha512PluginObj = &HmacSha512{
	internalTypes.Plugin{
		Name: "HMAC-SHA512",
		Mode: 205,
	},
}

func (p HmacSha512) Execute(result *internalTypes.EncryptionTaskResult) error {
	h := hmac.New(sha512.New, []byte(result.TargetPassword))
	result.EncryptedHash = hex.EncodeToString(h.Sum(nil))
	return nil
}
