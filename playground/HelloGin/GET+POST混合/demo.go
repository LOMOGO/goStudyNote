package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func main() {
	router := gin.Default()
	/*示例：
	POST /post?id=10086&status=1 HTTP/1.1
	Content-Type: application/x-www-form-urlencoded

	name=lomogo&age=80
	*/
	router.POST("/post", func(c *gin.Context) {
		id := c.Query("id")
		status := c.Query("status")
		name := c.PostForm("name")
		age := c.PostForm("age")

		c.JSON(http.StatusOK, gin.H{
			"id": id,
			"status": status,
			"name": name,
			"age": age,
		})
	})

	router.Run(":8080")
}
