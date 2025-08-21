package problems

func maxVowels(s string, k int) int {
	isVowel := func(ch byte) bool {
		switch ch {
		case 'a', 'e', 'i', 'o', 'u':
			return true
		}
		return false
	}

	count := 0
	for i := 0; i < k; i++ {
		if isVowel(s[i]) {
			count++
		}
	}

	maxCount := count

	for left, right := 1, k; right < len(s); left, right = left+1, right+1 {
		if isVowel(s[left-1]) {
			count--
		}
		if isVowel(s[right]) {
			count++
		}
		if count > maxCount {
			maxCount = count
			if maxCount == k {
				return maxCount
			}
		}
	}

	return maxCount
}
