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
func kthToLast(head *ListNode, k int) int {

	first := head
	second := head

	for i:=0 ; i < k; i++ {
		first = first.Next
	}

	for first != nil {
		second = second.Next
		first = first.Next
	}

	return second.Val
}
//leetcode submit region end(Prohibit modification and deletion)


func TestKthNodeFromEndOfListLcci(t *testing.T) {
	// your test code here
	
}