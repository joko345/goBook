package models

import (
	"github.com/jinzhu/gorm"
	"github.com/joko345/goBook/pkg/config"
)

var dbLogin *gorm.DB

type User struct {
	ID       uint   `json:"id"`
	Username string `json:"username"`
	Password string `json:"password"`
}

func (u *User) CreateUser() *User {
	dbLogin.NewRecord(u)
	dbLogin.Create(&u)
	return u
}

func GetUserByUsername(username string) (User, error) {
	var user User
	if err := dbLogin.Where("username = ?", username).First(&user).Error; err != nil {
		return user, err
	}
	return user, nil
}

func init() {
	config.ConnectLogin()
	dbLogin = config.GetLogin()

	dbLogin.AutoMigrate(&User{})
}
