package crud

import "gorm.io/gorm"

type Crudable interface {
	Get() error
	Set() error
	Update() error
	Delete() error
}

type IDeleteParams struct {
	ID uint `query:"id"`
}

type IPaginateParams struct {
	Limit  uint `query:"limit"`
	Offset uint `query:"page"`
}

func NewPaginate(limit uint, page uint) *IPaginateParams {
	return &IPaginateParams{Limit: limit, Offset: page}
}

func (pg *IPaginateParams) PaginatedResult(db *gorm.DB) *gorm.DB {
	offset := (pg.Offset - 1) * pg.Limit
	return db.Offset(int(offset)).Limit(int(pg.Limit))
}
