package sorted

import (
	"fmt"
	"sync"
	"testing"
	"time"
	"unsafe"
)

func Test_sortedPow(t *testing.T) {
	fmt.Println(sorted2powSorted([]int{-3, -2, 0, 1, 6}))
}

func Test_align(t *testing.T) {
	var a int64
	fmt.Println(unsafe.Alignof(&a))
	var b int32
	fmt.Println(unsafe.Alignof(&b))
}

func Test_sema(t *testing.T) {
	var wg sync.WaitGroup
	var count int
	var ch = make(chan bool, 1)
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func() {
			ch <- true
			count++
			time.Sleep(time.Millisecond)
			count--
			<-ch
			wg.Done()
		}()
	}
	wg.Wait()
	fmt.Println(count)
}
