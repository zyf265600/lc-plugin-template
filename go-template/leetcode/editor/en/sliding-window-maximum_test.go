package leetcode_solutions

import (
	"container/list"
	"testing"
)

// leetcode submit region begin(Prohibit modification and deletion)
type MonotonicQueue struct {
	maxq list.List
}

func (mq *MonotonicQueue) push(n int) {
	// 将小于 n 的元素全部删除
	for mq.maxq.Len() > 0 && mq.maxq.Back().Value.(int) < n {
		mq.maxq.Remove(mq.maxq.Back())
	}
	// 然后将 n 加入尾部
	mq.maxq.PushBack(n)
}

func (mq *MonotonicQueue) max() int {
	return mq.maxq.Front().Value.(int)
}

func (mq *MonotonicQueue) pop(n int) {
	if n == mq.maxq.Front().Value.(int) {
		mq.maxq.Remove(mq.maxq.Front())
	}
	// 否则无需操作，已经被删除了
}

func maxSlidingWindow(nums []int, k int) []int {
	window := MonotonicQueue{list.List{}}
	res := make([]int, 0)

	for i := 0; i < len(nums); i++ {
		if i < k-1 {
			window.push(nums[i])
		} else {
			window.push(nums[i])
			res = append(res, window.max())
			window.pop(nums[i-k+1])
		}
	}
	return res
}

//leetcode submit region end(Prohibit modification and deletion)

func TestSlidingWindowMaximum(t *testing.T) {
	// your test code here
	//nums := []int{1, 3, -1, -3, 5, 3, 6, 7}
	//k := 3
	//fmt.Println(maxSlidingWindow(nums, k))
}
