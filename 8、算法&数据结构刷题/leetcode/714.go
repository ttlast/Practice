func maxProfit(prices []int, fee int) int {
    nlen := len(prices)
    dp := make([][2]int, nlen)

    dp[0][1] = -prices[0]
    for i:=1; i<nlen; i++ {
        dp[i][1] = maxInt(dp[i-1][1], dp[i-1][0]-prices[i])
        dp[i][0] =maxInt(dp[i-1][0], dp[i-1][1]+prices[i]-fee)
    }

    return dp[nlen-1][0]
}

func maxInt(a int, b int) int {
	if a > b {
		return a
	}
	return b
}