package can_jump

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_canJump2(t *testing.T) {
	assert.True(t, canJump2([]int{3, 2, 1}) == 1)
}
