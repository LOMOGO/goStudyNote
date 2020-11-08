/*又是到遇到不懂的东西去互联网搜寻答案的一天，然后一脸懵逼的点进一篇文章，又一脸懵逼的退出这篇文章，
我都怀疑有些文章作者是不是故意的，非得把一个简单的答案给你解释成你看不懂的样子，也不知道这个作者是真的懂
还是什么，非蠢即坏。
 */
package main

import "github.com/gin-gonic/gin"

func main() {
	r := gin.Default()

	//首先中间件是在客户端请求和服务器接口接收请求中间的一个程序，因此中间件可以做一些记录日志和校验请求的工作
	//然后，在gin中，中间件的执行顺序和中间件的注册顺序有关，中间件顺序越靠前越先执行，接下来就开始将Next方法和Abort方法
	/*Next方法是用在中间件中的，它的作用就是：如果有两个中间件按顺序前后注册的话，在第一个中间件程序的代码段中，Next方法出现在哪一行代码，那一行代码就执行后面的中间件程序，
	后面的中间件程序执行好之后再执行第一个中间件程序中Next方法出现的那一行代码后面的代码段。它的作用就是这么这样*/
	/*Abort方法也是用在中间件中，只不过它的作用是：如果有两个中间件按顺序前后注册的话，在第一个中间件程序的代码段中，只要Abort方法出现，那么后面注册的那个中间件程序就不会执行。*/

	// 接口含义：/m1m2的意思是使用中间件1、2,然后/123的意思是 /m1m2/123这个接口的输出结果是：123
	r.GET("/m1m2/123", middleware1(), middleware6())

	r.GET("m6m2/312", middleware6(), middleware2())
	r.GET("/m2m6/312", middleware2(), middleware6())

	r.GET("/m6m3/312", middleware6(), middleware3())
	r.GET("/m3m6/132", middleware3(), middleware6())

	r.GET("m6m4/312", middleware6(), middleware4())
	r.GET("/m4m6/123", middleware4(), middleware6())

	r.GET("/m5m6/12", middleware5(), middleware6())

	r.Run(":8080")
}

func middleware1() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.String(200, "1")
		c.String(200, "2")
	}
}

func middleware2() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Next()
		c.String(200, "1")
		c.String(200, "2")
	}
}

func middleware3() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.String(200, "1")
		c.Next()
		c.String(200, "2")
	}
}

func middleware4() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.String(200, "1")
		c.String(200, "2")
		c.Next()
	}
}

func middleware5() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Abort()
		c.String(200, "1")
		c.String(200, "2")
	}
}

func middleware6() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.String(200, "3")
	}
}
