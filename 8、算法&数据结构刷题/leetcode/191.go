func hammingWeight(num uint32) int {
    ret := 0
    for num > 0 {
        ret ++
        num &= (num-1)
    }
    return ret
}