package link

func reverseKGroup(head *ListNode, k int) *ListNode {
	dummy := &ListNode{}
	tail := dummy
	tail.Next = head
	for {
		i := 0
		node := head
		for ; node != nil && i < k-1; i++ {
			node = node.Next
		}
		if node == nil {
			return dummy.Next
		}
		nextHead := node.Next
		node.Next = nil

		lastNode := head
		head = reverse(head)
		tail.Next = head
		tail = lastNode
		head = nextHead
	}
	return dummy.Next
}

func reverse(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	node := head.Next
	first := reverse(head.Next)
	head.Next = nil
	node.Next = head
	return first
}
