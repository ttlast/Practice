func myPow(x float64, n int) float64 {
    var pown int
    pown = n
    if n < 0 {
        pown = -n
    }

    var result, tmp float64
    result = 1
    tmp = x
    var tt int
    tt = 1
    for i:=0; i<=31; i++ {
        if (pown&tt) == tt {
            result *= tmp
        }
        tmp *= tmp
        tt <<= 1
    }

    if n < 0 {
        return 1.0/result
    }

    return result
    
}