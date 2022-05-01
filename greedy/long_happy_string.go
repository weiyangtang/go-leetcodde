package greedy

import "sort"

func longestDiverseString(a int, b int, c int) string {

	chCnt := []struct {
		cnt int
		ch  byte
	}{{a, 'a'}, {b, 'b'}, {c, 'c'}}
	retBytes := make([]byte, 0)
	for {
		// 构造优先队列
		sort.Slice(chCnt, func(i, j int) bool {
			return chCnt[i].cnt > chCnt[j].cnt
		})
		hasNext := false
		for i := range chCnt {
			if chCnt[i].cnt > 0 && (len(retBytes) < 2 || !(retBytes[len(retBytes)-2] == chCnt[i].ch && retBytes[len(retBytes)-1] == chCnt[i].ch)) {
				retBytes = append(retBytes, chCnt[i].ch)
				chCnt[i].cnt--
				hasNext = true
				break
			}
		}
		if !hasNext {
			break
		}
	}
	res := ""
	for i := range retBytes {
		res += string(rune(retBytes[i]))
	}
	return res
}
