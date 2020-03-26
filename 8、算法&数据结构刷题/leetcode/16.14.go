func bestLine(points [][]int) []int {
    var maxSum int
    result := make([]int, 2)
    nlen := len(points)
    for i:=0; i<nlen; i++ {
        var maxk, s, e int
        mm := make(map[float64][]int)
        for j:=i+1; j<nlen; j++ {
            if points[i][0] == points[j][0] {
                maxk ++
                if maxk == 1 {
                    s, e = i,j
                }
            } else {
                kk := getSlope(points[i], points[j])
                if _, ok:=mm[kk]; !ok {
                    mm[kk] = []int{i, j, 1}
                } else {
                    mm[kk][2] ++
                }
            }
        }
        if maxk > maxSum {
            maxSum = maxk
            result = []int{s, e}
        }
        for _, val := range mm {
            if maxSum < val[2] {
                maxSum = val[2]
                result = []int{val[0], val[1]}
            } else if maxSum == val[2] && result[0] >= val[0] && result[1] >= val[1]{
                maxSum = val[2]
                result = []int{val[0], val[1]}
            }
        }
    }
    return result
}

func getSlope(a []int, b []int) float64 {
    var k float64
    k = float64((b[1]-a[1]))/float64(b[0]-a[0])
    return k
}