package single_flight

import (
	"fmt"
	"sync"
	"testing"
	"time"
)

func demo() (interface{}, error) {
	time.Sleep(time.Second)
	return nil, nil
}

func Test_singleFlight(t *testing.T) {
	g := new(Group)
	var w sync.WaitGroup
	w.Add(10)
	for i := 0; i < 10; i++ {
		go func(idx int) {
			defer w.Done()
			fmt.Printf("start process %d\n", idx)
			_, _ = g.Do("heheda", demo)
		}(i)
	}
	w.Wait()
}
