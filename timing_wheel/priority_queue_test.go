package timing_wheel

import (
	"container/heap"
	"sort"
	"testing"

	"github.com/stretchr/testify/assert"
)

type job struct {
	val int64
}

func (j *job) Priority() int64 {
	return j.val
}

func TestPriorityQueue(t *testing.T) {
	queue := NewPriorityQueue()
	vals := []int64{3, 9, 2, 8, 1}
	for _, val := range vals {
		heap.Push(queue, &job{val: val})
	}
	sort.Slice(vals, func(i, j int) bool { return vals[i] <= vals[j] })
	for idx := 0; queue.Len() != 0; idx++ {
		job := heap.Pop(queue).(*job)
		assert.Equal(t, job.val, vals[idx])
	}
}
