package _map

import (
	"fmt"
	"testing"
)

func TestLNRS(t *testing.T) {
	inputString := "pwwkew"
	res, lnrs := lengthOfLongestSubstring(inputString)
	fmt.Println(res)
	fmt.Println(lnrs)
}
