package router

import (
	"context"
	"github.com/gogf/gf/v2/net/ghttp"
	"xcoder/internal/controller/unit_test"
)

func UnitTest(ctx context.Context, group *ghttp.RouterGroup) {
	group.Group("/v1/unit_test", func(group *ghttp.RouterGroup) {
		group.Bind(unit_test.NewV1())
	})
}
