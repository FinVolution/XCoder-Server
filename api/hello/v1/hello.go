package v1

import (
	"github.com/gogf/gf/v2/frame/g"
)

type HelloReq struct {
	g.Meta `path:"/hello" tags:"Hello" method:"get" summary:"hello api"`
}
type HelloRes struct {
	g.Meta `mime:"text/html" example:"string"`
}

type HsReq struct {
	g.Meta `path:"/hs" tags:"Hello" method:"get" summary:"hs api"`
}

type HsRes struct {
	g.Meta `mime:"text/html" example:"string"`
}
