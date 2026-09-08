package leetcode_solutions

import (
	"math"
	"testing"
)

// leetcode submit region begin(Prohibit modification and deletion)
func minWindow(s string, t string) string {
	if len(t) > len(s) {
		return ""
	}

	window := map[rune]int{}
	needs := map[rune]int{}

	for _, v := range t {
		needs[v]++
	}

	ls := len(s)
	lt := len(needs) // unique keys of needs string

	left, right := 0, 0
	resStart, resLength := 0, math.MaxInt32
	var valid int

	for right < ls {
		c := rune(s[right])
		right++

		window[c]++
		if window[c] == needs[c] {
			valid++
		}

		for valid == lt {
			if resLength > right-left {
				resLength = right - left
				resStart = left
			}

			cDeleting := rune(s[left])
			if window[cDeleting] == needs[cDeleting] { // 必须放前面，也就是==判定，因为第一次进缩的loop时needs比window只多不少，先减window会多删valid
				valid--
			}
			window[cDeleting]--

			left++
		}
	}
	if resLength == math.MaxInt32 {
		return ""
	}
	return s[resStart : resStart+resLength]
}

//leetcode submit region end(Prohibit modification and deletion)

func TestMinimumWindowSubstring(t *testing.T) {
	// your test code here
	minWindow("aaaaaaaaaaaabbbbbcdd", "abcdd")
}
