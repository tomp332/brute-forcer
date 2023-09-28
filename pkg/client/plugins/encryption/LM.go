package encryption

import (
	"crypto/des"
	"github.com/tomp332/gobrute/pkg/internalTypes"
	"gorm.io/gorm/utils"
	"log"
	"strings"
)

type LmPlugin struct {
	internalTypes.Plugin
}

var LmPluginObj = &LmPlugin{
	internalTypes.Plugin{
		Name: "LM",
		Mode: 200,
	},
}

func (p LmPlugin) Execute(result *internalTypes.EncryptionTaskResult) error {
	upperData := strings.ToUpper(result.TargetPassword)
	// Pad up to 14 bytes
	if len(upperData) < 14 {
		upperData += strings.Repeat("\x00", 14-len(upperData))
	}
	// Split into two 7-byte chunks
	chunk1 := upperData[:7]
	chunk2 := upperData[7:]
	// Convert to bytes
	chunk1Bytes := []byte(chunk1)
	chunk2Bytes := []byte(chunk2)
	// Creates 2 64-bit DES keys (with the addition of a parity bit for every seven bits)
	// Add a null byte to the end of each chunk to make 8 bytes if the length is not 8 bytes
	if len(chunk1Bytes) != 8 {
		chunk1Bytes = append(chunk1Bytes, 0)
	}
	if len(chunk2Bytes) != 8 {
		chunk2Bytes = append(chunk2Bytes, 0)
	}
	key1, err := des.NewCipher(chunk1Bytes)
	if err != nil {
		log.Fatal("Failed to generate DES keys, error: ", err)
	}
	key2, err := des.NewCipher(chunk2Bytes)
	if err != nil {
		log.Fatal(err)
	}
	hash1 := make([]byte, 8)
	hash2 := make([]byte, 8)
	// Encrypt each string with "KGS!@#$%", which gives two 8-byte cipher values.
	key1.Encrypt(hash1, []byte("KGS!@#$%"))
	key2.Encrypt(hash2, []byte("KGS!@#$%"))
	// Concatenate the two cipher values to produce a 16-byte cipher value.
	result.EncryptedHash = utils.ToString(append(hash1, hash2...))

	return nil
}
