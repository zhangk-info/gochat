package service

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// @BasePath /
// GetIndex
// @Summary 首页
// @Schemes
// @Description 访问首页
// @Tags 首页
// @Accept json
// @Produce json
// @Success 200 {string} string "成功返回json {"message": "hello"}"
// @Router / [get]
func GetIndex(c *gin.Context) {

	c.JSON(http.StatusOK, gin.H{
		"message": "hello",
	})
}
