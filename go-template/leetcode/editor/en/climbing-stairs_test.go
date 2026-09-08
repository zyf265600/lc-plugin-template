package leetcode_solutions

import "testing"

// leetcode submit region begin(Prohibit modification and deletion)
//
//	func stepDown(n int, dp []int) int {
//		if n == 0 {
//			return 1
//		}
//		if n < 0 {
//			return 0
//		}
//		if dp[n] != 0 {
//			return dp[n]
//		}
//		dp[n] = stepDown(n-1, dp) + stepDown(n-2, dp)
//		return dp[n]
//	}
func climbStairs(n int) int {
	//dp := make([]int, n+1)
	//return stepDown(n, dp)
	if n <= 2 {
		return n
	}

	dpLow := 1
	dpHigh := 2
	dp := 0

	for i := 3; i<=n; i++ {
		dp = dpLow + dpHigh
		dpLow = dpHigh
		dpHigh = dp
	}

	return dp
}

//leetcode submit region end(Prohibit modification and deletion)

func TestClimbingStairs(t *testing.T) {
	// your test code here

}
