package pkg

import (
	"fmt"
	"github.com/tomp332/gobrute/pkg/types"
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
	err = MainDB.AutoMigrate(&types.SlaveDTO{})
	if err != nil {
		panic(fmt.Sprintf("Unable to migrate Slaves table: %s", err.Error()))
	}
	err = MainDB.AutoMigrate(&types.CredentialsDTO{})
	if err != nil {
		panic(fmt.Sprintf("Unable to migrate Credentials table: %s", err.Error()))
	}
	err = MainDB.AutoMigrate(&types.BruteForceDTO{})
	if err != nil {
		panic(fmt.Sprintf("Unable to migrate Brute Force table: %s", err.Error()))
	}
	fmt.Println("Successfully connected to database")
}
