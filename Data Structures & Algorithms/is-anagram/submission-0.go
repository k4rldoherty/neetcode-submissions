func isAnagram(s string, t string) bool {
    if len(s) != len(t) {
        return false
    }
    sMap := map[byte]int{}
    tMap := map[byte]int{}
    for i := range s {
        if _, ok := sMap[s[i]]; ok {
            sMap[s[i]]++
        } else {
            sMap[s[i]] = 1
        }

        if _, ok := tMap[t[i]]; ok {
            tMap[t[i]]++
        } else {
            tMap[t[i]] = 1
        }
    }

    for k, v := range sMap {
        if tMap[k] != v {
            return false
        }
    }

    return true
}