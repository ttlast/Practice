func longestPalindrome(s string) int {
	numSlice := make([]int, 52)
	for _, char := range s {
		var pos int
		if char >= 'a' && char <= 'z' {
			pos = 26 + int(char - 'a')
		} else {
			pos = int(char - 'A')
		}
		numSlice[pos] ++
	}

	var ret int
	for _, val := range numSlice {
		if val&1 > 0 {
			ret += val^1
		} else {
			ret += val
		}
	}
    if ret < len(s) {
        return ret+1
    }
	return ret
}