func maxArea(heights []int) int {
	s := 0
	e := len(heights) - 1
	res := -1
	for s < e {
		a := int(math.Min(float64(heights[s]), float64(heights[e]))) * (e - s)
		if a > res {
			res = a
		}
		if heights[s] >= heights[e] {
			e--
		} else {
			s++
		}
	}
	return res
}
