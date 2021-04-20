package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func one(context *gin.Context) {
	context.String(http.StatusOK, "aaaa")
}

func main() {

	router := gin.Default()

	v1 := router.Group("/v1")

	{
		v1.GET("one", one)
	}

	//设置html路径
	router.LoadHTMLGlob("template/**/*")
	//第一个别名 第二个是实际的路径
	router.Static("/static", "static")

	router.Run(":8081")
}
