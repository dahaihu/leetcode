package cyclicbarrier

import (
	"context"
	"fmt"
	"sync"
	"testing"
	"time"
)

func Test_channelClosed(t *testing.T) {
	c := make(chan struct{})
	go func() {
		time.Sleep(time.Millisecond)
		close(c)
	}()
	// val := <-c
	// fmt.Printf("value is %v\n", val)
	val, ok := <-c
	fmt.Printf("val is %v, open is %v\n", val, ok)
}

const (
	mutexLocked = 1 << iota // mutex is locked
	mutexWoken
	mutexStarving
	mutexWaiterShift = iota
)

func Test_con(t *testing.T) {
	fmt.Printf("%b\n", -100)
	fmt.Println(mutexLocked, mutexWoken, mutexStarving, mutexWaiterShift)
}

func Test_cyclicBarrier(t *testing.T) {
	workers := 10
	wg := sync.WaitGroup{}
	b := New(workers)
	for i := 0; i < workers; i++ {
		wg.Add(1)
		workerIdx := i
		go func() {
			defer wg.Done()
			for i := 0; i < 100; i++ {
				fmt.Printf("worker %d print %d\n", workerIdx, i)
				err := b.Await(context.Background())
				if err != nil {
					break
				}
			}
		}()
	}
	wg.Wait()
}

func Test_closedChannel(t *testing.T) {
	c := make(chan struct{})
	close(c)
	<-c
	fmt.Println("hello world")
}
