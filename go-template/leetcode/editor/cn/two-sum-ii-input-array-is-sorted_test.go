package leetcode_solutions

import "testing"

// leetcode submit region begin(Prohibit modification and deletion)
func twoSumIiInputArrayIsSorted(numbers []int, target int) []int {
	first := 0
	last := len(numbers) - 1

	for first < last {
		if numbers[first]+numbers[last] == target {
			return []int{first + 1, last + 1}
		} else if numbers[first]+numbers[last] > target {
			last--
		} else {
			first++
		}
	}

	return nil
}

//leetcode submit region end(Prohibit modification and deletion)

func TestTwoSumIiInputArrayIsSorted(t *testing.T) {
	// your test code here

}
