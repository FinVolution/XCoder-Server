package v1

import (
	"github.com/gogf/gf/v2/frame/g"
	unittestin "xcoder/internal/model/input/unit_testin"
)

type UnitTestSseGenerateReq struct {
	g.Meta `path:"/sse" method:"post" tags:"单元测试" summary:"单元测试生成接口"`
	unittestin.UnitTestSseGenerateReq
}

type UnitTestSseGenerateRes struct{}
