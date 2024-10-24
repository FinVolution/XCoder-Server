package v1

import (
	"github.com/gogf/gf/v2/frame/g"
	"xcoder/internal/model/input/chatin"
)

type ChatSseGenerateReq struct {
	g.Meta `path:"/sse" method:"post" tags:"智能 Chat" summary:"智能 Chat 生成接口"`
	chatin.ChatSseGenerateReq
}

type ChatSseGenerateRes struct{}

type ChatResult struct {
	Error error  `json:"error"`
	Data  string `json:"data"`
}
