func searchRange(nums []int, target int) []int {
    result := make([]int, 2)
    result[0] = bsearch(nums, target, true)
    result[1] = bsearch(nums, target, false)
    return result
}

func bsearch(nums []int, target int, left bool) int {
    ret := -1
    l, r := 0, len(nums)-1
    for l <= r {
        mid := (l+r)>>1
        if nums[mid] == target {
            ret = mid
            if left {
                r = mid-1
            } else {
                l = mid+1
            }
        } else if nums[mid]> target {
            r = mid-1
        } else {
            l = mid+1
        }
    }
    return ret
}