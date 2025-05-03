package models

import (
	"LibrarySystem/common"
	"strconv"
)

type Book struct {
	ID          uint64 `json:"id" gorm:"primaryKey"`
	Title       string `json:"title" gorm:"unique;not null"`
	Author      string `json:"author" gorm:"not null"`
	Description string `json:"description" gorm:"not null"`
}

type Books []Book

func (b *Book) Save() error {
	result := common.DB.Create(&b)

	return result.Error
}

func GetBooksWithPagination(limit int, offset int, order string) (*Books, error) {
	var books Books

	result := common.DB.Offset(offset).Limit(limit).Order(order).Find(&books)

	if result.Error != nil {
		return nil, result.Error
	}

	return &books, nil
}

func GetBooks() (*Books, error) {
	var books Books

	result := common.DB.Find(&books)

	if result.Error != nil {
		return nil, result.Error
	}

	return &books, nil
}

func (b *Book) Update(updates map[string]any) error {
	result := common.DB.Model(&Book{}).Where("id = ?", strconv.FormatUint(b.ID, 10)).Updates(updates)
	return result.Error
}

func (b *Book) Delete() error {
	result := common.DB.Delete(&b)
	return result.Error
}

func GetBookByID(id int64) (*Book, error) {
	var book Book
	result := common.DB.First(&book, id)
	if result.Error != nil {
		return nil, result.Error
	}

	return &book, nil
}
