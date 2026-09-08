package leetcode_solutions

import "testing"

// leetcode submit region begin(Prohibit modification and deletion)
func findAnagrams(s string, p string) []int {
	l := len(p)
	n := len(s)
	res := make([]int, 0, n)

	window := make(map[rune]int)
	needs := make(map[rune]int)

	for _, v := range p {
		needs[v]++
	}

	left, right := 0, 0
	valid := 0

	for right < n {
		c := rune(s[right])
		right++
		window[c]++
		if window[c] == needs[c] {
			valid++
		}

		for right-left >= l {
			if valid == len(needs) {
				res = append(res, left)
			}
			d := rune(s[left])
			if window[d] == needs[d] {
				valid--
			}
			window[d]--
			left++
		}
	}
	return res
}

//leetcode submit region end(Prohibit modification and deletion)

func TestFindAllAnagramsInAString(t *testing.T) {
	// your test code here
	findAnagrams("cbaebabacd", "abc")
}
