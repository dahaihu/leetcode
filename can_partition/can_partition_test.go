package can_partition

import (
	"fmt"
	"testing"
)

func Test_canPartition(t *testing.T) {
	fmt.Println(canPartition([]int{1, 5, 5, 11}) == true)
	fmt.Println(canPartition([]int{1, 2, 3, 4, 5}) == true)
}
