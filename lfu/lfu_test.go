package lfu

import (
	"fmt"
	"testing"
)

func TestLFU(t *testing.T) {
	lfu := LFUConstructor(5)
	lfu.Put(10, 13)
	lfu.Put(3, 17)
	lfu.Put(6, 11)
	lfu.Put(10, 5)
	lfu.Put(9, 10)
	fmt.Println(lfu)
	fmt.Println(lfu.Get(13) == -1)
	lfu.Put(2, 19)
	fmt.Println(lfu)
	lfu.Put(2, 18)
	fmt.Println(lfu)
	lfu.Put(4, 20)
	fmt.Println(lfu)
}
