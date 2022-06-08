package models

import (
	"github.com/jinzhu/gorm"
	"github.com/sean-cunniffe/go-bookstore/src/pkg/config"
)

var Db *gorm.DB

type Book struct {
	gorm.Model
	Name        string `gorm:"" json:"name"`
	Author      string `json:"author"`
	Publication string `json:"publication"`
}

func init() {
	config.Connect()
	Db = config.GetDb()
	Db.AutoMigrate(&Book{})
}

func (b *Book) CreateBook() *Book {
	// check to see if new book is already in Db
	Db.NewRecord(b)
	Db.Create(&b)
	return b
}

func GetAllBooks() []Book {
	var Books []Book
	Db.Find(&Books)
	return Books
}

func GetBookById(id int) *Book {
	var book Book
	Db.Where("ID=?", id).Find(&book)
	return &book
}

func DeleteBook(id int) (book Book) {
	Db.Where("ID=?", id).Delete(&book)
	return
}

func SaveBook(bookDetails *Book) {
	Db.Save(bookDetails)
}
