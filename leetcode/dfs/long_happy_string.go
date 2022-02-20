package dfs

var happyStrs = make([]string, 0)

const MinValue int = 1e+8

func longestDiverseString(a int, b int, c int) string {
	dfs("", a+b+c, []int{a, b, c})
	maxHappyString := ""
	for i := range happyStrs {
		if len(maxHappyString) < len(happyStrs[i]) {
			maxHappyString = happyStrs[i]
		}
	}

	return maxHappyString

}

func dfs(s string, cnt int, rest []int) bool {
	if len(s) > cnt {
		return false
	}
	nextLoop := 3
	for i := 0; i < 3; i++ {
		s = s + string(rune('a'+i))
		if (len(s) >= 3 && (s[len(s)-3:] == "aaa" || s[len(s)-3:] == "bbb" || s[len(s)-3:] == "ccc")) || rest[i] == 0 {
			nextLoop--
			s = s[:len(s)-1]
			continue
		}
		rest[i]--
		if !dfs(s, cnt, rest) {
			happyStrs = append(happyStrs, s)
			nextLoop--
		}
		s = s[:len(s)-1]
		rest[i]++
	}
	return nextLoop > 0
}
