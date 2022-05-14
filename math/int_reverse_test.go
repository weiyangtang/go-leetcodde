package math

import (
	"fmt"
	"math"
	"testing"
)

func TestReverse(t *testing.T) {
	fmt.Println(math.MaxInt32)
	fmt.Println(reverse(2147483647))
}

func TestCalc(t *testing.T) {
	express := []float64{
		100 / 1396.234 * 2 * 204.24 / 0.98,
		100 / 1396.234 * 2 * 98 / 0.85,
		100 / 1396.234 * 1 * 99.174 / 0.99,
		100 / 1396.234 * 0.01 * 249.69 / 0.99,
		100 / 1396.234 * 0.01 * 194,
	}
	fmt.Printf("100g结果\t50g结果\n")
	var res float64 = 0
	for _, val := range express {
		fmt.Printf("%.3f \t %.3f \n", val, val/2)
		res += val
	}
	fmt.Printf("%.3f", 100-res)

}
