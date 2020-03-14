func majorityElement(nums []int) int {
	var retval, valtimes int

	for _, val := range nums {
		if valtimes <= 0 {
			retval = val
			valtimes ++
			continue
		}

		if retval == val {
			valtimes ++
			continue
		}
		
		valtimes --
	}

	return retval
}