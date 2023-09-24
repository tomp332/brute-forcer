package plugins

import (
	"github.com/tomp332/gobrute/src/cli/cliTypes"
	"github.com/tomp332/gobrute/src/cli/plugins/encoding"
)

var EncodingPluginsMap = map[string]cliTypes.EncodingPlugin{
	"MD5":    encoding.Md5PluginObj,
	"BASE64": encoding.Base64PluginObj,
}

var EncryptionPluginsMap = map[string]cliTypes.EncryptionPlugin{}
