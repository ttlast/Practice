func countCharacters(words []string, chars string) int {
    var ret int
    hashchars := make([]int, 26)
    for _, char := range chars {
        hashchars[char-'a'] ++
    }

    for _, word := range words {
        sign := true
        tmphash := make([]int, 26)
        for _, char := range word {
            tmphash[char-'a'] ++
        }

        for i:=0; i<26; i++ {
            if tmphash[i] > hashchars[i] {
                sign = false
                break
            }
        }
        if sign {
            ret += len(word)
        }
    }

    return ret
}