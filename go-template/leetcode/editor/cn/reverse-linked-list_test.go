package leetcode_solutions

import "testing"

//leetcode submit region begin(Prohibit modification and deletion)
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func reverseList(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}

	last := reverseList(head.Next)

	head.Next.Next = head
	head.Next = nil

	return last
	//if head == nil {
	//	return nil
	//}
	//
	//fast := head.Next
	//slow := head
	//
	//for fast != nil && slow != nil {
	//	fnext := fast.Next
	//	snext := fast
	//
	//	fast.Next = slow
	//	fast = fnext
	//	slow = snext
	//}
	//
	//head.Next = nil
	//
	//return slow
}

//leetcode submit region end(Prohibit modification and deletion)

func TestReverseLinkedList(t *testing.T) {
	// your test code here

}
