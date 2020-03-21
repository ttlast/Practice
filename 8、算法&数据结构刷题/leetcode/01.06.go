func compressString(S string) string {
	var retStrSlice []string
    
    var preChar rune
    var num int
    for _, char := range S {
        if preChar == char {
            num ++
            continue
        }
        if num > 0 {
            retStrSlice = append(retStrSlice, string(preChar))
            retStrSlice = append(retStrSlice, strconv.Itoa(num))
        }
        preChar = char
        num = 1
    }
    if num > 0 {
        retStrSlice = append(retStrSlice, string(preChar))
        retStrSlice = append(retStrSlice, strconv.Itoa(num))
    }

    newStr := strings.Join(retStrSlice, "")
    if len(newStr) >= len(S) {
        return S
    }
	return newStr
}