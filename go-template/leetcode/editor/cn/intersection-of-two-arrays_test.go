package leetcode_solutions

import "testing"

// leetcode submit region begin(Prohibit modification and deletion)
func intersection(nums1 []int, nums2 []int) []int {
	res := make([]int, 0)
	count := make(map[int]bool)

	for _, v := range nums1 {
		count[v] = true
	}
	for _, v := range nums2 {
		if count[v] == true {
			res = append(res, v)
			count[v] = false
		}
	}
	return res
}

//leetcode submit region end(Prohibit modification and deletion)

func TestIntersectionOfTwoArrays(t *testing.T) {
	// your test code here

}
