package eat

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func Hello(context *gin.Context) {
	context.String(http.StatusOK, "Hello")
}

func Router(router *gin.RouterGroup) {
	{
		router.GET("/world", Hello)
	}
}
