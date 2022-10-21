package tunny

import (
	"fmt"
	"sync"
	"testing"
	"time"
)

func Test_tunny(t *testing.T) {
	p := NewFunc(10, func(param interface{}) interface{} {
		fmt.Printf("start Processing %v\n", param)
		time.Sleep(time.Second)
		fmt.Printf("done Processing %v\n", param)
		return nil
	})
	var w sync.WaitGroup
	count := 100

	p.SetSize(100)
	w.Add(count)
	for i := 0; i < count; i++ {
		go func(param int) {
			defer w.Done()
			p.Process(param)
		}(i)
	}
	w.Wait()
}
