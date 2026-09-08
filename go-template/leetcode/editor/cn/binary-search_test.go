package leetcode_solutions

import "testing"

// leetcode submit region begin(Prohibit modification and deletion)
func search(nums []int, target int) int {
	left := 0
	right := len(nums)-1

	for left <= right {
		mid := (left + right) >> 1
		if nums[mid] == target {
			return mid
		} else if nums[mid] < target {
			left = mid + 1
		} else {
			right = mid - 1
		}
	}
	return -1
}

//leetcode submit region end(Prohibit modification and deletion)

func TestBinarySearch(t *testing.T) {
	// your test code here

}
