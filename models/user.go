package models

import (
	"fmt"
	"gochat/utils"

	"gorm.io/gorm"
)

type User struct {
	// gorm.Model创建了4个默认字段
	gorm.Model
	UserName string `valid:"emial"`
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

func CreateUser(user User) *gorm.DB {
	return utils.InitMysql().Create(&user)
}

func DeleteUser(user User) *gorm.DB {
	return utils.InitMysql().Delete(&user)
}

func Updateuser(user User) *gorm.DB {
	updateUser := User{UserName: user.UserName, Password: user.Password, Avatar: user.Avatar}
	// https://gorm.io/docs/update.html#Updates-multiple-columns
	return utils.InitMysql().Model(&user).Updates(updateUser)
}
