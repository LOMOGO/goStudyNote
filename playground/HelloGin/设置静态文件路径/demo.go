package main

import (
	"github.com/gin-gonic/gin"
)

func main()  {
	r := gin.Default()

	//设置静态文件路径在浏览器缓存了静态文件之后，当你下一次改变静态文件函数中的相对文件路径的时候，它会因为301永久重定向导致你下次测试这个新的文件路径时，显示的内容依旧是以前静态文件路径，，，
	r.Static("/data", "./logo")
	r.StaticFS("/staticfs", gin.Dir("./vscode", true))
	r.StaticFile("/favicon.ico", "./logo/favorite.ico")
	r.Run(":8080")
}
