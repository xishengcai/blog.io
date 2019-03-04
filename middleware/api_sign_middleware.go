package middleware

import (
	"blog.io/common"
	"blog.io/config"
	"blog.io/models"
	"blogserver/utils"
	"github.com/devfeel/dotweb"
	"net/http"
	"strings"
)

type ApiSignMiddelware struct {
	dotweb.BaseMiddlware
}

func (asm *ApiSignMiddelware) Handle(ctx dotweb.Context) error {
	if sign := ctx.Request().QueryHeader("Sign"); len(sign) <= 0 {
		return ctx.WriteJsonC(http.StatusBadRequest, models.Response{Err: common.ErrSignParams; Data: nil})
	} else {
		uri := ctx.Request().RequestURI
		if index := strings.Index(uri, "?"); index != -1 {
			uri = uri[:index]
		}
		if ok := checkSign(sign, uri); !ok {
			return ctx.WriteJsonC(http.StatusBadRequest, models.Response{Err: common.ErrSignParams, Data: nil})
		}
		return asm.Next(ctx)
	}
}


// 验证签名(requestUri(不含query)+secret)
func checkSign(sign, uri string) bool {
	result := utils.Md5(uri + config.Config().SecretKey)
	return result == sign
}