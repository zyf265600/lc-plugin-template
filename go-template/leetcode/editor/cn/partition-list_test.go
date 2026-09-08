package leetcode_solutions

import (
	"testing"
)

//leetcode submit region begin(Prohibit modification and deletion)
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func partition(head *ListNode, x int) *ListNode {
	if head == nil {
		return nil
	}

	less := &ListNode{-1, nil}
	elarge := &ListNode{-1, nil}

	l := less
	el := elarge
	scan := head

	for scan != nil {
		if scan.Val < x {
			l.Next = scan
			l = l.Next
		} else {
			el.Next = scan
			el = el.Next
		}
		scan = scan.Next
		l.Next = nil
		el.Next = nil
	}

	l.Next = elarge.Next

	return less.Next
}

//leetcode submit region end(Prohibit modification and deletion)

func TestPartitionList(t *testing.T) {
	// your test code here
	l := CreateHead([]int{1, 4, 3, 2, 5, 2})
	x := 3

	result := partition(l, x)
	PrintList(result)
}
