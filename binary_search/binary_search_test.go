package binary_search

import (
	"testing"

	"github.com/go-playground/assert/v2"
)

func Test_binarySearch(t *testing.T) {
	assert.Equal(t, search([]int{1, 0, 1, 1, 1}, 0), true)
}
