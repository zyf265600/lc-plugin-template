package leetcode_solutions

import "testing"

//leetcode submit region begin(Prohibit modification and deletion)
func removeElement(nums []int, val int) int {
    if len(nums) == 0 {
		return 0
	}

	fast := 0
	slow := 0

	for fast < len(nums) {
		if nums[fast] != val{
			nums[slow] = nums[fast]
			slow++
		}
		fast++
	}

	return slow
}
//leetcode submit region end(Prohibit modification and deletion)


func TestRemoveElement(t *testing.T) {
	// your test code here

}