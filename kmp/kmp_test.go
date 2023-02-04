package kmp

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_buildTable(t *testing.T) {
	s := "abcdabd"
	fmt.Println(buildJumpTable(s))
}

func Test_kmp(t *testing.T) {
	assert.Equal(t, 15, kmp("BBC ABCDAB ABCDABCDABDE", "ABCDABD"))
}
