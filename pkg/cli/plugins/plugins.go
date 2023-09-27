package plugins

import (
	"bufio"
	log "github.com/sirupsen/logrus"
	"github.com/tomp332/gobrute/pkg/cli/plugins/encoding"
	"github.com/tomp332/gobrute/pkg/cli/plugins/encryption"
	"github.com/tomp332/gobrute/pkg/internalTypes"
	"os"
)

var GoBrutePlugins = map[int16]internalTypes.GoBrutePlugin{
	100: encoding.Md5PluginObj,
	101: encoding.Base64PluginObj,
	200: encryption.LmPluginObj,
}

func GetPlugin(mode int16) internalTypes.GoBrutePlugin {
	return GoBrutePlugins[mode]
}

func DecryptWrapper(t *internalTypes.Task) (string, error) {
	f, err := os.Open(t.WordlistPath)
	scanner := bufio.NewScanner(f)
	if err != nil {
		log.Fatal(err)
	}
	defer func(f *os.File) {
		err := f.Close()
		if err != nil {
			log.Fatal(err)
		}
	}(f)
	currentPlugin := GetPlugin(t.Mode)
	for scanner.Scan() {
		line := scanner.Text()
		if len(line) == 0 {
			// skip blank lines
			continue
		}
		t.PlaintText = line
		log.WithFields(log.Fields{"hash": t.Hash, "password": line}).Debug("Executing")
		encryptedHash, err := currentPlugin.Execute(t)
		if err != nil {
			continue
		}
		if encryptedHash == t.Hash {
			return encryptedHash, nil
		}
	}
	return "", nil
}
