package service

import (
	"context"
	"xcoder/internal/model/input/chatin"
	"xcoder/internal/model/input/editin"
)

type IEdit interface {
	SseGenerate(ctx context.Context, in *editin.EditSseGenerateRequest, out chan *chatin.ChatResult) error
}

var localEdit IEdit

func Edit() IEdit {
	if localEdit == nil {
		panic("implement not found for interface IEdit, forgot register?")
	}
	return localEdit
}

func RegisterEdit(i IEdit) {
	localEdit = i
}
