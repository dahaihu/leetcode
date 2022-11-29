package reorderlist

import (
	"fmt"
	"testing"
)

func Test_reorderList(t *testing.T) {
	dummy := &Node[int]{}
	cur := dummy

	for i := 0; i < 5; i++ {
		cur.Next = &Node[int]{Val: i}
		cur = cur.Next
	}
	fmt.Println(dummy.Next)
	fmt.Println(reorderList(dummy.Next))
}
