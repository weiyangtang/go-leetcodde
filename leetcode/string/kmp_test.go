package string

import (
	"fmt"
	"testing"
)

func TestBuildNext(t *testing.T) {
	next := buildNext("abcabdddabcabc")
	for _, val := range next {
		fmt.Printf("%d\t", val)
	}
}

func TestIsSubString(t *testing.T) {
	// "mississippi"
	//"issip
	isSubString := isSubString("mississippi", "issip")
	fmt.Println(isSubString)
}
