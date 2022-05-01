package link

import (
	"fmt"
	"testing"
)

func TestReverseLink(t *testing.T) {
	arr := []int{1, 2, 3, 4, 5}
	dummy := &ListNode{}
	fmt.Printf("dummy=%#v\n", dummy)
	fmt.Printf("&dummy=%#v\n", &dummy)
	//fmt.Printf("*dummy=%#v\n", *dummy)
	node := dummy
	for _, val := range arr {
		node.Next = &ListNode{Val: val}
		node = node.Next
	}
	node = dummy.Next
	for node != nil {
		fmt.Printf("node=%#v\n", node)
		node = node.Next
	}
	head := reverseList(dummy.Next)
	node = head.Next
	for node != nil {
		fmt.Printf("node=%#v\n", node)
		node = node.Next
	}
}
