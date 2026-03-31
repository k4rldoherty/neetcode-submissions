func productExceptSelf(nums []int) []int {
	// 1 2 4 6 8
    pre := make([]int, len(nums))
	pre[0] = 1
	pre[1] = nums[0]
	// pre = 1 1 2 8 48
	for i := 2; i < len(pre); i++ {
		pre[i] = nums[i-1] * pre[i-1]
	}

    suf := make([]int, len(nums))
	suf[len(nums) - 1] = 1
	suf[len(nums) - 2] = nums[len(nums) - 1]
	for i := len(nums) - 3; i >= 0; i-- {
		suf[i] = nums[i+1] * suf[i+1]
	}
    res := make([]int, len(nums))
    for i := range res {		
		res[i] = pre[i] * suf[i]        
    }
    return res
}
