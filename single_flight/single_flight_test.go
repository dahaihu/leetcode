package single_flight

import (
	"fmt"
	"sync"
	"testing"
	"time"

	"golang.org/x/sync/singleflight"
)

func demo() (interface{}, error) {
	time.Sleep(time.Second)
	return nil, nil
}

func Test_singleFlight(t *testing.T) {
	g := new(singleflight.Group)
	var w sync.WaitGroup
	for i := 0; i < 10; i++ {
		w.Add(1)
		go func(idx int) {
			defer w.Done()
			fmt.Printf("start process %d\n", idx)
			_, _ = g.Do("heheda", demo)
		}(i)
	}
	w.Wait()
}
