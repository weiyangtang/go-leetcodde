package sort

// 基数排序
// @param arr 待排序数组
func radixSort(arr []int) {

	var (
		bucketSize  = 10
		buckets     = make([]int, bucketSize)
		n           = len(arr)
		tempSortArr = make([]int, n)
	)
	maxVal := max(arr...)
	for i := 1; i <= maxVal; i *= 10 {
		// 计数排序使用桶置空
		buckets = make([]int, bucketSize)
		// 待排序数字填0对齐后，遍历数组数字取第i位数字进行计数排序
		for j := 0; j < n; j++ {
			unSortNumByte := arr[j] / i % 10
			buckets[unSortNumByte]++
		}

		// buckets[j] 表示 第i位为j的数字最大排序位置为buckets[j]
		for j := 1; j < bucketSize; j++ {
			buckets[j] += buckets[j-1]
		}

		// 计数排序
		for j := n - 1; j >= 0; j-- {
			unSortNumByte := arr[j] / i % 10
			tempSortArr[buckets[unSortNumByte]-1] = arr[j]
			buckets[unSortNumByte]--
		}

		// 将上一轮排序后的结论赋值arr
		copy(arr, tempSortArr)
	}
}

func max(arr ...int) int {
	res := arr[0]
	for i := range arr {
		if arr[i] > res {
			res = arr[i]
		}
	}

	return res
}
