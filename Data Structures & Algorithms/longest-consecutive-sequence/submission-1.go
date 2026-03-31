func longestConsecutive(nums []int) int {
	if len(nums) == 0 { return 0 }
	
	sort.Slice(nums, func(i, j int) bool {
		return nums[i] < nums[j]
	})

	max := 1
	curr := 1

	for i := 1; i < len(nums); i++ {
		if nums[i-1] == (nums[i] - 1) {
			curr++
		} else if nums[i-1] == nums[i] {
			continue
		} else {
			if curr > max { max = curr }
			curr = 1
		}
	}

	if curr > max { max = curr }

	return max
}
