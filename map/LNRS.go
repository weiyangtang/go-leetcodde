package _map

// 最长无重复子串
func lengthOfLongestSubstring(s string) (int, string) {
	if s == "" {
		return 0, ""
	}
	hashMap := make(map[rune]int)
	res := 0
	lnrs := ""
	left := 0
	for i := 0; i < len(s); i++ {
		repeatIndex, ok := hashMap[rune(s[i])]
		if ok && repeatIndex >= left {
			left = repeatIndex + 1
		}
		if i-left+1 > res {
			lnrs = s[left : i+1]
			res = i - left + 1
		}
		//res = maxInt(res, i-left+1)

		hashMap[rune(s[i])] = i
	}
	return res, lnrs
}

func maxInt(x, y int) int {
	if x > y {
		return x
	}
	return y
}
