package router

import (
	"context"
	"github.com/gogf/gf/v2/net/ghttp"
	"xcoder/internal/controller/chat"
)

func Chat(ctx context.Context, group *ghttp.RouterGroup) {
	group.Group("/v1/chat", func(group *ghttp.RouterGroup) {
		group.Bind(chat.NewV1())
	})
}
