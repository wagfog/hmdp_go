package models

import (
	"strconv"
	"time"

	"github.com/jinzhu/gorm"
)

type User struct {
	ID         int64     `json:"id" gorm:"primaryKey;autoIncrement"`
	Phone      string    `json:"phone"`
	Password   string    `json:"password"`
	NickName   string    `json:"nickName"`
	Icon       string    `json:"icon"`
	CreateTime time.Time `json:"createTime"`
	UpdateTime time.Time `json:"updateTime"`
}

func GetUser(id int64) *User {
	var u User
	db.Table("tb_user").First(&u, id)
	return &u
}

func GetUserByPhone(phone string) *User {
	var u User
	db.Table("tb_user").Where("phone = ?", phone).First(&u)
	return &u
}

func CreateUser(phone string) (*User, error) {
	var u User
	err := db.Transaction(func(tx *gorm.DB) error {
		var count int64
		tx.Table("tb_user").Count(&count)
		u = User{
			Phone:      phone,
			NickName:   "user_" + strconv.Itoa(int(count)),
			CreateTime: time.Now(),
			UpdateTime: time.Now(),
		}
		Isok := tx.Table("tb_user").Create(&u)
		return Isok.Error
	})
	return &u, err
}
