package leetcode_solutions

import (
	"slices"
	"testing"
)

// leetcode submit region begin(Prohibit modification and deletion)
func sortedSquares(nums []int) []int {
	first := 0
	last := len(nums) - 1

	res := make([]int, 0, len(nums))

	for first <= last {
		squareFirst := nums[first] * nums[first]
		squareLast := nums[last] * nums[last]
		var newNum int

		if squareFirst > squareLast {
			newNum = squareFirst
			first++
		} else {
			newNum = squareLast
			last--
		}
		res = append(res, newNum)
	}

	slices.Reverse(res)

	return res
	//negative := make([]int, 0, len(nums))
	//mid := 0
	//
	//for i := range nums {
	//	square := nums[i] * nums[i]
	//	if nums[i] <= 0 {
	//		negative = append(negative, 0)
	//		copy(negative[1:], negative[0:])
	//		negative[0] = square
	//		nums[i] = 0
	//		mid++
	//	} else {
	//		nums[i] = square
	//	}
	//}
	//
	//head := 0
	//pn := 0
	//
	//for mid < len(nums) && pn < len(negative) {
	//	if nums[mid] < negative[pn] {
	//		nums[head] = nums[mid]
	//		mid++
	//	} else {
	//		nums[head] = negative[pn]
	//		pn++
	//	}
	//	head++
	//}
	//
	//for pn < len(negative) {
	//	nums[head] = negative[pn]
	//	pn++
	//	head++
	//}
	//for mid < len(nums) {
	//	nums[head] = nums[mid]
	//	mid++
	//	head++
	//}
	//
	//return nums
}

//leetcode submit region end(Prohibit modification and deletion)

func TestSquaresOfASortedArray(t *testing.T) {
	// your test code here

}
