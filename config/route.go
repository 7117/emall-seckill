package config

import (
	"gin_project/controller/create"
	"gin_project/controller/eat"
	"gin_project/controller/log"
	"gin_project/controller/login"
	"gin_project/controller/m"
	"gin_project/controller/mm"
	"gin_project/controller/one"
	"gin_project/controller/session"
	"gin_project/controller/user"
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

	login_router := router.Group("/login")
	login.Router(login_router)

	user_router := router.Group("/reg")
	user.Router(user_router)

	log_router := router.Group("/log")
	log.Router(log_router)

	sess_router := router.Group("/session")
	session.Router(sess_router)

	create_router := router.Group("/create")
	create.Router(create_router)

	one_router := router.Group("/one")
	one.Router(one_router)

	m_router := router.Group("/m")
	m.Router(m_router)

	mm_router := router.Group("/mm")
	mm.Router(mm_router)
}
