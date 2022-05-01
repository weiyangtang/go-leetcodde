package string

import (
	"fmt"
	"testing"
)

func TestLongestCommonPrefix(t *testing.T) {
	commonPrefix := longestCommonPrefix([]string{"abc", "abd", "ab"})
	fmt.Println(commonPrefix)
}
