package routers

import (
	"blogserver/middleware"
	"github.com/devfeel/dotweb"
	"github.com/jinzhu/gorm"
)

var DBConn *gorm.DB

type Router struct {
	server *dotweb.HttpServer
	group dotweb.Group
}

// 路由配置
func NewApiRouter(server *dotweb.HttpServer) *Router{
	router := &Router{server: server, group: server.Group("/api")}
	router.group.Use(middleware.NewCROSMiddleware())
	return router
}