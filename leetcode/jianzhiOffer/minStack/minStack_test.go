package minStack

import (
	"fmt"
	"testing"
)

func TestMinStack(t *testing.T) {
	minStack := Constructor()
	minStack.Push(0)
	minStack.Push(1)
	minStack.Push(0)
	fmt.Println(minStack.Min())
	minStack.Pop()
	fmt.Println(minStack.Top())
	fmt.Println(minStack.Min())
}
