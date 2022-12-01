package gas_station

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_gas(t *testing.T) {
	assert.Equal(t, gas([]int{1, 2, 3, 4, 5}, []int{3, 4, 5, 1, 2}), 3)
}
