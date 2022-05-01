package array

/*
螺旋矩阵，这道题特别考验coder对代码边界值的掌控能力
 m 行，n列
*/

func SpiralMatrix(m, n int) [][]int {
	matrix := make([][]int, m)
	for i := 0; i < m; i++ {
		matrix[i] = make([]int, n)
	}
	l, r, t, b := 0, n-1, 0, m-1
	// l, r, t,b := 0,7,0,0
	i := 1
	for i <= m*n {
		// 上
		for j := l; j <= r; j++ {
			matrix[t][j] = i
			i++
		}
		t++
		// 右
		for j := t; j <= b; j++ {
			matrix[j][r] = i
			i++
		}
		r--
		// 下
		for j := r; j >= l && t <= b; j-- {
			matrix[b][j] = i
			i++
		}
		b--
		// 左
		for j := b; j >= t && l <= r; j-- {
			matrix[j][l] = i
			i++
		}
		l++
	}
	return matrix
}
