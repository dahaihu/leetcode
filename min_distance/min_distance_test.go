package min_distance

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_minDistance(t *testing.T) {
	assert.Equal(t, minDistance("abc", "abb"), 1)
	assert.Equal(t, minDistance("horse", "ros"), 3)
	assert.Equal(t, minDistance("intention", "execution"), 5)
}
