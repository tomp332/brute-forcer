package crud

import (
	"github.com/tomp332/bruteForcer/src"
	"github.com/tomp332/bruteForcer/src/models"
)

// AddSlaves adds the given slaves to the database
func AddSlaves(slaves []*models.Slave) ([]*models.Slave, error) {
	result := src.MainDB.Create(slaves)
	if result.Error != nil {
		return slaves, result.Error
	}
	return slaves, nil
}

// GetSlaves returns all slaves from the database
func GetSlaves(limit, page int) ([]*models.Slave, error) {
	var slaves []*models.Slave
	err := src.MainDB.Scopes(src.NewPaginate(limit, page).PaginatedResult).Find(&slaves).Error
	return slaves, err
}
