package crud

import (
	"github.com/tomp332/gobrute/src"
	"github.com/tomp332/gobrute/src/models"
)

type SlavesCrud struct{}

var SlaveCrud = &SlavesCrud{}

func (s *SlavesCrud) Get(limit, offset uint) ([]*models.SlaveDTO, error) {
	var objSlice []*models.SlaveDTO
	err := src.MainDB.Scopes(NewPaginate(limit, offset).PaginatedResult).Find(&objSlice).Error
	if err != nil {
		return nil, err
	}
	return objSlice, err
}

func (s *SlavesCrud) Add(creds []*models.SlaveDTO) ([]*models.SlaveDTO, error) {
	result := src.MainDB.Create(creds)
	if result.Error != nil {
		return nil, result.Error
	}
	return creds, nil
}

func (s *SlavesCrud) Update(creds []*models.SlaveDTO) ([]*models.SlaveDTO, error) {
	result := src.MainDB.Save(creds)
	if result.Error != nil {
		return nil, result.Error
	}
	return creds, nil
}

func (s *SlavesCrud) Delete(creds []*models.SlaveDTO) ([]*models.SlaveDTO, error) {
	result := src.MainDB.Delete(creds)
	if result.Error != nil {
		return nil, result.Error
	}
	return creds, nil
}
