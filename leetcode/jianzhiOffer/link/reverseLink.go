package link

type ListNode struct {
	Val  int
	Next *ListNode
}

func reverseList(head *ListNode) *ListNode {
	dummyNode := ListNode{}
	node := head
	for node != nil {
		next := node.Next
		node.Next = dummyNode.Next
		dummyNode.Next = node
		node = next
	}
	head = dummyNode.Next
	return head
}
