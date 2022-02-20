package slideWindow

// 最小覆盖子串
func minWindow(s string, t string) string {
	sMap := make(map[byte]int)
	tMap := make(map[byte]int)
	for i := range t {
		tMap[t[i]]++
	}
	left, right := 0, 0
	size := len(s)
	validCnt := 0
	needCnt := len(t)
	resSet := make(map[string]bool)

	for right < size {
		sMap[s[right]]++
		if sMap[s[right]] <= tMap[s[right]] {
			validCnt++
		}
		if validCnt == needCnt {
			resSet[s[left:right+1]] = true
			validCnt--
			sMap[s[left]]--
			left++
			//continue
		}

		for left < right && sMap[s[left]] > tMap[s[left]] {
			sMap[s[left]]--
			left++
		}
		right++
	}

	var min int = 1e6
	res := ""
	for optionStr := range resSet {
		if len(optionStr) < min {
			min = len(optionStr)
			res = optionStr
		}
	}

	return res
}
