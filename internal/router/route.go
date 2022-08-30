package router

import (
	"github.com/gin-gonic/gin"
)

//路由
func NewRouter() *gin.Engine {
	r := gin.Default()
	/*
	 * 301:永久重定向
	 * 302：临时重定向
	*/
	r.GET("/index", func(c *gin.Context) {
		c.Redirect(301, "https://www.baidu.com")
	})

	//路由重定向
	r.GET("/a", func(c *gin.Context) {
		c.Request.URL.Path = "/b"
		r.HandleContext(c)
	})
	r.GET("/b", func(c *gin.Context) {
		c.String(200, "Hello World")
	})
	return r
}
