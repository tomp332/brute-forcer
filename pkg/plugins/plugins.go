package plugins

import (
	"github.com/tomp332/gobrute/pkg/internalTypes"
	"github.com/tomp332/gobrute/pkg/plugins/encoding"
)

var EncodingPluginsMap = map[string]internalTypes.EncodingPlugin{
	"MD5":    encoding.Md5PluginObj,
	"BASE64": encoding.Base64PluginObj,
}

var EncryptionPluginsMap = map[string]internalTypes.EncryptionPlugin{}
