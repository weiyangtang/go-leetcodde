package array

import (
	"fmt"
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
