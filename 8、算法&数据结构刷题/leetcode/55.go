func canJump(nums []int) bool {
    
    hash := make([]int, len(nums))

    hash[0] = 1;
    nlen := len(nums)
    for i, _ := range nums {
        for j:=1; j<=nums[i]; j++ {
            if (i+j >= nlen) {
                break;
            }
            hash[i+j] |= hash[i];
        }
    }
    if(hash[nlen-1] >= 1) {
        return true
    }
    return false
}