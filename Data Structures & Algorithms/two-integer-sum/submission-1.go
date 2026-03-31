func twoSum(nums []int, target int) []int {
    m := map[int]int{}
    for i, val := range nums {
        if _, ok := m[val]; ok {
            return []int{m[val], i}
        } else {

            m[target - val] = i
        }
    }
    return []int{-1, -1}
}
