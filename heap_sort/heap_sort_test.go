package heapsort

import (
	"fmt"
	"testing"
)

func Test_heapSort(t *testing.T) {
	nums := []int{1, 3, 2, 3, 10}
	heapSort(nums)
	fmt.Println(nums)
}

func Test_SmallHeap(t *testing.T) {
	q := newQuque([]int{3, 2, 1})
	fmt.Println(*q)
	q.push(10)
	fmt.Println(*q)
	for j := 3; j <= 10; j++ {
		q.push(j)
		fmt.Println(*q)
	}
}
