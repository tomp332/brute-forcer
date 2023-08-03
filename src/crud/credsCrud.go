package crud

import (
	"github.com/tomp332/bruteForcer/src"
	"github.com/tomp332/bruteForcer/src/models"
)

// AddCreds adds creds to the database
func AddCreds(creds []*models.CredsModel) ([]*models.CredsModel, error) {
	// Create a slice of CredsModel structs from the given ICredCreate structs
	result := src.MainDB.Create(creds)
	if result.Error != nil {
		return nil, result.Error
	}
	return creds, nil
}

// GetCreds returns creds from the database
func GetCreds(limit, page uint) ([]*models.CredsModel, error) {
	var creds []*models.CredsModel
	err := src.MainDB.Scopes(src.NewPaginate(limit, page).PaginatedResult).Find(&creds).Error
	return creds, err
}
