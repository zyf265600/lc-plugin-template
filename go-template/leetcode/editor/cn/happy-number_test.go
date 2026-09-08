package leetcode_solutions

import (
	"testing"
)

// leetcode submit region begin(Prohibit modification and deletion)
func isHappy(n int) bool {
	count := make(map[int]bool)
	for n != 1 && !count[n] {
		count[n] = true
		sum := 0
		for n > 0 {
			sum += n % 10 * (n % 10)
			n = n / 10
		}
		n = sum
	}
	return n==1
}

//leetcode submit region end(Prohibit modification and deletion)

func TestHappyNumber(t *testing.T) {

	// your test code here
	isHappy(19)
}
