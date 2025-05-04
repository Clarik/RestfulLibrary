package common

import (
	"fmt"
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var DB *gorm.DB

func InitDB() {
	var err error
	dsn := fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=%s sslmode=disable",
		os.Getenv("DB_HOST"),
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_NAME"),
		os.Getenv("DB_PORT"),
	)
	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{
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
