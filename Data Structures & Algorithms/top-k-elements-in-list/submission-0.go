import (
    "maps"
    "slices"
)

func topKFrequent(nums []int, k int) []int {
    freq := map[int]int{}
    for _, v := range nums {
        if _, ok := freq[v]; ok {
            freq[v]++
        } else {
            freq[v] = 1
        }
    }

    keys := slices.Collect(maps.Keys(freq))
    sort.Slice(keys, func(i, j int) bool {
        return freq[keys[i]] >= freq[keys[j]] 
    })

    return keys[:k]
}
