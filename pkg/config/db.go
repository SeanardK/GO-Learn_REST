package configDB

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var db *gorm.DB

func Connect() {
	d, err := gorm.Open(mysql.Open("root@tcp(127.0.0.1:3306)/go-bookstore"))
	if err != nil {
		panic("failed to connect database")
	}

	db = d
}

func GetDB() *gorm.DB {
	return db
}
