package prefixSum

type NumMatrix struct {
	Sums [][]int
}

func Constructor(matrix [][]int) NumMatrix {
	m := len(matrix)
	n := len(matrix[0])
	sums := make([][]int, m+1)
	for i := 0; i < m+1; i++ {
		sums[i] = make([]int, n+1)
	}

	for i := 1; i <= m; i++ {
		for j := 1; j <= n; j++ {
			sums[i][j] = sums[i-1][j] + sums[i][j-1] - sums[i-1][j-1] + matrix[i-1][j-1]
		}
	}
	return NumMatrix{Sums: sums}
}

func (m *NumMatrix) SumRegion(row1 int, col1 int, row2 int, col2 int) int {
	return m.Sums[row2+1][col2+1] - m.Sums[row1][col2+1] - m.Sums[row2+1][col1] + m.Sums[row1][col1]
}
