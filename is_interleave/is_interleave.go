package isinterleave

func isInterleave(s1, s2, s3 string) bool {
	s1len, s2len, s3len := len(s1), len(s2), len(s3)
	if s1len+s2len != s3len {
		return false
	}
	if s1len == 0 {
		return s2 == s3
	}
	if s2len == 0 {
		return s1 == s3
	}

	mark := make([][]bool, s1len+1)
	for i := 0; i <= s1len; i++ {
		mark[i] = make([]bool, s2len+1)
	}
	mark[0][0] = true
	for i := 1; i <= s1len; i++ {
		mark[i][0] = s1[:i] == s3[:i]
	}
	for i := 1; i <= s2len; i++ {
		mark[0][i] = s2[:i] == s3[:i]
	}
	for i := 1; i <= s1len; i++ {
		for j := 1; j <= s2len; j++ {
			mark[i][j] = (s1[i-1] == s3[i+j-1] && mark[i-1][j]) ||
				(s2[j-1] == s3[i+j-1] && mark[i][j-1])
		}
	}
	return mark[s1len][s2len]
}
