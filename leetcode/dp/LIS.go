package dp

func findNumberOfLIS(nums []int) int {
	n := len(nums)
	dp := make([]int, n)
	cnt := make([]int, n)
	dp[0] = 1
	cnt[0] = 1
	maxLen := 1
	for i := 1; i < n; i++ {
		dp[i] = 1
		cnt[i] = 1
		for j := i - 1; j >= 0; j-- {
			if nums[j] < nums[i] {
				if dp[j]+1 > dp[i] {
					dp[i] = dp[j] + 1
					cnt[i] = cnt[j]
				} else if dp[j]+1 == dp[i] {
					cnt[i] += cnt[j]
				}
			}
		}
		maxLen = max(maxLen, dp[i])
	}
	res := 0
	for i := 0; i < n; i++ {
		if dp[i] == maxLen {
			res += cnt[i]
		}
	}
	return res
}

func max(x, y int) int {
	if x >= y {
		return x
	}
	return y
}
