func sortString(s string) string {
	hash := make([]int, 26)

	lens := len(s)
	for _,cc := range s {
		hash[int(cc-'a')] ++
	}

	var chaSlice []rune
	for true {
		for i:=0; i<26; i++ {
			if hash[i] > 0 {
				chaSlice = append(chaSlice, rune(i+'a'))
				hash[i] --
			}
		}
		for i:=25; i>=0; i-- {
			if hash[i] > 0 {
				chaSlice = append(chaSlice, rune(i+'a'))
				hash[i] --
			}
		}
		if len(chaSlice) >= lens {
			break
		}
	}

	return string(chaSlice)
}