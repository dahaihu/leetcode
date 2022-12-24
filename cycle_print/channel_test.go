package cycle_print

import (
	"fmt"
	"os"
	"runtime"
	"runtime/trace"
	"sync"
	"sync/atomic"
	"testing"
	"time"
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
	ChannelPrint(3, 10)
}

func BenchmarkCondCyclePrint(t *testing.B) {
	for i := 0; i < t.N; i++ {
		CondCyclePrint(3, 1000)
	}
}

func Test_print(t *testing.T) {
	CondCyclePrint(3, 10)
}

func Test_range(t *testing.T) {
	wg := sync.WaitGroup{}
	wg.Add(3)
	words := []string{"foo", "bar", "baz"}
	for _, word := range words {
		word := word
		go func() {
			defer wg.Done()
			fmt.Println(word)
		}()
	}
	wg.Wait()
}

type Foo struct {
	a uint64
	// _ [56]byte
	b uint64
	// _ [56]byte
}

type Foomark struct {
	a uint64
	_ [56]byte
	b uint64
	_ [56]byte
}

func Benchmark_cpuPadding(t *testing.B) {
	var foo Foo
	var wg sync.WaitGroup
	wg.Add(2)
	go func() {
		defer wg.Done()
		for i := 0; i < 1000*1000; i++ {
			atomic.AddUint64(&foo.a, 1)
		}
	}()

	go func() {
		defer wg.Done()
		for i := 0; i < 1000*1000; i++ {
			atomic.AddUint64(&foo.b, 1)
		}
	}()
	wg.Wait()
}

func Benchmark_cpuPaddin1(t *testing.B) {
	var foo Foomark
	var wg sync.WaitGroup
	wg.Add(2)
	go func() {
		defer wg.Done()
		for i := 0; i < 1000*1000; i++ {
			atomic.AddUint64(&foo.a, 1)
		}
	}()

	go func() {
		defer wg.Done()
		for i := 0; i < 1000*1000; i++ {
			atomic.AddUint64(&foo.b, 1)
		}
	}()
	wg.Wait()
}

func Test_hehe(t *testing.T) {
	runtime.GOMAXPROCS(1)
	f, _ := os.Create("hehetrace.out")
	defer f.Close()
	trace.Start(f)
	defer trace.Stop()

	var wg sync.WaitGroup
	wg.Add(1)
	go func() {
		defer wg.Done()
		time.Sleep(time.Millisecond)
		fmt.Println("hello world")
	}()
	go func() {
		for {
		}
	}()
	wg.Wait()
}
