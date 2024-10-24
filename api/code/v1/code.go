package v1

import (
	"github.com/gogf/gf/v2/frame/g"
	"xcoder/internal/model/input/codein"
)

type CodeGenerateReq struct {
	g.Meta `path:"/generate" method:"post" tags:"代码生成" summary:"代码生成接口"`
	codein.CodeGenerateRequest
}

type CodeGenerateRes struct {
	Code string `json:"code"`
}
