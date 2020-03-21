
func search(nums []int, target int) int {
	nlen := len(nums)
	return bfind(nums, target, 0, nlen)
}

func bfind(nums []int, target int, l int, r int) int {
	if (l>=r) {
		return -1
	}

	mid := (l+r)>>1
	if (nums[mid] == target) {
		return mid
	}

    if nums[mid] > nums[r-1] {
        if target > nums[mid] {
            return bfind(nums, target, mid, r)
        }
        if target <= nums[r-1] {
            return bfind(nums, target, mid, r)
        }
        return bfind(nums, target, l, mid)
    }

    if target > nums[mid] {
        if target <= nums[r-1] {
            return bfind(nums, target, mid, r)
        }
        return bfind(nums, target, l, mid)
    }
    return bfind(nums, target, l, mid)

}