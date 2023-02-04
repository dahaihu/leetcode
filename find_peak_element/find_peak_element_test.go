package find_peak_element

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_findPeakElement(t *testing.T) {
	// 1 or 5
	assert.Equal(t, 5, findPeakElement([]int{1, 2, 1, 3, 5, 6, 4}))
}
