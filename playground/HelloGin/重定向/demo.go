package main

import (
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	r.GET("/test", func(c *gin.Context) {
		//重定向分为永久重定向和临时重定向，使用永久重定向的时候浏览器具有记忆性，会永久的将两个网址绑定在一起
		c.Redirect(302, "http://music.163.com")
	})

	r.Run(":8080")
}
