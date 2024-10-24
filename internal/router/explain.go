package router

import (
	"context"
	"github.com/gogf/gf/v2/net/ghttp"
	"xcoder/internal/controller/explain"
)

func Explain(ctx context.Context, group *ghttp.RouterGroup) {
	group.Group("/v1/code_explain", func(group *ghttp.RouterGroup) {
		group.Bind(explain.NewV1())
	})
}
