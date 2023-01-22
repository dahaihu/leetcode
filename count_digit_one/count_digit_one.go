package countdigitone

func countDigitOne(n int) int {
	cur, count := 1, 0
	for cur <= n {
		left, val, right := (n/cur)/10, (n/cur)%10, n%cur
		switch {
		case val > 1:
			count += left*cur + cur
		case val == 1:
			count += left*cur + right + 1
		case val == 0:
			count += left * cur
		}
		cur = cur * 10
	}
	return count
}
