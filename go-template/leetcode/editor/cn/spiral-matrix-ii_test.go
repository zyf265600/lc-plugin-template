package leetcode_solutions

import (
	//"fmt"
	"testing"
)

// leetcode submit region begin(Prohibit modification and deletion)
func elementGenerator(n int) []int {
	var nums []int = make([]int, 0, n)
	for i := 1; i <= n*n; i++ {
		nums = append(nums, i)
	}
	return nums
}
func generateMatrix(n int) [][]int {
	nums := elementGenerator(n)

	res := make([][]int, n)
	for i := 0; i < n; i++ {
		res[i] = make([]int, n)
	}

	np := 0
	left, top := 0, 0
	right, bottom := n-1, n-1

	for np < n*n {
		if top <= bottom {
			for i := left; i <= right; i++ {
				res[top][i] = nums[np]
				np++
			}
			top++
		}

		if left <= right {
			for i := top; i <= bottom; i++ {
				res[i][right] = nums[np]
				np++
			}
			right--
		}

		if top <= bottom {
			for i := right; i >= left; i-- {
				res[bottom][i] = nums[np]
				np++
			}
			bottom--
		}

		if left <= right {
			for i := bottom; i >= top; i-- {
				res[i][left] = nums[np]
				np++
			}
			left++
		}
	}

	return res
}

//leetcode submit region end(Prohibit modification and deletion)

func TestSpiralMatrixIi(t *testing.T) {
	// your test code here
	generateMatrix(3)
}
