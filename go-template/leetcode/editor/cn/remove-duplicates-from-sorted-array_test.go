package leetcode_solutions

//
//import "testing"
//
//// leetcode submit region begin(Prohibit modification and deletion)
//func removeDuplicates(nums []int) int {
//	if len(nums) == 1 {
//		return 1
//	}
//
//	fast := 0
//	slow := 0
//
//	for fast < len(nums) {
//		if nums[fast] != nums[slow] {
//			slow++
//			nums[slow] = nums[fast]
//		}
//		fast++
//	}
//
//	return slow + 1
//}
//
////leetcode submit region end(Prohibit modification and deletion)
//
//func TestRemoveDuplicatesFromSortedArray(t *testing.T) {
//	// your test code here
//
//}
