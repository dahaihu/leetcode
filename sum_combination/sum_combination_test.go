package sum_combination

import (
	"fmt"
	"testing"
)

func Test_sumCombination(t *testing.T) {
	fmt.Println(combinationSum([]int{1, 2, 3}, 3))
}

func Test_sumCombination2(t *testing.T) {
	fmt.Println(combinationSum2([]int{1, 1, 1, 2, 2, 3}, 3))
}

func _append(a []int) []int {
	a = append(a, 10)
	return a
}

func Test_copy(t *testing.T) {
	a := []int{1, 2, 3}
	_append(a[:])
	fmt.Println(a)
}
