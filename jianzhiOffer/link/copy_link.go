package link

type Node struct {
	Val    int
	Next   *Node
	Random *Node
}

func copyRandomList(head *Node) *Node {
	dummy := &Node{}
	node := dummy
	oldNode := head
	nodeMap := make(map[*Node]*Node)
	for oldNode != nil {
		node.Next = &Node{
			Val: oldNode.Val,
		}
		nodeMap[oldNode] = node
		node = node.Next
		oldNode = oldNode.Next
	}
	oldNode = head
	node = dummy.Next
	for oldNode != nil {
		node.Random = nodeMap[oldNode.Random]
		oldNode = oldNode.Next
		node = node.Next
	}
	return dummy.Next
}
