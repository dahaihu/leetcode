package next_number

import (
	"testing"

	"github.com/go-playground/assert/v2"
)

func Test_nextNumber(t *testing.T) {
	assert.Equal(t, nextNumber([]int{1, 2, 3, 4}, 225), 224)
	assert.Equal(t, nextNumber([]int{1, 4}, 24), 14)
	assert.Equal(t, nextNumber([]int{2, 4}, 21), 4)

}
