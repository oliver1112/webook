package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func main() {
	server := gin.Default()
	server.GET("/hello", func(c *gin.Context) {
		c.String(http.StatusOK, "hello, Go")
	})

	// 参数路由，路径参数
	server.GET("/users/:name", func(ctx *gin.Context) {
		name := ctx.Param("name")
		ctx.String(http.StatusOK, "hello, "+name)
	})

	// GET /order?id=123
	server.GET("/order", func(ctx *gin.Context) {
		id := ctx.Query("id")
		ctx.String(http.StatusOK, "订单 ID 是 "+id)
	})

	server.GET("/views/*apple", func(ctx *gin.Context) {
		view := ctx.Param("apple")
		ctx.String(http.StatusOK, "view 是 "+view)
	})

	server.POST("/post", func(ctx *gin.Context) {
		ctx.String(http.StatusOK, "hello, post method")
	})

	server.Run(":8080") // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
