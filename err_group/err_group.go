package err_group

import (
	"sync"
)

var token = struct{}{}

type Group struct {
	err     error
	errOnce sync.Once
	wg      sync.WaitGroup
	sema    chan struct{}
}

// New create err group with limited concurrency. when limited is less than 1,
// it is unlimited with active goroutine. Otherwise the most active goroutine
// number is concurrency
func New(concurrency int) *Group {
	g := new(Group)
	if concurrency > 0 {
		g.sema = make(chan struct{}, concurrency)
	}
	return g
}

func (g *Group) Go(f func() error) {
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
		if err := f(); err != nil {
			g.errOnce.Do(func() {
				g.err = err
			})
		}
	}()
}

func (g *Group) Wait() error {
	g.wg.Wait()
	return g.err
}
