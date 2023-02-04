package kmp

func buildJumpTable(s string) []int {
	table := make([]int, len(s))
	table[0] = 1
	for length := 2; length < len(table); length++ {
		value := s[:length]
		var maxMatchLength int
		for innerLen := length - 1; innerLen > 0; innerLen-- {
			if value[:innerLen] == value[len(value)-innerLen:] {
				maxMatchLength = innerLen
				break
			}
		}
		if maxMatchLength == 0 {
			table[length-1] = length
		} else {
			table[length-1] = length - maxMatchLength
		}
	}
	return table
}

func kmp(s string, target string) int {
	jumpTable := buildJumpTable(target)
	i, j := 0, 0
	for j < len(s) {
		if s[j] == target[j-i] {
			if matchedLen := j - i + 1; matchedLen == len(target) {
				return i
			}
			j++
		} else {
			if i == j {
				i += 1
				j += 1
			} else {
				i += jumpTable[j-i-1]
			}
		}
	}
	if matchedLen := j - i + 1; matchedLen == len(target) {
		return i
	}
	return -1
}
