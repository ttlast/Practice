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

    var retStr string
    for i := 0; i<=top; i++ {
        retStr =  retStr + string(sstack[i])
    }

	return retStr
}