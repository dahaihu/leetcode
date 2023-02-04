package binary_search

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_binarySearch(t *testing.T) {
	assert.Equal(t, search([]int{1, 0, 1, 1, 1}, 0), true)
}

func Test_findMin(t *testing.T) {
	assert.Equal(t, findMin([]int{3, 4, 5, 1, 2}), 1)
	assert.Equal(t, findMin([]int{4, 5, 6, 7, 0, 1, 2}), 0)
	assert.Equal(t, findMin([]int{11, 13, 15, 17}), 11)
}
