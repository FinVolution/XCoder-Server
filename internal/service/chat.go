package service

import (
	"context"
	"xcoder/internal/model/input/chatin"
)

type IChat interface {
	SseGenerate(ctx context.Context, in *chatin.ChatSseGenerateReq, out chan *chatin.ChatResult) error
	UpdateChatRecordsFields(ctx context.Context, in *chatin.ChatMessageUpdateMysqlReq) (out *chatin.ChatMessageUpdateMysqlRes, err error)
}

var localChat IChat

func Chat() IChat {
	if localChat == nil {
		panic("implement not found for interface IChat, forgot register?")
	}
	return localChat
}

func RegisterChat(i IChat) {
	localChat = i
}
