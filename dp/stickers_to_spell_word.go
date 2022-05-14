package dp

// dp[state] = min{dp[state|targetState[j]]+1} j in (0, len(stickers))
// state: target转成n位的比特位，target[i]能由贴纸拼接得到则第i位为1，否则为0
// dp[state] 表示需要把targe的state位拼接好最少需要的贴纸数
// targetState[j] 表示target和strickers[j]的交集，并且转成target的state位
// dp[0] = 0
var memo []int

const (
	INF = 10000
)

func minStickers(stickers []string, target string) int {
	n := len(target)
	m := (1 << n) - 1

	memo = make([]int, m)
	for i := range memo {
		memo[i] = -1
	}

	res := dfs(0, target, stickers)
	if res == INF {
		return -1
	}
	return res
}

func dfs(mask int, target string, stickers []string) int {
	n := len(target)
	m := (1 << n) - 1
	if mask == m {
		return 0
	}
	if memo[mask] != -1 {
		return memo[mask]
	}
	res := INF
	for _, s := range stickers {
		state := mask
		for i := 0; i < len(s); i++ {
			for j := range target {
				if target[j] == s[i] && (state>>j)&1 == 0 {
					state |= 1 << j
					break
				}
			}
		}

		if state != mask {
			left := dfs(state, target, stickers) + 1
			res = min(left, res)
		}
	}
	memo[mask] = res
	return res
}

func min(x, y int) int {
	if x < y {
		return x
	}
	return y
}
