package leetcode_solutions

import (
	"fmt"
	"testing"
)

// leetcode submit region begin(Prohibit modification and deletion)
func spiralOrder(matrix [][]int) []int {
	var res []int
	y := len(matrix)
	x := len(matrix[0])
	left := 0
	right := x - 1
	top := 0
	bottom := y - 1

	for len(res) < x*y {
		if bottom >= top {
			for i := left; i <= right; i++ {
				res = append(res, matrix[top][i])
			}
			top++
		}

		//fmt.Println(res)

		if right >= left {
			for i := top; i <= bottom; i++ {
				res = append(res, matrix[i][right])
			}
			right--
		}

		//fmt.Println(res)

		if bottom >= top {
			for i := right; i >= left; i-- {
				res = append(res, matrix[bottom][i])
			}
			bottom--
		}

		//fmt.Println(res)

		if right >= left {
			for i := bottom; i >= top; i-- {
				res = append(res, matrix[i][left])
			}
			left++
		}

		//fmt.Println(res)
	}

	return res
}

//leetcode submit region end(Prohibit modification and deletion)

func TestSpiralMatrix(t *testing.T) {
	// your test code here
	matrix := [][]int{{1, 2, 3}, {4, 5, 6}, {7, 8, 9}}
	fmt.Println(spiralOrder(matrix))
}
