package service

import (
	"context"
	"xcoder/internal/model/input/chatin"
	"xcoder/internal/model/input/commentin"
)

type IComment interface {
	SseGenerate(ctx context.Context, in *commentin.CodeCommentRequest, out chan *chatin.ChatResult) error
}

var localComment IComment

func Comment() IComment {
	if localComment == nil {
		panic("implement not found for interface IComment, forgot register?")
	}
	return localComment
}

func RegisterComment(i IComment) {
	localComment = i
}
