package middleware

import (
	"github.com/gogf/gf/v2/net/ghttp"
)

func (s *sMiddleware) RecoveryHandler(r *ghttp.Request) {
	r.Middleware.Next()
}
