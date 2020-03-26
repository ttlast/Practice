func permuteUnique(nums []int) [][]int {
    var resultSlice [][]int
    dfs(&nums, 0, len(nums), &resultSlice)
    return resultSlice
}

func dfs(nums *[]int, from int, len int, resultSlice *[][]int) {
    if (from == len) {
        newNum := make([]int, len)
        copy(newNum, (*nums)[0:])
        *resultSlice = append(*resultSlice, newNum)
        return 
	}
	mm := make(map[int]int)
    for i:=from; i<len; i++ {
		if _, ok := mm[(*nums)[i]]; ok {
			continue
		}
        (*nums)[from], (*nums)[i] = (*nums)[i], (*nums)[from]
        dfs(nums, from+1, len, resultSlice)
		(*nums)[from], (*nums)[i] = (*nums)[i], (*nums)[from]
		mm[(*nums)[i]] = 1
    }
}