package config

import (
	"gin_project/controller/eat"
	"github.com/gin-gonic/gin"
	_ "net/http"
)

func Router(router *gin.Engine) {

	//设置html路径
	router.LoadHTMLGlob("template/**/*")
	//第一个别名 第二个是实际的路径
	router.Static("/static", "static")

	eat_router := router.Group("/eat")

	eat.Router(eat_router)
}
