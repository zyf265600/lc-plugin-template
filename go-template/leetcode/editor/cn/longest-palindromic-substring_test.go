package leetcode_solutions

import "testing"

// leetcode submit region begin(Prohibit modification and deletion)
func longestPalindrome(s string) string {
	res := ""
	for i := 0; i < len(s); i++ {
		res1 := Palindrome(s, i, i)
		res2 := ""
		if i+1 < len(s) {
			res2 = Palindrome(s, i, i+1)
		}

		if len(res1) > len(res) {
			res = res1
		}
		if len(res2) > len(res) {
			res = res2
		}
	}

	return res
}

func Palindrome(s string, left int, right int) string {
	for left >= 0 && right < len(s) && s[left] == s[right] {
		left--
		right++
	}
	return s[left+1 : right]
}

//leetcode submit region end(Prohibit modification and deletion)

func TestLongestPalindromicSubstring(t *testing.T) {

	// your test code here
	longestPalindrome("cbbd")
}
