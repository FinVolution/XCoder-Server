package router

import (
	"context"
	"github.com/gogf/gf/v2/net/ghttp"
	"xcoder/internal/controller/optimize"
)

func Optimize(ctx context.Context, group *ghttp.RouterGroup) {
	group.Group("/v1/code_optimize", func(group *ghttp.RouterGroup) {
		group.Bind(optimize.NewV1())
	})
}
