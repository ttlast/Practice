func firstMissingPositive(nums []int) int {
	numLen := len(nums)
	for idx, val := range nums {
		nexIdx := idx
		for (val != nexIdx+1) {
			if val <= 0 || val > numLen {
				break
			}
			nexIdx = val-1
			next := nums[nexIdx]
			nums[nexIdx] = val
			val = next
		}
	}
	for idx, val := range nums {
		if idx+1 != val {
			return idx+1
		}
	}
	return numLen+1
}