package super_egg_drop

import (
	"fmt"
	"testing"
)

func Test_superEggDrop(t *testing.T) {
	for _, line := range superEggDrop(10, 100) {
		fmt.Println(line)
	}
}
