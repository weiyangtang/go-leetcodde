package binarySearch

// 普通的二分查找，适用于无重复的有序数组
func ordinaryBinSearch(arr []int, target int) int {
	left, right := 0, len(arr)-1
	for left <= right {
		mid := left + (right-left)/2
		if arr[mid] == target {
			return mid
		} else if arr[mid] > target {
			right = mid - 1
		} else {
			left = mid + 1
		}
	}
	return -1
}

func leftBoundBinSearch(arr []int, target int) int {
	left, right := 0, len(arr)-1
	for left <= right {
		mid := left + (right-left)/2
		if arr[mid] == target {
			right = mid - 1
		} else if arr[mid] > target {
			right = mid - 1
		} else {
			left = mid + 1
		}
	}

	if left == len(arr) || arr[left] != target {
		return -1
	}
	return left
}

func rightBoundBinSearch(arr []int, target int) int {
	left, right := 0, len(arr)-1
	for left <= right {
		mid := left + (right-left)/2
		if arr[mid] == target {
			left = mid + 1
		} else if arr[mid] > target {
			right = mid - 1
		} else {
			left = mid + 1
		}
	}

	if right < 0 || arr[right] != target {
		return -1
	}
	return right
}
