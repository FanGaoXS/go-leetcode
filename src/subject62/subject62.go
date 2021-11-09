package subject62

func uniquePaths(m int, n int) int {
	ints := [][]int{}
	for i := 0; i < m; i++ {
		dp := make([]int, n)
		dp[0] = 1
		ints = append(ints,dp)
	}
	for i := 1; i < n; i++ {
		ints[0][i] = 1
	}
	for i := 1; i < m; i++ {
		for j := 1; j < n; j++ {
			ints[i][j] = ints[i-1][j] + ints[i][j-1]
		}
	}
	return ints[m-1][n-1]
}
