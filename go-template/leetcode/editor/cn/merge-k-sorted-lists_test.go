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
func mergeKLists(lists []*ListNode) *ListNode {
	if lists == nil {
		return nil
	}

	amount := len(lists)

	for i := 0; i < amount; i++ {

	}

	return nil
}

//leetcode submit region end(Prohibit modification and deletion)

func TestMergeKSortedLists(t *testing.T) {
	// your test code here
	list := make([]*ListNode, 3)
	list[0] = CreateHead([]int{1, 4, 5})
	list[1] = CreateHead([]int{1, 3, 4})
	list[2] = CreateHead([]int{2, 6})

	result := mergeKLists(list)
	PrintList(result)
}
