package crud

import (
	"github.com/tomp332/bruteForcer/src"
)

type CRUD interface {
	Add() error
	Get() error
	Update() error
	Delete() error
}

// Get fetch data from the database
func Get(limit, page uint, model interface{}) (interface{}, error) {
	err := src.MainDB.Scopes(src.NewPaginate(limit, page).PaginatedResult).Find(&model).Error
	if err != nil {
		return nil, err
	}
	return model, err
}

// Add data to the database
func Add(model interface{}) (interface{}, error) {
	result := src.MainDB.Create(model)
	if result.Error != nil {
		return nil, result.Error
	}
	return model, nil
}
