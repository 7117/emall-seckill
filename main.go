package main

import (
	"gin_project/config"
	"gin_project/controller/eat"
	"github.com/gin-gonic/gin"
)

func main() {

	router := gin.New()

	router.Use(gin.Logger(), gin.Recovery())
	router.Use(eat.MiddlewareA())

	config.Router(router)

	router.Run(":8081")
}
