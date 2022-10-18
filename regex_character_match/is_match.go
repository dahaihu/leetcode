package regex_character_match

func isMatch(s string, p string) bool {
	mark := make([][]bool, len(p)+1)
	for i := 0; i <= len(p); i++ {
		mark[i] = make([]bool, len(s)+1)
	}
	mark[0][0] = true
	for i := 2; i <= len(p); i++ {
		if p[i-1] == '*' {
			mark[i][0] = mark[i-2][0]
		}
	}
	for i := 1; i <= len(p); i++ {
		for j := 1; j <= len(s); j++ {
			switch p[i-1] {
			case '*':
				mark[i][j] = mark[i-2][j] || mark[i-1][j] || ((p[i-2] == s[j-1] || p[i-2] == '.') && mark[i][j-1])
			default:
				mark[i][j] = mark[i-1][j-1] && (p[i-1] == '.' || p[i-1] == s[j-1])
			}
		}
	}
	return mark[len(p)][len(s)]
}
