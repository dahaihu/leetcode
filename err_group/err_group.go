package err_group

import (
	"context"
	"sync"
)

var token = struct{}{}

type Group struct {
	sema   chan struct{}
	wg     sync.WaitGroup
	err    error
	once   sync.Once
	ctx    context.Context
	cancel context.CancelFunc
}

func New(ctx context.Context, limited int64) *Group {
	g := &Group{}
	if limited > 0 {
		g.sema = make(chan struct{}, limited)
	}
	g.ctx, g.cancel = context.WithCancel(ctx)
	return g
}

func (g *Group) Go(f func(context.Context) error) {
	g.wg.Add(1)
	go func() {
		defer func() {
			if g.sema != nil {
				<-g.sema
			}
			g.wg.Done()
		}()
		if g.sema != nil {
			g.sema <- token
		}
		if err := f(g.ctx); err != nil {
			g.once.Do(func() {
				g.err = err
				g.cancel()
			})
		}
	}()
}

func (g *Group) Wait() error {
	g.wg.Wait()
	// cancel can be repeatedly canceled
	g.cancel()
	return g.err
}
