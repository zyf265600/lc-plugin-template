package leetcode_solutions

//
//import (
//	"fmt"
//	"testing"
//)
//
//// leetcode submit region begin(Prohibit modification and deletion)
//type MyLinkedList struct {
//	head *ListNode
//	size int
//}
//
//func Constructor() MyLinkedList {
//	return MyLinkedList{&ListNode{}, 0} // dummy/哨兵 节点
//}
//
//func (this *MyLinkedList) Get(index int) int {
//	if index >= this.size || index < 0 {
//		return -1
//	}
//	p := this.head
//	for i := 0; i < index+1; i++ {
//		p = p.Next
//	}
//	return p.Val
//}
//
//func (this *MyLinkedList) AddAtHead(val int) {
//	this.AddAtIndex(0, val)
//	//var newHead *ListNode
//	//if this.size == 0 {
//	//	newHead = &ListNode{val, nil}
//	//} else {
//	//	newHead = &ListNode{val, this.head}
//	//}
//	//this.head = newHead
//	//this.size++
//}
//
//func (this *MyLinkedList) AddAtTail(val int) {
//	this.AddAtIndex(this.size, val)
//	//var newTail *ListNode = &ListNode{val, nil}
//	//if this.size == 0 {
//	//	this.head = newTail
//	//} else {
//	//	p := this.head
//	//	for p.Next != nil {
//	//		p = p.Next
//	//	}
//	//	p.Next = newTail
//	//}
//	//this.size++
//}
//
//func (this *MyLinkedList) AddAtIndex(index int, val int) {
//	newNode := &ListNode{val, nil}
//
//	if index > this.size || index < 0 { // invalid index
//		return
//	}
//
//	p := this.head
//	for i := 0; i < index; i++ {
//		p = p.Next
//	}
//	newNode.Next = p.Next
//	p.Next = newNode
//
//	this.size++
//}
//
//func (this *MyLinkedList) DeleteAtIndex(index int) {
//	if index >= this.size || index < 0 { // invalid index
//		return
//	}
//
//	p := this.head
//	for i := 0; i < index; i++ {
//		p = p.Next
//	}
//	p.Next = p.Next.Next
//
//	this.size--
//}
//
///**
// * Your MyLinkedList object will be instantiated and called as such:
// * obj := Constructor();
// * param_1 := obj.Get(index);
// * obj.AddAtHead(val);
// * obj.AddAtTail(val);
// * obj.AddAtIndex(index,val);
// * obj.DeleteAtIndex(index);
// */
////leetcode submit region end(Prohibit modification and deletion)
//
//func TestDesignLinkedList(t *testing.T) {
//	// your test code here
//	newll := Constructor()
//	newll.AddAtHead(1)     // 1
//	newll.AddAtTail(3)     // 1 3
//	newll.AddAtIndex(1, 2) // 1 2 3
//	newll.Get(1)           // 2
//	newll.DeleteAtIndex(1) // 1 3
//	newll.Get(1)           // 3
//	p := newll.head
//	for p != nil && newll.size != 0 {
//		fmt.Println(p.Val)
//		p = p.Next
//	}
//}
