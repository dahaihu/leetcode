package leastinterval

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_leastInterval(t *testing.T) {
	assert.Equal(t, 16, leastInterval([]byte{'A', 'A', 'A', 'A', 'A', 'A', 'B', 'C', 'D', 'E', 'F', 'G'}, 2))
}
