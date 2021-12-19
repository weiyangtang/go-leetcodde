package tree

func BuildBTree(nodeValues []int) *TreeNode {
	if nodeValues == nil || len(nodeValues) == 0 {
		return nil
	}
	n := len(nodeValues)
	nodeQueue := make([]*TreeNode, n)
	root := &TreeNode{Val: nodeValues[0]}
	nodeQueue[0] = root
	left, right := 0, 1
	for i := 1; i < n; {
		for nodeQueue[left] == nil {
			left++
		}
		parent := nodeQueue[left]
		if nodeValues[i] == -1 {
			parent.Left = nil
		} else {
			parent.Left = &TreeNode{Val: nodeValues[i]}
		}
		nodeQueue[right] = parent.Left
		i++
		right++
		if i < n {
			if nodeValues[i] == -1 {
				parent.Right = nil
			} else {
				parent.Right = &TreeNode{Val: nodeValues[i]}
			}
			nodeQueue[right] = parent.Right
			i++
			right++
		}
		left++
	}
	return root
}
