package link

import "testing"

func TestReverseK(t *testing.T) {
	dummy := &DListNode{}
	arr := []int{1, 2, 3}
	node := dummy
	for i := 0; i < len(arr); i++ {
		node.Next = &DListNode{Val: arr[i]}
		node = node.Next
	}
	reverseKGroup(dummy.Next, 2)
}
