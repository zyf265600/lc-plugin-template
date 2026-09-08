package leetcode_solutions

import (
	"testing"
)

// leetcode submit region begin(Prohibit modification and deletion)
func lengthOfLongestSubstring(s string) int {
	res := 0
	window := make(map[rune]int)

	left, right := 0, 0

	for right < len(s) {
		c := rune(s[right])
		right++
		window[c]++
		if len(window) == right-left {
			if res < right-left {
				res = right - left
			}
		}
		for len(window) < right-left {
			d := rune(s[left])
			left++
			window[d]--
			if window[d] == 0 {
				delete(window, d)
			}
		}
	}

	return res
}

//leetcode submit region end(Prohibit modification and deletion)

func TestLongestSubstringWithoutRepeatingCharacters(t *testing.T) {
	// your test code here

}
