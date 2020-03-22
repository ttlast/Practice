func maxProfit(k int, prices []int) int {
    nlen := len(prices)
    if  nlen <= 1 || k <= 0 {
        return 0
    }
    if k > (nlen>>1) {
        k = nlen>>1
    }
    var dp [2][2][]int
    for i:=0; i<2; i++ {
        dp[i][0] = make([]int, k+1)
        dp[i][1] = make([]int, k+1)
    }
	
    
    for j:=1; j<=k; j++ {
        dp[0][0][j] = -1000000
        dp[0][1][j] = -1000000
	}
	dp[0][1][1] = -prices[0]
	pre := 0
	next := 1

    for i:=1; i<nlen; i++ {
        dp[next][1][0] = maxInt(dp[pre][1][0], -prices[i])
        for j:=1; j<=k; j++ {
			// 买入
			dp[next][1][j] = maxInt(dp[pre][1][j], dp[pre][0][j-1]-prices[i])
			// 卖出
			dp[next][0][j] = maxInt(dp[pre][0][j], dp[pre][1][j]+prices[i])
		}
		next ^= 1
		pre ^= 1
	}
    ret := 0
    for j:=0; j<=k; j++ {
        ret = maxInt(ret, dp[pre][0][j])
    }
    
    return ret
}

func maxInt(a int, b int) int {
	if a > b {
		return a
	}
	return b
}