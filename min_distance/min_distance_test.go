package min_distance

import (
	"testing"

	"github.com/go-playground/assert/v2"
)

func Test_minDistance(t *testing.T) {
	assert.Equal(t, minDistance("abc", "abb"), 1)
	assert.Equal(t, minDistance("horse", "ros"), 3)
	assert.Equal(t, minDistance("intention", "execution"), 5)
}
