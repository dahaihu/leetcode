package scheduler

import (
	"fmt"
	"strconv"
	"testing"
	"time"
)

func unixMill() func() int64 {
	return func() int64 {
		return time.Now().UnixMilli()
	}
}

func Test_scheduleAsTime(t *testing.T) {
	scheduler := NewSchduler(100, time.Millisecond, unixMill())
	for i := 0; i < 100; i++ {
		var executeTime int64
		if i < 10 {
			executeTime = time.Now().UnixMilli() + int64(i)*10
		} else {
			executeTime = time.Now().UnixMilli() + 100
		}
		scheduler.Push(&demoJob{
			name:        strconv.Itoa(int(executeTime)),
			executeTime: int(executeTime),
		})
	}
	time.Sleep(time.Second * 10)
}

func Test_stopScheduler(t *testing.T) {
	scheduler := NewSchduler(10, time.Millisecond, unixMill())
	for i := 0; i < 10; i++ {
		var executeTime int64
		if i < 10 {
			executeTime = time.Now().UnixMilli() + int64(i+1)*10000
		} else {
			executeTime = time.Now().UnixMilli() + 100
		}
		scheduler.Push(&demoJob{
			name:        strconv.Itoa(int(executeTime)),
			executeTime: int(executeTime),
		})
	}
	scheduler.Stop(func(j IJob) {
		fmt.Println(j.ExecuteTime())
	})
}
