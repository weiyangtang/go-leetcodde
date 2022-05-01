package link

import "testing"

func TestLinkSort(t *testing.T) {
	dummy := &ListNode{}
	arr := []int{4,2,1,3}
	node := dummy
	for i := 0; i < len(arr); i++ {
		node.Next = &ListNode{Val: arr[i]}
		node = node.Next
	}
	sortList(dummy.Next)
}
