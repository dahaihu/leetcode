package cycle_print

import (
	"fmt"
	"sync"
)

func CondCyclePrint(times int, k int) {
	cond := sync.NewCond(&sync.Mutex{})
	var (
		val int
		wg  sync.WaitGroup
	)
	for i := 0; i < k; i++ {
		wg.Add(1)
		i := i
		go func() {
			defer wg.Done()
			for time := 0; time < times; time++ {
				cond.L.Lock()
				for val%k != i {
					cond.Wait()
				}
				fmt.Printf("worker %d print %d\n", i, val)
				val += 1
				cond.Broadcast()
				cond.L.Unlock()
			}
		}()
	}
	wg.Wait()
}
