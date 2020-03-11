func removeDuplicates(S string) string {
	var sstack = make([]rune, len(S))
    top := -1
	for _, char := range S {
        
        if top >= 0 && sstack[top] == char {
            top -= 1
            continue
        }

        top += 1
        sstack[top] = char
	}

	return string(sstack[0:top+1])
}