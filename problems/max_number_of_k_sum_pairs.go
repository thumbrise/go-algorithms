package problems

func maxOperations(nums []int, k int) int {
	result := 0

	summands := make(map[int]int)
	for _, num := range nums {
		if summand, ok := summands[num]; ok && summand > 0 {
			result++
			summands[num]--
			continue
		}

		summands[k-num]++
	}

	return result
}
