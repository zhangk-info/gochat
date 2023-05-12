# Go项目实战

1. go mod init gochat

## 切换国内可用代理地址
* go env -w GOPROXY=https://goproxy.cn,direct

## 组件

### gin HTTP Web 框架

```go
// 这下面的代码要拆成3份 1. Router() *gin.Engine 2. service.GetIndex() 3. Run()
r := gin.Default()
// 访问http://localhost:8080/ping 输出{"message":"pong"}
r.GET("/ping", func(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "pong",
	})
})
r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
```

### gorm 全功能 ORM

### viper 配置文件

### gin-swagger 支持gin的swagger文档
* go install github.com/swaggo/swag/cmd/swag@v1.6.0
* swag init 每次更新后执行
* 访问http://localhost/swagger/index.html

```
@Summary	摘要
@Produce	API 可以产生的 MIME 类型的列表，MIME 类型你可以简单的理解为响应类型，例如：json、xml、html 等等
@Param	参数格式，从左到右分别为：参数名、入参类型、数据类型、是否必填、注释
@Success	响应成功，从左到右分别为：状态码、参数类型、数据类型、注释
@Success 200 {string} string "成功"
@Failure	响应失败，从左到右分别为：状态码、参数类型、数据类型、注释
@Router	路由，从左到右分别为：路由地址，HTTP 方法
```

```
@Summary	摘要
@Produce	API 可以产生的 MIME 类型的列表，MIME 类型你可以简单的理解为响应类型，例如：json、xml、html 等等
@Param title body string false "文章标题"
@Success 200 {object} models.User "成功"
@Success 200 {object} string "成功"
@Success 200 {string} string "成功"
@Failure 400 {object} string "请求错误"
@Router /api/v1/articles/{id} [put]
```