package stack2queue

import (
	"fmt"
	"testing"
)

func TestStack2queue(t *testing.T) {
	obj := Constructor()
	obj.AppendTail(1)
	obj.AppendTail(2)
	obj.AppendTail(3)
	top := obj.DeleteHead()
	fmt.Println(top)
	fmt.Println(obj.DeleteHead())
	obj.AppendTail(4)
	fmt.Println(obj.DeleteHead())
	fmt.Println(obj.DeleteHead())
	fmt.Println(obj.DeleteHead())
	fmt.Println(obj.DeleteHead())

}