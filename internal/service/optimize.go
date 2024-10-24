package service

import (
	"context"
	"xcoder/internal/model/input/chatin"
	"xcoder/internal/model/input/optimizein"
)

type IOptimize interface {
	SseGenerate(ctx context.Context, in *optimizein.OptimizeSseGenerateRequest, out chan *chatin.ChatResult) error
}

var localOptimize IOptimize

func Optimize() IOptimize {
	if localOptimize == nil {
		panic("implement not found for interface IOptimize, forgot register?")
	}
	return localOptimize
}

func RegisterOptimize(i IOptimize) {
	localOptimize = i
}
