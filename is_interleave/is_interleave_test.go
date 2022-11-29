package isinterleave

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_isInterleave(t *testing.T) {
	assert.Equal(t, false, isInterleave("", "", "a"))
}
