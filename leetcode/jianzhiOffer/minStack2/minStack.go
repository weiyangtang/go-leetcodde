package minStack2

/*
minQueue 最小栈
*/

type MinStack struct {
	stack    []int
	minStack []int
}

/** initialize your data structure here. */
func Constructor() MinStack {
	return MinStack{}
}

func (this *MinStack) Push(x int) {
	this.stack = append(this.stack, x)
	if len(this.minStack) == 0 || this.minStack[len(this.minStack)-1] >= x {
		this.minStack = append(this.minStack, x)
	}
}

func (this *MinStack) Pop() {
	top := this.stack[len(this.stack)-1]
	this.stack = this.stack[:len(this.stack)-1]
	if this.minStack[len(this.minStack)-1] == top {
		this.minStack = this.minStack[:len(this.minStack)]
	}
}

func (this *MinStack) Top() int {
	return this.stack[len(this.stack)-1]
}

func (this *MinStack) Min() int {
	return this.minStack[len(this.minStack)-1]
}
