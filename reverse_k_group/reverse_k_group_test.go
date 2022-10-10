package reverse_k_group

import (
	"fmt"
	"testing"
)

func Test_reverseKGroup(t *testing.T) {
	l1 := &ListNode{Val: 0}
	cur := l1
	for i := 1; i < 10; i++ {
		next := &ListNode{Val: i}
		cur.Next = next
		cur = next
	}
	fmt.Println(l1)
	fmt.Println(reverseKGroup(l1, 2))
}
