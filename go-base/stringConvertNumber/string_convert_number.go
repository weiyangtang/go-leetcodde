package stringConvertNumber

import (
	"fmt"
	"strconv"
)

func stringConvertNumber(strNumber string) int {
	parseInt, err := strconv.ParseInt(strNumber, 10, 0)
	if err != nil {
		fmt.Errorf("parse int error: %s", err)
	}

	return int(parseInt)
}
func isHappy(n int) bool {
	hashMap := make(map[int]bool, 0)
	for _, ok := hashMap[n]; ok; {

	}
}
