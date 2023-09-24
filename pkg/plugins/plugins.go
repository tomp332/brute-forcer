package plugins

import (
	"github.com/tomp332/gobrute/pkg/plugins/encoding"
	"github.com/tomp332/gobrute/pkg/types"
)

var EncodingPluginsMap = map[string]types.EncodingPlugin{
	"MD5":    encoding.Md5PluginObj,
	"BASE64": encoding.Base64PluginObj,
}

var EncryptionPluginsMap = map[string]types.EncryptionPlugin{}
