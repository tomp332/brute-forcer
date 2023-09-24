package internalTypes

type Plugin struct {
	Name string
}

type PluginBase interface{}

type EncryptionPlugin interface {
	Encrypt(data string) (string, error)
	Decrypt(data string) (string string)
	PluginBase
}

type EncodingPlugin interface {
	Encode(data string) (string, error)
	Decode(data string) (string, error)
	PluginBase
}
