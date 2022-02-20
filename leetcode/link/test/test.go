package test

type LRUCache struct {
	m     map[int]*ListNode
	list  MyList
	cap   int
	limit int
}

func Constructor(capacity int) LRUCache {
	m := make(map[int]*ListNode)
	list := MyList{}
	return LRUCache{m: m, list: list, cap: 0, limit: capacity}
}

func (this *LRUCache) Get(key int) int {
	if val, ok := this.m[key]; ok {
		return val.Val
	}
	return -1
}

func (this *LRUCache) Put(key int, value int) {
	if val, ok := this.m[key]; ok {
		this.list.Delete(val)
		this.cap--
	}
	node := &ListNode{Val: value, Key: key}
	this.list.AddFirst(node)
	this.m[key] = node
	this.cap++

	if this.cap > this.limit {
		delete(this.m, this.list.Tail.Key)
		this.list.Delete(this.list.Tail)
		this.cap--
	}
}

type ListNode struct {
	Val  int
	Key  int
	Next *ListNode
	Prev *ListNode
}

type MyList struct {
	Head *ListNode
	Tail *ListNode
}

func (this *MyList) AddFirst(node *ListNode) {
	if this.Head == nil {
		this.Head = node
		this.Tail = node
		return
	}
	node.Next = this.Head
	this.Head.Prev = node
	this.Head = node
}

func (this *MyList) Delete(node *ListNode) {
	if node.Next == nil {
		this.Tail = this.Tail.Prev
		return
	}
	node.Next.Prev = node.Prev
	node.Prev.Next = node.Next
}
