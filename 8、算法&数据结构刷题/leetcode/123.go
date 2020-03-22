func maxProfit(prices []int) int {
    n := len(prices)
    if n <= 1 {
        return 0
    }
    dp := make([][3][2]int, n)

    dp[0][0][0] = 0
    dp[0][0][1] = -prices[0]
    dp[0][1][0] = -100000
    dp[0][1][1] = -100000
    dp[0][2][0] = -100000

    for i:=1; i<n; i ++ {
        dp[i][0][1] = maxInt(dp[i-1][0][1], dp[i-1][0][0]-prices[i])
        dp[i][1][0] = maxInt(dp[i-1][1][0], dp[i-1][0][1]+prices[i])
        dp[i][1][1] = maxInt(dp[i-1][1][1], dp[i-1][1][0]-prices[i])
        dp[i][2][0] = maxInt(dp[i-1][2][0], dp[i-1][1][1]+prices[i])
    }
    return maxInt(maxInt(dp[n-1][0][0], dp[n-1][1][0]), dp[n-1][2][0])
}

func maxInt(a int, b int) int {
	if a > b {
		return a
	}
	return b
}