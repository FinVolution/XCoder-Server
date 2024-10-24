package v1

import (
	"github.com/gogf/gf/v2/frame/g"
	"xcoder/internal/model/input/editin"
)

type EditSseGenerateReq struct {
	g.Meta `path:"/sse" method:"post" tags:"代码编辑" summary:"代码编辑生成接口"`
	editin.EditSseGenerateRequest
}

type EditSseGenerateRes struct{}
