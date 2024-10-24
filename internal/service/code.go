package service

import (
	"context"
	codeV1 "xcoder/api/code/v1"
	"xcoder/internal/model/input/codein"
)

type ICode interface {
	Generate(ctx context.Context, req *codein.CodeGenerateRequest) (res *codeV1.CodeGenerateRes, err error)
	UpdateGenerationInfo(ctx context.Context, in *codein.CodeUpdateGenerationInfoReq) (out *codein.CodeUpdateGenerationInfoRes, err error)
}

var localCode ICode

func Code() ICode {
	if localCode == nil {
		panic("implement not found for interface ICode, forgot register?")
	}
	return localCode
}

func RegisterCode(i ICode) {
	localCode = i
}
