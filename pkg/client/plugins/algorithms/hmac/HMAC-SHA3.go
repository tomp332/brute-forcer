package hmac

import (
	"crypto/hmac"
	"encoding/hex"
	log "github.com/sirupsen/logrus"
	"github.com/tomp332/gobrute/pkg/internalTypes"
	"golang.org/x/crypto/sha3"
	"hash"
)

type HmacSha3 struct {
	internalTypes.Plugin
}

var HmacSha3PluginObj = &HmacSha3{
	internalTypes.Plugin{
		Name: "HMAC-SHA3",
		Mode: 203,
	},
}

func (p HmacSha3) Execute(result *internalTypes.EncryptionTaskResult) error {
	var f func() hash.Hash
	switch result.MetaData.Size {
	case 224:
		f = sha3.New224
	case 256:
		f = sha3.New256
	case 384:
		f = sha3.New384
	case 512:
		f = sha3.New512
	default:
		log.Fatal("Error receiving SHA3 size to use")
	}
	h := hmac.New(f, []byte(result.TargetPassword))
	h.Write([]byte(result.TextMessage))
	result.EncryptedHash = hex.EncodeToString(h.Sum(nil))
	return nil
}
