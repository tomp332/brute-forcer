package internalTypes

type Plugin struct {
	Name string
	Mode int16
}

type PluginBase interface{}

type GoBrutePlugin interface {
	Execute(result *EncryptionTaskResult) error
	PluginBase
}
