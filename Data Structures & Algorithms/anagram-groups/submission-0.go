func groupAnagrams(strs []string) [][]string {
    if len(strs) == 1 {
        return [][]string{strs}
    } else if len(strs) == 2 {
        if isAnagram(strs[0], strs[1]) {
            return [][]string{
                []string{strs[0], strs[1]},
            }
        } else {
            return [][]string{
                []string{strs[0]}, 
                []string{strs[1]},
            }
        }
    }


    res := [][]string{}
    anagrams := map[string][]string{}
    for i := range strs {
        checked := false
        for k := range anagrams {
            if k == strs[i] {
                checked = true
            }
            for _, v := range anagrams[k] {
                if v == strs[i] {
                    checked = true
                }
            }
        }

        if checked {
            continue
        }

        foundSomething := false
        for j := i + 1; j < len(strs); j++ {
            if isAnagram(strs[i], strs[j]) {
                anagrams[strs[i]] = append(anagrams[strs[i]], strs[j])
                foundSomething = true
            }
        }
        if !foundSomething {
            anagrams[strs[i]] = []string{}
        }
    }

    for k, v := range anagrams {
        toAdd := []string{}
        toAdd = append(toAdd, k)
        for _, val := range v {
            toAdd = append(toAdd, val)
        }
        res = append(res, toAdd)
    }
    return res
}

func isAnagram(s, t string) bool {
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
