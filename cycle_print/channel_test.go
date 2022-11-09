package cycle_print

import (
	"fmt"
	"testing"
)

func ChannelPrint(workers int, target int) {
	var val int
	cycle := New(workers, func(workeridx int, _ int) bool {
		val++
		fmt.Printf("worker %d print %d\n", workeridx, val)
		return val < target
	})
	cycle.Start()
	<-cycle.Done()
}

func BenchmarkChannelCyclePrint(t *testing.B) {
	for i := 0; i < t.N; i++ {
		ChannelPrint(3, 1000)
	}
}

func BenchmarkCondCyclePrint(t *testing.B) {
	for i := 0; i < t.N; i++ {
		CondCyclePrint(3, 1000)
	}
}
