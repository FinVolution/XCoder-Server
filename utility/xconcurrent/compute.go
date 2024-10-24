package xconcurrent

import (
	"context"
	"errors"
	"fmt"
	"github.com/gogf/gf/v2/frame/g"
	"time"
)

type Base struct {
	name string
	m    chan struct{}
	Err  error
}

func NewBase(name string) *Base {
	return &Base{
		name: fmt.Sprintf("dao/%s", name),
		m:    make(chan struct{}),
	}
}

func (b *Base) Name() string {
	return b.name
}

func (b *Base) Wait() {
	<-b.m
}

func (b *Base) Close() {
	close(b.m)
}

func (b *Base) Done() <-chan struct{} {
	return b.m
}

func (b *Base) Compute(ctx context.Context, f func(context.Context) error) {
	go func() {
		defer func() {
			if r := recover(); r != nil {
				errMsg := fmt.Sprintf("panic in %s, err: %v", b.name, r)
				g.Log().Error(ctx, errMsg)
			}
		}()
		if err := f(ctx); err != nil {
			err := fmt.Errorf("error in %s, err: %w", b.name, err)
			b.Err = err
			errMsg := err.Error()
			if errors.Is(err, context.Canceled) {
				g.Log().Infof(ctx, "time: %v event cancel error.kind, Dao Compute Canceled, "+
					"message: %s", time.Now(), errMsg,
				)
				return
			}
			g.Log().Error(ctx, errMsg)
		}
	}()
}
