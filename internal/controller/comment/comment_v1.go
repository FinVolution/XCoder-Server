package comment

import (
	"context"
	"fmt"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/util/gconv"
	"xcoder/api/comment/v1"
	"xcoder/internal/controller/utils"
	"xcoder/internal/model/input/chatin"
	"xcoder/internal/model/input/commentin"
	"xcoder/internal/service"
	"xcoder/utility/xruntime"
)

type ControllerV1 struct{}

func NewV1() *ControllerV1 {
	return &ControllerV1{}
}

func (c *ControllerV1) SseGenerate(ctx context.Context, req *v1.CommentSseGenerateReq) (
	resp *v1.CommentSseGenerateRes, err error) {
	r := utils.GetStreamingChatReq(ctx)

	result := make(chan *chatin.ChatResult)
	// 异步获取 chat 结果
	go func() {
		defer close(result)
		defer func() {
			if r := recover(); r != nil {
				errMsg := fmt.Sprintf("panic in %s, err: %v", "chat.Call", r)
				mstack := xruntime.MStack(2, 5)
				errMsgWithStack := fmt.Sprintf("%s, stack:\n%s", errMsg, mstack)
				g.Log().Error(ctx, errMsgWithStack)
			}
		}()

		in := &commentin.CodeCommentRequest{}
		err := gconv.Struct(req, in)
		if err != nil {
			return
		}
		err = service.Comment().SseGenerate(ctx, in, result)
		if err != nil {
			g.Log().Errorf(ctx, "ChatGenerate Call err: %v", err)
			result <- &chatin.ChatResult{Error: err, Data: ""}
			return
		}
	}()

	utils.ParseStreamingChatResp(ctx, r, result)

	return nil, nil
}
