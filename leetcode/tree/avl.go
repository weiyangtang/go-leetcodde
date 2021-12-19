package tree

import (
	"math"
)

var (
	nodeVal = math.MinInt
)

func isAvl(root *TreeNode) bool {
	nodeVal = math.MinInt
	//return judgeAvl(root) >= 0
	return isBST(root) && isBalance(root)
}

func judgeAvl(root *TreeNode) int {
	if root == nil {
		return 0
	}
	left := judgeAvl(root.Left)
	if left == -1 || root.Val <= nodeVal {
		return -1
	}
	nodeVal = root.Val
	right := judgeAvl(root.Right)
	if right == -1 || math.Abs(float64(right-left)) > 1 {
		return -1
	}
	return int(math.Max(float64(left), float64(right)) + 1)
}

func isBST(root *TreeNode) bool {
	if root == nil {
		return true
	}
	if !isBST(root.Left) || root.Val <= nodeVal {
		return false
	}
	nodeVal = root.Val
	return isBST(root.Right)
}

func isBalance(root *TreeNode) bool {
	return height(root) >= 0
}

func height(root *TreeNode) int {
	if root == nil {
		return 0
	}
	lh := height(root.Left)
	if lh == -1 {
		return -1
	}
	rh := height(root.Right)
	if rh == -1 || abs(lh-rh) > 1 {
		return -1
	}
	return max(lh, rh) + 1
}

func abs(x int) int {
	if x > 0 {
		return x
	}
	return -x
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}
