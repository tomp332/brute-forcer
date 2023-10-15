package hmac

import (
	"crypto/hmac"
	"encoding/hex"
	"github.com/tomp332/gobrute/pkg/internalTypes"
	"golang.org/x/crypto/sha3"
)

type HmacSha384 struct {
	internalTypes.Plugin
}

var HmacSha384PluginObj = &HmacSha384{
	internalTypes.Plugin{
		Name: "HMAC-SHA384",
		Mode: 202,
	},
}

func (p HmacSha384) Execute(result *internalTypes.EncryptionTaskResult) error {
	h := hmac.New(sha3.New384, []byte(result.TargetPassword))
	result.EncryptedHash = hex.EncodeToString(h.Sum(nil))
	return nil
}
