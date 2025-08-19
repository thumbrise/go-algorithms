package problems

import (
	"strconv"
)

func compress(chars []byte) int {
	if len(chars) == 0 {
		return 0
	}

	j := 0
	count := 1

	for i := 1; i <= len(chars); i++ {
		if i < len(chars) && chars[i] == chars[i-1] {
			count++
		} else {
			chars[j] = chars[i-1]
			j++

			if count == 1 {
				continue
			}

			for _, c := range []byte(strconv.Itoa(count)) {
				chars[j] = c
				j++
			}

			count = 1

		}
	}

	return j
}
