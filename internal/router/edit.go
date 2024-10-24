package router

import (
	"context"
	"github.com/gogf/gf/v2/net/ghttp"
	"xcoder/internal/controller/edit"
)

func Edit(ctx context.Context, group *ghttp.RouterGroup) {
	group.Group("/v1/code_edit", func(group *ghttp.RouterGroup) {
		group.Bind(edit.NewV1())
	})
}
