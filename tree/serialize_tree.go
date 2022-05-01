package tree

import (
	"fmt"
	"strconv"
	"strings"
)

type Codec struct {
}

func Constructor() Codec {
	return Codec{}
}

// Serializes a tree to a single string.
func (this *Codec) serialize(root *TreeNode) string {
	if root == nil {
		return "#"
	}
	queue := make([]*TreeNode, 0)
	node := root
	serialStr := ""
	for len(queue) > 0 || node != nil {
		serialStr = fmt.Sprintf("%s,%d", serialStr, node.Val)
		queue = append(queue, node.Left)
		queue = append(queue, node.Right)

		node = queue[0]
		queue = queue[1:]
		for len(queue) > 0 && node == nil {
			serialStr = fmt.Sprintf("%s,#", serialStr)
			node = queue[0]
			queue = queue[1:]
		}
	}
	return strings.TrimLeft(serialStr,",")
}

// Deserializes your encoded data to tree.
func (this *Codec) deserialize(data string) *TreeNode {
	nodeVals := strings.Split(data, ",")
	if len(nodeVals) == 0 {
		return nil
	}
	queue := make([]*TreeNode, 0)
	val, _ := strconv.Atoi(nodeVals[0])
	node := &TreeNode{Val: val}
	dummy := node
	for i := 1; i < len(nodeVals); {
		if nodeVals[i] == "#" {
			node.Left = nil
		} else {
			val, _ = strconv.Atoi(nodeVals[i])
			node.Left = &TreeNode{Val: val}
		}
		queue = append(queue, node.Left)
		i++
		if i < len(nodeVals) {
			if nodeVals[i] == "#" {
				node.Right = nil
			} else {
				val, _ = strconv.Atoi(nodeVals[i])
				node.Right = &TreeNode{Val: val}
			}
			queue = append(queue, node.Right)
			i++
			node = queue[0]
			queue = queue[1:]
		}

	}
	return dummy
}
