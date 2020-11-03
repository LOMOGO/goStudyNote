package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func main() {
	router := gin.Default()

	//匹配的url格式：http://localhost:8080/welcome?firstname=心&lastname=木
	router.GET("/welcome", func(c *gin.Context) {
		//如果get参数里面没有这个参数，那么就使用默认值 稀客
		firstname := c.DefaultQuery("firstname", "稀客")
		lastname := c.Query("lastname")

		c.String(http.StatusOK, "欢迎%s%s",lastname, firstname)
	})

	router.Run(":8080")
}
