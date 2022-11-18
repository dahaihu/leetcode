package threeappear

import (
	"fmt"
	"testing"
)

func Test_threeAppearTimes(t *testing.T) {
	fmt.Println(maxThreeAppearTimes([]string{"a", "c", "b", "d", "a", "e", "e"}))
}
