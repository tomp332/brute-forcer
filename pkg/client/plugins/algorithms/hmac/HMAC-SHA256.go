package hmac

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/hex"
	"github.com/tomp332/gobrute/pkg/internalTypes"
)

type HmacSha256 struct {
	internalTypes.Plugin
}

var HmacSha256PluginObj = &HmacSha256{
	internalTypes.Plugin{
		Name: "HMAC-SHA256",
		Mode: 202,
	},
}

func (p HmacSha256) Execute(result *internalTypes.EncryptionTaskResult) error {
	h := hmac.New(sha256.New, []byte(result.TargetPassword))
	result.EncryptedHash = hex.EncodeToString(h.Sum(nil))
	return nil
}
