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
	var (
		M         = 3169.963
		m float64 = 100
	)

	express := []float64{
		m / M * 60 * 1 / 0.4,
		m / M * 2 * 0.015 * 78 / 0.99,
		m / M * 0.34 * 40 * 2 / 0.96,
		m / M * 0.11 * 56 * 2 / 0.85,
		m / M * 0.3 * 1 * 362.19 / 0.98,
		m / M * 0.6 * 1 * 147.26 / 0.35,
	}
	fmt.Printf("100g结果\t200g结果\n")
	var res float64 = 0
	for _, val := range express {
		fmt.Printf("%.3f \t %.3f \n", val, val*2)
		res += val
	}
	ans := 100 - res
	fmt.Printf("%.3f \t %.3f \n", ans, ans*2)

}
