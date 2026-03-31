func productExceptSelf(nums []int) []int {
    res := make([]int, len(nums))
    for i := range res {
        prod := 1
        for j := range nums {
            if j == i {
                continue
            }
            prod *= nums[j]
        }
        res[i] = prod
    }
    return res
}
