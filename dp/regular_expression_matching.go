package dp

import "fmt"

// dp[i][j] = dp[i-1][j-1] && (s1[i] == s2[j] || s2[j] == '.')
// dp[i][j] = (dp[i-1][j] || dp[i-1][j-2]) && (s1[i] == s2[j-1] || s2[j-1] == '.') && s2[j] == '*'
// dp[i][j] = false // i == 0 || j == 0
// dp[1][1] = s1[1] == s2[1] || s2[1] == '.'
// dp[i][1] = false
func isMatch(s string, p string) bool {
	m,n := len(s),len(p)
	dp := make([][]bool, m+1)
	for i := 0; i < m+1; i++ {
		dp[i] = make([]bool, n+1)
	}
	dp[0][0] = true
	dp[1][1] = s[0] == p[0] || p[0] == '.'

	for j := 2; j <= n; j++ {
		dp[0][j] = dp[0][j-2] && (p[j-1] == '*')
	}
	for i := 1; i <= m; i++ {
		for j := 2; j <= n; j++ {
			dp[i][j] = (dp[i-1][j-1] && (s[i-1] == p[j-1] || p[j-1] == '.'))||((dp[i-1][j] || dp[i-1][j-2]) && (s[i-1] == p[j-2] || p[j-2] == '.') && p[j-1] == '*') || (dp[i][j-2] || p[j-1] == '*')
		}
	}
	fmt.Println(dp)
	return dp[m][n]
}