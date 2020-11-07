###无中间件启动
使用
`r := gin.New()`
代替
```
//默认启动方式，包含Logger、Recovery 中间件
r := gin.Default()
```
###有中间件启动
```
package main

import "github.com/gin-gonic/gin"

func main() {
	router := gin.New()

	//全局中间件
	//使用Logger中间件
	router.Use(gin.Logger())
	//使用Recovery中间件
	router.Use(gin.Recovery())

	//路由添加中间件，可以添加任意多个
	router.GET("/benchmark", MyBenchlogger(), benchEndPoint)

	//路由组中添加中间件
	//group1 := router.Group("/", AuthRequired())
	//上面这句用法和下面的写法等价
	group1 := router.Group("/")
	group1.Use(AuthRequired())
	{
		group1.POST("/login", loginEndpoint)
		group1.POST("/submit", submitEndpoint)
		group1.POST("/read", readEndpoint)

		// 分组的嵌套
		testing := group1.Group("testing")
		testing.GET("/analytics", analyticsEndpoint)
	}
}

```