package prefixSum

import (
	"fmt"
	"testing"
)

func TestNumMatrix_SumRegion(t *testing.T) {
	numMatrix := Constructor([][]int{{3, 0, 1, 4, 2}, {5, 6, 3, 2, 1}, {1, 2, 0, 1, 5}, {4, 1, 0, 1, 7}, {1, 0, 3, 0, 5}})
	sumRegion := numMatrix.SumRegion(1, 1, 2, 2)
	fmt.Println(sumRegion)
}
