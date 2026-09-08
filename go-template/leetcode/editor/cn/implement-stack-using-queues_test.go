package leetcode_solutions

import "testing"

// leetcode submit region begin(Prohibit modification and deletion)
type MyStack struct {
	queue      []int
	topElement int
}

func Constructor() MyStack {
	return MyStack{make([]int, 0), 0}
}

func (this *MyStack) Push(x int) {
	this.queue = append(this.queue, x)
	this.topElement = x
}

func (this *MyStack) Pop() int {
	size := len(this.queue)
	for i := 2; i < size; i++ {
		temp := this.queue[0]
		copy(this.queue[0:size-1], this.queue[1:])
		this.queue[size-1] = temp
	}
	this.topElement = this.queue[0]
	copy(this.queue[0:size-1], this.queue[1:])
	this.queue[size-1] = this.topElement
	res := this.queue[0]
	this.queue = this.queue[1:]
	return res
}

func (this *MyStack) Top() int {
	return this.topElement
}

func (this *MyStack) Empty() bool {
	return len(this.queue) == 0
}

/**
 * Your MyStack object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Push(x);
 * param_2 := obj.Pop();
 * param_3 := obj.Top();
 * param_4 := obj.Empty();
 */
//leetcode submit region end(Prohibit modification and deletion)

func TestImplementStackUsingQueues(t *testing.T) {
	// your test code here
	//stack := Constructor()
}
