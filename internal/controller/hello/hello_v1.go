package hello

import (
	"context"
	"github.com/gogf/gf/v2/frame/g"
	v1 "xcoder/api/hello/v1"
)

type ControllerV1 struct{}

func NewV1() *ControllerV1 {
	return &ControllerV1{}
}

func (c *ControllerV1) Hello(ctx context.Context, req *v1.HelloReq) (res *v1.HelloRes, err error) {
	g.RequestFromCtx(ctx).Response.Writeln("Hello World!")
	return
}

func (c *ControllerV1) Hs(ctx context.Context, req *v1.HsReq) (res *v1.HsRes, err error) {
	g.RequestFromCtx(ctx).Response.Writeln("ok")
	return
}
