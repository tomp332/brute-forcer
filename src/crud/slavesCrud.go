package crud

import (
	"github.com/tomp332/bruteForcer/src"
	"github.com/tomp332/bruteForcer/src/models"
)

// AddSlaves adds the given slaves to the database
func AddSlaves(slaves []*models.SlaveModel) ([]*models.SlaveModel, error) {
	result := src.MainDB.Create(slaves)
	if result.Error != nil {
		return slaves, result.Error
	}
	return slaves, nil
}

// GetSlaves returns all slaves from the database
func GetSlaves(limit, page uint) ([]*models.SlaveModel, error) {
	var slaves []*models.SlaveModel
	err := src.MainDB.Scopes(src.NewPaginate(limit, page).PaginatedResult).Find(&slaves).Error
	return slaves, err
}
