package monotone_queue

func maxSlidingWindow(nums []int, k int) []int {
	queue := make([]int, 0)
	maxArr := make([]int, 0)
	for i := 0; i < k; i++ {
		num := nums[i]
		for len(queue) > 0 && queue[len(queue)-1] <= num {
			queue = queue[:len(queue)-1]
		}
		queue = append(queue, num)
	}
	maxArr = append(maxArr, queue[0])
	for i := k; i < len(nums); i++ {
		num := nums[i]
		for len(queue) > 0 && queue[len(queue)-1] <= num {
			queue = queue[:len(queue)-1]
		}
		queue = append(queue, num)
		if queue[0] == nums[i-k] {
			queue = queue[1:]
		}
		maxArr = append(maxArr, queue[0])
	}
	return maxArr
}
