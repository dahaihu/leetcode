package reverse_k_group

import (
	"bytes"
	"fmt"
	"strconv"
)

type ListNode struct {
	Val  int
	Next *ListNode
}

func (l *ListNode) String() string {
	buf := bytes.Buffer{}
	cur := l
	for cur != nil {
		if cur.Next != nil {
			buf.WriteString(fmt.Sprintf("%d=>", cur.Val))
		} else {
			buf.WriteString(strconv.Itoa(cur.Val))
		}
		cur = cur.Next
	}
	return buf.String()
}

func reverseKGroup(head *ListNode, k int) *ListNode {
	dummy := &ListNode{}
	pre := dummy
	for head != nil {
		cur := head
		for i := 0; i < k; i++ {
			if cur == nil {
				pre.Next = head
				return dummy.Next
			}
			cur = cur.Next
		}
		cur = head
		var segmentPre *ListNode
		for i := 0; i < k; i++ {
			next := cur.Next
			cur.Next = segmentPre
			segmentPre = cur
			cur = next
		}
		pre.Next = segmentPre
		pre = head
		head = cur
	}
	return dummy.Next
}
