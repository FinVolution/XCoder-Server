package router

import (
	"context"
	"github.com/gogf/gf/v2/net/ghttp"
	"xcoder/internal/controller/code"
)

func Code(ctx context.Context, group *ghttp.RouterGroup) {
	group.Group("/v1/code", func(group *ghttp.RouterGroup) {
		group.Bind(code.NewV1())
	})
}
