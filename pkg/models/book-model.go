package models

import (
	"encoding/json"
	"net/http"

	configDB "github.com/SeanardK/GO-REST_API/pkg/config"
	"gorm.io/gorm"
)

type Book struct {
	Name      string `json:"name"`
	Price     uint   `json:"price"`
	Published string `json:"published_date"`
	Status    string `json:"status"`
}

var db *gorm.DB

func init() {
	configDB.Connect()

	db = configDB.GetDB()
	db.AutoMigrate(&Book{})
}

func BookCreate(r *http.Request) Book {
	var book Book

	json.NewDecoder(r.Body).Decode(&book)

	db.Create(&book)
	return book
}

func BookGetAll() []Book {
	var Books []Book
	db.Find(&Books)

	return Books
}
