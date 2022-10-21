package longest_palindrome

func longestPalindrome(s string) string {
	if len(s) == 0 {
		return ""
	}
	result := s[:1]
	mark := make([][]bool, len(s))
	for i := 0; i < len(s); i++ {
		mark[i] = make([]bool, len(s))
	}
	for i := 0; i < len(s); i++ {
		mark[0][i] = true
	}
	for i := 0; i < len(s)-1; i++ {
		mark[1][i] = s[i] == s[i+1]
		if len(result) == 1 && mark[1][i] {
			result = s[i : i+2]
		}
	}
	for length := 2; length < len(s); length++ {
		for start := 0; start < len(s)-length; start++ {
			end := start + length
			mark[length][start] = s[start] == s[end] && mark[length-2][start+1]
			if mark[length][start] && length+1 > len(result) {
				result = s[start : start+length+1]
			}
		}
	}
	return result
}
