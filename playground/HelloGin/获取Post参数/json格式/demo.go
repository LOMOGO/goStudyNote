package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type UserInfo struct {
	Name string `json:"name"`
	Age uint8 `json:"age"`
	Slogan string `json:"slogan"`
}

func main()  {
	router := gin.Default()

	router.POST("/json_post", func(c *gin.Context) {
		var userinfo UserInfo
		c.BindJSON(&userinfo)
		c.JSON(http.StatusOK, userinfo)
	})

	router.Run(":8080")
}
