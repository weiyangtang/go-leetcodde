package tree

import (
	"fmt"
	"testing"
)

func TestBuildBTree(t *testing.T) {
	arr := []int{1, 2, 3, 4, 5, -1, 6, 7}
	root := BuildBTree(arr)
	fmt.Println(root.Val)
}
