package countdigitone

func countDigitOne(n int) int {
	cur := 1
	count := 0
	for cur <= n {
		left, val, right := n/(10*cur), (n/cur)%10, n%cur
		switch {
		case val == 0:
			count += left * cur
		case val == 1:
			count += left*cur + right + 1
		case val > 1:
			count += left*cur + cur
		}
		cur = cur * 10
	}
	return count
}
