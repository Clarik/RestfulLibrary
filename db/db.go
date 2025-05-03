package db

import (
	"LibrarySystem/common"
	"LibrarySystem/models"
)

func InitDB() {
	common.InitDB()
	common.DB.AutoMigrate(&models.Book{})
}

func InitDummyDB() {
	common.InitDummyDB()
	common.DB.AutoMigrate(&models.Book{})
}

func CloseDummyDB() {
	common.CloseDummyDB()
}
