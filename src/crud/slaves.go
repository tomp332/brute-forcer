package crud

import (
	"github.com/tomp332/bruteForcer/src"
	"github.com/tomp332/bruteForcer/src/models"
)

func AddSlaves(slaves []*models.Slave) ([]*models.Slave, error) {
	result := src.MainDB.Create(slaves)
	if result.Error != nil {
		return slaves, result.Error
	}
	return slaves, nil
}
