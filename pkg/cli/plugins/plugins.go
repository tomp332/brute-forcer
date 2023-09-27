package plugins

import (
	"github.com/tomp332/gobrute/pkg/cli/plugins/encoding"
	"github.com/tomp332/gobrute/pkg/cli/plugins/encryption"
	"github.com/tomp332/gobrute/pkg/internalTypes"
)

var GoBrutePlugins = map[int16]internalTypes.GoBrutePlugin{
	100: encoding.Md5PluginObj,
	101: encoding.Base64PluginObj,
	200: encryption.LmPluginObj,
}

func GetPlugin(mode int16) internalTypes.GoBrutePlugin {
	return GoBrutePlugins[mode]
}
