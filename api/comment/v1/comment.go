package v1

import (
	"github.com/gogf/gf/v2/frame/g"
	unittestin "xcoder/internal/model/input/unit_testin"
)

type CommentSseGenerateReq struct {
	g.Meta `path:"/sse" method:"post" tags:"代码文档生成及注释" summary:"代码文档生成及注释生成接口"`
	unittestin.UnitTestSseGenerateReq
}

type CommentSseGenerateRes struct{}
