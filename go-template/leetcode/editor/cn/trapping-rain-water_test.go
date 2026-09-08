package leetcode_solutions

import (
	"fmt"
	"testing"
)

// leetcode submit region begin(Prohibit modification and deletion)
func trap(height []int) int {
	leftToRight := make([]int, len(height))
	rightToLeft := make([]int, len(height))
	res := 0

	stack := make([]int, 0)

	for i := 0; i < len(height); i++ {
		for len(stack) > 0 && height[i] > stack[len(stack)-1] {
			stack = stack[:len(stack)-1]
		}
		if len(stack) == 0 {
			rightToLeft[i] = -1
		} else {
			rightToLeft[i] = stack[len(stack)-1]
		}
		stack = append(stack, height[i])
	}

	fmt.Println(rightToLeft)

	stack = []int{}

	for i := len(height) - 1; i >= 0; i-- {
		for len(stack) > 0 && height[i] > stack[len(stack)-1] {
			stack = stack[:len(stack)-1]
		}
		if len(stack) == 0 {
			leftToRight[i] = -1
		} else {
			leftToRight[i] = stack[len(stack)-1]
		}
		stack = append(stack, height[i])
	}

	fmt.Println(leftToRight)

	for i := 0; i < len(height); i++ {
		if leftToRight[i] == -1 || rightToLeft[i] == -1 {
			continue
		}
		level := min(leftToRight[i], rightToLeft[i]) - height[i]
		res += level
	}

	return res
}

// leetcode submit region end(Prohibit modification and deletion)
func TestTrappingRainWater(t *testing.T) {
	// your test code here

}
