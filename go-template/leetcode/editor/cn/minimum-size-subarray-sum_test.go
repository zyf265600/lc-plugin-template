package leetcode_solutions

import (
	"testing"
)

// leetcode submit region begin(Prohibit modification and deletion)
func minSubArrayLen(target int, nums []int) int {
	n := len(nums)

	res := n + 1
	window := 0
	left, right := 0, 0

	for right < n {
		window += nums[right]
		right++

		for window >= target {
			if right-left < res {
				res = right - left
			}
			window -= nums[left]
			left++
		}
	}
	if res == n+1 {
		res = 0
	}
	return res
}

//leetcode submit region end(Prohibit modification and deletion)

func TestMinimumSizeSubarraySum(t *testing.T) {
	// your test code here

}
