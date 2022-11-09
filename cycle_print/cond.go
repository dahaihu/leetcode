package cycle_print

import (
	"sync"
)

func CondCyclePrint(workers int, target int) {
	cond := sync.NewCond(&sync.Mutex{})
	var (
		val int
		wg  sync.WaitGroup
	)
	for i := 0; i < workers; i++ {
		wg.Add(1)
		i := i
		go func() {
			defer wg.Done()
			for {
				cond.L.Lock()
				for val%workers != i {
					cond.Wait()
				}
				// fmt.Printf("worker %d print %d\n", i, val)
				val += 1
				cond.Broadcast()
				cond.L.Unlock()
				if val >= target {
					break
				}
			}
		}()
	}
	wg.Wait()
}
