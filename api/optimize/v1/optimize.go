package v1

import (
	"github.com/gogf/gf/v2/frame/g"
	unittestin "xcoder/internal/model/input/unit_testin"
)

type OptimizeSseGenerateReq struct {
	g.Meta `path:"/sse" method:"post" tags:"代码优化" summary:"代码优化生成接口"`
	unittestin.UnitTestSseGenerateReq
}

type OptimizeSseGenerateRes struct{}