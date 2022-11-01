package sema

import (
	"context"
	"sync"
	"testing"
	"time"
)

func Test_sema(t *testing.T) {
	sem := NewWeighted(100)
	var wg sync.WaitGroup
	n := 100
	wg.Add(n)
	for i := 0; i < n; i++ {
		i := i
		go func() {
			defer wg.Done()
			sem.Acquire(context.Background(), int64(i))
			time.Sleep(time.Millisecond * 1)
			sem.Release(int64(i))
		}()
	}
	wg.Wait()
}
