package median_sliding_window

import (
	"fmt"
	"sort"
	"testing"
)

func Test_medianSlidingWindow(t *testing.T) {
	fmt.Println(medianSlidingWindow([]int{1, 2, 3, 4}, 3))
}

func Test_search(t *testing.T) {
	a := []int{1, 3, 5, 10}
	fmt.Println(sort.Search(len(a), func(i int) bool { return 11 <= a[i] }))
}
