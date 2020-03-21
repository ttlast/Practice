func isPowerOfThree(n int) bool {
	tn := 1
	for true {
		if tn == n {
			return true
		}
		tn *= 3
		if tn > n || tn < 0 {
			break
		}
	}
	return false
}