package leetcode_solutions

import "testing"

// leetcode submit region begin(Prohibit modification and deletion)
type NumArray struct {
	preSum []int
}

func Constructor(nums []int) NumArray {
	newArray := NumArray{}
	preSum := make([]int, len(nums)+1)
	for i := 1; i <= len(nums); i++ {
		preSum[i] = preSum[i-1] + nums[i-1]
	}

	newArray.preSum = preSum
	return newArray
}

func (this *NumArray) SumRange(left int, right int) int {
	return this.preSum[right+1] - this.preSum[left]
}

/**
 * Your NumArray object will be instantiated and called as such:
 * obj := Constructor(nums);
 * param_1 := obj.SumRange(left,right);
 */
//leetcode submit region end(Prohibit modification and deletion)

func TestRangeSumQueryImmutable(t *testing.T) {
	// your test code here

}
