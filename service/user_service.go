package service

import (
	"gochat/models"
	"net/http"

	"github.com/asaskevich/govalidator"
	"github.com/gin-gonic/gin"
)

// GetUserList
// @Summary 用户列表
// @Schemes
// @Description 用户列表
// @Tags 用户
// @Accept json
// @Produce json
// @Success 200 {object} models.User "成功"
// @Router /users [get]
func GetUserList(c *gin.Context) {
	data := models.GetUserList()
	c.JSON(http.StatusOK, gin.H{
		"message": data,
	})
}

// CreateUser
// @Summary 创建用户
// @Schemes
// @Description 创建用户
// @Tags 用户
// @Accept json
// @Produce json
// @parm name query string true "用户名"
// @parm password query string true "密码"
// @parm rePassword query string false "确认密码"
// @Success 200 {object} models.User "成功"
// @Router /users [post]
func CreateUser(c *gin.Context) {
	user := models.User{}
	user.UserName = c.Query("userName")
	password := c.Query("password")
	rePassword := c.Query("rePassword")
	if password != rePassword {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "密码和确认密码不一致",
		})
	}
	user.Password = password
	// 进行校验
	govalidator.ValidateStruct(user)
	// 保存
	models.CreateUser(user)
	c.JSON(http.StatusOK, gin.H{
		"message": "新增用户成功",
	})
}
