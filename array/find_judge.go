package array

func findJudge(n int, trust [][]int) int {
	trustSlice := make([][]int, n+1)
	for i := range trust {
		trustSlice[i] = make([]int, 2)
		a, b := trust[i][0], trust[i][1]
		// 被信任+1
		trustSlice[b][0]++
		// 信任别人+1
		trustSlice[a][1]++
	}
	cnt := 0
	idx := 0
	for i := range trustSlice {
		if trustSlice[i][0] == n-1 && trustSlice[i][1] == 0 {
			cnt++
			idx = i
		}
	}
	if cnt == 1 {
		return idx
	}
	return -1
}
