package src

import (
	"fmt"
	"github.com/tomp332/gospray/src/models"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var MainDB *gorm.DB

func InitDB() {
	var err error
	MainDB, err = gorm.Open(sqlite.Open(GlobalSettings.DBFilePath), &gorm.Config{})
	if err != nil {
		panic(fmt.Sprintf("failed to connect database: %s", err.Error()))
	}
	err = MainDB.AutoMigrate(&models.SlaveDTO{})
	if err != nil {
		panic(fmt.Sprintf("Unable to migrate MainDB table: %s", err.Error()))
	}
	err = MainDB.AutoMigrate(&models.CredentialsDTO{})
	if err != nil {
		panic(fmt.Sprintf("Unable to migrate MainDB table: %s", err.Error()))
	}
	fmt.Println("Successfully connected to database")
}
