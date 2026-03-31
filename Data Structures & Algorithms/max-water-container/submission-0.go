func maxArea(heights []int) int {
	a := -1
	for i := 0; i < len(heights) - 1; i++ {
		for j := len(heights) - 1; j > i; j-- {
			var area int
			if heights[i] >= heights[j] {
				area = (heights[j]) * (j - i)
			} else {
				area = (heights[i]) * (j - i)
			}

			if area > a {
				a = area
			}
		}
	}
	return a
}
