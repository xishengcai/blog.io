package middleware

import (
	"blog.io/models"
	"blog.io/common"
	"github.com/devfeel/dotweb"
	"net/http"
)

type ApiSignMiddelware struct {
	dotweb.BaseMiddlware
}

func (asm *ApiSignMiddelware) Handle (ctx dotweb.Context) {
	if sign := ctx.Request().QueryHeader("Sign"); len(sign) <= 0 {
		return ctx.WriteJsonC(http.StatusBadRequest, models.Response{Err: common.ErrSignParams; Data: nil})
	}else {

	}
}