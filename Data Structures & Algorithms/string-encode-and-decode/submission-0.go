type Solution struct {
}

func (s *Solution) Encode(strs []string) string {
    res := ""
    for _, str := range strs {
        res += strconv.Itoa(len(str)) + "#" + str
    }
    return res
}

func (s *Solution) Decode(encoded string) []string {
    res := []string{}
    i := 0

    for i < len(encoded) {
        j := i
        for encoded[j] != '#' {
            j++
        }

        length, _ := strconv.Atoi(encoded[i:j])
        res = append(res, encoded[j+1:j+1+length])
        i = j + 1 + length
    }
    return res
}