package merge_channel

import (
	"fmt"
	"testing"
)

func makeChannel(values []int) <-chan int {
	out := make(chan int)
	go func() {
		for _, value := range values {
			out <- value
		}
		close(out)
	}()
	return out
}

func Test_mergeChannel(t *testing.T) {
	chs1 := makeChannel([]int{1, 2, 3})
	chs2 := makeChannel([]int{4, 5, 6, 8})
	merged := merge(chs1, chs2)
	for value := range merged {
		fmt.Println(value)
	}
}
