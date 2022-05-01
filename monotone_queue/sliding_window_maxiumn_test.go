package monotone_queue

import (
	"fmt"
	"testing"
)

func TestMaxSlideWindow(t *testing.T) {
	var (
		nums = []int{1, 3, -1, -3, 5, 3, 6, 7}
		k    = 3
	)

	slidingWindow := maxSlidingWindow(nums, k)
	fmt.Println(slidingWindow)
}
