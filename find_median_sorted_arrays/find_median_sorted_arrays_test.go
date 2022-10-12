package find_median_sorted_arrays

import (
	"fmt"
	"testing"
)

func Test_splitArrays(t *testing.T) {
	a := []int{100}
	b := []int{2, 3, 10}
	aMid, bMid := splitArrays(a, b)
	fmt.Println(aMid, bMid)
}

func Test_findMedianSortedArrays(t *testing.T) {
	fmt.Println(findMedianSortedArrays([]int{1, 3}, []int{2}))
}
