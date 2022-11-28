package queueforstack

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_queue(t *testing.T) {
	q := Queue[int]{}
	count := 10
	for i := 0; i < count; i++ {
		q.Push(i)
	}
	for i := 0; i < count; i++ {
		assert.Equal(t, q.Pop(), i)
	}
	assert.Equal(t, q.Empty(), true)
}

func Test_stack(t *testing.T) {
	q := Stack[int]{}
	count := 10
	for i := 0; i < count; i++ {
		q.Push(i)
	}
	for i := count - 1; i >= 0; i-- {
		assert.Equal(t, q.Pop(), i)
	}
	assert.Equal(t, q.Empty(), true)
}
