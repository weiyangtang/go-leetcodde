package stack

import (
	"fmt"
	"testing"
)

func TestValid(t *testing.T) {
	valid := isValid("()")
	fmt.Println(valid)
}
