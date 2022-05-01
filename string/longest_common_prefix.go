package string

// 最长公共前缀
func longestCommonPrefix(strs []string) string {
	res := ""
	var current rune
	for i := 0; i < 200; i++ {
		for j := range strs {
			if len(strs[j]) <= i {
				return res
			}
			if j == 0 {
				current = rune(strs[j][i])
			} else if current != rune(strs[j][i]) {
				return res
			}
		}
		res += string(current)
	}
	return res
}
