package v1

import (
	"github.com/gogf/gf/v2/frame/g"
	unittestin "xcoder/internal/model/input/unit_testin"
)

type ExplainSseGenerateReq struct {
	g.Meta `path:"/sse" method:"post" tags:"代码解释" summary:"代码解释生成接口"`
	unittestin.UnitTestSseGenerateReq
}

type ExplainSseGenerateRes struct{}