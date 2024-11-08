package models

import (
	"github.com/jinzhu/gorm"
	"github.com/joko345/goBook/pkg/config"
)

var db *gorm.DB

type Book struct {
	gorm.Model
	Name   string `gorm: "json: "name"`
	Author string `json: "author"`
	Rilis  string `json: "rilis"`
}

func init() {
	config.Connect() //dari file config untuk konek db
	db = config.GetDB()
	db.AutoMigrate(&Book{})
}

func (b *Book) CreateBook() *Book {
	db.NewRecord(b) //b merupakan reasign Book
	db.Create(&b)
	return b
}

func GetBook() []Book {
	var Books []Book //deklarasi var Books dan return value ke []Book
	db.Find(&Books)
	return Books
}

func GetBookById(Id int64) (*Book, *gorm.DB) {
	var getBook Book
	db := db.Where("ID=?", Id).Find(&getBook) //var ID akan mencari Id dalam db
	return &getBook, db
}

func DeleteBook(ID int64) Book {
	var book Book
	db.Where("ID=?", ID).Delete(book)
	return book
}
