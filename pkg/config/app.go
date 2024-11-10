package config

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

var (
	db      *gorm.DB
	dbLogin *gorm.DB
)

func Connect() {
	d, err := gorm.Open("mysql", "root:gambut01@tcp(localhost:3306)/gobook?charset=utf8&parseTime=True&loc=Local")

	if err != nil {
		panic(err)
	}
	db = d
}

func GetDB() *gorm.DB {
	return db
}

func ConnectLogin() {
	dLogin, err := gorm.Open("mysql", "root:gambut01@tcp(localhost:3306)/goLogin?charset=utf8&parseTime=True&loc=Local")

	if err != nil {
		panic(err)
	}
	dbLogin = dLogin
}

func GetLogin() *gorm.DB {
	return dbLogin
}
