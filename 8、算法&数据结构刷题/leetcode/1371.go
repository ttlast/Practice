func findTheLongestSubstring(s string) int {
	dp := make([]int, 32)
	for i:=0; i<len(dp); i++ {
		dp[i] = -1
	}

	var maxLen int
	curStep := 0
	for i, cc := range s {
		switch cc {
		case 'a':
			curStep ^= 1
		case 'e':
			curStep ^= 2
		case 'i':
			curStep ^= 4
		case 'o':
			curStep ^= 8
		case 'u':
			curStep ^= 16
		}
		if dp[curStep] < 0 {
			dp[curStep] = i
			continue
		}

		maxLen = maxInt(maxLen, i-dp[curStep]+1)
	}
	
	return maxLen
}

func MaxInt(a int, b int) int {
	if a > b {
		return a
	}
	return b
}