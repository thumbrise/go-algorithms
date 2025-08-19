package problems

func modifyBlocksSum(data []byte) int {
	if len(data) == 0 {
		return 0
	}
	j := 0

	var count = 1
	for i := 1; i <= len(data); i++ {
		if i < len(data) && data[i] == data[i-1] {
			count++
		} else {
			if count == 1 {
				data[j] = data[i-1]
				j++
				continue
			}

			data[j] = byte(int(data[i-1]) + count)
			j++

			count = 1
		}
	}

	return j
}
