func massage(nums []int) int {
    switch len(nums) {
        case 1:
            return nums[0]
        case 0:
            return 0
    }
    dp := make([]int, len(nums))
    dp[0], dp[1] = nums[0], nums[0]
    if dp[1] < nums[1] {
        dp[1] = nums[1]
    }
    for i:=2; i<len(nums); i++ {
        dp[i] = dp[i-1]
        if dp[i] < dp[i-2] + nums[i] {
            dp[i] = dp[i-2] + nums[i]
        }
    }
    return dp[len(nums)-1]
}