func combine(n int, k int) [][]int {
	tmpSlice := make([]int, k)
	var resultSlice [][]int
	dfs(&tmpSlice, 0, 1, n, k, &resultSlice)
	return resultSlice
}

func dfs(slice *[]int, level int, start int, len int, k int, resultSlice *[][]int) {
	if level == k {
		tmpSlice := make([]int, k)
		copy(tmpSlice, (*slice)[0:])
		*resultSlice = append(*resultSlice, tmpSlice)
		return 
	}

	for i:=start; i<=len; i++ {
		(*slice)[level] = i
		dfs(slice, level+1, i+1, len, k, resultSlice)
	}
}