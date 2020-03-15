func luckyNumbers (matrix [][]int) []int {
    var retList []int
    for i := range matrix {
        sign := 1
        mind := 0
        for j := range matrix[0] {
            if matrix[i][j] < matrix[i][mind] {
                mind = j
            }
        }
        
        for k := range matrix {
            if matrix[k][mind] > matrix[i][mind] {
                sign = 0
            }
        }
        
        if sign >= 1 {
            retList = append(retList, matrix[i][mind])
        }
    }
    return retList
}