package string

func buildNext2(pattern string) []int {
	n := len(pattern)
	next := make([]int, n)
	next[0] = -1
	for i := 1; i < n; i++ {
		j := next[i-1]
		if pattern[i] == pattern[j+1] {
			next[i] = next[i-1] + 1
		} else {
			k := next[i-1]
			for k >= 0 && pattern[k+1] != pattern[i] {
				k = next[k]
			}
			// -1 或者 ==
			if pattern[k+1] == pattern[i] {
				k++
			}
			next[i] = k
		}
	}
	return next
}

// next[i] = next[i-1]+1; pattern[i-1] == pattern[next[i-1]]
// next[i] =

func contains(source, pattern string) bool {
	next := buildNext2(pattern)
	i, j := 0, 0
	for i < len(source) && j < len(pattern) {
		for j > 0 && source[i] != pattern[j] {
			j = next[j-1] + 1
		}
		if source[i] == pattern[j] {
			j++
		}
		i++
	}
	if j >= len(pattern) {
		return true
	}
	return false
}
