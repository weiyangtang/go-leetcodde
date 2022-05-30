package sort

func findKthLargest(nums []int, k int) int {
	n := len(nums)
	left, right := 0, n-1
	for {
		mid := divide(nums, left, right)
		if mid == n-k {
			return nums[mid]
		}
		if mid > n-k {
			right = mid - 1
		} else {
			left = mid + 1
		}
	}
}

func divide(nums []int, left, right int) int {
	midVal := nums[left]
	for left < right {
		for left < right && nums[right] >= midVal {
			right--
		}
		nums[left] = nums[right]

		for left < right && nums[left] < midVal {
			left++
		}
		nums[right] = nums[left]
	}
	nums[left] = midVal

	return left
}
