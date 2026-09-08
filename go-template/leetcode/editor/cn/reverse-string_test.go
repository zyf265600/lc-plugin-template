package leetcode_solutions

import "testing"

//leetcode submit region begin(Prohibit modification and deletion)
func reverseString(s []byte)  {
    first := 0
	last := len(s)-1

	for first < last {
		temp := s[last]
		s[last] = s[first]
		s[first] = temp
		first++
		last--
	}
}
//leetcode submit region end(Prohibit modification and deletion)


func TestReverseString(t *testing.T) {
	// your test code here
	
}