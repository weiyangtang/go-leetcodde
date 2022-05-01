package link

type DListNode struct {
	Val  int
	Key  int
	Next *DListNode
	Prev *DListNode
}

type List struct {
	Head *DListNode
	Tail *DListNode
}

func (this *List) AddFirst(node *DListNode) {
	if this.Head == nil {
		this.Head = node
		this.Tail = node
		return
	}
	node.Next = this.Head
	this.Head.Prev = node
	this.Head = node
}

func (this *List) Delete(node *DListNode) {
	if node.Next == nil {
		this.Tail = this.Tail.Prev
		return
	}
	node.Next.Prev = node.Prev
	node.Prev.Next = node.Next
}
