package leetcode_solutions

import "testing"

// leetcode submit region begin(Prohibit modification and deletion)
func checkInclusion(s1 string, s2 string) bool {
	n := len(s2)
	count := make(map[rune]int)
	window := make(map[rune]int)
	valid := 0
	for _, v := range s1 {
		count[v]++
	}
	left, right := 0, 0
	for right < n {
		c := rune(s2[right])
		right++
		window[c]++
		if count[c] == window[c] {
			valid++
		}

		for right-left >= len(s1) {
			if right-left == len(s1) && valid == len(count) {
				return true
			}
			d := rune(s2[left])
			if count[d] == window[d] {
				valid--
			}
			window[d]--

			left++
		}
	}
	return false
}

//leetcode submit region end(Prohibit modification and deletion)

func TestPermutationInString(t *testing.T) {
	// your test code here

}
