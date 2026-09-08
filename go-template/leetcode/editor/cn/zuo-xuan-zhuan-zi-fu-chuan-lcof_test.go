package leetcode_solutions

import "testing"

//leetcode submit region begin(Prohibit modification and deletion)
func dynamicPassword(password string, target int) string {
	if target == len(password) {
		return password
	}
    first := password[target:]
	last := password[:target]
	return first + last
}
//leetcode submit region end(Prohibit modification and deletion)


func TestZuoXuanZhuanZiFuChuanLcof(t *testing.T) {
	// your test code here
	
}