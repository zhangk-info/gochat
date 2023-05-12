package models

import (
	"fmt"
	"gochat/utils"
	"gorm.io/gorm"
)

type User struct {
	// gorm.Model创建了4个默认字段
	gorm.Model
	UserName string
	Password string
	Avatar   string
}

func (table User) TableName() string {
	return "user"
}

func NewUser() User {
	return User{}
}

func GetUserList() []*User {
	data := make([]*User, 10)
	utils.InitMysql().Find(&data)
	for _, v := range data {
		fmt.Println(v)
	}
	return data
}
