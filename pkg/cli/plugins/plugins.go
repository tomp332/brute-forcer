package plugins

import (
	encoding2 "github.com/tomp332/gobrute/pkg/cli/plugins/encoding"
	"github.com/tomp332/gobrute/pkg/cli/plugins/encryption"
	"github.com/tomp332/gobrute/pkg/internalTypes"
)

var EncodingPluginsMap = map[string]internalTypes.EncodingPlugin{
	"MD5":    encoding2.Md5PluginObj,
	"BASE64": encoding2.Base64PluginObj,
}

var EncryptionPluginsMap = map[string]internalTypes.EncryptionPlugin{
	"LM": encryption.LmPluginObj,
}
