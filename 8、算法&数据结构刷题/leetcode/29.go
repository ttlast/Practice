func divide(dividend int, divisor int) int {
    dd := divisor
    if dd < 0 {
        dd = -dd
    }

    int64Dividend := int64(dividend)
    if int64Dividend < 0 {
        int64Dividend = -int64Dividend
    }
    var ret int64
    for true {
        vv := int64(1)
        tt := int64(dd)
        if int64Dividend < tt {
            break
        }
        for true {
            if int64Dividend < tt {
                break
            }
            tt <<= 1
            vv <<= 1
        }
        int64Dividend -= tt>>1
        ret += vv>>1
    }
    if divisor < 0 {
        ret = -ret
    }
    if dividend < 0 {
        ret = -ret
    }

    if ret >= 2147483648 {
        return 2147483647
    }
    return int(ret)
}