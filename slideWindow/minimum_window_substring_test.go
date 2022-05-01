package slideWindow

import (
	"fmt"
	"testing"
)

func TestMinWindow(test *testing.T) {
	//"ADOBECODEBANC"
	//"ABC"
	s := "ADOBECODEBANC"
	t := "ABC"
	res := minWindow(s, t)
	fmt.Println(res)
}
