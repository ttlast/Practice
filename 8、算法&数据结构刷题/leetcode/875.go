func minEatingSpeed(piles []int, H int) int {
    var ans int
    l,r := 1, 1000000000
    for true {
        if r < l {
            break
        }
        mid := (l+r)>>1
        if bfind(piles, mid) <= int64(H) {
            ans = mid
            r = mid-1
        } else {
            l = mid+1
        }
    }
    return ans;
}

func bfind(piles []int, k int) int64 {

    var ret int64
    for _,v := range piles {
        ret += int64((v+k-1)/k)
    }

    return ret
}