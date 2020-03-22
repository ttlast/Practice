func maxProfit(prices []int) int {
	var ans int
	if len(prices) <= 1 {
		return 0
	}
	for idx := 1; idx < len(prices); idx ++ {
		if (prices[idx] > prices[idx-1]) {
			ans += prices[idx] - prices[idx-1]
		}
	}
	return ans
}