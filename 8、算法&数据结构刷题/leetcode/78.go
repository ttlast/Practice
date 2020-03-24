func subsets(nums []int) [][]int {
    var tmpSlict []int
    var resultSlice [][]int
    dfs(&tmpSlict, nums, 0, 0, len(nums), &resultSlice)
    return resultSlice
}

func dfs(slice *[]int, nums []int, level int, start int, tlen int, resultSlice *[][]int) {
    tmpSlict := make([]int, len(*slice))
    copy(tmpSlict, (*slice)[0:])
    *resultSlice = append(*resultSlice, tmpSlict)

	for i:=start; i<tlen; i++ {
		(*slice) = append((*slice), nums[i])
		dfs(slice, nums, level+1, i+1, tlen, resultSlice)
        (*slice) = (*slice)[0:level]
	}
}