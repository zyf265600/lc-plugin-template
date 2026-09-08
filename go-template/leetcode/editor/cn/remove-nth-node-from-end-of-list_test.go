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
func removeNthFromEnd(head *ListNode, n int) *ListNode {
	dummy := &ListNode{Val: -1, Next: head}

	fast := dummy
	slow := dummy

	for i := 0; i < n+1; i++ {
		if fast != nil {
			fast = fast.Next
		}
	}

	for slow != nil && fast != nil {
		fast = fast.Next
		slow = slow.Next
	}

	slow.Next = slow.Next.Next

	return dummy.Next
}

//leetcode submit region end(Prohibit modification and deletion)

func TestRemoveNthNodeFromEndOfList(t *testing.T) {
	// your test code here

}
