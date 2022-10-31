package quick_sort

import (
	"fmt"
	"testing"
)

func Test_quickSort(t *testing.T) {
	nums := []int{1, 3, 2, 2, 6, -1}
	quickSort(nums, 0, len(nums)-1)
	fmt.Println(nums)
}
