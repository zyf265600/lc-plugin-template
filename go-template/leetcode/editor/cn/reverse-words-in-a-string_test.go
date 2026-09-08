package leetcode_solutions

import (
	"slices"
	"strings"
	"testing"
)

// leetcode submit region begin(Prohibit modification and deletion)
func reverseWords(s string) string {
	words := strings.Fields(s)
	slices.Reverse(words)
	return strings.Join(words, " ")
}

//leetcode submit region end(Prohibit modification and deletion)

func TestReverseWordsInAString(t *testing.T) {
	// your test code here

}
