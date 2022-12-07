package longest_valid_parentheses

func longestValidParentheses(s string) int {
	dp := make([]int, len(s))
	longest := 0
	for i := 1; i < len(s); i++ {
		if s[i] != ')' {
			continue
		}
		if s[i-1] == '(' {
			dp[i] = 2
			if i > 2 {
				dp[i] += dp[i-2]
			}
		} else if dp[i-1] > 0 {
			preStart := i - dp[i-1]
			if preStart-1 >= 0 && s[preStart-1] == '(' {
				dp[i] = 2 + dp[i-1]
				if preStart-2 >= 0 && dp[preStart-2] >= 0 {
					dp[i] += dp[preStart-2]
				}
			}
		}
		if dp[i] > longest {
			longest = dp[i]
		}
	}
	return longest
}
