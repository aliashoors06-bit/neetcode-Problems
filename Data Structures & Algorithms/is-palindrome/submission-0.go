func isPalindrome(s string) bool {
    s = strings.ToLower(s)
    i, j := 0, len(s)-1

    for i < j {
        if !isAlnum(s[i]) {
            i++
            continue
        }
        if !isAlnum(s[j]) {
            j--
            continue
        }
        if s[i] != s[j] {
            return false
        }
        i++
        j--
    }
    return true
}

func isAlnum(c byte) bool {
    return (c >= 'a' && c <= 'z') || (c >= '0' && c <= '9')
}