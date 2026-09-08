package leetcode_solutions

import "testing"

//leetcode submit region begin(Prohibit modification and deletion)
func isValid(s string) bool {
	stack := make([]byte, 0, len(s))
	matchTable := map[byte]byte{
		')': '(',
		'}': '{',
		']': '[',
	}

	for i := range s {
		if s[i] == '(' || s[i] == '{' || s[i] == '[' {
			stack = append(stack, s[i])
		} else {
			if len(stack) == 0 {
				return false
			}
			mark := stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			if matchTable[s[i]] != mark{
				return false
			}
		}
	}

	if len(stack) != 0 {
		return false
	}

	return true
}

//leetcode submit region end(Prohibit modification and deletion)

func TestValidParentheses(t *testing.T) {
	// your test code here

}
