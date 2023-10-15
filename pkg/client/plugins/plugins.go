package plugins

import (
	"github.com/tomp332/gobrute/pkg/client/plugins/algorithms"
	"github.com/tomp332/gobrute/pkg/client/plugins/algorithms/hmac"
	"github.com/tomp332/gobrute/pkg/internalTypes"
)

var GoBrutePlugins = map[int16]internalTypes.GoBrutePlugin{
	100: algorithms.Md5PluginObj,
	101: algorithms.Base64PluginObj,
	102: algorithms.Md4PluginObj,
	103: algorithms.Md4WindowsPluginObj,
	105: algorithms.Sha1PluginObj,
	200: algorithms.LmPluginObj,
	201: hmac.HmacSha1PluginObj,
	202: hmac.HmacSha3PluginObj,
	203: hmac.HmacSha224PluginObj,
	204: hmac.HmacSha256PluginObj,
	205: hmac.HmacSha384PluginObj,
	206: hmac.HmacSha512PluginObj,
}

func GetPlugin(mode int16) internalTypes.GoBrutePlugin {
	return GoBrutePlugins[mode]
}
