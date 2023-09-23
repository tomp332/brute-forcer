package plugins

import (
	"github.com/tomp332/gobrute/src/models"
	"github.com/tomp332/gobrute/src/plugins/encoding"
)

var EncodingPluginsMap = map[string]models.EncodingPlugin{
	"MD5":    encoding.Md5PluginObj,
	"BASE64": encoding.Base64PluginObj,
}

var EncryptionPluginsMap = map[string]models.EncryptionPlugin{}
