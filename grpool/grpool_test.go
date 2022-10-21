package grpool

import (
	"fmt"
	"sync"
	"testing"
)

func Test_grPool(t *testing.T) {
	p := NewPool(10, 5)
	count := 100
	var w sync.WaitGroup
	w.Add(count)
	for i := 0; i < 100; i++ {
		p.JobQueue <- func(idx int) func() {
			return func() {
				defer w.Done()
				fmt.Printf("job %d\n", idx)
			}
		}(i)
	}
	w.Wait()
}
