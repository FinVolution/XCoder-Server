package service

import (
	"github.com/gogf/gf/v2/net/ghttp"
)

type IMiddleware interface {
	// RecoveryHandler 恢复中间件
	RecoveryHandler(r *ghttp.Request)
	// ErrResponseHandler HTTP响应及错误返回处理中间件
	ErrResponseHandler(r *ghttp.Request)
}

var localMiddleware IMiddleware

func Middleware() IMiddleware {
	if localMiddleware == nil {
		panic("implement not found for interface IMiddleware, forgot register?")
	}
	return localMiddleware
}

func RegisterMiddleware(i IMiddleware) {
	localMiddleware = i
}
