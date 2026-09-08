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
func deleteDuplicates(head *ListNode) *ListNode {
    if head == nil {
		return nil
	}

	fast := head
	slow := head

	for fast != nil {
		if fast.Val != slow.Val {
			slow.Next = fast
			slow = slow.Next
		}

		fast = fast.Next
	}

	slow.Next = nil

	return head
}
//leetcode submit region end(Prohibit modification and deletion)


func TestRemoveDuplicatesFromSortedList(t *testing.T) {
	// your test code here
	
}