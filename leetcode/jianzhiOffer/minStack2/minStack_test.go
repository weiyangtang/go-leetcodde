package minStack2

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

func TestSlice(t *testing.T) {
	x := make([]int, 0)
	x = append(x, 1, 2, 3, 4, 5, 6)
	x = x[:len(x)-1]
	fmt.Println(x)
	fmt.Println(len(x))
	fmt.Println(cap(x))
	x = append(x, 100)
	fmt.Println(x)

}
