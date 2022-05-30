package sort

import (
	"fmt"
	"testing"
)

func TestRadixSort(t *testing.T) {
	//arr := []int{3, 16, 9, 34, 8021, 342, 2334, 532, 91, 99, 92, 91}
	arr := []int{1, 10000000}
	radixSort(arr)
	fmt.Println(arr)
}
