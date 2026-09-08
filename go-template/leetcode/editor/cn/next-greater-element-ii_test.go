package leetcode_solutions

import "testing"

// leetcode submit region begin(Prohibit modification and deletion)
func nextGreaterElements(nums []int) []int {
	n := len(nums)

	stack := make([]int, 0)
	res := make([]int, n)

	for i := 2 * n - 1; i >= 0; i-- {
		for len(stack) > 0 && nums[i%n] >= stack[len(stack)-1] {
			stack = stack[:len(stack)-1]
		}
		if len(stack) == 0{
			res[i%n] = -1
		} else {
			res[i%n] = stack[len(stack)-1]
		}
		stack = append(stack, nums[i%n])
	}
	return res
}

//leetcode submit region end(Prohibit modification and deletion)

func TestNextGreaterElementIi(t *testing.T) {
	// your test code here

}
