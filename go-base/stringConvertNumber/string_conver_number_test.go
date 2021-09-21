package stringConvertNumber

import (
	"fmt"
	"strconv"
	"testing"
)

func TestParseInt(t *testing.T) {
	number := stringConvertNumber("38909405")
	fmt.Println(number)
	itoa := strconv.Itoa(9303)
	fmt.Println(itoa)
}
