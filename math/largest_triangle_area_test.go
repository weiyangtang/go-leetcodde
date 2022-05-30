package math

import (
	"fmt"
	"testing"
)

func TestLargestTriangleArea(t *testing.T) {
	res := largestTriangleArea([][]int{{0, 0}, {0, 1}, {1, 0}, {0, 2}, {2, 0}})
	fmt.Println(res)
}

func TestArea(t *testing.T) {
	res := area([][]int{{4, 6}, {6, 5}, {3, 1}}...)
	fmt.Println(res)
}
