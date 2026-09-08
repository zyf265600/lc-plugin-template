package leetcode_solutions

import "testing"

// leetcode submit region begin(Prohibit modification and deletion)
func searchLBond(scores []int, target int) (int, bool) {
	left := 0
	right := len(scores) - 1
	valid := false
	for left <= right {
		mid := (left + right) >> 1
		if scores[mid] == target {
			right = mid - 1
			valid = true
		} else if scores[mid] > target {
			right = mid - 1
		} else {
			left = mid + 1
		}
	}
	return left, valid
}
func searchRBond(scores []int, target int) (int, bool) {
	left := 0
	right := len(scores) - 1
	valid := false
	for left <= right {
		mid := (left + right) >> 1
		if scores[mid] == target {
			left = mid + 1
			valid = true
		} else if scores[mid] > target {
			right = mid - 1
		} else {
			left = mid + 1
		}
	}
	return right, valid
}
func countTarget(scores []int, target int) int {
	resLeft, validLeft := searchLBond(scores, target)
	resRight, validRight := searchRBond(scores, target)

	if validRight && validLeft {
		return resRight - resLeft + 1
	}

	return 0
}

//leetcode submit region end(Prohibit modification and deletion)

func TestZaiPaiXuShuZuZhongChaZhaoShuZiLcof(t *testing.T) {
	// your test code here

}
