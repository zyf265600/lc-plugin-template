package leetcode_solutions

import "testing"

// leetcode submit region begin(Prohibit modification and deletion)
func moveZeroes(nums []int) {
	fast := 0
	slow := 0

	for fast < len(nums) {
		if nums[fast] != 0 {
			nums[slow] = nums[fast]
			slow++
		}
		fast++
	}

	for slow < len(nums) {
		nums[slow] = 0
		slow++
	}
}

//leetcode submit region end(Prohibit modification and deletion)

func TestMoveZeroes(t *testing.T) {
	// your test code here

}
