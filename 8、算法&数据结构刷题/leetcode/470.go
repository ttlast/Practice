func rand10() int {
    var val int
	for true {
		val = rand7()+(rand7()-1)*7
		if val <= 40 {
			break
		}
	}
	return val%10+1
}