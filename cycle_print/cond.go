package cycle_print

import (
	"fmt"
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
			loop := true
			for loop {
				cond.L.Lock()
				for val%workers != i {
					cond.Wait()
				}
				if val < target {
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
