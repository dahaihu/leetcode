package wiggle_max_length

import (
	"testing"

	"github.com/go-playground/assert/v2"
)

func Test_wiggleMaxLength(t *testing.T) {
	assert.Equal(t, 3, wiggleMaxLength([]int{1, 3, 2}))
}
