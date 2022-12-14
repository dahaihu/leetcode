package isinterleave

func isInterleave(s1, s2, s3 string) bool {
	if len(s1)+len(s2) != len(s3) {
		return false
	}
	mark := make([][]bool, len(s1)+1)
	for i := 0; i < len(s1)+1; i++ {
		mark[i] = make([]bool, len(s2)+1)
	}
	mark[0][0] = true
	for i := 1; i < len(s1)+1; i++ {
		mark[i][0] = s1[:i] == s3[:i]
	}
	for i := 1; i < len(s2)+1; i++ {
		mark[0][i] = s2[:i] == s3[:i]
	}
	for i := 1; i < len(s1)+1; i++ {
		for j := 1; j < len(s2)+1; j++ {
			s3idx := i + j
			mark[i][j] = (s1[i-1] == s3[s3idx-1] && mark[i-1][j]) ||
				(s2[j-1] == s3[s3idx-1] && mark[i][j-1])
		}
	}
	return mark[len(s1)][len(s2)]
}
