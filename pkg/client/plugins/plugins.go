package plugins

import (
	"github.com/tomp332/gobrute/pkg/client/plugins/encoding"
	"github.com/tomp332/gobrute/pkg/client/plugins/encryption"
	"github.com/tomp332/gobrute/pkg/internalTypes"
)

var GoBrutePlugins = map[int16]internalTypes.GoBrutePlugin{
	100: encoding.Md5PluginObj,
	101: encoding.Base64PluginObj,
	102: encoding.Md4PluginObj,
	103: encoding.Md4WindowsPluginObj,
	200: encryption.LmPluginObj,
}

func GetPlugin(mode int16) internalTypes.GoBrutePlugin {
	return GoBrutePlugins[mode]
}
