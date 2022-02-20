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

func TestBuildNext2(t *testing.T) {
	next := buildNext2("abcabdddabcabc")
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

func TestKmpContains(t *testing.T) {
	contains := contains("mississippi", "issip")
	fmt.Println(contains)
}

func TestSlice(t *testing.T) {
	value := make([]int, 1)
	fmt.Println(cap(value))
	fmt.Println(len(value))
}
