package max_number

import (
	"bytes"
	"sort"
	"strconv"
)

type cursor struct {
	first  string
	second string
	idx    int
}

func (c *cursor) ends() bool {
	return c.idx == len(c.first)+len(c.second)
}

func (c *cursor) value() byte {
	idx := c.idx
	c.idx++
	if idx >= 0 && idx < len(c.first) {
		return c.first[idx]
	}
	return c.second[idx%len(c.first)]
}

func larger(items0, items1 string) bool {
	c1 := &cursor{first: items0, second: items1, idx: 0}
	c2 := &cursor{first: items1, second: items0, idx: 0}
	for !c1.ends() {
		c1v, c2v := c1.value(), c2.value()
		switch {
		case c1v > c2v:
			return true
		case c1v < c2v:
			return false
		}
	}
	return true
}

func largestNumber(nums []int) string {
	items := make([]string, 0, len(nums))
	allZero := true
	for _, num := range nums {
		if num != 0 {
			allZero = false
		}
		items = append(items, strconv.Itoa(num))
	}
	if allZero {
		return "0"
	}
	sort.Slice(items, func(i, j int) bool { return larger(items[i], items[j]) })
	var result bytes.Buffer
	for _, item := range items {
		for i := 0; i < len(item); i++ {
			result.WriteByte(item[i])
		}
	}
	return result.String()
}
