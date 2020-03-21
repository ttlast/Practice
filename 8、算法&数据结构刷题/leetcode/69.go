func mySqrt(x int) int {
    if x <= 0 {
        return 0
    }
	if x <= 3 {
		return 1
	}

	int64x := int64(x)
	var l, r int
	l = 2
	r = x/2
	for l<r {
		mid := int64((l+r+1)>>1)
        tmp := mid*mid
        if tmp == int64x {
            return int(mid)
        }
		if tmp < int64x {
			l = int(mid)
            continue
		}
		r =int(mid-1)
	}
	return l
}