package middleware

import (
	"blog.io/models"
	"blogserver/common"
	"blogserver/repositories"
	"github.com/devfeel/dotweb"
	"net/http"
	"strings"
)

//修改jwt源码 jwt.go defaultCheckJWT() 如果是sessions 登录登出接口 不做处理
// 不如此处理的话. dotweb的中间件调用貌似有点问题. /api/的中间件会影响所有的中间件
type CrosMiddleware struct {
	dotweb.BaseMiddlware
}

func (cm *CrosMiddleware) Handle(ctx dotweb.Context) error {
	// 插入ip统计
	repo := new(repositories.RedisRepository)
	repo.InsertIp(ctx.Request().QueryHeader("X-Real-IP"))
	ctx.Response().SetHeader("Access-Controller-Allow-Origin", "*")
	ctx.Response().SetHeader("Access-Control-Allow-Headers", "Context-Type,Authorization, Sign")
	ctx.Response().SetHeader("Access-Control-Allow-Methods", "GET, POST, PATCH, DELETE, OPTIONS")

	if strings.Contains(ctx.Request().RequestURI, "v1") && ctx.Request().Method != "OPTIONS" {
		if sign := ctx.Request().QueryHeader("Sign"); len(sign) <= 0 {
			return ctx.WriteJsonC(http.StatusBadRequest,models.Response{Err: common.ErrSignParams, Data: nil})
		}else {
			uri := ctx.Request().RequestURI
			if index := strings.Index(uri, "?"); index != -1 {
				uri = uri[:index]
			}
			if ok := checkSign(sign, uri); !ok {

			}
		}
	}
}

func NewCROSMiddleware() dotweb.Middleware{
	return &CrosMiddleware{}
}