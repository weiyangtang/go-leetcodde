package stack

import (
	"fmt"
	"testing"
)

func TestStack2Queue(t *testing.T) {
	queue := Constructor()
	queue.Push(1)
	queue.Push(2)
	queue.Push(3)
	queue.Push(4)
	pop := queue.Pop()
	fmt.Println(pop)
	queue.Push(5)
	fmt.Println(queue.Pop())
	fmt.Println(queue.Pop())
	fmt.Println(queue.Pop())
	fmt.Println(queue.Pop())
}
