package main

import (
	"gochat/router"
	"gochat/utils"
)

func main() {
	utils.InitConfig()
	utils.InitMysql()
	// 得到*gin.Engine并执行run启动8080端口
	router.Router().Run(":80")
}
