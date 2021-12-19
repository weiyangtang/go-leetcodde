package doublePointer

import "sort"

func threeSum(nums []int) [][]int {
	sort.Ints(nums)
	println(nums)
	n := len(nums)
	res := make([][]int, 0)
	for i := 0; i < n-2; i++ {
		if i > 0 && nums[i-1] == nums[i] {
			continue
		}
		j, k := i+1, n-1
		target := -nums[i]
		for j < k {
			if j > i+1 && nums[j-1] == nums[j] {
				continue
			}
			if nums[j]+nums[k] == target {
				res = append(res, []int{nums[i], nums[j], nums[k]})
				j++
			} else if nums[j]+nums[k] < target {
				j++
			} else {
				k--
			}
		}
	}
	return res
}
