package cycle_print

import (
	"fmt"
	"sync"
)

func CondCyclePrint(workers int, target int) {
	// cond.Wait()
	// cond.Signal()
	// cond.Broadcast()
	// cond.L.Lock()
	// cond.L.UnLock()
	cond := sync.NewCond(&sync.Mutex{})
	var (
		val int
		wg  sync.WaitGroup
	)
	wg.Add(workers)
	for i := 0; i < workers; i++ {
		i := i
		go func() {
			defer wg.Done()

			loop := true
			for loop {
				cond.L.Lock()
				for val%workers != i {
					cond.Wait()
				}
				if val <= target {
					fmt.Printf("worker %d print %d\n", i, val)
				} else {
					loop = false
				}
				val += 1
				cond.Broadcast()
				cond.L.Unlock()
			}
		}()
	}
	wg.Wait()
}
