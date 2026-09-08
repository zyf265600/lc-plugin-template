package leetcode_solutions

import "testing"

// leetcode submit region begin(Prohibit modification and deletion)
func merge(nums1 []int, m int, nums2 []int, n int) {
	tail := len(nums1) - 1
	p1 := m - 1
	p2 := n - 1

	for p1 >= 0 && p2 >= 0 {
		n1 := nums1[p1]
		n2 := nums2[p2]
		if n1 > n2 {
			nums1[tail] = n1
			p1--
		} else {
			nums1[tail] = n2
			p2--
		}
		tail--
	}

	for p2 >= 0 {
		nums1[tail] = nums2[p2]
		p2--
		tail--
	}
}

//leetcode submit region end(Prohibit modification and deletion)

func TestMergeSortedArray(t *testing.T) {
	// your test code here

}
