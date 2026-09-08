package leetcode_solutions

import "testing"

// leetcode submit region begin(Prohibit modification and deletion)
func searchLeafBond(nums []int, target int) int {
	left := 0
	right := len(nums) - 1

	for left <= right {
		mid := (right + left) >> 1
		if nums[mid] == target {
			right = mid - 1
		} else if nums[mid] > target {
			right = mid - 1
		} else {
			left = mid + 1
		}
	}
	if left < 0 || left >= len(nums) {
		return -1;
	}
	if nums[left] != target {
		return -1
	}

	return left
}
func searchRightBond(nums []int, target int) int {
	left := 0
	right := len(nums) - 1

	for left <= right {
		mid := (right + left) >> 1
		if nums[mid] == target {
			left = mid + 1
		} else if nums[mid] > target {
			right = mid - 1
		} else {
			left = mid + 1
		}
	}
	if right < 0 || right >= len(nums) {
		return -1;
	}
	if nums[right] != target {
		return -1
	}

	return right
}
func searchRange(nums []int, target int) []int {
	return []int{searchLeafBond(nums, target), searchRightBond(nums, target)}
}

//leetcode submit region end(Prohibit modification and deletion)

func TestFindFirstAndLastPositionOfElementInSortedArray(t *testing.T) {
	// your test code here

}
