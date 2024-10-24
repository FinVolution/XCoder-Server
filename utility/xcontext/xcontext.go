package xcontext

import (
	"context"
	"time"
)

// WithProtect return a copy of the parent context but if the parent context is cancelled, the returned context will not.
func WithProtect(parent context.Context) (ctx context.Context) {
	if parent == nil {
		panic("cannot create context from nil parent")
	}
	return valueOnlyContext{parent}
}

type valueOnlyContext struct {
	context.Context
}

func (valueOnlyContext) Deadline() (deadline time.Time, ok bool) {
	return
}

func (valueOnlyContext) Done() <-chan struct{} {
	return nil
}

func (valueOnlyContext) Err() error {
	return nil
}
