package hmac

import (
	"crypto/hmac"
	"crypto/sha1"
	"encoding/hex"
	"github.com/tomp332/gobrute/pkg/internalTypes"
)

type HmacSha1 struct {
	internalTypes.Plugin
}

var HmacSha1PluginObj = &HmacSha1{
	internalTypes.Plugin{
		Name: "HMAC-SHA1",
		Mode: 201,
	},
}

func (p HmacSha1) Execute(result *internalTypes.EncryptionTaskResult) error {
	h := hmac.New(sha1.New, []byte(result.TargetPassword))
	result.EncryptedHash = hex.EncodeToString(h.Sum(nil))
	return nil
}
