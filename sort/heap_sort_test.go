package sort

import (
	"fmt"
	"testing"
)

func TestAdjustHeap(t *testing.T) {
	arr := []int{3, 4, 1, 5, 6, 4}
	adjustMaxHeap(arr, 0, len(arr)-1)
	fmt.Println(arr)
}

func TestHeapSort(t *testing.T) {
	arr := []int{3, 4, 1, 5, 6, 4}
	heapSort(arr)
	fmt.Println(arr)
}

func TestFindKth(t *testing.T) {
	arr := []int{3, 2, 1, 5, 6, 4}
	k := 2
	kth := findKth(arr, k)
	fmt.Println(kth)
}
