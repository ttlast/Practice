func trap(height []int) int {
    nlen := len(height)
    
    leftMaxHeight := make([]int, nlen)
    rightMaxHeight := make([]int, nlen)

    for i:=1; i<nlen; i++ {
        leftMaxHeight[i] = leftMaxHeight[i-1]
        if leftMaxHeight[i] < height[i-1] {
            leftMaxHeight[i] = height[i-1]
        }
    }

    for i:=nlen-2; i>=0; i-- {
        rightMaxHeight[i] = rightMaxHeight[i+1]
        if rightMaxHeight[i] < height[i+1] {
            rightMaxHeight[i] = height[i+1]
        }
    }

    var ret int
    for i:=0; i<nlen; i++ {
        maxHeight := leftMaxHeight[i]
        if maxHeight > rightMaxHeight[i] {
            maxHeight = rightMaxHeight[i]
        }
        if maxHeight > height[i] {
            ret += maxHeight - height[i]
        }
    }

    return ret
}