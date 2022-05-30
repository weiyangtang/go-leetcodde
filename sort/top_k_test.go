package sort

import (
	"fmt"
	"testing"
)

func TestTopKDivide(t *testing.T) {
	nums := []int{3, 2, 1, 5, 6, 4, 0, 1}
	mid := divide(nums, 0, len(nums)-1)
	fmt.Println(nums[:mid])
	fmt.Println(nums[mid:])
	fmt.Println(mid)
}

func TestTopK(t *testing.T) {
	nums := []int{3, 2, 3, 1, 2, 4, 5, 5, 6}
	kth := findKthLargest(nums, 4)
	fmt.Println(kth)
}
