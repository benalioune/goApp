package models

import (
	"example/web-service-gin/database"
)

type Book struct {
	ID     uint    `json:"id" gorm:"primaryKey"`
	Title  string  `json:"title"`
	Author string  `json:"author"`
	Price  float64 `json:"price"`
}

func GetAllBooks() ([]Book, error) {
	var books []Book
	if err := database.DB.Find(&books).Error; err != nil {
		return nil, err
	}
	return books, nil
}

func AddNewBook(book *Book) error {
	if err := database.DB.Create(&book).Error; err != nil {
		return err
	}
	return nil
}

func GetBookByID(id uint) (*Book, error) {
	var book Book
	if err := database.DB.First(&book, id).Error; err != nil {
		return nil, err
	}
	return &book, nil
}
