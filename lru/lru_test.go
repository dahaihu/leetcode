package lru

import (
	"fmt"
	"testing"
)

func Test_LRU(t *testing.T) {
	l := NewLRU[int, int](10)
	for i := 0; i < 10; i++ {
		l.Set(i, i)
	}
	fmt.Println(l.List)
	for i := 10; i < 100; i++ {
		l.Set(i, i)
		fmt.Println(l.List)
		fmt.Printf("size is %d\n", l.Len())
	}
}
