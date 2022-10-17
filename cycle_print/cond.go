package cycle_print

import (
	"fmt"
	"sync"
)

func CondCyclePrint(times int, k int) {
	cond := sync.NewCond(&sync.Mutex{})
	var (
		t int
		w sync.WaitGroup
	)
	w.Add(k)
	for i := 0; i < k; i++ {
		go func(printValue int) {
			defer w.Done()
			for i := 0; i < times; i++ {
				cond.L.Lock()
				for t != printValue {
					cond.Wait()
				}
				fmt.Printf("%d_%d\n", i, printValue)
				t = (t + 1) % k
				cond.Broadcast()
				cond.L.Unlock()
			}
		}(i)
	}
	w.Wait()
}
