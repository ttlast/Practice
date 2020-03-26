func maxSubArray(nums []int) int {
    if len(nums) <= 0 {
        return 0
    }

    ret := nums[0]
    tmpSum, minSum := nums[0], nums[0]

    for i:=1; i<len(nums); i++ {
        val := nums[i]
        tmpSum += val

        if ret < tmpSum {
            ret = tmpSum
        }
        if ret < tmpSum-minSum {
            ret = tmpSum-minSum
        }

        if minSum > tmpSum {
            minSum = tmpSum
        }
    }

    return ret
}