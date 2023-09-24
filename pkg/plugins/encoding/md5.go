package encoding

import (
	"crypto/md5"
	"encoding/hex"
	"github.com/tomp332/gobrute/pkg/types"
)

type Md5Plugin struct {
	types.Plugin
}

var Md5PluginObj = &Md5Plugin{
	types.Plugin{
		Name: "MD5",
	},
}

func (p Md5Plugin) Encode(data string) (string, error) {
	md5Hash := md5.New()
	_, _ = md5Hash.Write([]byte(data))
	return hex.EncodeToString(md5Hash.Sum(nil)), nil
}

func (p Md5Plugin) Decode(data string) (string, error) {
	return p.Encode(data)
}
