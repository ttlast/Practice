func minIncrementForUnique(A []int) int {
    var ret int
    if len(A) <= 0 {
        return 0
    }
    sort.Ints(A)
    pre := A[0]
    for i:=1; i<len(A); i++ {
        if A[i] <= pre {
            ret += pre+1-A[i]
            A[i] = pre+1
        }
        pre = A[i]
    }
    return ret
}