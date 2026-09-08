package leetcode_solutions

import (
	"slices"
	"testing"
)

// leetcode submit region begin(Prohibit modification and deletion)
func nextGreaterElement(nums1 []int, nums2 []int) []int {
	n2 := len(nums2)

	nextLarge := make([]int, n2)
	stack := make([]int, 0)
	res := make([]int, 0)

	for i := n2 - 1; i >= 0; i-- {
		number := nums2[i]
		for len(stack) > 0 && number > stack[len(stack)-1] {
			stack = stack[:len(stack)-1]
		}
		if len(stack) == 0 {
			nextLarge[i] = -1
		} else {
			nextLarge[i] = stack[len(stack)-1]
		}
		stack = append(stack, number)
	}

	for i := range nums1 {
		index := slices.Index(nums2, nums1[i])
		res = append(res, nextLarge[index])
	}

	return res
}

//leetcode submit region end(Prohibit modification and deletion)

func TestNextGreaterElementI(t *testing.T) {
	// your test code here

}
