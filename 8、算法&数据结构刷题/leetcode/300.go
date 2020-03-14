func lengthOfLIS(nums []int) int {
	maxLen := make([]int, len(nums))

	for idx,_ := range maxLen {
		maxLen[idx] = 1
	}

	var retVal int
	for i,_ := range nums {
		for j:=0; j<i; j++ {
			if (nums[i] > nums[j]) {
				maxLen[i] = MaxInt(maxLen[i], maxLen[j] + 1)
			}
		}
		retVal = MaxInt(retVal, maxLen[i])
	}

	return retVal
}

func MaxInt(a int, b int) int {
	if a > b {
		return a
	}
	return b
}