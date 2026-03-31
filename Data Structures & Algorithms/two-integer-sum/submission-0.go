func twoSum(nums []int, target int) []int {
    // m = {[4,0], []}
    m := map[int]int{}
    for i, val := range nums {
        // 7 - 4 = 3
        diff := target - val
        if _, ok := m[val]; ok {
            return []int{m[val], i}
        } else {
            m[diff] = i
        }
    }
    return []int{-1, -1}
}
