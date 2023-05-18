package router

import (
	docs "gochat/docs"
	"gochat/service"

	"github.com/gin-gonic/gin"
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func Router() *gin.Engine {
	r := gin.Default()

	// 引入生成的swagger文件并配置
	docs.SwaggerInfo.BasePath = "/" // /api/v1
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))

	r.GET("/", service.GetIndex)
	r.GET("/users", service.GetUserList)
	r.POST("/users", service.CreateUser)
	return r
}
