package array

import (
	"fmt"
	"math"
	"testing"
)

func TestSpiralMatrix(t *testing.T) {
	m, n := 1, 7
	matrix := SpiralMatrix(m, n)
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			fmt.Printf("%d\t", matrix[i][j])
		}
		fmt.Println()
	}
}

func TestCalc(t *testing.T) {
	N, n := 3, 14
	M := (N - 1) * int(math.Ceil(float64(n*1.0/2/(n-1))))
	fmt.Println(M)
	fmt.Println(float64(n*1.0/2/(n-1)))
}
