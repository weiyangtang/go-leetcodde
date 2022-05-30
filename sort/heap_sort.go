package sort

func findKth(arr []int, k int) int {
	n := len(arr)
	// 构造k个节点的堆
	for i := k/2 - 1; i >= 0; i-- {
		adjustMinHeap(arr, i, k)
	}

	for i := k; i < n; i++ {
		if arr[i] > arr[0] {
			arr[i], arr[0] = arr[0], arr[i]
			adjustMinHeap(arr, 0, k)
		}
	}
	return arr[0]
}

func heapSort(arr []int) {
	// 构造大顶堆
	n := len(arr)
	// 最后一个非叶子节点下标
	for k := n/2 - 1; k >= 0; k-- {
		adjustMaxHeap(arr, k, n)
	}

	for i := 0; i < n; i++ {
		// 每次将堆顶arr[0]和未排序部分堆尾arr[n-1-i]
		arr[n-1-i], arr[0] = arr[0], arr[n-1-i]
		adjustMaxHeap(arr, 0, n-1-i)
	}
}

// 调整以arr[left]为堆顶的，构造成大顶堆，每个节点的值大于或等于左右孩子节点的值
func adjustMaxHeap(arr []int, left, right int) {
	num := arr[left]
	i := left
	for 2*i+1 < right {
		childIndex := 2*i + 1
		if childIndex+1 < right && arr[childIndex] < arr[childIndex+1] {
			childIndex++
		}
		if num < arr[childIndex] {
			arr[i] = arr[childIndex]
			i = childIndex
		} else {
			break
		}
	}
	arr[i] = num
}

func adjustMinHeap(arr []int, left, right int) {
	num := arr[left]
	i := left
	for 2*i+1 < right {
		childIndex := 2*i + 1
		if childIndex+1 < right && arr[childIndex] > arr[childIndex+1] {
			childIndex++
		}
		if num > arr[childIndex] {
			arr[i] = arr[childIndex]
			i = childIndex
		} else {
			break
		}
	}
	arr[i] = num
}
