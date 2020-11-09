//服务启动之后在另一个终端中输入 kill -1 pid 即可重启服务
package main

import (
	"github.com/fvbock/endless"
	"github.com/gin-gonic/gin"
	"log"
	"syscall"
	"time"
)

func main() {
	endless.DefaultReadTimeOut = 5 * time.Second
	endless.DefaultWriteTimeOut = 20 * time.Second
	endless.DefaultMaxHeaderBytes = 1 << 20

	r := gin.Default()
	r.GET("/hello", func(c *gin.Context) {
		c.String(200, "如何优雅的重启和停止服务器")
	})
	server := endless.NewServer(":8080", r)
	server.BeforeBegin = func(add string) {
		log.Printf("Actual pid is %d", syscall.Getpid())
	}

	err := server.ListenAndServe()
	if err != nil {
		log.Printf("服务错误：%v", err)
	}
}