package crud

import (
	"github.com/tomp332/gobrute/pkg"
	"github.com/tomp332/gobrute/pkg/types"
)

type ISlavesCrud struct{}

var SlavesCrud = &ISlavesCrud{}

func (s *ISlavesCrud) Get(limit, offset uint) ([]*types.SlaveDTO, error) {
	var objSlice []*types.SlaveDTO
	err := pkg.MainDB.Scopes(NewPaginate(limit, offset).PaginatedResult).Find(&objSlice).Error
	if err != nil {
		return nil, err
	}
	return objSlice, err
}

func (s *ISlavesCrud) Add(creds []*types.SlaveDTO) ([]*types.SlaveDTO, error) {
	result := pkg.MainDB.Create(creds)
	if result.Error != nil {
		return nil, result.Error
	}
	return creds, nil
}

func (s *ISlavesCrud) Update(creds []*types.SlaveDTO) ([]*types.SlaveDTO, error) {
	result := pkg.MainDB.Save(creds)
	if result.Error != nil {
		return nil, result.Error
	}
	return creds, nil
}

func (s *ISlavesCrud) Delete(creds []*types.SlaveDTO) ([]*types.SlaveDTO, error) {
	result := pkg.MainDB.Delete(creds)
	if result.Error != nil {
		return nil, result.Error
	}
	return creds, nil
}
