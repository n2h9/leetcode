package solution

func maxArea(height []int) int {
	min := func(a,b int) int {
		if a < b {
			return a
		}
		return b
	}
	maxAreaVal := 0
	for i := 0; i < len(height); i++ {
		currMaxAreaVal := 0
		for j := i+1; j < len(height); j++ {
			h := min(height[i], height[j])
			if v := h * (j-i); v > currMaxAreaVal {
				currMaxAreaVal = v
			}
		}
		if currMaxAreaVal > maxAreaVal {
			maxAreaVal = currMaxAreaVal
		}
	}
	return maxAreaVal
}
