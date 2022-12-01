package longest_valid_parentheses

func longestValidParentheses(s string) int {
	mark := make([]int, len(s))
	result := 0
	for i := 1; i < len(s); i++ {
		if s[i] == ')' {
			if s[i-1] == '(' {
				if i > 2 {
					mark[i] = mark[i-2] + 2
				} else {
					mark[i] = 2
				}
			} else if i > mark[i-1] && s[i-mark[i-1]-1] == '(' {
				mark[i] = mark[i-1] + 2
				if i-mark[i-1]-2 > 0 {
					mark[i] += mark[i-mark[i-1]-2]
				}
			}
			if mark[i] > result {
				result = mark[i]
			}
		}
	}
	return result
}
