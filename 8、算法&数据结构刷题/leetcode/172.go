func trailingZeroes(n int) int {
    var ret int
    for n > 0 {
        ret += n/5
        n/=5
    }
    return ret
}