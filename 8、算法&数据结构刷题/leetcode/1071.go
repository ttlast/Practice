func gcdOfStrings(str1 string, str2 string) string {
    str1Len := len(str1)
    str2Len := len(str2)
    
    var m,n int
    if str1Len > str2Len {
        m = str1Len
        n = str2Len
    } else {
        n = str1Len
        m = str2Len
    }

    for m%n != 0 {
        p := m%n
        m = n
        n = p
    }

    if str1 + str2 == str2 + str1 {
        return str1[0:n]
    }

    return ""
}