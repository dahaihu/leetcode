package timing_wheel

import (
	"context"
	"sort"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPriorityQueue(t *testing.T) {
	queue := NewDelayQueue(context.Background(), WithBuffer(10))
	vals := []int64{3, 9, 2, 8, 1}
	for _, val := range vals {
		queue.Offer(&slot{expire: val})
	}
	sort.Slice(vals, func(i, j int) bool { return vals[i] <= vals[j] })
	for idx := 0; queue.Len() != 0; idx++ {
		job := <-queue.C
		assert.Equal(t, job.ExecuteTime(), vals[idx])
	}
}
