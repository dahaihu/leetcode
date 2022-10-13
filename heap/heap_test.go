package heap

import (
	"container/heap"
	"fmt"
	"testing"
)

func Test_heapSort(t *testing.T) {
	h := new(HeapInt)
	*h = []int{1, 3, 2, 6, 4}
	heap.Init(h)
	fmt.Printf("%+v", h)
}
