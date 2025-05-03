package common

import (
	"os"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var DB *gorm.DB

func InitDB() {
	var err error
	DB, err = gorm.Open(sqlite.Open("library.db"), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})
	if err != nil {
		panic("Could not connect to database")
	}
}

func InitDummyDB() {
	var err error
	DB, err = gorm.Open(sqlite.Open("dummy.db"), &gorm.Config{})
	if err != nil {
		panic("Could not connect to database")
	}

}

func CloseDummyDB() {
	sqlDB, err := DB.DB()
	if err != nil {
		panic("failed to get sql.DB")
	}

	err = sqlDB.Close()
	if err != nil {
		panic("failed to close database connection")
	}
	os.Remove("dummy.db")
}
