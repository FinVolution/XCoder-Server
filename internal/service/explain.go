package service

import (
	"context"
	"xcoder/internal/model/input/chatin"
	"xcoder/internal/model/input/explainin"
)

type IExplain interface {
	SseGenerate(ctx context.Context, in *explainin.ExplainSseGenerateRequest, out chan *chatin.ChatResult) error
}

var localExplain IExplain

func Explain() IExplain {
	if localExplain == nil {
		panic("implement not found for interface IExplain, forgot register?")
	}
	return localExplain
}

func RegisterExplain(i IExplain) {
	localExplain = i
}
