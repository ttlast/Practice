func getLeastNumbers(arr []int, k int) []int {
	var retSlice []int
	hashLen := make([]int, 10001)

	for _, val := range arr {
		hashLen[val] ++
	}
	for idx, val := range hashLen {
		for val > 0 {
			retSlice = append(retSlice, idx)
			val --
			k --
			if k <= 0 {
				break
			}
		}
		if k <= 0 {
			break
		}
	}
	
	return retSlice
}