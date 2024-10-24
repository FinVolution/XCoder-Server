package router

import (
	"context"
	"github.com/gogf/gf/v2/net/ghttp"
	"xcoder/internal/controller/comment"
)

func Comment(ctx context.Context, group *ghttp.RouterGroup) {
	group.Group("/v1/code_comment", func(group *ghttp.RouterGroup) {
		group.Bind(comment.NewV1())
	})
}
