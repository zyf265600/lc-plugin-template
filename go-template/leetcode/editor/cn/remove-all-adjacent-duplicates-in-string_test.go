package leetcode_solutions

import "testing"

// leetcode submit region begin(Prohibit modification and deletion)
func removeDuplicates(s string) string {
	stack := make([]byte, 0, len(s))

	for i := range s {
		c := s[i]
		if len(stack) == 0 {
			stack = append(stack, c)
		} else if stack[len(stack)-1] != c {
			stack = append(stack, c)
		} else {
			stack = stack[:len(stack)-1]
		}
	}

	return string(stack)
}

//leetcode submit region end(Prohibit modification and deletion)

func TestRemoveAllAdjacentDuplicatesInString(t *testing.T) {

	// your test code here
	removeDuplicates("abbaca")
}
