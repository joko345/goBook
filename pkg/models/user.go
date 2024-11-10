package models

import (
	"github.com/jinzhu/gorm"
	"github.com/joko345/goBook/pkg/config"
)

var dbLogin *gorm.DB

// Model User untuk menangani data pengguna
type User struct {
	ID       uint   `json:"id"`
	Username string `json:"username"`
	Password string `json:"password"` // Menyimpan password yang sudah di-hash untuk keamanan
}

// Fungsi untuk membuat user baru
func (u *User) CreateUser() *User {
	dbLogin.NewRecord(u) // Membuat record baru
	dbLogin.Create(&u)   // Menyimpan user ke database
	return u
}

// Fungsi untuk mendapatkan user berdasarkan username
func GetUserByUsername(username string) (User, error) {
	var user User
	if err := dbLogin.Where("username = ?", username).First(&user).Error; err != nil {
		return user, err
	}
	return user, nil
}

// Inisialisasi dan koneksi ke database
func init() {
	config.ConnectLogin()       // Menghubungkan ke database
	dbLogin = config.GetLogin() // Mendapatkan instance DB dari config
	// Auto-migrate untuk model User (hanya sekali saja pada aplikasi pertama)
	dbLogin.AutoMigrate(&User{})
}
