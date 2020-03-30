func minimumLengthEncoding(words []string) int {
    var result int
    nlen := len(words)
    for i:=0; i<nlen; i++ {
        words[i] = reverse(words[i])
    }
    sort.Strings(words)
    for i:=0; i<nlen; i++ {
        next := i+1
        if next == nlen {
            result += len(words[i])+1
            continue
        }
        if len(words[i]) <= len(words[next]) && words[i] == words[next][0:len(words[i])] {
            continue
        } else {
            result += len(words[i])+1
        }
    }
    return result
}

func reverse(word string) string {
    sl := []rune(word)
    for i,j := 0, len(word)-1; i<j; i,j = i+1,j-1 {
        sl[i],sl[j] = sl[j], sl[i]
    }
    return string(sl)
}