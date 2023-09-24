package crud

import (
	manager "github.com/tomp332/gobrute/src/ manager"
	"github.com/tomp332/gobrute/src/ manager/managerTypes"
)

type SlavesCrud struct{}

var SlaveCrud = &SlavesCrud{}

func (s *SlavesCrud) Get(limit, offset uint) ([]*managerTypes.SlaveDTO, error) {
	var objSlice []*managerTypes.SlaveDTO
	err := manager.MainDB.Scopes(NewPaginate(limit, offset).PaginatedResult).Find(&objSlice).Error
	if err != nil {
		return nil, err
	}
	return objSlice, err
}

func (s *SlavesCrud) Add(creds []*managerTypes.SlaveDTO) ([]*managerTypes.SlaveDTO, error) {
	result := manager.MainDB.Create(creds)
	if result.Error != nil {
		return nil, result.Error
	}
	return creds, nil
}

func (s *SlavesCrud) Update(creds []*managerTypes.SlaveDTO) ([]*managerTypes.SlaveDTO, error) {
	result := manager.MainDB.Save(creds)
	if result.Error != nil {
		return nil, result.Error
	}
	return creds, nil
}

func (s *SlavesCrud) Delete(creds []*managerTypes.SlaveDTO) ([]*managerTypes.SlaveDTO, error) {
	result := manager.MainDB.Delete(creds)
	if result.Error != nil {
		return nil, result.Error
	}
	return creds, nil
}
