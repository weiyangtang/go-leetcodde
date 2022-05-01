package link

import (
	"fmt"
)

type ListNode struct {
	Val  int
	Next *ListNode
}

func sortList(head *ListNode) *ListNode {
	return mergeSort(head)
}

func mergeSort(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	fmt.Println(head)
	fast, slow := head, head
	for fast != nil && fast.Next != nil && fast.Next.Next != nil {
		fast = fast.Next.Next
		slow = slow.Next
	}
	node1 := head
	node2 := slow.Next
	slow.Next = nil
	head1 := mergeSort(node1)
	head2 := mergeSort(node2)
	return merge(head1, head2)
}

func merge(head1, head2 *ListNode) *ListNode {
	if head1 == nil || head2 == nil {
		if head1 == nil {
			return head2
		}
		return head1
	}
	dummy := &ListNode{}
	node := dummy
	for head1 != nil && head2 != nil {
		if head1.Val < head2.Val {
			node.Next = head1
			node = node.Next
			head1 = head1.Next
		} else {
			node.Next = head2
			node = node.Next
			head2 = head2.Next
		}
	}
	if head1 != nil {
		node.Next = head1
	} else {
		node.Next = head2
	}
	return dummy.Next
}
