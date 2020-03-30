func removeOuterParentheses(S string) string {
    var top = 0
    var result []rune
    vstack := make([]int, len(S))
    for i, char := range(S) {
        if char == '(' {
            vstack[top] = i
            top ++
            if top > 1 {
                result = append(result, rune(char))
            }
        } else {
            top --
            if top >= 1 {
                result = append(result, rune(char))
            }
        }
    }

    return string(result)
}