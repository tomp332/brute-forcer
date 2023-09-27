package internalTypes

type Plugin struct {
	Name string
	Mode int16
}

type PluginResult struct {
	Password string
	Hash     string
}

type PluginBase interface{}

type GoBrutePlugin interface {
	Execute(result *PluginResult) error
	PluginBase
}
