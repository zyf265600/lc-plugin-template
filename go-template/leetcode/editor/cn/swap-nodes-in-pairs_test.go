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
func swapPairs(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}

	dummy := &ListNode{-1, head}

	p := dummy

	for p.Next != nil && p.Next.Next != nil {
		node1 := p.Next.Next
		node2 := p.Next
		p.Next = node1
		node2.Next = node1.Next
		node1.Next = node2
		p = node2
	}

	return dummy.Next
}

//leetcode submit region end(Prohibit modification and deletion)

func TestSwapNodesInPairs(t *testing.T) {
	// your test code here

}
