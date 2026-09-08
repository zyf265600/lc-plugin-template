package leetcode_solutions

import "testing"

// leetcode submit region begin(Prohibit modification and deletion)
func twoSum(nums []int, target int) []int {
	count := make(map[int]int)
	for i, n := range nums {
		index, ok := count[target-n]
		if ok {
			return []int{index, i}
		}
		count[n] = i
	}
	return nil
}

//leetcode submit region end(Prohibit modification and deletion)

func TestTwoSum(t *testing.T) {
	// your test code here

}
