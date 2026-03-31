type Solution struct{}

func (s *Solution) Encode(strs []string) string {
    sb := strings.Builder{}
    fmt.Fprintf(&sb, "%v%v", len(strs), "#")
    for i := range strs {
        fmt.Fprintf(&sb, "%v%v%v", len(strs[i]), "#", strs[i])
    }
    return sb.String()
}

func (s *Solution) Decode(encoded string) []string {
    walker := strings.Index(encoded, "#")
    l := encoded[:walker]
    strsLen, err := strconv.Atoi(l)
    if err != nil {
        return []string{}
    }
    walker++
    subStr := encoded[walker:]
    res := make([]string, strsLen)
    for i := range res {
        walker = strings.Index(subStr, "#")
        l := subStr[:walker]
        wordLen, err := strconv.Atoi(l)
        if err != nil {
            return []string{}
        }
        walker++
        res[i] = subStr[walker:walker+wordLen]
        subStr = subStr[walker+wordLen:]
    }
    return res
}