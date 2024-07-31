// @Title main.go
// @Description $END$
// @Author Zero - 2024/7/31 22:16:14

package main

import "github.com/gin-gonic/gin"

func main() {
	r := gin.New()
	gin.SetMode(gin.ReleaseMode)
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "Ok",
		})
	})
	r.Run() // 监听并在 0.0.0.0:8080 上启动服务
}