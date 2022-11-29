package queueforstack

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_queue(t *testing.T) {
	q := Queue[int]{}
	count := 10
	for i := 0; i < count; i++ {
		q.push(i)
	}
	for i := 0; i < count; i++ {
		assert.Equal(t, q.pop(), i)
	}
	assert.Equal(t, q.empty(), true)
}

func Test_stack(t *testing.T) {
	q := queueStack[int]{}
	count := 10
	for i := 0; i < count; i++ {
		q.push(i)
	}
	for i := count - 1; i >= 0; i-- {
		assert.Equal(t, q.pop(), i)
	}
	assert.Equal(t, q.empty(), true)
}

func Test_stackQueue(t *testing.T) {
	s := stackQueue[int]{}
	count := 10
	times := 2
	i := 0
	{
		for ; i < count/times; i++ {
			s.push(i)
		}
		for j := 0; j < count/times/times; j++ {
			assert.Equal(t, j, s.pop())
		}
	}

	{
		for ; i < count; i++ {
			s.push(i)
		}

		for j := count / times / times; j < count; j++ {
			assert.Equal(t, j, s.pop())
		}
	}
}
