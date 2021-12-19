package string

// 模式串是否source串的子串
func isSubString(source, pattern string) bool {
	next := buildNext(pattern)
	// sourceIndex, patternIndex
	i, j := 0, 0
	for i < len(source) && j < len(pattern) {
		if source[i] == pattern[j] {
			i++
			j++
		} else if j > 0 {
			j = next[j-1]
		} else {
			i++
		}
	}
	if j == len(pattern) {
		return true
	}
	return false
}

func buildNext(pattern string) []int {
	n := len(pattern)
	next := make([]int, n)
	l, r := 0, 1
	for r < n {
		if pattern[l] == pattern[r] { // 前后缀相等，next =l + 1
			next[r] = l + 1
			l++
			r++
		} else if l > 0 {
			// 前后缀不完全相等，但p[0:l-1]和p[r-l:r-1]是相等的，如果直接从p[0]和p[r-1]开始比较就太可惜了，// abcabxxxxabcad   abca和abca 分别找两个相等子串的最大相等前后缀
			// 其实就是求其中一个子串的最大相等前后缀不就是 next[l-1]嘛
			l = next[l-1]
		} else {
			// 如果p[0:l-1]和p[r-l:r-1] 两个子串长度为0，也就不存在任何相等的前后缀了
			r++
		}
	}
	return next
}
