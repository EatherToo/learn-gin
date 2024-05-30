package router

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {
	router := gin.Default()

	// 省略其他代码
	// 添加 Get 请求路由
	router.GET("/", func(context *gin.Context) {
		context.String(http.StatusOK, "hello gin get method")
	})
	// 添加 Post 请求路由
	router.POST("/", func(context *gin.Context) {
		context.String(http.StatusOK, "hello gin post method")
	})
	// 添加 Put 请求路由
	router.PUT("/", func(context *gin.Context) {
		context.String(http.StatusOK, "hello gin put method")
	})
	// 添加 Delete 请求路由
	router.DELETE("/", func(context *gin.Context) {
		context.String(http.StatusOK, "hello gin delete method")
	})
	// 添加 Patch 请求路由
	router.PATCH("/", func(context *gin.Context) {
		context.String(http.StatusOK, "hello gin patch method")
	})
	// 添加 Head 请求路由
	router.HEAD("/", func(context *gin.Context) {
		context.String(http.StatusOK, "hello gin head method")
	})
	// 添加 Options 请求路由
	router.OPTIONS("/", func(context *gin.Context) {
		context.String(http.StatusOK, "hello gin options method")
	})
	// 省略其他代码

	return router
}
