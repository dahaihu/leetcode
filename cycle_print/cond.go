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

var token = struct{}{}

type cycleChannel []chan struct{}

func newCycleChannel(n int) cycleChannel {
	chs := make(cycleChannel, n)
	for i := 0; i < n; i++ {
		chs[i] = make(chan struct{})
	}
	return chs
}

func (c cycleChannel) start() {
	c.notify(0)
}

func (c cycleChannel) end() {
	for _, ch := range c {
		close(ch)
	}
}

func (c cycleChannel) notify(i int) {
	if i == len(c) {
		i = 0
	}
	c[i] <- token
}

func (c cycleChannel) wait(i int) (ok bool) {
	_, ok = <-c[i]
	return ok
}

func (c cycleChannel) cyclePrint(idx int, val *int, target int) {
	for {
		ok := c.wait(idx)
		if !ok {
			break
		}
		*val = *val + 1
		fmt.Printf("worker %d print %d\n", idx, *val)
		if *val >= target {
			c.end()
			break
		}
		c.notify(idx + 1)
	}
}

func ChannelPrint(workers int, target int) {
	channels := newCycleChannel(workers)
	var (
		cur int
		wg  sync.WaitGroup
	)
	wg.Add(workers)
	for i := 0; i < workers; i++ {
		workerIdx := i
		go func() {
			defer wg.Done()

			channels.cyclePrint(workerIdx, &cur, target)
		}()
	}
	channels.start()
	wg.Wait()
}
