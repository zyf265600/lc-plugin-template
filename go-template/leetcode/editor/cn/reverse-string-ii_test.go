package leetcode_solutions

import (
	//"fmt"
	"testing"
)

// leetcode submit region begin(Prohibit modification and deletion)
func reverseStr(s string, k int) string {
	ss := []byte(s)
	n := len(s)

	for i := 0; i < n; i += 2*k {
		first := i
		last := i + k - 1
		if last >= n {
			last = n - 1
		}

		for first < last {
			ss[first], ss[last] = ss[last], ss[first]
			first++
			last--
		}
	}

	return string(ss)
}

//leetcode submit region end(Prohibit modification and deletion)

func TestReverseStringIi(t *testing.T) {
	// your test code here

}
