package tree

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestAvl(t *testing.T) {
	nodeValues := []int{10, 5, 15, -1, -1, 6, 20}
	root := BuildBTree(nodeValues)
	res := isAvl(root)
	assert.Equal(t, false, res)

	nodeValues1 := []int{10, 5, 15, 4}
	root1 := BuildBTree(nodeValues1)
	res1 := isAvl(root1)
	assert.Equal(t, true, res1)

}

func TestAvl2(t *testing.T) {

	//nodeValues := []int{10, 5, 15, 4}
	//root := BuildBTree(nodeValues)
	//res := isAvl(root)
	//assert.Equal(t, true, res)

	nodeValues1 := []int{10, 5, 15, 4}
	root1 := BuildBTree(nodeValues1)
	res1 := isAvl(root1)
	assert.Equal(t, true, res1)
}
