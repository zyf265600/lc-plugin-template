package leetcode_solutions

import (
	"strconv"
	"strings"
	"testing"
)

// leetcode submit region begin(Prohibit modification and deletion)
func evalRPN(tokens []string) int {
	calculation := "+-*/"
	stack := make([]int, 0)
	for i := range tokens {
		c := tokens[i]
		if strings.Contains(calculation, c) {
			right := stack[len(stack)-1]
			left := stack[len(stack)-2]
			stack = stack[:len(stack)-2]
			switch c {
			case "+":
				stack = append(stack, left+right)
			case "-":
				stack = append(stack, left-right)
			case "*":
				stack = append(stack, left*right)
			case "/":
				stack = append(stack, left/right)
			}
		} else {
			num, _ := strconv.Atoi(c)
			stack = append(stack, num)
		}
	}
	return stack[0]
}

//leetcode submit region end(Prohibit modification and deletion)

func TestEvaluateReversePolishNotation(t *testing.T) {
	// your test code here

}
