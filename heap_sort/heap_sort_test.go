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
