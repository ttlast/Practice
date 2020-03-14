func getPermutation(n int, k int) string {

    sign := make([]int, n+1)
    sum := make([]int, n+1)
    sum[0] = 1
    for i:=1; i<=n; i++ {
        sum[i] = sum[i-1]*i
    }

    var retString string
    for i:=n-1; i>=0; i-- {
        for j:=1; j<=n; j++ {
            if sign[j] == 1 {
                continue
            }
            if k <= sum[i] {
                retString = retString + strconv.Itoa(j)
                sign[j] = 1
                break
            }
            k -= sum[i]
        }
    }
    
    return retString
}