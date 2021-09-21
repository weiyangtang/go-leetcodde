package link

import (
	"fmt"
	"testing"
)

func TestCopyLink(t *testing.T) {
	//arr := [][]int{{1, 2}, {3, 5}, {7, 8}}
	//fmt.Println(arr[2][1])
	//fmt.Println(arr)
	//arr1 := [3][3]int{{1, 2, 3}, {3, 5}, {7, 8}}
	//fmt.Println(arr1[0][2])
	nodes := [...][2]int{{7, -1}, {13, 7}, {11, 1}, {10, 11}, {1, 7}}
	dummy := &Node{}
	node := dummy
	nodeMap := make(map[int]*Node)
	for i := 0; i < len(nodes); i++ {
		node.Next = &Node{Val: nodes[i][0]}
		node = node.Next
		nodeMap[nodes[i][0]] = node
	}
	node = dummy.Next
	for i := 0; i < len(nodes); i++ {
		if nodes[i][1] == -1 {
			node = node.Next
			continue
		}
		node.Random = nodeMap[nodes[i][1]]
		node = node.Next
	}
	node = dummy.Next
	//for node != nil {
	//	fmt.Printf("%#v\n", node)
	//	node = node.Next
	//}
	node = copyRandomList(dummy.Next)
	for node != nil {
		fmt.Printf("%#v\n", node)
		node = node.Next
	}
}
