package longest_valid_parentheses

func longestValidParentheses(s string) int {
	dp := make([]int, len(s))
	var out int
	for i := 1; i < len(s); i++ {
		k := s[i]
		if k != ')' {
			continue
		}
		if s[i-1] == '(' {
			dp[i] = 2
			if i >= 2 {
				dp[i] += dp[i-2]
			}
		} else {
			dp[i] = dp[i-1]
			if preStart := i - dp[i-1]; preStart > 0 && s[preStart-1] == '(' {
				dp[i] += 2
				if preStart-2 >= 0 {
					dp[i] += dp[preStart-2]
				}
			}
		}
		if dp[i] > out {
			out = dp[i]
		}
	}
	return out
}
