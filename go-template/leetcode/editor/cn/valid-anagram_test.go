package leetcode_solutions

import "testing"

// leetcode submit region begin(Prohibit modification and deletion)
func isAnagram(s string, t string) bool {
	if len(s) != len(t) {
		return false
	}

	count := make(map[rune]int)

	for _, v := range s {
		count[v]++
	}
	for _, v := range t {
		count[v]--
		if count[v] < 0 {
			return false
		}
	}
	return true
}

//leetcode submit region end(Prohibit modification and deletion)

func TestValidAnagram(t *testing.T) {
	// your test code here

}
