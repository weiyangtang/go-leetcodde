package backtrace

import (
	"fmt"
	"testing"
)

var res [][]int
func permute(nums []int) [][]int {
	res = make([][]int,0)
	visited := make([]bool, len(nums))
	backtrace(nums,[]int{},visited)
	fmt.Println(res)
	return res
}

func backtrace(nums []int, selected []int, visited []bool) {
	if len(nums) == len(selected) {
		res = append(res, append([]int{},selected...))
		fmt.Println(res)
		return
	}
	for i := range nums {
		if !visited[i] {
			selected = append(selected, nums[i])
			visited[i] = true
			backtrace(nums, selected, visited)
			selected = selected[:len(selected)-1]
			visited[i] = false
		}
	}
}

func TestPermutations(t *testing.T) {
	result := permute([]int{1, 2, 3})
	fmt.Println(result)
}
