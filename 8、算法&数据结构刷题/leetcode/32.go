func longestValidParentheses(s string) int {
    var sstack = make([]int, len(s))

    var result = 0;
    top := -1
	for idx, char := range s {
        if char == ')' && top >= 0 && s[sstack[top]] == '(' {
            top -= 1
            if (top >= 0) {
                result = MaxInt(result, idx - sstack[top])
            } else {
                result = MaxInt(result, idx + 1)
            }
            continue
        }

        top += 1
        sstack[top] = idx
    }
    
    return result
}

func MaxInt(a int, b int) int {
	if a > b {
		return a
	}
	return b
}