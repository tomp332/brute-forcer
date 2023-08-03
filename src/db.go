package src

import (
	"fmt"
	"github.com/tomp332/bruteForcer/src/models"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var MainDB *gorm.DB

type Paginate struct {
	Limit uint `query:"limit"`
	Page  uint `query:"page"`
}

func NewPaginate(limit uint, page uint) *Paginate {
	return &Paginate{Limit: limit, Page: page}
}

func (pg *Paginate) PaginatedResult(db *gorm.DB) *gorm.DB {
	offset := (pg.Page - 1) * pg.Limit

	return db.Offset(int(offset)).
		Limit(int(pg.Limit))
}

func InitDB() {
	var err error
	MainDB, err = gorm.Open(sqlite.Open(GlobalSettings.DBFilePath), &gorm.Config{})
	if err != nil {
		panic(fmt.Sprintf("failed to connect database: %s", err.Error()))
	}
	err = MainDB.AutoMigrate(&models.SlaveModel{})
	if err != nil {
		panic(fmt.Sprintf("Unable to migrate MainDB table: %s", err.Error()))
	}
	err = MainDB.AutoMigrate(&models.CredsModel{})
	if err != nil {
		panic(fmt.Sprintf("Unable to migrate MainDB table: %s", err.Error()))
	}
	fmt.Println("Successfully connected to database")
}
