package service

import (
	"gochat/models"
	"net/http"

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
