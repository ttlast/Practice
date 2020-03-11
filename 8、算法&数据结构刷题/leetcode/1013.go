func canThreePartsEqualSum(A []int) bool {
    sumA := make([]int, len(A))

    var Pre int
    for i, val := range A {
        sumA[i] = Pre + val
        Pre = sumA[i]
    }

    if (Pre%3 != 0) {
        return false
    }

    result := Pre/3
    
    for idx,val := range sumA {
        if val == result {
            idx ++
            for idx < len(sumA)-1 {
                if sumA[idx] == result*2 {
                    return true
                }
                idx ++
            }
        }
    }
    return false
}