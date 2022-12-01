package find_peak_element

import (
	"testing"

	"github.com/go-playground/assert/v2"
)

func Test_findPeakElement(t *testing.T) {
	assert.Equal(t, 2, findPeakElement([]int{1, 2, 3, 1, 0}))
}
