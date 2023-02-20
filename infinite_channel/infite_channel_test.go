package infinite_channel

import (
	"sort"
	"sync"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_InfiniteChannel(t *testing.T) {
	var swg sync.WaitGroup
	ich := New()
	for i := 0; i < 100; i++ {
		swg.Add(1)
		cur := i
		go func() {
			defer swg.Done()
			for j := cur * 10; j < (cur+1)*10; j++ {
				ich.In <- j
			}
		}()
	}
	var rwg sync.WaitGroup
	var result []int
	l := &sync.Mutex{}
	for i := 0; i < 10; i++ {
		rwg.Add(1)
		go func() {
			defer rwg.Done()
			for {
				val, ok := <-ich.Out
				if !ok {
					break
				}
				l.Lock()
				result = append(result, val.(int))
				l.Unlock()
			}
		}()
	}
	swg.Wait()
	ich.Close()
	rwg.Wait()

	sort.Ints(result)
	for idx, val := range result {
		assert.Equalf(t, idx, val, "idx %d val is not %d", idx, val)
	}
}
