package binarySearch

import "fmt"

func search(nums []int, target int) int {
	return rotatedSearch(nums, target, 0, len(nums)-1)
}

func rotatedSearch(a []int, target, left, right int) int {
	fmt.Println(left,right)
	if left > right {
		return -1
	}
	mid := right + (left-right)/2
	if a[mid] == target {
		return mid
	}
	if a[mid] < target && target <= a[right] {
		return binarySearch(a,target,mid+1, right)
	}
	if a[left] <= target && target < a[mid] {
		return binarySearch(a,target,left,mid-1)
	}
	if a[left] < a[mid] {
		return rotatedSearch(a,target, mid+1,right)
	}
	return rotatedSearch(a,target,left,mid-1)
}

func binarySearch(a []int, target, left, right int) int {
	fmt.Println("BinarySearch", left, right)
	for left <= right {
		mid := right + (left-right)/2
		if a[mid] == target {
			return mid
		} else if a[mid] > target {
			right = mid - 1
		} else {
			left = mid + 1
		}
	}
	return -1
}
