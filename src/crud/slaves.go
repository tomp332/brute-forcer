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
func GetSlaves() ([]*models.Slave, error) {
	var slaves []*models.Slave
	result := src.MainDB.Find(&slaves)
	if result.Error != nil {
		return slaves, result.Error
	}
	return slaves, nil
}
