package cmd

import (
	"context"
	"xcoder/internal/router"
	"xcoder/internal/service"

	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/net/ghttp"
	"github.com/gogf/gf/v2/os/gcmd"
)

var (
	Main = gcmd.Command{
		Name:        "main",
		Usage:       "main",
		Brief:       "start http server",
		Description: "默认启动所有服务",
		Func: func(ctx context.Context, parser *gcmd.Parser) (err error) {
			g.Log().Debug(ctx, "start http server")

			s := g.Server()

			s.BindMiddleware("/*any", []ghttp.HandlerFunc{
				ghttp.MiddlewareCORS,
				service.Middleware().RecoveryHandler,
				service.Middleware().ErrResponseHandler,
			}...)

			s.Group("/", func(group *ghttp.RouterGroup) {
				router.Hs(ctx, group)
				router.Code(ctx, group)
				router.Chat(ctx, group)
				router.UnitTest(ctx, group)
				router.Explain(ctx, group)
				router.Optimize(ctx, group)
				router.Edit(ctx, group)
				router.Comment(ctx, group)
			})

			s.Run()
			return nil
		},
	}
)
