package crud

import (
	"github.com/tomp332/gobrute/pkg/internalTypes"
	"github.com/tomp332/gobrute/pkg/manager"
)

type ISlavesCrud struct{}

var SlavesCrud = &ISlavesCrud{}

func (s *ISlavesCrud) Get(limit, offset uint) ([]*internalTypes.SlaveDTO, error) {
	var objSlice []*internalTypes.SlaveDTO
	err := manager.MainDB.Scopes(NewPaginate(limit, offset).PaginatedResult).Find(&objSlice).Error
	if err != nil {
		return nil, err
	}
	return objSlice, err
}

func (s *ISlavesCrud) Add(creds []*internalTypes.SlaveDTO) ([]*internalTypes.SlaveDTO, error) {
	result := manager.MainDB.Create(creds)
	if result.Error != nil {
		return nil, result.Error
	}
	return creds, nil
}

func (s *ISlavesCrud) Update(creds []*internalTypes.SlaveDTO) ([]*internalTypes.SlaveDTO, error) {
	result := manager.MainDB.Save(creds)
	if result.Error != nil {
		return nil, result.Error
	}
	return creds, nil
}

func (s *ISlavesCrud) Delete(creds []*internalTypes.SlaveDTO) ([]*internalTypes.SlaveDTO, error) {
	result := manager.MainDB.Delete(creds)
	if result.Error != nil {
		return nil, result.Error
	}
	return creds, nil
}
