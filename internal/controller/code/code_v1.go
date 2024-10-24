package code

import (
	"context"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/util/gconv"
	codeV1 "xcoder/api/code/v1"
	"xcoder/internal/controller/utils/response"
	"xcoder/internal/model/input/codein"
	"xcoder/internal/service"
)

type ControllerV1 struct{}

func NewV1() *ControllerV1 {
	return &ControllerV1{}
}

func (c *ControllerV1) Generate(ctx context.Context, req *codeV1.CodeGenerateReq) (
	res *codeV1.CodeGenerateRes, err error) {
	var (
		in = &codein.CodeGenerateRequest{}
		r  = g.RequestFromCtx(ctx)
	)

	err = gconv.Struct(req, in)
	if err != nil {
		return nil, err
	}

	generate, err := service.Code().Generate(ctx, in)
	if err != nil {
		return nil, err
	}

	response.JsonSuccess(r, generate)

	return
}
