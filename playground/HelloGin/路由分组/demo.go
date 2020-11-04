package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

func main() {
	router := gin.Default()

	g1 := router.Group("/g1")
	{
		g1.GET("/add/:num1/:num2", func(c *gin.Context) {
			num1, _ := strconv.Atoi(c.Param("num1"))
			num2, _ := strconv.Atoi(c.Param("num2"))

			c.String(http.StatusOK, "num1 + num2 = %d", num1 + num2)
		})
		g1.GET("/subtract/:num1/:num2", func(c *gin.Context) {
			num1, _ := strconv.Atoi(c.Param("num1"))
			num2, _ := strconv.Atoi(c.Param("num2"))

			c.String(http.StatusOK, "num1 - num2 = %d", num1 - num2)
		})
	}

	g2 := router.Group("/g2")
	{
		g2.GET("/square/:num1", func(c *gin.Context) {
			num1, _ := strconv.Atoi(c.Param("num1"))
			c.String(http.StatusOK, "num1^2 = %d", num1 * num1)
		})
		g2.GET("/double/:num1", func(c *gin.Context) {
			num1, _ := strconv.Atoi(c.Param("num1"))
			c.String(http.StatusOK, "num1 * 2 = %d", num1 * 2)
		})
	}

	router.Run(":8080")
}
