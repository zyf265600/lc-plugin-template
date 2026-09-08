package leetcode_solutions

import "testing"

//leetcode submit region begin(Prohibit modification and deletion)
func canConstruct(ransomNote string, magazine string) bool {
	if len(ransomNote) > len(magazine) {
		return false
	}

	count := make(map[rune]int)
    for _, v := range magazine {
		count[v] += 1
	}
	for _, v :=  range ransomNote {
		count[v] -= 1
		if count[v] < 0 {
			return false
		}
	}
	return true
}
//leetcode submit region end(Prohibit modification and deletion)


func TestRansomNote(t *testing.T) {
	// your test code here
	
}