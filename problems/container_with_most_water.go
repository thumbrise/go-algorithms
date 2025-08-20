package problems

func maxArea(height []int) int {
	result := 0
	left, right := 0, len(height)-1
	for left < right {
		w := right - left

		if height[left] <= height[right] {
			result = max(result, w*height[left])
			left++
		} else {
			result = max(result, w*height[right])
			right--
		}
	}

	return result
}
