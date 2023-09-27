package encoding

import (
	"crypto/md5"
	"encoding/hex"
	"github.com/tomp332/gobrute/pkg/internalTypes"
)

type Md5Plugin struct {
	internalTypes.Plugin
}

var Md5PluginObj = &Md5Plugin{
	internalTypes.Plugin{
		Name: "MD5",
		Mode: 100,
	},
}

func (p Md5Plugin) Execute(password *string, result *string) error {
	hasher := md5.New()
	hasher.Write([]byte(*password))
	*result = hex.EncodeToString(hasher.Sum(nil))
	return nil

}
