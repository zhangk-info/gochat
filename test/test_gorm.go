package main

import (
	"fmt"
	"gochat/models"
	"gochat/utils"
)

func main() {

	db := utils.InitMysql()

	// 迁移 schema
	db.AutoMigrate(&models.User{})

	// Create
	user1 := &models.User{UserName: "admin", Password: "456"}
	db.Create(user1)

	// Read
	var user models.User
	db.First(&user, 1) // 根据整型主键查找
	fmt.Println(user)
	//db.First(&user, "code = ?", "D42") // 查找 code 字段值为 D42 的记录

	// Update - 将 User 的 price 更新为 200
	db.Model(&user).Update("Password", 123)
	// Update - 更新多个字段
	db.Model(&user).Updates(models.User{UserName: "admin", Password: "123"}) // 仅更新非零值字段
	db.Model(&user).Updates(map[string]interface{}{"Avatar": "https://"})

	// Delete - 删除 User
	//db.Delete(&user, 2)
}
