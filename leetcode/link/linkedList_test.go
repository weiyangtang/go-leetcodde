package link

import "testing"

func TestInitList(t *testing.T) {
	list := List{}
	tmp := &DListNode{Val: 3}
	list.AddFirst(&DListNode{Val: 1})
	list.AddFirst(&DListNode{Val: 2})
	list.AddFirst(tmp)
	list.AddFirst(&DListNode{Val: 4})
	list.Delete(tmp)
}
