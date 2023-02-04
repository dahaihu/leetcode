package wiggle_max_length

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_wiggleMaxLength(t *testing.T) {
	assert.Equal(t, 3, wiggleMaxLength([]int{1, 3, 2}))
}
