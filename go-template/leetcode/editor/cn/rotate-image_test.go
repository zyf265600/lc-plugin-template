package leetcode_solutions

import (
	"testing"
)

// leetcode submit region begin(Prohibit modification and deletion)
func rotate(matrix [][]int) {
	n := len(matrix)
	for y := 0; y < n; y++ {
		for x := y + 1; x < n; x++ {
			matrix[x][y], matrix[y][x] = matrix[y][x], matrix[x][y]
		}
	}

	for i := 0; i < n; i++ {
		//slices.Reverse(matrix[i])
		first := 0
		last := n - 1
		for first < last {
			matrix[i][first], matrix[i][last] = matrix[i][last], matrix[i][first]
			first++
			last--
		}
	}
}

//leetcode submit region end(Prohibit modification and deletion)

func TestRotateImage(t *testing.T) {
	// your test code here

}
