package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func main() {
	router := gin.Default()

	router.GET("/blblpic", func(c *gin.Context) {
		response, err := http.Get("https://i0.hdslb.com/bfs/archive/9829be9ed3da9001d05cc4208129d22ffd84ab39.png@880w_388h_1c_95q")
		if err != nil || response.StatusCode != http.StatusOK {
			c.Status(http.StatusServiceUnavailable)
			return
		}

		pic := response.Body
		picSize := response.ContentLength
		picMIME := response.Header.Get("Content-Type")
		extraHeaders := map[string]string{
			"Content-Disposition": `attachment; filename="gopher.png"`,
		}

		c.DataFromReader(http.StatusOK, picSize, picMIME, pic, extraHeaders)
	})
	router.Run(":8080")
}
