package service

import (
	"context"
	"xcoder/internal/model/input/chatin"
	"xcoder/internal/model/input/unit_testin"
)

type IUnit interface {
	SseGenerate(ctx context.Context, in *unit_testin.UnitTestSseGenerateReq, out chan *chatin.ChatResult) error
}

var localUnit IUnit

func Unit() IUnit {
	if localUnit == nil {
		panic("implement not found for interface IUnit, forgot register?")
	}
	return localUnit
}

func RegisterUnit(i IUnit) {
	localUnit = i
}
