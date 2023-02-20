package merge_channel

import "sync"

func merge(chs ...<-chan int) <-chan int {
	out := make(chan int)
	var wg sync.WaitGroup
	for _, ch := range chs {
		wg.Add(1)
		go func(c <-chan int) {
			defer wg.Done()
			for val := range c {
				out <- val
			}
		}(ch)
	}
	go func() {
		wg.Wait()
		close(out)
	}()
	return out
}
