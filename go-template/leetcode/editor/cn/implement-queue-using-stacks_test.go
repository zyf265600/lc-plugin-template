package leetcode_solutions

//
//import (
//	"fmt"
//	"testing"
//)
//
//// leetcode submit region begin(Prohibit modification and deletion)
//type MyQueue struct {
//	stack1 []int
//	stack2 []int
//}
//
//func Constructor() MyQueue {
//	return MyQueue{stack1: make([]int, 0), stack2: make([]int, 0)}
//}
//
//func (this *MyQueue) Push(x int) {
//	this.stack1 = append(this.stack1, x)
//}
//
//func (this *MyQueue) Pop() int { // 中心思想，s2 == 空 再进行stack1中值的转移
//	this.Peek()
//	x := this.stack2[len(this.stack2)-1]
//	this.stack2 = this.stack2[:len(this.stack2)-1]
//	return x
//}
//
//func (this *MyQueue) Peek() int {
//	if len(this.stack2) == 0 {
//		for len(this.stack1) > 0 {
//			size := len(this.stack1)
//			this.stack2 = append(this.stack2, this.stack1[size-1])
//			this.stack1 = this.stack1[:size-1]
//		}
//	}
//	return this.stack2[len(this.stack2)-1]
//}
//
//func (this *MyQueue) Empty() bool {
//	if len(this.stack2) == 0 && len(this.stack1) == 0 {
//		return true
//	}
//	return false
//}
//
///**
// * Your MyQueue object will be instantiated and called as such:
// * obj := Constructor();
// * obj.Push(x);
// * param_2 := obj.Pop();
// * param_3 := obj.Peek();
// * param_4 := obj.Empty();
// */
////leetcode submit region end(Prohibit modification and deletion)
//
//func TestImplementQueueUsingStacks(t *testing.T) {
//	// your test code here
//	queue := Constructor()
//	fmt.Println(queue)
//	queue.Push(1)
//	queue.Push(2)
//	queue.Push(3)
//	fmt.Println(queue)
//	queue.Peek()
//	fmt.Println(queue)
//	queue.Push(4)
//	queue.Push(5)
//	queue.Push(6)
//	fmt.Println(queue)
//	fmt.Println(queue.Pop())
//	fmt.Println(queue)
//}
