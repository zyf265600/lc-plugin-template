package leetcode_solutions

import (
	"testing"
)

// leetcode submit region begin(Prohibit modification and deletion)
func pathEncryption(path string) string {
	p := []byte(path)
	for i := range path {
		if p[i] == byte('.') {
			p[i] = byte(' ')
		}
	}
	return string(p)
}

//leetcode submit region end(Prohibit modification and deletion)

func TestTiHuanKongGeLcof(t *testing.T) {
	// your test code here

}
