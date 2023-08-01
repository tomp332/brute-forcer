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

func GetCreds(limit, page int) ([]*models.Cred, error) {
	var creds []*models.Cred
	err := src.MainDB.Scopes(src.NewPaginate(limit, page).PaginatedResult).Find(&creds).Error
	return creds, err
}
