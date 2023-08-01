package crud

import (
	"github.com/tomp332/bruteForcer/src"
	"github.com/tomp332/bruteForcer/src/models"
)

func AddCreds(creds []*models.Cred) ([]*models.Cred, error) {
	result := src.MainDB.Create(creds)
	if result.Error != nil {
		return creds, result.Error
	}
	return creds, nil
}

func GetCreds(limit int, err error) ([]*models.Cred, error) {
	var creds []*models.Cred
	result := src.MainDB.Find(&creds)
	if result.Error != nil {
		return creds, result.Error
	}
	return creds, nil
}
