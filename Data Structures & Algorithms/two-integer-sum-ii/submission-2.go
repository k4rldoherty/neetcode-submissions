func twoSum(numbers []int, target int) []int {
	s := 0
	e := len(numbers) - 1
	for s < e {
		if numbers[s] + numbers[e] > target { 
			e-- 
		} else if numbers[s] + numbers[e] < target { 
			s++ 
		} else { return []int{ s+1, e+1 }}
	}
	return []int{}
}
