func maxWidthRamp(A []int) int {
    istack := make([]int, len(A))

    var result int
    var top := -1
    
    for idx, val := range A {

        leftMind := bfind(A, istack, 0, top, val)
        if (leftMind >= 0) {
            result = MaxInt(result, idx-istack[leftMind])
        }

        if top >= 0 && A[istack[top]] < val {
            continue
        }

        top ++
        istack[top] = idx
    }

    return result
}

// [6,0]
func bfind(A []int, list []int, left int, right int, val int) int {
    if (right < left) {
        return -1;
    }
    if (right == left) {
        if (val >= A[list[left]]) {
            return left
        }
        return -1;
    }

    mid := (left+right)>>1;

    if (val < A[list[mid]]) {
        return bfind(A, list, mid+1, right, val)
    }

    return bfind(A, list, left, mid, val)
}

func MaxInt(a int, b int) int {
	if a > b {
		return a
	}
	return b
}

//[6,0,8,2,1,5]
//[6,0]

//[50000, 49999, 49998 ..... 1]