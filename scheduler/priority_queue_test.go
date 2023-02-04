package scheduler

import (
	"fmt"
	"math/rand"
	"strconv"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

type demoJob struct {
	executeTime int
	name        string
}

func (j *demoJob) Do() {
	fmt.Printf("executeTime is %d\n", j.executeTime)
	time.Sleep(time.Duration(rand.Int63n(100)) * time.Millisecond)
}

func (j *demoJob) ExecuteTime() int64 {
	return int64(j.executeTime)
}

func shuffle(start, end int) []int {
	rand.Seed(time.Now().Unix())
	nums := make([]int, end-start)
	for i := 0; i < len(nums); i++ {
		nums[i] = i + start
	}
	for i := len(nums) - 1; i >= 0; i-- {
		swapIdx := rand.Intn(i + 1)
		nums[i], nums[swapIdx] = nums[swapIdx], nums[i]
	}
	return nums
}

func Test_priorityQueue(t *testing.T) {
	queue := new(PriorityQueue)
	count := 100
	nums := shuffle(0, count)
	for _, num := range nums {
		num := num
		Push(queue, NewJobWrapper(&demoJob{
			name:        strconv.Itoa(num),
			executeTime: num,
		}))
	}
	for i := 0; i < count; i++ {
		head := Pop(queue).(*JobWrapper)
		assert.Equal(t, head.job.ExecuteTime(), int64(i))
	}
	assert.Equal(t, queue.Len(), 0)
}
