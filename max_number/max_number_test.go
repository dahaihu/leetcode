package max_number

import (
	"fmt"
	"testing"
)

func Test_maxNumber(t *testing.T) {
	fmt.Println(largestNumber([]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 0}))
}

func Test_bigger(t *testing.T) {
	nums := []string{"332", "331"}
	fmt.Println(larger(nums[0], nums[1]))
}

func Test_char(t *testing.T) {
	fmt.Println('3', '2')
}
