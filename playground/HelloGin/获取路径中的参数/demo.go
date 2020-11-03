package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func main() {
	//使用默认中间件创建一个gin路由器，logger and recovery (crash-free) 中间件
	router := gin.Default()

	//此规则能匹配/user/lomogo这种格式，但是不能匹配/user/或者/user这种格式
	router.GET("/user/:name", func(c *gin.Context) {
		//此处的获取的name参数就是上面路径中冒号后面的name
		name := c.Param("name")
		c.String(http.StatusOK, "hello,%s",name)
	})

	//此规则既能匹配 /user/lomogo/ 这种格式，又能匹配 /user/lomogo/帅 这种格式
	//如果没有其他路由器匹配 /user/lomogo，它将重定向到 /user/lomogo/
	router.GET("/user/:name/*adjective", func(c *gin.Context) {
		name := c.Param("name")
		adjective := c.Param("adjective")
		c.String(http.StatusOK, "%s,你真%s", name, adjective)
	})

	//默认启动的是8080端口，也可以启动其他端口
	router.Run()
	// router.Run("8000") 比如启动在8000端口
}
