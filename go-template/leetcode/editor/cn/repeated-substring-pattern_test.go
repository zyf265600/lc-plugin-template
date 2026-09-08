package leetcode_solutions

import (
	"bytes"
	"testing"
)

// leetcode submit region begin(Prohibit modification and deletion)
func repeatedSubstringPattern(s string) bool {
	ss := []byte(s)                  // abab
	l := len(s)                      // 4
	for i := 1; i <= len(ss)/2; i++ { // 2
		sub := ss[:i] // ab
		if l%len(sub) == 0 {
			for j := 0; j < l/len(sub); j++ { // 0-1
				first := 0 + len(sub)*j       // 2
				last := len(sub) + len(sub)*j // 4
				if !bytes.Equal(sub, ss[first:last]) {
					break
				}
				if last == l {
					return true
				}
			}
		}
	}
	return false
}

//leetcode submit region end(Prohibit modification and deletion)

func TestRepeatedSubstringPattern(t *testing.T) {
	// your test code here
	repeatedSubstringPattern("abab")
}
