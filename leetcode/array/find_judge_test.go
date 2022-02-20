package array

import (
	"fmt"
	"testing"
)

func TestFindJudge(t *testing.T) {
	judge := findJudge(3, [][]int{{1, 2}})
	fmt.Println(judge)
}
