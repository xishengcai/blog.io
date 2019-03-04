package routers

import (
	"blog.io/config"
	"blog.io/middleware"
	"github.com/devfeel/dotweb"
	"github.com/jinzhu/gorm"
)

var DBConn *gorm.DB

type Router struct {
	server *dotweb.HttpServer
	group  dotweb.Group
}

// 路由配置
func NewApiRouter(server *dotweb.HttpServer) *Router {
	router := &Router{server: server, group: server.Group("/api")}
	router.group.Use(middleware.NewCROSMiddleware())
	return router
}

func (r *Router) V1() {
	v1 := r.group.Group("/v1")

	if config.Config().EnvProd {
		v1.Use(&middleware.ApiSignMiddleware{})
	}
}
