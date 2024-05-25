package main

// 可以任意切換使用不同的database

import (
	"fmt"
	"log"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

const (
	USERNAME = "dbuser"
	PASSWORD = "ecc"
	NETWORK  = "tcp"
	SERVER   = "127.0.0.1"
	PORT     = 3306
	DATABASE = "demo"
)

type User struct {
	ID       int64  `json:"id" gorm:"primary_key;auto_increase'"`
	Username string `json:"username"`
	Password string `json:""`
}

func CreateUser(db *gorm.DB, user *User) error {
	return db.Create(user).Error
}

func FindUser(db *gorm.DB, id int64) (*User, error) {
	user := new(User)
	user.ID = id
	err := db.First(&user).Error
	return user, err
}

func main() {
	dsn := fmt.Sprintf("%s:%s@%s(%s:%d)/%s?charset=utf8mb4&parseTime=True&loc=Local", USERNAME, PASSWORD, NETWORK, SERVER, PORT, DATABASE)
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("使用 gorm 連線 DB 發生錯誤，原因為 " + err.Error())
	}
	if err := db.AutoMigrate(new(User)); err != nil {
		panic("資料庫 Migrate 失敗，原因為 " + err.Error())
	}
	user := &User{
		Username: "test",
		Password: "test",
	}
	if err := CreateUser(db, user); err != nil {
		panic("新增 user 失敗，原因為 " + err.Error())
	}
	if user, err := FindUser(db, 1); err == nil {
		log.Println("查詢到 User 為 ", user)
	} else {
		panic("查詢 user 失敗，原因為 " + err.Error())
	}
}
