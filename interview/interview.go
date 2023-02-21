package interview

import (
	"bytes"
	"container/list"
	"fmt"
)

func merge(nums1, nums2 []int) []int {
	var result []int
	for i, j := 0, 0; i < len(nums1) && j < len(nums2); {
		num1, num2 := nums1[i], nums2[j]
		switch {
		case num1 == num2:
			result = append(result, num1)
			i++
			j++
		case num1 < num2:
			i++
		case num1 > num2:
			j++
		}
	}
	return result
}

func slidingWindowMaxK(nums []int, k int) []int {
	var result []int
	var queue []int
	for i := 0; i < k; i++ {
		for len(queue) != 0 && nums[queue[len(queue)-1]] < nums[i] {
			queue = queue[:len(queue)-1]
		}
		queue = append(queue, i)
	}
	result = append(result, nums[queue[0]])
	for i := k; i < len(nums); i++ {
		for len(queue) != 0 && queue[0] <= i-k {
			queue = queue[1:]
		}
		for len(queue) != 0 && nums[queue[len(queue)-1]] < nums[i] {
			queue = queue[:len(queue)-1]
		}
		queue = append(queue, i)
		result = append(result, nums[queue[0]])
	}
	return result
}

type LRU struct {
	items    map[string]*list.Element
	queue    *list.List
	capacity int
}

type node struct {
	key   string
	value interface{}
}

func (l *LRU) Add(key string, value interface{}) {
	element, ok := l.items[key]
	if ok {
		node := element.Value.(*node)
		node.value = value
		l.queue.MoveToFront(element)
		return
	}
	if length := l.queue.Len(); length == l.capacity {
		last := l.queue.Back()
		l.queue.Remove(last)
		node := last.Value.(*node)
		delete(l.items, node.key)
	}
	node := &node{key: key, value: value}
	element = l.queue.PushFront(node)
	l.items[key] = element
}

func (l *LRU) Get(key string) (val interface{}, ok bool) {
	element, ok := l.items[key]
	if ok {
		l.queue.MoveToFront(element)
		return element.Value.(*node).value, ok
	}
	return nil, false
}

func (l *LRU) String() string {
	var buf bytes.Buffer
	head := l.queue.Front()
	for head != nil {
		node := head.Value.(*node)
		buf.WriteString(fmt.Sprintf("->%s(%+v)", node.key, node.value))
		head = head.Next()
	}
	return buf.String()
}

func longest(nums []int) int {
	mark := make([]int, len(nums))
	mark[0] = 1
	result := 1
	for i := 1; i < len(nums); i++ {
		var cur int
		for j := i - 1; j >= 0; j-- {
			if nums[i] >= nums[j] && mark[j] > cur {
				cur = mark[j]
			}
		}
		mark[i] = cur + 1
		if mark[i] > result {
			result = mark[i]
		}
	}
	return result
}

func asChan(values []int) chan int {
	c := make(chan int)
	go func() {
		for _, v := range values {
			c <- v
			// time.Sleep(time.Duration(rand.Intn(1000)) * time.Millisecond)
		}
		close(c)
	}()
	return c
}

func mergeChan(a, b chan int) chan int {
	c := make(chan int)
	go func() {
		for {
			select {
			case av, an := <-a:
				if an {
					c <- av
				} else {
					for val := range b {
						c <- val
					}
					close(c)
					return
				}
			case bv, bn := <-b:
				if bn {
					c <- bv
				} else {
					for val := range a {
						c <- val
					}
					close(c)
					return
				}
			}
		}
	}()
	return c
}
