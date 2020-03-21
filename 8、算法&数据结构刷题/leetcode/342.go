func isPowerOfFour(num int) bool {
	if num <=0 || num&(num-1) > 0 {
		return false
	}
	if num&0xAAAAAAAA > 0 {
		return false
	}
	return true
}