package leetcode_solutions

import "testing"

// leetcode submit region begin(Prohibit modification and deletion)
//
//	func dp(n int, memo []int) int {
//		if n == 0 || n == 1 {
//			return n
//		}
//
//		if memo[n] != -1 {
//			return memo[n]
//		}
//
//		memo[n] = dp(n-1, memo) + dp(n-2, memo)
//
//		return memo[n]
//	}
func fib(n int) int {
	if n == 0 || n == 1 {
		return n
	}

	dp0 := 0
	dp1 := 1
	dp := 0

	for i := 2; i <= n; i++ {
		dp = dp0 + dp1
		dp0 = dp1
		dp1 = dp
	}

	return dp
}

//leetcode submit region end(Prohibit modification and deletion)

func TestFibonacciNumber(t *testing.T) {
	// your test code here

}
