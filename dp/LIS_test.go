package dp

import (
	"fmt"
	"testing"
)

func TestLIS(t *testing.T) {
	lis := findNumberOfLIS([]int{1, 2, 4, 3, 5, 4, 7, 2})
	fmt.Println(lis)
}
