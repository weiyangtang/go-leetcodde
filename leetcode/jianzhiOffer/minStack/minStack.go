package minStack

import "container/list"

/*
minQueue 最小栈
*/

type MinStack struct {
	stack, minQueue list.List
}

/** initialize your data structure here. */
func Constructor() MinStack {
	return MinStack{}
}

func (this *MinStack) Push(x int) {
	this.stack.PushBack(x)
	if this.minQueue.Len() > 0 && this.minQueue.Front().Value.(int) < x {
		return
	}
	this.minQueue.PushFront(x)
}

func (this *MinStack) Pop() {
	min := this.stack.Remove(this.stack.Back())
	if this.minQueue.Len() > 0 && this.minQueue.Front().Value.(int) == min {
		this.minQueue.Remove(this.minQueue.Front())
	}
}

func (this *MinStack) Top() int {
	if this.stack.Len() == 0 {
		return -1
	}
	return this.stack.Back().Value.(int)
}

func (this *MinStack) Min() int {
	return this.minQueue.Front().Value.(int)
}
